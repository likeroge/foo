import { useState } from "react";
import { LinkList, OpenListButton } from "../../../widgets/link-list";
import style from "./PageLayout.module.css";

export const PageLayout = ({ children }: { children: React.ReactNode }) => {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className={style.pageTabbarContainer}>
      {/* <MenuWidget /> */}
      <OpenListButton setIsOpen={setIsOpen} />
      {isOpen && <LinkList />}

      <div className={style.pageContainer} onClick={() => setIsOpen(false)}>
        {children}
      </div>
    </div>
  );
};
