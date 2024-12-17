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
  mutation LoginUser($name: String!, $password: String!) {
    loginUser(name: $name, password: $password) {
      userid
      name
      token
    }
  }
`;

export const CREATE_SCHEDULE = gql`
  mutation createschedule($input: NewSchedule!) {
    createschedule(input: $input) {
      subjectname
      classroomname
      dayofweek
      period
    }
  }
`;

export const GET_SCHEDULES = gql`
  query GetSchedules {
    schedules {
      userid
      subjectname
      classroomname
      dayofweek
      period
    }
  }
`;

export const GET_SUBJECTS = gql`
  query GetSubjects {
    subjects {
      subjectid
      subjectname
    }
  }
`;

// Todoを作成するミューテーション
export const CREATE_TODO = gql`
  mutation CreateTodo($input: NewTodo!) {
    createTodo(input: $input) {
      todoid
      userid
      title
      completed
      subjectname
      period
      classroomname
    }
  }
`;

export const GET_TODOS = gql`
  query GetTodos($userid: ID!) {
    todos(userid: $userid) {
      id
      todoid
      userid
      title
      completed
      subjectname
      classroomname
      period
    }
  }
`;
