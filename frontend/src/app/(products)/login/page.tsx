import React from "react";

export default function Home() {
  return (
    <div className="min-h-screen flex bg-beige">
      {/* color bar  */}
      <div className="w-1/6 bg-brown h-screen"></div>
      <div className="w-1/6 bg-pink h-screen"></div>

      {/* Right side (form section) */}
      <div className="flex-1 flex items-center justify-center bg-beige p-10">
        <div>
          <h1 className="text-6xl font-bold text-textbrown mb-8">Time Table</h1>
          <form className="space-y-6">
            <div>
              <label className="block text-textbrown text-lg">名前</label>
              <input
                type="text"
                className="w-full px-4 py-2 border border-brown rounded-md"
              />
            </div>
            <div>
              <label className="block text-textbrown text-lg">パスワード</label>
              <input
                type="password"
                className="w-full px-4 py-2 border border-brown rounded-md"
              />
            </div>
            <div className="flex space-x-4">
              <button
                type="submit"
                className="bg-pink px-6 py-2 rounded-md text-textbrown"
              >
                登録
              </button>
              <button
                type="button"
                className="bg-pink px-6 py-2 rounded-md text-textbrown"
              >
                ログイン
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
