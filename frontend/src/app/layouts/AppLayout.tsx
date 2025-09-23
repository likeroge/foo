import style from "./AppLayout.module.css";

export const AppLayout = ({ children }: { children: React.ReactNode }) => {
  return <div className={style.appContainer}>{children}</div>;
};
