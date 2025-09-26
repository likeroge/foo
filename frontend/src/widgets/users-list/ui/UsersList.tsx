import type { User } from "../../../entities/User";
import { DeleteUserButton } from "./DeleteUserButton";
import styles from "./UsersList.module.css";

export const UsersList = ({ users }: { users: User[] }) => {
  return (
    <div className={styles.usersList}>
      <h2>Users:</h2>
      <div className={styles.usersListContainer}>
        {users.map((user) => (
          <div key={user.id} className={styles.userListElement}>
            <p>
              id: {user.id}, name: {user.name}
            </p>
            <DeleteUserButton userId={user.id} />
          </div>
        ))}
      </div>
    </div>
  );
};
