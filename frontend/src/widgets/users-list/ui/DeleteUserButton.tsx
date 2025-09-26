import { deleteUser } from "../api/delete-user-button";
import styles from "./DeleteUserButton.module.css";

export const DeleteUserButton = ({ userId }: { userId: number }) => {
  return (
    <button
      onClick={() => deleteUser(userId)}
      type="button"
      className={styles.deleteUserButton}
    >
      Delete user
    </button>
  );
};
