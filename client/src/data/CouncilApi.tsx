import axios from "axios"
import { ICouncil } from "@/lib/typesNew";
import { serverUrl } from "@/lib/utils";


export const getCouncils = async (): Promise<ICouncil[]> => {
  const response = await axios.get(`${serverUrl}/api/v1/council/members/get/all`)
    .then((response) => {
      if (!response.data.data) {
        return []
      }
      return response.data.data;
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return response;
}


export const getCouncilById = async (id: string): Promise<ICouncil> => {
  const response = await axios.get(`${serverUrl}/api/v1/council/members/get/${id}`)
    .then((response) => {
      return response["data"]["member"];
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return response;
}