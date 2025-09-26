import { useState } from "react";
import styles from "./FindUserForm.module.css";
import {
  findUserByEmail,
  findUserById,
  findUserByName,
} from "../api/find-users-handler";
import type { User } from "../../../entities/User";

export const FindUserForm = ({
  setSearchResults,
}: {
  setSearchResults: (result: User) => void;
}) => {
  const [searchParam, setSearchParam] = useState("id");
  const [searchValue, setSearchValue] = useState("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    let result;
    switch (searchParam) {
      case "id":
        result = await findUserById(Number(searchValue));
        break;
      case "name":
        result = await findUserByName(searchValue);
        break;
      case "email":
        result = await findUserByEmail(searchValue);
        break;
    }
    console.log(result);
    setSearchResults(result);
  };

  const handleSearchParamChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    e.preventDefault();
    setSearchParam(e.currentTarget.value);
  };

  return (
    <form className={styles.findUserForm} onSubmit={handleSubmit}>
      <div className={styles.findUserInputsContainer}>
        <input
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
          className={styles.findUserInput}
          type="text"
          placeholder="Enter user id or name or email"
        />
        <select
          className={styles.findUserInput}
          value={searchParam}
          onChange={handleSearchParamChange}
        >
          <option value="id">id</option>
          <option value="name">name</option>
          <option value="email">email</option>
        </select>
      </div>

      <button type="submit" className={styles.findUserButton}>
        Find user
      </button>
    </form>
  );
};
