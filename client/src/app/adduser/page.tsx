"use client";

import { useState } from "react";

const AddUser = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [gender, setGender] = useState("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch("http://localhost:8080/adduser", {
        method: "POST",
        body: JSON.stringify({ name, email, phoneNumber, gender }),
      });
      if (response.ok) {
        const data = await response.json();
        console.log(data);
      } else {
        console.error("Failed to add user");
      }
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className="flex align-middle flex-col w-60 justify-center mt-10 m-auto">
      <h1>Add User</h1>
      <form onSubmit={handleSubmit} className="flex flex-col gap-4">
        <input
          type="text"
          placeholder="Name"
          onChange={(e) => {
            setName(e.target.value);
          }}
          className="text-black"
        />
        <input
          type="email"
          placeholder="Email"
          onChange={(e) => {
            setEmail(e.target.value);
          }}
          className="text-black"
        />
        <input
          type="text"
          placeholder="Phone Number"
          onChange={(e) => {
            setPhoneNumber(e.target.value);
          }}
          className="text-black"
        />
        <input
          type="text"
          placeholder="Gender"
          onChange={(e) => {
            setGender(e.target.value);
          }}
          className="text-black"
        />
        <button type="submit" className="p-2 bg-white text-black rounded">
          Add User
        </button>
      </form>
    </div>
  );
};

export default AddUser;
