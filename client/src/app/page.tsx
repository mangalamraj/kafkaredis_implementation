"use client";

import { useGetUsers } from "../../hooks/getUsers";

export default function Home() {
  const { users, loading } = useGetUsers();
  console.log("users", users);

  return (
    <main className="min-h-screen p-8">
      {loading ? (
        <div>Loading...</div>
      ) : (
        <>
          <h1 className="text-2xl font-bold mb-6 text-white">User Directory</h1>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {users.users?.length > 0 ? (
              users.users.map((user) => (
                <div
                  key={user.id}
                  className="bg-gray-800 p-4 rounded-lg text-white"
                >
                  <h2 className="text-xl font-semibold">{user.name}</h2>
                  <p className="text-gray-300">Email: {user.email}</p>
                  <p className="text-gray-300">Phone: {user.phone_number}</p>
                  <p className="text-gray-300">Gender: {user.gender}</p>
                </div>
              ))
            ) : (
              <div>No users found</div>
            )}
          </div>
        </>
      )}
    </main>
  );
}
