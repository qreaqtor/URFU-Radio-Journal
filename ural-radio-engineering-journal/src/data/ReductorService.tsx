import axios from "axios";
import { IReductor } from "@/lib/typesNew";
import { serverUrl } from "@/lib/utils";


export const getReductors = async (): Promise<IReductor[]> => {
  const response = await axios.get(`${serverUrl}/api/v1/redaction/members/get/all`)
    .then((response) => {
      return response.data.data;
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return response;
}

export const getReductorById = async (id: string): Promise<IReductor> => {
  const response = await axios.get(`${serverUrl}/api/v1/redaction/members/get/${id}`)
    .then((response) => {
      return response["data"]["member"];
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return response;
}