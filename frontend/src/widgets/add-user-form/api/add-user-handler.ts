import { api } from "../../../api/api";
import type { CreateUserDto } from "../model/CreateUserDto";

export const addUser = async (user: CreateUserDto) => {
  try {
    const { data: result } = await api.post("/user", user);
    console.log(result);
  } catch (error) {
    console.error("ERROR- AAAAAAAA: ", (error as Error).message);
  }
};
