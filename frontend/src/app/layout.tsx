"use client";
import { ReactNode } from "react";
import "./globals.css";
import { ApolloProvider } from "@apollo/client";
import client from "../lib/apollo";

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="ja">
      <body>
        <ApolloProvider client={client}>{children}</ApolloProvider>
      </body>
    </html>
  );
}
