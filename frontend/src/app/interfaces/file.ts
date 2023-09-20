export type uploadFile = {
  uid: string;
  name: string;
  size: number;
  type: string;
  lastModified?: number;
  lastModifiedDate?: Date;
  url?: string;
  thumbUrl?: string;
  originFileObj?: File;
}