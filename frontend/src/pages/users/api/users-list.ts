import { api } from "../../../api/api";

export const getAllUsers = async () => {
  try {
    const { data } = await api.get("/user/all");
    return data;
  } catch (error) {
    console.error((error as Error).message);
  }
};
