import styles from "./OpenListButton.module.css";

export const OpenListButton = ({
  setIsOpen,
}: {
  setIsOpen: (isOpen: boolean) => void;
}) => {
  const handleOpenList = () => {
    setIsOpen(true);
  };
  return (
    <button className={styles.openListButton} onClick={handleOpenList}>
      Open list
    </button>
  );
};
