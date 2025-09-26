import { api } from "../../../api/api";

export const deleteUser = async (id: number) => {
  try {
    const { data: result } = await api.delete(`/user/${id}`);
    console.log(result);
  } catch (error) {
    console.error("ERROR- AAAAAAAA: ", (error as Error).message);
  }
};
