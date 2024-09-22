import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";

// HTTPリンクの作成
const httpLink = createHttpLink({
  uri: "http://localhost:8080/graphql", // GraphQLエンドポイントを指定
});

// 認証トークンをリクエストに追加するためのリンク
const authLink = setContext((_, { headers }) => {
  if (typeof window !== "undefined") {
    const token = localStorage.getItem("authToken"); // ローカルストレージからトークンを取得
    return {
      headers: {
        ...headers,
        authorization: token ? `Bearer ${token}` : "", // トークンが存在すればヘッダーに追加
      },
    };
  }
  return { headers };
});

// ApolloClientの作成
const client = new ApolloClient({
  link: authLink.concat(httpLink), // 認証ヘッダーをHTTPリンクに追加
  cache: new InMemoryCache(),
});

export default client;
