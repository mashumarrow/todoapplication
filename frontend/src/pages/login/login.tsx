import React {useState} from 'react';
import {useMutation, gql} from '@apollo/client'

const REGISTER_USER = gql`
    RegisterUser($input: NewUser!) {
    registerUser(input: $input) {
      userid
      name
    }
  }
`;

const GET_USERS = gql`
  query Users {
    users {
      userid
      name
    }
  }
`;
const Login: React.FC = () => {
    const [userid, setUserid] = useState('');
    const [name, setName] = useState('');
    const [password, setPassword] = useState('');
    const [registerUser] = useMutation(REGISTER_USER, {
      refetchQueries: [{ query: GET_USERS }], // 新しいユーザーを登録後にユーザー一覧を再取得
    });
    const { loading, error, data } = useQuery(GET_USERS);
  
    const handleSubmit = async (e: React.FormEvent) => {
      e.preventDefault();
      try {
        await registerUser({
          variables: {
            input: { userid, name, password },
          },
        });
        setUserid('');
        setName('');
        setPassword('');
      } catch (err) {
        console.error('Failed to register user', err);
      }
    };

    return(
        <div>

        </div>
    );}