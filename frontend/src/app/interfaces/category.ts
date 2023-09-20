export type Category = {
  _id: string,
  name: string,
  type: number,
  created_at: {
    seconds: number;
    nanos: number;
  },
  updated_at: {
    seconds: number;
    nanos: number;
  },
}