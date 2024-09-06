import { useEffect } from "react";
import { ApolloClient, InMemoryCache, gql } from "@apollo/client";

const client = new ApolloClient({
  uri: "http://localhost:8080/playground", //apolloclientが通信するgraphqlのエンドポイント
  cache: new InMemoryCache(),
});
