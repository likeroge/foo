import { api } from "../../../api/api";

export const findUserById = async (id: number) => {
  try {
    const { data: result } = await api.get(`/user/find/id/${id}`);
    return result;
  } catch (error) {
    console.error("ERROR- AAAAAAAA: ", (error as Error).message);
  }
};

export const findUserByName = async (name: string) => {
  try {
    const { data: result } = await api.get(`/user/find/name/${name}`);
    return result;
  } catch (error) {
    console.error("ERROR- AAAAAAAA: ", (error as Error).message);
  }
};

export const findUserByEmail = async (email: string) => {
  try {
    const { data: result } = await api.get(`/user/find/email/${email}`);
    return result;
  } catch (error) {
    console.error("ERROR- AAAAAAAA: ", (error as Error).message);
  }
};
