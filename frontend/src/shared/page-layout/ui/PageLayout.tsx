import style from "./PageLayout.module.css";

export const PageLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      {/* <MenuWidget /> */}
      <div className={style.pageContainer}>{children}</div>
    </>
  );
};
