import { useEffect, useState } from "react";
import { PageLayout } from "../../../shared/page-layout";
import type { User } from "../../../entities/User";
import { getAllUsers } from "../api/users-list";
import { UsersList } from "../../../widgets/users-list";
// import styles from "./AllUsersPage.module.css";

export const AllUsersPage = () => {
  const [users, setUsers] = useState<User[]>([]);
  useEffect(() => {
    (async () => {
      try {
        const userList = await getAllUsers();
        setUsers(userList);
      } catch (error) {
        console.error((error as Error).message);
      }
    })();
  }, [users]);
  console.log(users);

  return (
    <PageLayout>
      <UsersList users={users} />
    </PageLayout>
  );
};
