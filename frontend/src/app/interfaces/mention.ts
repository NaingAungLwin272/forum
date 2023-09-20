export type MentionModalData = {
  user: string[]
}

export type MentionFeeds = {
  marker: string;
  feed: string[];
}

export type EditorConfig = {
  toolbar: {
    items: string[];
    shouldNotGroupWhenFull: boolean;
  };
  mention: {
    dropdownLimit: number;
    feeds: MentionFeeds[];
  };
}