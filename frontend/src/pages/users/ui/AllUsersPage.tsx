import { useEffect, useState } from "react";
import { PageLayout } from "../../../shared/page-layout";
import type { User } from "../../../entities/User";
import { getAllUsers } from "../api/users-list";
import { UsersList } from "../../../widgets/users-list";
import { PageTitle } from "../../../shared/page-title";
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
  }, []);

  return (
    <PageLayout>
      <PageTitle title="AllUsersPage" />
      <UsersList users={users} />
    </PageLayout>
  );
};
