import { useMemo, useState } from "react";
import { IComment } from "../service/types/typesNew";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { approveComment, deleteComment, updateComment } from "../service/api/CommentService";
import { commentSchema, TCommentSchema } from "../service/types/schemes";
import { transformDate } from "../service/utils/utils";

type CommentComponentProps = {
  comment: IComment
}

const CommentComponent = ({comment}: CommentComponentProps) => {
  const [isApproving, setIsApproving] = useState(false);
  const date = useMemo(() => transformDate(comment.date), [comment.date]);
  const queryClient = useQueryClient();

  const {register, handleSubmit} = useForm<TCommentSchema>({resolver: zodResolver(commentSchema)});

  const articleMutation = useMutation({
    mutationFn: async (data: TCommentSchema) => {
      await updateComment(comment.id, data.contentRu, data.contentEng);
      if (!comment.isApproved) {
        await approveComment(comment.id, data.contentEng);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["comments", String(comment.articleId)]});
      queryClient.invalidateQueries({queryKey: ["commentsApproved", String(comment.articleId)]});
    }
  });

  const deleteArticleMutation = useMutation({
    mutationFn: async () => await deleteComment(comment.id),
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["comments", String(comment.articleId)]});
      queryClient.invalidateQueries({queryKey: ["commentsApproved", String(comment.articleId)]});
    }
  });

  const onSubmitHandle = (data: TCommentSchema) => {
    articleMutation.mutate(data);
    setIsApproving(false);
  }

  const onClickApproveButton = () => {
    setIsApproving(true);
  }

  const onClickNotApproveButton = () => {
    setIsApproving(false);
  }

  const onClickDeleteButton = () => {
    deleteArticleMutation.mutate();
  }

  

  if (isApproving) {
    return (
      <div className="article__comment">
        <form onSubmit={handleSubmit(onSubmitHandle)}>
          <p className="comment__approve-text">{comment.author}</p>
          <input {...register("contentRu")} className="comment__input" type="text" defaultValue={comment.content.Ru} placeholder="Текст комментария на русском..."/>
          <input {...register("contentEng")} className="comment__input" type="text" defaultValue={comment.content.Eng} placeholder="Текст комментария на английском..."/>
          <div className="comment__approve-buttons">
            <button className="small-button" type='submit'>{comment.isApproved ? "Сохранить": "Одобрить"}</button>
            <button className="small-button-delete" onClick={onClickNotApproveButton} type="button">Отменить</button>
          </div>
        </form>
      </div>
    )
  }

  return (
    <div className="article__comment">
      <div className="comment__description">
        <p className="comment__text">{comment.author}</p>
        <p className="comment__text">{comment.content.Ru}</p>
        <p className="comment__date">{date}</p>
      </div>
      <div className="comment__approve-buttons">
        <button className="small-button" onClick={onClickApproveButton}>{comment.isApproved ? "Изменить": "Одобрить"}</button>
        <button className="small-button-delete" onClick={onClickDeleteButton}>Удалить</button>
      </div>
      
    </div>
  );
};

export default CommentComponent;