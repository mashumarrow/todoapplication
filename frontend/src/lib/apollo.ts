import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";

// HTTPリンクの作成
const httpLink = createHttpLink({
  uri: "http://localhost:8080/graphql", // GraphQLエンドポイントを指定
});

// 認証トークンをリクエストヘッダーに追加
const authLink = setContext((_, { headers }) => {
  // ローカルストレージからトークンを取得
  const token = localStorage.getItem("authToken");

  // ヘッダーにトークンを追加
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    },
  };
});

// Apollo Clientのインスタンス作成
const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
});

export default client;
