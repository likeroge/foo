import styles from "./PageTitle.module.css";

export const PageTitle = ({ title }: { title: string }) => {
  return <h1 className={styles.title}>{title}</h1>;
};
