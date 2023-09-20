export type Comment = {
  _id: string;
  user_id: string;
  question_id: string;
  parent_id: string;
  sort: number;
  description: string;
  vote_count: number;
  view_count: number;
  is_solution: {
    value: boolean;
  };
  is_deleted: {
    value: boolean;
  };
  created_at: {
    seconds: number;
    nanos: number;
  };
  updated_at: {
    seconds: number;
    nanos: number;
  };
}

export type UserProfile = {
  user_profile: string;
}

export type Vote = {
  comment: Comment;
  user_profile?: string;
  user_id?: string;
}

export type CreateUserVote = {
  user_id: string;
  question_id: string;
  comment_id: string;
  noti_token: string;
}

export type UserVote = {
  _id: string;
  user_id: string;
  question_id: string;
  comment_id: string;
  created_at: {
    seconds: number;
    nanos: number;
  };
  updated_at: {
    seconds: number;
    nanos: number;
  };
}

export type CreateView = {
  user_id: string;
  question_id: string;
}
