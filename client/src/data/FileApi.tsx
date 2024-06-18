import axios from "axios";
import { serverUrl } from "@/lib/utils";

export const downloadFile = async (id: string) => {
  const response = await axios.get(`${serverUrl}/api/v1/files/download/${id}`)
    .then((response) => {
      return response["data"]["id"];
    }).catch((error: Error) => {
      console.log(error.message);
      throw new Error(error.message);
    }); 
  return response;
}

// export const postSearch = async (id: string) => {
//   const response = await axios.get(`https://journa-token.onrender.com/files/get/${id}?download=false`)
//     .then(() => {
//       return `https://journa-token.onrender.com/files/get/${id}?download=true`;
//     }).catch((error: Error) => {
//       throw new Error(error.message);
//     }); 
//   return response;
// }