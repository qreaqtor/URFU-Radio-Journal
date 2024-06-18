import axios from "axios";
import { IReductor } from "../types/typesNew";
import { serverUrl } from "../utils/serverPaths";


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


export const createReductor = async (
  nameRu: string,
  nameEng: string,
  email: string,
  imageId: string,
  descriptionRu: string,
  descriptionEng: string,
  contentRu: string,
  contentEng: string,
  rank: string,
  locationRu: string,
  locationEng: string,
  dateJoin: string
): Promise<IReductor> => {
  const response = await axios.post(`${serverUrl}/api/v1/redaction/members/create`, {
    "name" : {
      "Ru": nameRu,
      "Eng": nameEng
    },
    "email" : email,
    "imageId" : imageId,
    "description" : {
        "Ru": descriptionRu,
        "Eng": descriptionEng
    },
    "content" :  {
        "Ru": contentRu,
        "Eng": contentEng
    },
    "rank": rank,
    "location": {
        "Ru": locationRu,
        "Eng": locationEng
    },
    "dateJoin": dateJoin
  }, {
    headers: {
      "Authorization": `${window.localStorage.getItem("token")}`
    }
  }).then((response) => {
      return response["data"];
    }).catch((error) => {
      console.log(error);
      throw new Error(error.message);
    }); 
  return response;
}


export const updateReductor = async (
  id: number,
  nameRu: string,
  nameEng: string,
  email: string,
  imageId: string,
  descriptionRu: string,
  descriptionEng: string,
  contentRu: string,
  contentEng: string,
  rank: string,
  locationRu: string,
  locationEng: string,
  dateJoin: string
): Promise<IReductor> => {
  const response = await axios.put(`${serverUrl}/api/v1/redaction/members/update/${id}`, {
    "name" : {
      "Ru": nameRu,
      "Eng": nameEng
    },
    "email" : email,
    "imageId" : imageId,
    "description" : {
        "Ru": descriptionRu,
        "Eng": descriptionEng
    },
    "content" :  {
        "Ru": contentRu,
        "Eng": contentEng
    },
    "rank": rank,
    "location": {
        "Ru": locationRu,
        "Eng": locationEng
    },
    "dateJoin": dateJoin
  }, {
    headers: {
      "Authorization": `${window.localStorage.getItem("token")}`
    }
  }).then((response) => {
      return response["data"];
    }).catch((error) => {
      console.log(error);
      throw new Error(error.message);
    }); 
  return response;
}


export const deleteReductor = async (id: string) => {
  const response = await axios.delete(`${serverUrl}/api/v1/redaction/members/delete/${id}`, {
    headers: {
      "Authorization": `${window.localStorage.getItem("token")}`
    }
  }).then((response) => {
      return response["data"]["data"];
    }).catch((error: Error) => {
      console.log(error.message);
      throw new Error(error.message);
    }); 
  return response;
}