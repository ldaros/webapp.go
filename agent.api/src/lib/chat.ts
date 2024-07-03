export enum ChatRole {
  System = "system",
  Assistant = "assistant",
  User = "user",
}

export type ChatMessage = {
  role: ChatRole;
  content: string;
};
