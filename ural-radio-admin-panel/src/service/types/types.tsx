export type IRuEng = {
  Ru: string;
  Eng: string;
}


export interface IIssue {
  id: string;
  coverPathId: string;
  filePathId: string;
  videoPathId: string;
  title: IRuEng;
  date: string;
}


export interface IArticle {
  id: string;
  filePathId: string;
  editionId: string;
  title: IRuEng;
  authors: string[];
  reference: IRuEng;
  content: IRuEng;
  keywords: IRuEng[];
  literature: string[];
}


export interface IEditor {
  id: string;
  imagePathId: string;
  name: IRuEng;
  email: string;
  content: IRuEng;
  description: IRuEng;
  location: IRuEng;
  rank: string;
  scopus: string;
}

export interface IComment {
  id: string;
  articleId: string;
  date: string;
  content: IRuEng;
  isApproved: string;
}