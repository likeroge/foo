import { api } from "../../../api/api";

export const deleteUser = async (id: number) => {
  try {
    const { data: result } = await api.delete(`/user/delete/${id}`);
    console.log(result);
    window.location.reload();
    return result;
  } catch (error) {
    console.error("ERROR- AAAAAAAA: ", (error as Error).message);
  }
};
