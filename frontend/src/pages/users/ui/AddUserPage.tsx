import { PageLayout } from "../../../shared/page-layout";
import { PageTitle } from "../../../shared/page-title";
import { AddUserForm } from "../../../widgets/add-user-form";

export const AddUserPage = () => {
  return (
    <PageLayout>
      <PageTitle title="AddUserPage" />
      <AddUserForm />
    </PageLayout>
  );
};
