import { useState } from "react";
import { PageLayout } from "../../../shared/page-layout";
import { PageTitle } from "../../../shared/page-title";
import { sendOFP } from "../api/ofp-handlers";
import { Button } from "../../../shared/button";
import { Form } from "../../../shared/form";

export const OFPLoaderPage = () => {
  const [file, setFile] = useState<File>();
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log(e.target.files);
    if (!e.target.files) return;
    if (e.target.files.length === 0) return;
    if (e.target.files.length > 1) return;

    setFile(e.target.files[0]);
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!file) return;
    if (!file.name.endsWith(".pdf")) {
      console.error("File is not .pdf");
      return;
    }
    // if (!file.name.endsWith(".txt")) {
    // console.log("File is not .txt");
    // return;
    // }
    const result = await sendOFP(file);
    console.log(result);
  };

  return (
    <PageLayout>
      <PageTitle title="OFPLoaderPage" />
      <Form onSubmit={handleSubmit}>
        <input type="file" name="file" onChange={handleFileChange} />
        {file && file.name && <Button name="Send" colorscheme="primary" />}
      </Form>
    </PageLayout>
  );
};
