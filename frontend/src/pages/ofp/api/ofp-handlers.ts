import { api } from "../../../api/api";
import type { ApiError } from "../../../entities/ApiError";
// import type { ApiError } from "../../../entities/ApiError";

export const sendOFP = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file);
  try {
    const response = await api.post("/ofp/send", formData);
    return response.data;
    // console.log(response.data);
    // const fileName = response.data.fileName || "download.txt";
    // const blob = new Blob([response.data], {
    //   type: response.data.type || "application/octet-stream",
    // });
    // const url = window.URL.createObjectURL(blob);
    // const a = document.createElement("a");
    // a.href = url;
    // a.download = fileName;
    // document.body.appendChild(a);
    // a.click();
    // a.remove();
    // window.URL.revokeObjectURL(url);
  } catch (error) {
    throw error as ApiError;
    // if (typeof error === ApiError) {
    //   console.log(error);
    // }
  }
};
