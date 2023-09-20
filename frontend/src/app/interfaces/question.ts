import { Comment } from './comment';

export type Question = {
  _id: string;
  user_id: string;
  title: string;
  CommentId: string;
  language_ids: string[];
  tag_ids: string[];
  view_count: number,
  vote_count: number,
  reply_count: number,
  reply: number,
  solution_count: number,
  user_ids: string[];
  is_solution: {
    value: boolean;
  }
  created_at: {
    seconds: number;
    nanos: number;
  },
}

export type FilterQuestion = {
  language_ids: string[];
  tag_ids: string[];
  user_id: string[];
  title: string;
}

export type QuestionDetail = {
  _id: string;
  user_id: string;
  title: string;
  language_ids: string[];
  tag_ids: string[];
  view_count: number,
  vote_count: number,
  reply_count: number,
  created_at: {
    seconds: number;
    nanos: number;
  },
  updated_at: {
    seconds: number;
    nanos: number;
  },
  comments: Comment[]
}