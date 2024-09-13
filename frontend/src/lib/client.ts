import React from "react";
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  HttpLink,
  from,
} from "@apollo/client";
import dynamic from "next/dynamic";
//import Login from "./pages/login/login";
import { onError } from "@apollo/client/link/error";
import Getusers from "./components/Getusers";

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

const link = from([
  errorLink,
  new HttpLink({ uri: "http://localhost:8080/graphql" }),
]);

// const errorLink = on Error(({graphqlErrors, networkError}) => {
// if(graphqlErrors){
//     graphqlErrors.map(({message, location, path}) =>{
//         alert(`GraphQL error ${message}`);

//     })
// }
// })
// Apollo Client を初期化
const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: link,
});

//const Login = dynamic(() => import("./pages/login/login"), { ssr: false });

const App: React.FC = () => {
  return (
    <ApolloProvider client={client}>
      <div className="App">
        <Getusers />
      </div>
    </ApolloProvider>
  );
};

export default App;
