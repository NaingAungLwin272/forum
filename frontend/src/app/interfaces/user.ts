export type SigninUser = {
  email: string;
  password: string;
  is_remember_me: boolean;
}

export type Email = {
  email: string;
}
export type User = {
  _id: string;
  staff_id: string;
  name: string;
  email: string;
  phone: string;
  profile: string;
  password: string;
  address: string;
  dob: any;
  display_name: string;
  role: string;
  team_id: string;
  department_id: string;
  about_me: string;
  mail_subscribe: boolean;
  last_post: any;
  last_login: any;
  noti_token?: string;
  created_at: {
    seconds: number;
    nanos: number;
  };
}

export type UserSummary = {
  questions: number;
  answers: number;
  votes: number;
  solved: number;
  bookmarks: number;
  badges: number;
  mentions: number;
  notifications: number;
  messages: number;
}

export type UserModalData = {
  user: User[];
  user_id: string,
  question_id: string,
  parent_id: string,
  sort: number,
  comment_id?: string;
  description?: string;
  type: 'create' | 'update';
  count: number;
  question_detail_user_id?: string;
}
