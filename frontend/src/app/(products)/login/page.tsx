"use client";
import { ApolloProvider } from "@apollo/client";
import client from "../../../lib/apollo"; // apollo.tsからclientをインポート
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import "../../globals.css";
import { useMutation } from "@apollo/client";
import { REGISTER_USER, LOGIN_USER } from "../../../graphql/queries";

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
      console.log("Login response:", loginResponse); // ここでログを確認

      // トークンがある場合はローカルストレージに保存
      if (loginResponse?.loginUser?.token) {
        localStorage.setItem("authToken", loginResponse.loginUser.token);
        console.log("保存されたトークン:", localStorage.getItem("authToken"));
        alert("ログイン成功！");
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
