import { useState, useEffect } from "react";

export type Tusers = {
  message: string;
  users: {
    id: number;
    name: string;
    email: string;
    phone_number: string;
    gender: string;
  }[];
};

export const useGetUsers = () => {
  const [users, setUsers] = useState<Tusers>({ message: "", users: [] });
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchUsers = async () => {
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/users");
      if (!response.ok)
        throw new Error(`HTTP error! status: ${response.status}`);
      const data: Tusers = await response.json(); // Ensure it matches Tusers type
      console.log("yellow", data);
      console.log("yellow", data.message);
      setUsers(data); // Set the entire object
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return { users, loading, error };
};
