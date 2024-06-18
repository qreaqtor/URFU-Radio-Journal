export type IRuEng = {
  Ru: string,
  Eng: string
}

export type IIssue = {
  id: number,
  year: number,
  number: number,
  volume: number,
  documentID: string,
  imageID: string,
  date: string
}

export type IAuthor = {
  fullname: IRuEng,
  affilation: string,
  email: string
}

export type IArticle = {
  id: number,
  editionId: number,
  title: IRuEng,
  authors: IAuthor[],
  content: IRuEng,
  keywords: IRuEng[],
  documentID: string,
  literature: string[],
  reference: IRuEng,
  videoID: string,
  dateReceipt: string,
  dateAcceptance: string,
  doi: string
}

export type IComment = {
  id: number,
  articleId: number,
  content: IRuEng,
  date: string,
  isApproved: boolean,
  author: string
}

export type ICouncil = {
  id: number,
  name: IRuEng,
  email: string,
  imageID: string,
  scopus: string,
  description: IRuEng,
  content: IRuEng,
  rank: string,
  location: IRuEng,
  dateJoin: string
}

export type IReductor = {
  id: number,
  name: IRuEng,
  email: string,
  imageID: string,
  description: IRuEng,
  content: IRuEng,
  rank: string,
  location: IRuEng,
  dateJoin: string
}