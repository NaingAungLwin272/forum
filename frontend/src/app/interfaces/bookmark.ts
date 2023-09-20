export type Bookmark = {
  _id: string;
  user_id: string;
  comment_id: string;
  question_id: string;
  created_at: {
    seconds: number;
    nanos: number;
  };
  updated_at: {
    seconds: number;
    nanos: number;
  };
}

export type CreateBookmark = {
  user_id: string;
  comment_id: string;
  question_id: string;
}
