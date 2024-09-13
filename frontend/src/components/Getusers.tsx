import React, { useEffect } from "react";
import { useQuery, gql } from "@apollo/client";
import { ALLGET_USERS } from "../GraphQL/Queries";

function Getusers() {
  const { error, loading, data } = useQuery(ALLGET_USERS);

  useEffect(() => {
    console.log(data);
  }, [data]);
  return <div></div>;
}

export default Getusers;
