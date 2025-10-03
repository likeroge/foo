import { appRouter } from "../../../app/router/AppRouter";
import styles from "./LinkList.module.css";
import { LinkListElement } from "./LinkListElement";

export const LinkList = () => {
  return (
    <div className={styles.linkList}>
      <h1>MENU</h1>
      {appRouter.routes
        .filter((route) => route.path !== "/")
        .filter((route) => route.path !== "*")

        .map((route) => (
          <LinkListElement key={route.path} title={route.path} />
        ))}
    </div>
  );
};
