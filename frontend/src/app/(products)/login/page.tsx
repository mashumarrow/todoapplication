"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import "../../globals.css";
import { useMutation, useQuery } from "@apollo/client";
import { REGISTER_USER, LOGIN_USER } from "../../../graphql/queries";

export default function Home() {
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [registerUser, { loading: registerLoading, error: registerError }] =
    useMutation(REGISTER_USER, {
      errorPolicy: "all",
    });
  const {
    data: loginData,
    refetch: loginUser,
    loading: loginLoading,
    error: loginError,
  } = useQuery(LOGIN_USER, {
    variables: { name, password },
    skip: true, // クエリの自動実行をスキップ
    errorPolicy: "all",
  });

  const router = useRouter();

  // ユーザー登録＆自動ログイン
  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      console.log("送信データ（登録）:", { name, password });

      // ユーザーを登録
      const { data: registerData } = await registerUser({
        variables: { input: { name, password } },
      });
      alert("登録が成功しました！");

      // 登録後、ログイン処理を実行
      const { data: loginResponse } = await loginUser(); // refetchでクエリを実行

      // ログインが成功したらトークンをlocalStorageに保存
      if (loginResponse?.loginUser?.token) {
        localStorage.setItem("authToken", loginResponse.loginUser.token);
        alert("ログイン成功！");
        router.push("/subject"); // ログイン成功後にページ遷移
      } else {
        alert("登録後の自動ログインに失敗しました。");
      }

      // 必要に応じてフォームをリセット
      setName("");
      setPassword("");
    } catch (err) {
      if (err instanceof Error) {
        console.error("登録またはログインエラー:", err.message);
        alert(`エラー: ${err.message}`);
      } else {
        console.error("予期しないエラー:", err);
        alert("予期しないエラーが発生しました。");
      }
    }
  };

  // ユーザーログイン
  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      console.log("送信データ（ログイン）:", { name, password });

      const { data: loginResponse } = await loginUser({
        variables: { name, password },
      });

      if (loginResponse?.loginUser?.token) {
        localStorage.setItem("authToken", loginResponse.loginUser.token); // トークンをlocalStorageに保存
        alert("ログイン成功！");
        router.push("/subject"); // ログイン成功後にページ遷移
      } else {
        alert("ユーザー名またはパスワードが間違っています。");
      }
    } catch (err) {
      if (err instanceof Error) {
        console.error("ログインエラー:", err.message);
        alert(`ログインに失敗しました: ${err.message}`);
      } else {
        console.error("予期しないエラー:", err);
        alert("予期しないエラーが発生しました。");
      }
    }
  };

  return (
    <div className="min-h-screen flex bg-beige">
      {/* color bar  */}
      <div className="w-1/6 bg-brown h-screen"></div>
      <div className="w-1/6 bg-pink h-screen"></div>

      {/* Right side  */}
      <div className="flex-1 flex items-center justify-center bg-beige p-10">
        <div>
          <h1 className="text-6xl font-bold text-textbrown mb-8">Time Table</h1>
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
              <label className="block text-textbrown text-lg">パスワード</label>
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
  );
}
