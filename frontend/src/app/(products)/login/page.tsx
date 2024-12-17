"use client";
import { ApolloProvider, useLazyQuery } from "@apollo/client";
import client from "../../../lib/apollo"; // apollo.tsからclientをインポート
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import "../../globals.css";
import { useMutation } from "@apollo/client";
import { REGISTER_USER, LOGIN_USER, GET_TODOS } from "../../../graphql/queries";

export default function Home() {
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  // ユーザー登録用のMutation
  const [registerUser, { loading: registerLoading, error: registerError }] =
    useMutation(REGISTER_USER, {
      errorPolicy: "all",
    });

  // ユーザーログイン用のMutation
  const [loginUser, { loading: loginLoading, error: loginError }] = useMutation(
    LOGIN_USER,
    {
      errorPolicy: "all",
    }
  );
  // ユーザーのTodoデータをフェッチするLazyQuery
  const [
    getTodos,
    { data: todosData, loading: todosLoading, error: todosError },
  ] = useLazyQuery(GET_TODOS);

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      console.log("送信データ（登録）:", { name, password });
      await registerUser({ variables: { input: { name, password } } });
      alert("登録が成功しました！");
      setName("");
      setPassword("");
    } catch (err) {
      console.error("登録エラー:", err);
    }
  };

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      // loginUser ミューテーションを実行
      const { data: loginResponse } = await loginUser({
        variables: { name, password },
      });

      // loginResponse が正しく返ってきているか確認
      console.log("Login response:", JSON.stringify(loginResponse, null, 2));

      // トークンがある場合はローカルストレージに保存
      if (loginResponse?.loginUser?.token) {
        localStorage.setItem("authToken", loginResponse.loginUser.token);
        localStorage.setItem("userid", loginResponse.loginUser.userid);
        console.log("保存されたトークン:", localStorage.getItem("authToken"));
        console.log("保存されたユーザーID:", localStorage.getItem("userid"));

        alert("ログイン成功");

        // ユーザーIDを取得し、Todoデータをフェッチ
        const userid = loginResponse.loginUser.userid;
        console.log("取得したユーザーID:", userid);

        // getTodos クエリを実行してデータを取得
        const { data: todosResponse } = await getTodos({
          variables: { userid },
        });

        // データが正しく取得されているか確認
        if (todosResponse?.todos) {
          console.log("取得したTodoデータ:", todosResponse.todos);
        } else {
          console.error("Todoデータが取得できませんでした:", todosResponse);
        }

        // ページ遷移
        router.push("/subject");
      } else {
        alert("ログイン失敗");
      }
    } catch (err) {
      console.error("ログインエラー:", err);
    }
  };

  return (
    <ApolloProvider client={client}>
      <div className="min-h-screen flex bg-beige">
        <div className="w-1/6 bg-brown h-screen"></div>
        <div className="w-1/6 bg-pink h-screen"></div>
        <div className="flex-1 flex items-center justify-center bg-beige p-10">
          <div>
            <h1 className="text-6xl font-bold text-textbrown mb-8">
              Time Table
            </h1>
            <form className="space-y-6" onSubmit={handleRegister}>
              <div>
                <label className="block text-textbrown text-lg">名前</label>
                <input
                  type="text"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  className="w-full px-4 py-2 border border-brown rounded-md"
                />
              </div>
              <div>
                <label className="block text-textbrown text-lg">
                  パスワード
                </label>
                <input
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full px-4 py-2 border border-brown rounded-md"
                />
              </div>
              <div className="flex space-x-4">
                <button
                  type="submit"
                  className="bg-pink px-6 py-2 rounded-md text-textbrown"
                  disabled={registerLoading}
                >
                  登録
                </button>
                <button
                  type="button"
                  className="bg-pink px-6 py-2 rounded-md text-textbrown"
                  onClick={handleLogin}
                  disabled={loginLoading}
                >
                  ログイン
                </button>
              </div>
              {(registerError || loginError) && (
                <p className="text-textbrown-500">エラーが発生しました</p>
              )}
            </form>
          </div>
        </div>
      </div>
    </ApolloProvider>
  );
}
