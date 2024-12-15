"use client"; // 必須

import React, { useState, useEffect } from "react";
import { useMutation, useLazyQuery } from "@apollo/client";
import { CREATE_SCHEDULE } from "../../../graphql/queries";
import { CREATE_TODO } from "../../../graphql/queries";
import { ApolloError } from "@apollo/client";
import { title } from "process";

type ScheduleEntry = {
  subject: string;
  classroom: string;
  todos?: string[]; // Todoリストを追加
};

type Schedule = {
  [key: string]: ScheduleEntry;
};

const days = [
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
];
const periods = [1, 2, 3, 4, 5, 6];

export default function TimeTable() {
  const [schedule, setSchedule] = useState<Schedule>({});
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isTodoModalOpen, setIsTodoModalOpen] = useState(false); // Todoリスト用のモーダル
  const [selectedCell, setSelectedCell] = useState<{
    dayofweek: string;
    period: number;
  }>({
    dayofweek: "",
    period: 0,
  });
  const [subject, setSubject] = useState("");
  const [classroom, setClassroom] = useState("");

  const [todos, setTodos] = useState<string[]>([]);
  const [newTodo, setNewTodo] = useState(""); // 新しいTodoアイテム用
  const [savedToken, setSavedToken] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("authToken");
    setSavedToken(token);
    if (!token) {
      console.log("トークンが存在しません。ログインが必要です。");
    }
  }, []);

  // Apollo Client: ミューテーションとクエリ
  const [createSchedule] = useMutation(CREATE_SCHEDULE, {
    context: {
      headers: {
        Authorization: savedToken ? `Bearer ${savedToken}` : undefined,
      },
    },
  });
  const [createTodo] = useMutation(CREATE_TODO, {
    context: {
      headers: {
        Authorization: savedToken ? `Bearer ${savedToken}` : undefined,
      },
    },
  });

  //const [getTodos, { data }] = useLazyQuery(GET_TODOS); // Todoリスト取得クエリ

  const generateTodoId = (dayofweek: string, period: number) =>
    `${dayofweek}-${period}`;

  // モーダルを開く関数（教科名・教室名の追加用）
  const openModal = (dayofweek: string, period: number) => {
    setSelectedCell({ dayofweek, period });
    const key = `${dayofweek}${period}`;
    const cellData = schedule[key] || {};
    setSubject(cellData.subject || "");
    setClassroom(cellData.classroom || "");
    setIsModalOpen(true);
  };

  // Todoリスト用モーダルを開く関数
  const openTodoModal = (dayofweek: string, period: number) => {
    const todoid = generateTodoId(dayofweek, period);
    setSelectedCell({ dayofweek, period });
    const key = `${dayofweek}${period}`;
    const cellData = schedule[key] || {};
    setTodos(cellData.todos || []);
    setIsTodoModalOpen(true);
    console.log("Generated Todo ID:", todoid);
  };

  // モーダルを閉じる関数
  const closeModal = () => {
    setIsModalOpen(false);
    setIsTodoModalOpen(false);
    setSelectedCell({ dayofweek: "", period: 0 });
    setSubject("");
    setClassroom("");
    setTodos([]);
  };

  // Todoリストに新しいアイテムを追加
  const addTodo = async () => {
    if (newTodo.trim() !== "") {
      try {
        // if (!savedToken) {
        //   console.error("トークンが見つかりません。ログインが必要です。");
        //   return;
        // }

        const { data } = await createTodo({
          variables: {
            input: {
              title: newTodo,
              completed: false,
              //dayofweek: selectedCell.dayofweek, // selectedCell から取得
              //period: selectedCell.period, // selectedCell から取得
              todoid: `${selectedCell.dayofweek}-${selectedCell.period}`,
              period: parseInt(selectedCell.period.toString(), 10),
              //subjectname: subject,
              //classroomname: classroom,
            },
          },
        });

        if (data && data.createTodo) {
          const key = `${selectedCell.dayofweek}${selectedCell.period}`;

          // scheduleを更新
          setSchedule((prevSchedule) => ({
            ...prevSchedule,
            [key]: {
              ...prevSchedule[key],
              todos: [...(prevSchedule[key]?.todos || []), newTodo],
            },
          }));

          // ローカルのtodosも即時更新
          setTodos((prevTodos) => [...prevTodos, newTodo]);
          setNewTodo(""); // 入力フィールドをクリア
        }
      } catch (error) {
        console.error("Todo作成エラー:", error);
        alert("Todoの作成中にエラーが発生しました");
      }
    }
  };

  // Todoリストからアイテムを削除
  const removeTodo = (index: number) => {
    const key = `${selectedCell.dayofweek}${selectedCell.period}`;

    setSchedule((prevSchedule) => {
      if (!prevSchedule[key] || !prevSchedule[key].todos) {
        console.warn(`Key "${key}" or todos is undefined.`);
        return prevSchedule; // 何も変更せずに返す
      }

      const updatedTodos = prevSchedule[key].todos.filter(
        (_, i) => i !== index
      );

      return {
        ...prevSchedule,
        [key]: {
          ...prevSchedule[key],
          todos: updatedTodos,
        },
      };
    });

    // ローカルのtodosも更新
    setTodos((prevTodos) => prevTodos.filter((_, i) => i !== index));
  };
  // スケジュールの保存処理
  const handleSave = async () => {
    const key = `${selectedCell.dayofweek}${selectedCell.period}`;

    try {
      const { data } = await createSchedule({
        variables: {
          input: {
            dayofweek: selectedCell.dayofweek,
            period: parseInt(selectedCell.period.toString(), 10),
            subjectname: subject,
            classroomname: classroom,
          },
        },
      });

      if (data) {
        setSchedule((prevSchedule) => ({
          ...prevSchedule,
          [key]: {
            subject,
            classroom,
          },
        }));
        closeModal();
      }
    } catch (error) {
      console.error("エラー:", error);
      alert("エラーが発生しました");
    }
  };

  return (
    <div className="min-h-screen flex flex-col items-center bg-beige p-10">
      <h1 className="text-4xl font-bold text-textbrown mb-8">Time Table</h1>
      <table className="table-auto border-collapse border border-gray-300">
        <thead>
          <tr>
            <th className="border border-gray-300 w-16 h-16 bg-brown"></th>
            {days.map((day, index) => (
              <th
                key={index}
                className="border border-gray-300 w-24 h-16 bg-pink text-textbrown"
              >
                {day}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {periods.map((period) => (
            <tr key={period}>
              <td className="border border-gray-300 w-16 h-16 bg-pink text-center text-textbrown">
                {period}
              </td>
              {days.map((day, index) => {
                const cellKey = `${day}${period}`;
                const cellData = schedule[cellKey] || {};
                return (
                  <td
                    key={index}
                    className="border border-gray-300 w-24 h-24 bg-cream relative text-center cursor-pointer"
                    onClick={() => openModal(day, period)} // セルをクリックで教科名・教室名を追加
                  >
                    <div className="text-textbrown">{cellData.subject}</div>
                    <div className="text-textbrown">{cellData.classroom}</div>
                    <button
                      className="absolute bottom-1 right-1 text-gray-600"
                      onClick={(e) => {
                        e.stopPropagation(); // セルクリックイベントを止める
                        openTodoModal(day, period); // Todoリスト追加用モーダルを開く
                      }}
                    >
                      ＋
                    </button>
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>

      {/* 教科名・教室名追加用モーダル */}
      {isModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50">
          <div className="bg-cream p-6 rounded-md w-1/3">
            <h2 className="text-2xl mb-4 text-textbrown">
              {selectedCell.dayofweek} {selectedCell.period} 時限
            </h2>
            <div className="mb-4">
              <label className="block text-textbrown mb-2">教科名</label>
              <input
                type="text"
                className="w-full px-3 py-2 border border-gray-300 rounded"
                value={subject}
                onChange={(e) => setSubject(e.target.value)}
              />
            </div>
            <div className="mb-4">
              <label className="block text-textbrown mb-2">教室場所</label>
              <input
                type="text"
                className="w-full px-3 py-2 border border-gray-300 rounded"
                value={classroom}
                onChange={(e) => setClassroom(e.target.value)}
              />
            </div>
            <div className="flex justify-end">
              <button
                className="bg-pink px-4 py-2 rounded mr-2 text-textbrown"
                onClick={handleSave}
              >
                決定
              </button>
              <button
                className="bg-gray-300 px-4 py-2 rounded text-textbrown"
                onClick={closeModal}
              >
                キャンセル
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Todoリスト追加用モーダル */}
      {isTodoModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50">
          <div className="bg-cream p-6 rounded-md w-1/3">
            <h2 className="text-2xl mb-4 text-textbrown">
              {selectedCell.dayofweek} {selectedCell.period} 時限 Todoリスト
            </h2>
            <div className="mb-4">
              <ul>
                {todos.map((todo, index) => (
                  <li key={index} className="flex justify-between items-center">
                    <span>{todo}</span>
                    <button
                      className="text-red-500 ml-2"
                      onClick={() => removeTodo(index)}
                    >
                      削除
                    </button>
                  </li>
                ))}
              </ul>
              <input
                type="text"
                className="w-full px-3 py-2 border border-gray-300 rounded mt-2"
                placeholder="新しいTodoを追加"
                value={newTodo}
                onChange={(e) => setNewTodo(e.target.value)}
              />
              <button
                onClick={addTodo}
                className="bg-blue-500 text-white px-4 py-2 rounded mt-2"
              >
                追加
              </button>
            </div>
            <div className="flex justify-end">
              <button
                className="bg-gray-300 px-4 py-2 rounded text-textbrown"
                onClick={closeModal}
              >
                キャンセル
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
