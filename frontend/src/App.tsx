import React from "react";
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import dynamic from "next/dynamic";
import Login from "./pages/login/login";
import { onError } from "@apollo/client/link/error";
import { ApolloLink } from "@apollo/client";

const errorLink = onError(({ graphQLErrors, networkError }) => {
  if (graphQLErrors) {
    graphQLErrors.forEach(({ message, locations, path }) => {
      console.error(
        `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
      );
    });
  }
  if (networkError) {
    console.error(`[Network error]: ${networkError}`);
  }
});
// Apollo Client を初期化
const client = new ApolloClient({
  uri: "http://localhost:8080/graphql", // GraphQL のエンドポイント
  cache: new InMemoryCache(),
  link: ApolloLink.from([errorLink]),
});

const App: React.FC = () => {
  return (
    <ApolloProvider client={client}>
      <div className="App">
        <Login /> {/* Loginコンポーネントを表示 */}
      </div>
    </ApolloProvider>
  );
};

// dynamic を使って SSR を無効にする
export default dynamic(() => Promise.resolve(App), { ssr: false });
