import { useState } from "react";
import { PageLayout } from "../../../shared/page-layout";
import { FindUserForm } from "../../../widgets/find-user-form";
import styles from "./FindUserPage.module.css";
import type { User } from "../../../entities/User";

export const FindUserPage = () => {
  const [searchResults, setSearchResults] = useState<User>();
  return (
    <PageLayout>
      <div className={styles.findUserPageContainer}>
        <div className={styles.findUserFormContainer}>
          <h3>Please use name, email or id to find user:</h3>
          <FindUserForm setSearchResults={setSearchResults} />
        </div>
        <div className={styles.findUserResultsContainer}>
          <h3>Results:</h3>
          <div>{searchResults ? searchResults.name : "No results"}</div>
        </div>
      </div>
    </PageLayout>
  );
};
