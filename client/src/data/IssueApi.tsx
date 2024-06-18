import axios from "axios"
import { IIssue } from "@/lib/typesNew";
import { serverUrl } from "@/lib/utils";
import { getArticles } from "./AticleApi";
import { IArticle } from "@/lib/typesNew";


export const getIssues = async (offset: number, limit: number): Promise<{allCount: number, data: IIssue[]}> => {
  let allCount = 0;
  const response = await axios.get(`${serverUrl}/api/v1/editions/get/all?offset=${offset}&limit=${limit}`)
    .then((response) => {
      allCount = response.data.all_count;
      return response["data"]["data"].sort((issue: IIssue, issue2: IIssue) => {
        return issue2.year - issue.year || issue2.volume - issue.volume || issue2.number - issue.number
      });
    }).catch((error: Error) => {
      console.log(error);
      throw new Error(error.message);
    }); 
    console.log(response);
  return {allCount: allCount, data: response};
}

export const getIssueById = async (id: string): Promise<IIssue> => {
  const response = await axios.get(`${serverUrl}/api/v1/editions/get/${id}`)
    .then((response) => {
      return response["data"]["edition"];
    }).catch((error: Error) => {
      throw new Error(error.message);
    }); 
  return response;
}

export const getArticlesBySearch = async (search: string): Promise<IArticle[]> => {
  if (search == '') {
    return []
  }
  const responseIssues = await axios.get(`${serverUrl}/api/v1/editions/get/all`)
    .then((response) => {
      return response["data"]["data"];
    }).catch((error: Error) => {
      console.log(error.message);
      throw new Error(error.message);
    }); 
  
  let articles = [];

  for (let i = 0; i < responseIssues.length; i++) {
    const issueArticles = await getArticles(String(responseIssues[i].id));
    articles.push(...issueArticles);
  }

  articles = articles.filter(article => {
    const titleRu = article.title.Ru.toLowerCase().includes(search.toLowerCase());
    const titleEng = article.title.Eng.toLowerCase().includes(search.toLowerCase());
    const authorsRu = article.authors.map(author => author.fullname.Ru.toLowerCase()).join(', ').includes(search.toLowerCase());
    const authorsEng = article.authors.map(author => author.fullname.Eng.toLowerCase()).join(', ').includes(search.toLowerCase());
    const keyWordsRu = article.keywords.map(word => word.Ru.toLowerCase()).join(', ').includes(search.toLowerCase());
    const keyWordsEng = article.keywords.map(word => word.Eng.toLowerCase()).join(', ').includes(search.toLowerCase());

    return titleRu || titleEng || authorsRu || authorsEng || keyWordsRu || keyWordsEng;
  });
  console.log(articles);
  return articles;
}


fetch('http://158.160.173.39:8080/api/v1/editions/get/9', {
  method: 'GET',
headers: {
 'Content-Type': 'application/json',
}})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error(error));
