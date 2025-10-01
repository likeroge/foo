import { api } from "../../../api/api";

export const sendOFP = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file);
  try {
    const response = await api.post("/ofp/send", formData, {
      responseType: "arraybuffer",
    });
    console.log(response.data);

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
    console.error((error as Error).message);
    return;
  }
};
