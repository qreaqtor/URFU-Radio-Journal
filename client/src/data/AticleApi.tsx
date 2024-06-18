import axios from "axios";
import { IArticle } from "@/lib/typesNew";
import { serverUrl } from "@/lib/utils";


export const getArticles = async (editionId: string): Promise<IArticle[]> => {
  const response = await axios.get(`${serverUrl}/api/v1/article/get/all?editionId=${editionId}`)
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
// offset=${offset}&limit=${limit}
export const getArticlesSearch = async (
  search: string, 
  offset: number,
  limit: number,
  title: boolean, 
  keywords: boolean, 
  authors: boolean, 
  annotation: boolean
): Promise<{allCount: number, data: IArticle[]}> => {
  let allCount = 0;
  const response = await axios.get(`${serverUrl}/api/v1/article/get/all?search=${search}&offset=${offset}&limit=${limit}&title=${title}&authors=${authors}&keywords=${keywords}&affilation=${annotation}`)
    .then((response) => {
      if (response.data.all_count == 0) {
        return [];
      }
      allCount = response.data.all_count;
      return response.data.data;
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return {allCount, data: response};
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