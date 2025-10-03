import { useState } from "react";
import { PageLayout } from "../../../shared/page-layout";
import { PageTitle } from "../../../shared/page-title";
import { sendOFP } from "../api/ofp-handlers";
import { Button } from "../../../shared/button";
import { Form } from "../../../shared/form";
import type { ApiError } from "../../../entities/ApiError";
import { Flex } from "../../../shared/flex";

export const OFPLoaderPage = () => {
  const [file, setFile] = useState<File>();
  const [apiError, setApiError] = useState<ApiError>();
  const [apiResponse, setApiResponse] = useState<string>();
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
    try {
      const result = await sendOFP(file);
      console.log(result);
      setApiResponse(result);
    } catch (error) {
      setApiError(error as ApiError);
    }
  };

  return (
    <PageLayout>
      <PageTitle title="OFPLoaderPage" />
      <Flex direction="row" justifyContent="spaceBetween">
        <Flex direction="column">
          <Form onSubmit={handleSubmit}>
            <input type="file" name="file" onChange={handleFileChange} />
            {file && file.name && <Button name="Send" colorscheme="primary" />}
          </Form>
          {apiError && <p>{apiError.message}</p>}
        </Flex>
        <Flex
          direction="column"
          style={{
            maxWidth: "50%",
            textWrap: "wrap",
          }}
        >
          {apiResponse && (
            <p style={{ whiteSpace: "pre-wrap", wordWrap: "break-word" }}>
              {JSON.stringify(apiResponse)}
            </p>
          )}
        </Flex>
      </Flex>
    </PageLayout>
  );
};
