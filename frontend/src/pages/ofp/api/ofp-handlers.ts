import { api } from "../../../api/api";
import type { ApiError } from "../../../entities/ApiError";
// import type { ApiError } from "../../../entities/ApiError";

export const sendOFP = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file);
  try {
    const response = await api.post("/ofp/send", formData, {
      responseType: "blob",
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    console.log(response);

    const fileName =
      response.headers["content-disposition"].split("=")[1] ?? "download.txt";

    const blob = new Blob([response.data]);
    const url = window.URL.createObjectURL(blob);

    const a = document.createElement("a");

    a.href = url;
    a.download = fileName;
    document.body.appendChild(a);
    a.click();
    a.remove();
    window.URL.revokeObjectURL(url);

    return {
      message: `File ${fileName} uploaded successfully`,
    };
  } catch (error) {
    throw error as ApiError;
  }
};
