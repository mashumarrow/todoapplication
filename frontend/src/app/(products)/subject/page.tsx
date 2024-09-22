"use client"; // 必須

import React, { useState, useEffect } from "react";
import { useMutation, useQuery, ApolloError } from "@apollo/client";
import { CREATE_SCHEDULE, GET_SCHEDULES } from "../../../graphql/queries";

type ScheduleEntry = {
  subject: string;
  classroom: string;
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
  const [selectedCell, setSelectedCell] = useState<{
    dayofweek: string;
    period: number;
  }>({ dayofweek: "", period: 0 });
  const [subject, setSubject] = useState("");
  const [classroom, setClassroom] = useState("");

  // Apollo Client: データ取得
  const { data: scheduleData, refetch } = useQuery(GET_SCHEDULES);

  // Apollo Client: ミューテーション
  const [createSchedule] = useMutation(CREATE_SCHEDULE, {
    update(cache, { data: { createschedule } }) {
      // キャッシュを手動で更新
      const existingSchedules: any = cache.readQuery({ query: GET_SCHEDULES });
      const newScheduleEntry = createschedule;

      cache.writeQuery({
        query: GET_SCHEDULES,
        data: {
          schedules: [...existingSchedules.schedules, newScheduleEntry],
        },
      });
    },
  });

  useEffect(() => {
    if (scheduleData && scheduleData.schedules) {
      const newSchedule: Schedule = {};

      scheduleData.schedules.forEach(
        (item: {
          dayofweek: string;
          period: number;
          subjectname: string;
          classroomname: string;
        }) => {
          const key = `${item.dayofweek}${item.period}`;
          newSchedule[key] = {
            subject: item.subjectname,
            classroom: item.classroomname,
          };
        }
      );

      setSchedule(newSchedule);
    }
  }, [scheduleData]);

  // モーダルを開く関数
  const openModal = (dayofweek: string, period: number) => {
    setSelectedCell({ dayofweek, period });
    const key = `${dayofweek}${period}`;
    setSubject(schedule[key]?.subject || "");
    setClassroom(schedule[key]?.classroom || "");
    setIsModalOpen(true);
  };

  // モーダルを閉じる関数
  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedCell({ dayofweek: "", period: 0 });
    setSubject(""); // フィールドをリセット
    setClassroom(""); // フィールドをリセット
  };

  // スケジュールの更新
  const handleSave = async () => {
    const key = `${selectedCell.dayofweek}${selectedCell.period}`;

    try {
      // ミューテーションの実行
      await createSchedule({
        variables: {
          input: {
            dayofweek: selectedCell.dayofweek,
            period: parseInt(selectedCell.period.toString(), 10),
            subjectname: subject,
            classroomname: classroom,
          },
        },
      });

      // 手動で状態を更新してセルに反映
      setSchedule((prevSchedule) => ({
        ...prevSchedule,
        [key]: { subject, classroom },
      }));

      closeModal();
    } catch (error) {
      // errorがApolloErrorかどうかを確認
      if (error instanceof ApolloError) {
        if (error.graphQLErrors) {
          console.error("GraphQL Errors:", error.graphQLErrors);
        }
        if (error.networkError) {
          console.error("Network Error:", error.networkError);
        }
      } else {
        // ApolloError以外のエラーの場合
        console.error("An unknown error occurred:", error);
      }
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
                    className="border border-gray-300 w-24 h-24 bg-cream text-center cursor-pointer"
                    onClick={() => openModal(day, period)}
                  >
                    <div className="text-textbrown">{cellData.subject}</div>
                    <div className="text-textbrown">{cellData.classroom}</div>
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>

      {/* モーダル */}
      {isModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50">
          <div className="bg-cream p-6 rounded-md w-1/3">
            <h2 className="text-2xl mb-4 text-textbrown">
              {selectedCell.dayofweek} {selectedCell.period}
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
    </div>
  );
}
