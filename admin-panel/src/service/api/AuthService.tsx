import axios from "axios";
import { serverUrl } from "../utils/serverPaths";


export const login = async (username: string, password: string) => {
  const response = await axios.post(`${serverUrl}/api/v1/admin/auth/login`, {
    "username": username,
    "password": password
  }).then((response) => {
      console.log(response.data.Token);
      window.localStorage.setItem('token', response.data.Token)
      return response["data"];
    }).catch((error) => {
      console.log(error);
      throw new Error(error.message);
    }); 
  return response;
}