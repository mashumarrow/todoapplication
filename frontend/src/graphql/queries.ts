import { gql } from "@apollo/client";

export const REGISTER_USER = gql`
  mutation RegisterUser($input: NewUser!) {
    registerUser(input: $input) {
      userid
      name
      password
    }
  }
`;

export const LOGIN_USER = gql`
  query LoginUser($name: String!, $password: String!) {
    loginUser(name: $name, password: $password) {
      userid
      name
    }
  }
`;
