import { api } from "../../../api/api";
import type { User } from "../../../entities/User";

export const getAllUsers = async () => {
  try {
    const { data } = await api.get<User[]>("/user/all");
    if (!data) return [];
    return data;
  } catch (error) {
    console.error((error as Error).message);
    return [];
  }
};
