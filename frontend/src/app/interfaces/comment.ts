export type ReplyComment = {
  _id: string;
  user_id: string;
  parent_id: string;
  question_id: string;
  description: string;
  sort: number;
  vote_count: number;
  solutionCount?: number;
  voteCount?: number;
  replyCount?: number;
  solutionLinks?: string[];
  is_solution: {
    value: boolean;
  };
  vote: boolean;
  bookmark: boolean;
  isVote?: boolean;
  isBookmark?: boolean;
  created_at: {
    seconds: number;
    nanos: number;
  },
  updated_at: {
    seconds: number;
    nanos: number;
  },
}

export type Comment = {
  _id: string;
  user_id: string;
  question_id: string;
  parent_id: string;
  description: string;
  sort: number;
  vote_count: number;
  solutionCount?: number;
  solutionLinks?: string[];
  voteCount?: number;
  replyCount?: number;
  is_solution: {
    value: boolean;
  };
  vote: boolean;
  bookmark: boolean;
  isVote?: boolean;
  isBookmark?: boolean;
  created_at: {
    seconds: number;
    nanos: number;
  },
  updated_at: {
    seconds: number;
    nanos: number;
  },
  reply_comments?: ReplyComment[]
}

export type RequestComment = {
  user_id: string,
  question_id: string,
  parent_id: string,
  sort: number,
  description: string;
  is_solution?: boolean;
  noti_token?: string;
}
