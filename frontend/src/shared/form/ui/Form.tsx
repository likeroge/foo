import styles from "./Form.module.css";

export const Form = ({
  children,
  onSubmit,
}: {
  children: React.ReactNode;
  onSubmit?: React.FormEventHandler<HTMLFormElement>;
}) => {
  return (
    <form className={styles.form} onSubmit={onSubmit}>
      {children}
    </form>
  );
};
