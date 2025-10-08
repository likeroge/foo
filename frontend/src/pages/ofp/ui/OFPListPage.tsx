import { PageLayout } from "../../../shared/page-layout";
import { PageTitle } from "../../../shared/page-title";
import { OFPListComponent } from "../../../widgets/ofp-list";

export const OFPListPage = () => {
  return (
    <PageLayout>
      <PageTitle title="OFPListPage" />
      <OFPListComponent />
    </PageLayout>
  );
};
