import axios from "axios";
import { serverUrl } from "../utils/serverPaths";
import { IAuthor, IRuEng, IArticle } from "../types/typesNew";


export const getArticles = async (editionId: string): Promise<IArticle[]> => {
  const response = await axios.get(`${serverUrl}/api/v1/article/get/all?editionId=${editionId}`)
    .then((response) => {
      return response.data.data;
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return response;
}


export const getArticleById = async (id: string): Promise<IArticle> => {
  const response = await axios.get(`${serverUrl}/api/v1/article/get/${id}`)
    .then((response) => {
      return response.data.article;
    }).catch((error: Error) => {
      console.log(error.message);
      throw new Error(error.message);
    }); 
  return response;
}


export const createArticle = async (
  editionId: number, 
  titleRu: string,
  titleEng: string,
  authors: IAuthor[], 
  contentRu: string,
  contentEng: string,
  keywords: IRuEng[],
  documentID: string,
  literature: string[],
  referenceRu: string,
  referenceEng: string,
  videoID: string,
  dateReceipt: string,
  dateAcceptance: string,
  doi: string,
): Promise<IArticle> => {
  const response = await axios.post(`${serverUrl}/api/v1/article/create`, {
    "editionId": editionId,
    "title": {
      "Ru": titleRu,
      "Eng": titleEng
    },
    "authors": authors,
    "content": {
      "Ru": contentRu,
      "Eng": contentEng
    },
    "keywords": keywords,
    "documentID": documentID,
    "literature": literature,
    "reference": {
      "Ru": referenceRu,
      "Eng": referenceEng
    },
    "VideoID": videoID,
    "dateReceipt": dateReceipt,
    "DateAcceptance": dateAcceptance,
    "doi": doi
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


export const updateArticle = async (
  id: number,
  titleRu: string,
  titleEng: string,
  authors: IAuthor[], 
  contentRu: string,
  contentEng: string,
  keywords: IRuEng[],
  documentId: string,
  literature: string[],
  referenceRu: string,
  referenceEng: string,
  videoId: string,
  dateReceipt: string,
  dateAcceptance: string,
  doi: string,
): Promise<IArticle> => {
  const response = await axios.put(`${serverUrl}/api/v1/article/update`, {
    "id": id,
    "title": {
      "Ru": titleRu,
      "Eng": titleEng
    },
    "authors": authors,
    "content": {
      "Ru": contentRu,
      "Eng": contentEng
    },
    "keywords": keywords,
    "documentID": documentId,
    "literature": literature,
    "reference": {
      "Ru": referenceRu,
      "Eng": referenceEng
    },
    "videoID": videoId,
    "dateReceipt": dateReceipt,
    "DateAcceptance": dateAcceptance,
    "doi": doi
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

export const deleteArticle = async (id: string) => {
  const response = await axios.delete(`${serverUrl}/api/v1/article/delete/${id}`, {
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

