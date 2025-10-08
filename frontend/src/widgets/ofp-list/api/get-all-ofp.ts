import { api } from "../../../api/api";

export const getAllOFP = async () => {
  try {
    const { data: result } = await api.get(`/ofp/all`);
    console.log(result);
    return result;
  } catch (error) {
    throw (error as Error).message;
  }
};
