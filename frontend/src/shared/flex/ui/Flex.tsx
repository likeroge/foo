import styles from "./Flex.module.css";

export const Flex = ({
  children,
  direction,
  style,
  justifyContent = "flexStart",
}: {
  children: React.ReactNode;
  direction: "row" | "column";
  style?: React.CSSProperties;
  justifyContent?: "flexStart" | "flexEnd" | "center" | "spaceBetween";
}) => {
  if (direction === "row")
    return (
      <div
        style={style}
        className={styles.flexRow + " " + styles[justifyContent]}
      >
        {children}
      </div>
    );
  return (
    <div
      style={style}
      className={styles.flexColumn + " " + styles[justifyContent]}
    >
      {children}
    </div>
  );
};
