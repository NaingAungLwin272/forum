export type Badge = {
  _id: string;
  name: string,
  description: string;
  type: number;
  level: number;
  created_at: {
    seconds: number;
    nanos: number;
  },
  updated_at: {
    seconds: number;
    nanos: number;
  },
}

export type UserBadge = {
  _id: string;
  user_id: string,
  badge_id: string;
  created_at: {
    seconds: number;
    nanos: number;
  },
  updated_at: {
    seconds: number;
    nanos: number;
  },
}

export type UserPoint = {
  _id?: string;
  user_id?: string;
  reaction_level?: number;
  qa_level?: number;
  question_count?: number;
  answer_count?: number;
  solved_count?: number;
  created_at?: {
    seconds: number;
    nanos: number;
  },
  updated_at?: {
    seconds: number;
    nanos: number;
  }
}