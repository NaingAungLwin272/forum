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
  profile: string;
  display_name: string;
  role: string;
  team_id?: string;
  team_name: string;
  department_id?: string;
  department_name: string;
  about_me: string;
  phone: string;
  address: string;
  last_login: any;
  last_post: any;
  dob: any;
  created_at: {
    seconds: number;
    nanos: number;
  };
}

export type CsvData = {
  staff_id: string;
  name: string;
  display_name: string;
  email: string;
  role: string;
  phone: string;
  address: string;
  department_name: string;
  team_name: string;
}

export type UserSummary = {
  questions: number;
  answers: number;
  votes: number;
  solved: number;
  bookmarks: number;
  badges: number;
  notifications: number;
  messages: number;
}

export type FilterUser = {
  name: string;
  department_id: string;
  team_id: string;
  email: string;
}
