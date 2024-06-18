import axios from "axios";
import { serverUrl } from "../utils/serverPaths";
import { IComment } from "../types/typesNew";



export const getComments = async (id: number, isApproved: boolean): Promise<IComment[]> => {
  const response = await axios.get(`${serverUrl}/api/v1/comments/get/all?onlyApproved=${isApproved}&articleId=${id}`)
    .then((response) => {
      return response.data.data;
    }).catch((error: Error) => {
      console.log(error);
      throw new Error(error.message);
    }); 
  return response;
}


export const createComment = async (
  articleId: number, 
  content: string, 
  author: string, 
  date: string, 
): Promise<IComment> => {
  const response = await axios.post(`${serverUrl}/api/v1/comments/create`, {
    "articleId": articleId,
    "content": content,
    "author": author,
    "date": date,
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

export const updateComment = async (
  commentId: number, 
  contentRu: string, 
  contentEng: string,
): Promise<IComment> => {
  const response = await axios.patch(`${serverUrl}/api/v1/comments/update`, {
    "id" : commentId,
    "content" : {
        "Ru": contentRu,
        "Eng": contentEng
    }
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


export const approveComment = async (
  commentId: number, 
  content: string, 
): Promise<IComment> => {
  const response = await axios.patch(`${serverUrl}/api/v1/comments/approve`, {
    "id" : commentId,
    "content" : content
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

export const deleteComment = async (
  commentId: number, 
): Promise<IComment> => {
  const response = await axios.delete(`${serverUrl}/api/v1/comments/delete/${commentId}`, {
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