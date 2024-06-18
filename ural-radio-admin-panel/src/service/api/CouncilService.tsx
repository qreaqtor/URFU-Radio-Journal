import axios from "axios";
import { ICouncil } from "../types/typesNew";
import { serverUrl } from "../utils/serverPaths";

export const getCouncils = async (): Promise<ICouncil[]> => {
  const response = await axios.get(`${serverUrl}/api/v1/council/members/get/all`)
    .then((response) => {
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


export const createCouncil = async (
  nameRu: string,
  nameEng: string,
  email: string,
  imageId: string,
  scopus: string,
  descriptionRu: string,
  descriptionEng: string,
  contentRu: string,
  contentEng: string,
  rank: string,
  locationRu: string,
  locationEng: string,
  dateJoin: string
): Promise<ICouncil> => {
  const response = await axios.post(`${serverUrl}/api/v1/council/members/create`, {
    "name" : {
      "Ru": nameRu,
      "Eng": nameEng
    },
    "email" : email,
    "imageId" : imageId,
    "scopus" : scopus,
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


export const updateCouncil = async (
  id: number,
  nameRu: string,
  nameEng: string,
  email: string,
  imageId: string,
  scopus: string,
  descriptionRu: string,
  descriptionEng: string,
  contentRu: string,
  contentEng: string,
  rank: string,
  locationRu: string,
  locationEng: string,
  dateJoin: string
): Promise<ICouncil> => {
  const response = await axios.put(`${serverUrl}/api/v1/council/members/update/${id}`, {
    "name" : {
      "Ru": nameRu,
      "Eng": nameEng
    },
    "email" : email,
    "imageId" : imageId,
    "scopus" : scopus,
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


export const deleteCouncil = async (id: string) => {
  const response = await axios.delete(`${serverUrl}/api/v1/council/members/delete/${id}`, {
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