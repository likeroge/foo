import styles from "./Button.module.css";

export const Button = ({
  name,
  colorscheme,
}: {
  name: string;
  colorscheme: "primary" | "secondary" | "tertiary";
}) => {
  if (colorscheme === "primary")
    return (
      <button className={styles.primary + " " + styles.button}>{name}</button>
    );
  if (colorscheme === "secondary")
    return (
      <button className={styles.secondary + " " + styles.button}>{name}</button>
    );
  if (colorscheme === "tertiary")
    return (
      <button className={styles.tertiary + " " + styles.button}>{name}</button>
    );
};
