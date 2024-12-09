"use client";

import { useGetUsers, Tusers } from "../../../hooks/getUsers";

const GetUser = () => {
  const { users } = useGetUsers();

  return (
    <div className="text-white flex align-middle flex-col w-60 justify-center mt-10 m-auto">
      <div>
        {users.users.length > 0
          ? users.users.map((u) => <div key={u.id}>{u.name}</div>)
          : "No users found"}
      </div>
      <button className="bg-white text-black p-2 rounded-md">Get Users</button>
    </div>
  );
};

export default GetUser;
