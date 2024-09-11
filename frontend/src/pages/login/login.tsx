import React, { useState, useEffect } from "react";
import { useMutation, gql } from "@apollo/client";
import dynamic from "next/dynamic";

// GraphQL ミューテーション
const REGISTER_USER = gql`
  mutation RegisterUser($input: NewUser!) {
    registerUser(input: $input) {
      name
      password
    }
  }
`;

const Login: React.FC = () => {
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [clientReady, setClientReady] = useState(false); // クライアント側のみでApollo操作を行うためのフラグ

  const [registerUser] = useMutation(REGISTER_USER);

  useEffect(() => {
    // クライアントサイドでフラグをセット
    setClientReady(true);
  }, []);

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!clientReady) return; // サーバーサイドでの実行を防ぐ

    try {
      await registerUser({
        variables: {
          input: { name, password },
        },
      });
      setName("");
      setPassword("");
    } catch (err) {
      console.error("Failed to register user", err);
    }
  };

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Login functionality not yet implemented.");
    // ここにログイン処理を実装
  };

  return (
    <div className="min-h-screen flex">
      {/* カラーバー */}
      <div className="w-1/3 bg-gray-400"></div>
      <div className="w-1/3 bg-pink-200"></div>
      <div className="w-1/3 bg-cream-200"></div>

     
    </div>
  );
};

// dynamicを使ってSSRを無効にする
export default dynamic(() => Promise.resolve(Login), { ssr: false });
