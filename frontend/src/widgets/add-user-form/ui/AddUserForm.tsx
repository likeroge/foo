import { useState } from "react";
import type { CreateUserDto } from "../model/CreateUserDto";
import styles from "./AddUserForm.module.css";
import { addUser } from "../api/add-user-handler";

export const AddUserForm = () => {
  const [userForm, setUserForm] = useState<CreateUserDto>({
    name: "",
    email: "",
  });
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!userForm.name || !userForm.email) return;
    try {
      await addUser(userForm);
    } catch (error) {
      console.error((error as Error).message);
    }
  };
  return (
    <form onSubmit={handleSubmit} className={styles.addUserForm}>
      <label htmlFor="name">Name:</label>
      <input
        type="text"
        id="name"
        value={userForm.name}
        onChange={(e) => {
          setUserForm({ ...userForm, name: e.target.value });
        }}
      />

      <label htmlFor="email">Email:</label>
      <input
        type="email"
        id="email"
        name="email"
        value={userForm.email}
        onChange={(e) => setUserForm({ ...userForm, email: e.target.value })}
      />

      <button className={styles.addUserButton} type="submit">
        Add user
      </button>
    </form>
  );
};
