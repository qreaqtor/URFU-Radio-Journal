import axios from "axios"
import { serverUrl } from "../utils/serverPaths.tsx";

export const uploadFile = async (form: FormData) => {
  const response = await axios.post(`${serverUrl}/api/v1/files/upload/`, form, {
    headers: {
      "Content-Type": 'multipart/form-data',
      "Authorization": `${window.localStorage.getItem('token')}`
    }
  }).then((response) => {
      return response.data.id;
    }).catch((error) => {
      console.log(error);
      return null;
    }); 

    
  return response;
}

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

export const deleteFile = async (id: string) => {
  const response = await axios.delete(`${serverUrl}/api/v1/files/delete/${id}`, {
    headers: {
      "Authorization": `${window.localStorage.getItem('token')}`
    }
  })
  .then((response) => {
    return response["data"]["id"];
  }).catch((error: Error) => {
    console.log(error.message);
    throw new Error(error.message);
  }); 
return response;
}


//730565d3-74f5-40a2-83c1-6c93bba47533