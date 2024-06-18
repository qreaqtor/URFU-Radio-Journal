import { z } from "zod";

export const issueSchema = z.object({
  year: z.string().min(4, 'Недопустимая длинна').max(4, 'Недопустимая длинна'),
  number: z.string(),
  volume: z.string(),
  date: z.string()
}).refine(data => {
  const date = new Date(`${data.date}T12:00:00Z`);
  return !isNaN(date.valueOf());
}, {
  message: "неверно введена дата",
  path: ["date"]
}).refine(data => {
  const year = parseInt(data.year);
  return !isNaN(year.valueOf());
}, {
  message: "Некорректный ввод года",
  path: ["year"]
}).refine(data => {
  const year = parseInt(data.number);
  return !isNaN(year.valueOf());
}, {
  message: "Некорректный ввод номера выпуска",
  path: ["number"]
}).refine(data => {
  const year = parseInt(data.volume);
  return !isNaN(year.valueOf());
}, {
  message: "Некорректный ввод тома",
  path: ["volume"]
});

export type TIssueSchema = z.infer<typeof issueSchema>;


export const createArticleSchema = z.object({
  titleRu: z.string().min(1, "Это поле обязательно для заполнения"),
  titleEng: z.string().min(1, "Это поле обязательно для заполнения"),
  contentRu: z.string().min(1, "Это поле обязательно для заполнения"),
  contentEng: z.string().min(1, "Это поле обязательно для заполнения"),
  keywordsRu: z.string(),
  keywordsEng: z.string(),
  referenceRu: z.string().min(1, "Это поле обязательно для заполнения"),
  referenceEng: z.string().min(1, "Это поле обязательно для заполнения"),
  dateReceipt: z.string(),
  dateAcceptance: z.string(),
  doi: z.string()
}).refine(data => {
  const date = new Date(`${data.dateReceipt}T12:00:00Z`);
  return !isNaN(date.valueOf());
}, {
  message: "Неверно введена дата",
  path: ["dateReceipt"]
}).refine(data => {
  const date = new Date(`${data.dateAcceptance}T12:00:00Z`);
  return !isNaN(date.valueOf());
}, {
  message: "Неверно введена дата",
  path: ["dateAcceptance"]
});

export type TCreateArticleSchema = z.infer<typeof createArticleSchema>;


export const councilSchema = z.object({
  nameRu: z.string().min(1, "Это поле обязательно для заполнения"),
  nameEng: z.string().min(1, "Это поле обязательно для заполнения"),
  email: z.string().email("Некорректный ввод email"),
  scopus: z.string().url("Некорректный ввод Scopus"),
  descriptionRu: z.string().min(1, "Это поле обязательно для заполнения").max(1000, "Максимальная длинна: 1000 символов"),
  descriptionEng: z.string().min(1, "Это поле обязательно для заполнения").max(1000, "Максимальная длинна: 1000 символов"),
  contentRu: z.string().min(1, "Это поле обязательно для заполнения"),
  contentEng: z.string().min(1, "Это поле обязательно для заполнения"),
  rank: z.string().min(1, "Это поле обязательно для заполнения"),
  locationRu: z.string().min(1, "Это поле обязательно для заполнения"),
  locationEng: z.string().min(1, "Это поле обязательно для заполнения"),
  dateJoin: z.string()
}).refine(data => {
  const date = new Date(`${data.dateJoin}T12:00:00Z`);
  return !isNaN(date.valueOf());
}, {
  message: "Неверно введена дата",
  path: ["dateJoin"]
}).refine(data => {
  return data.rank == "Главный редактор" 
          || data.rank == "Заместитель главного редактора" 
          || data.rank == "Зав. редакцией" 
          || data.rank == "Член редакционного совета";
}, {
  message: "Это поле может быть равно только: Главный редактор, Заместитель главного редактора, Зав. редакцией или Член редакционного совета",
  path: ["rank"]
});

export type TCouncilSchema = z.infer<typeof councilSchema>;


export const reductorSchema = z.object({
  nameRu: z.string().min(1, "Это поле обязательно для заполнения"),
  nameEng: z.string().min(1, "Это поле обязательно для заполнения"),
  email: z.string().email("Некорректный ввод email"),
  descriptionRu: z.string().min(1, "Это поле обязательно для заполнения").max(1000, "Максимальная длинна: 1000 символов"),
  descriptionEng: z.string().min(1, "Это поле обязательно для заполнения").max(1000, "Максимальная длинна: 1000 символов"),
  contentRu: z.string().min(1, "Это поле обязательно для заполнения"),
  contentEng: z.string().min(1, "Это поле обязательно для заполнения"),
  rank: z.string().min(1, "Это поле обязательно для заполнения"),
  locationRu: z.string().min(1, "Это поле обязательно для заполнения"),
  locationEng: z.string().min(1, "Это поле обязательно для заполнения"),
  dateJoin: z.string()
}).refine(data => {
  const date = new Date(`${data.dateJoin}T12:00:00Z`);
  return !isNaN(date.valueOf());
}, {
  message: "Неверно введена дата",
  path: ["dateJoin"]
});

export type TReductorSchema = z.infer<typeof reductorSchema>;

export const commentSchema = z.object({
  contentRu: z.string(),
  contentEng: z.string(),
});

export type TCommentSchema = z.infer<typeof commentSchema>;