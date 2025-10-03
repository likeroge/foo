import styles from "./ErrorMessage.module.css";

export const ErrorMessage = ({ message }: { message?: string }) => {
  return <div className={styles.errorText}>{message}</div>;
};
