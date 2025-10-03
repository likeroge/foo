import { useNavigate } from "react-router";
import styles from "./LinkListElement.module.css";

export const LinkListElement = ({ title }: { title: string | undefined }) => {
  const navigate = useNavigate();
  const onClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    if (!title) return;
    navigate(title);
  };
  return (
    <div className={styles.linkListElement} onClick={(e) => onClick(e)}>
      {title}
    </div>
  );
};
