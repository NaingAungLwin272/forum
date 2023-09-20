export type Noti = {
  _id: string,
  user_id: string,
  description: string,
  link: string,
  name: string,
  created_at:string,
  status: boolean,
  type: number
}

export type MapNoti = {
  title: string,
  description: string,
  type: number,
  link: string,
  created_at?: string
}