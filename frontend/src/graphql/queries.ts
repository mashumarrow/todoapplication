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
    }
  }
`;

export const CREATE_SCHEDULE = gql`
  mutation createschedule($input: NewSchedule!) {
    createschedule(input: $input) {
      userid
      subjectid
      subjectname
      classroomid
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
      subjectid
      subjectname
      classroomname
      classroomid
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
