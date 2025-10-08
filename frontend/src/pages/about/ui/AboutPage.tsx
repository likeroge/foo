import { PageLayout } from "../../../shared/page-layout";
import { PageTitle } from "../../../shared/page-title";
import styles from "./AboutPage.module.css";

export const AboutPage = () => {
  return (
    <PageLayout>
      <PageTitle title="AboutPage" />
      <div className={styles.aboutPageContainer}>
        <p>This is OCC APP v0.1</p>
        <p>Functions: </p>
        <ul>
          <li>Get all users</li>
          <li>Get all OFPs</li>
          <li>Send OFP</li>
        </ul>
      </div>
    </PageLayout>
  );
};
