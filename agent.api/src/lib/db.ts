import { ChatMessage, ChatRole } from "./chat";
import * as fs from "fs/promises";

type Store = {
  messages: Map<string, ChatMessage[]>;
};

const FILE_PATH = "db.json";

export class JsonDatabase {
  private static instance: JsonDatabase;
  private store: Store = { messages: new Map() };

  private constructor() {}

  public static async getInstance(): Promise<JsonDatabase> {
    if (!JsonDatabase.instance) {
      JsonDatabase.instance = new JsonDatabase();
      await JsonDatabase.instance.loadStore();
    }
    return JsonDatabase.instance;
  }

  public async addMessage(
    chatId: string,
    message: ChatMessage
  ): Promise<ChatMessage[]> {
    const chatMessages = this.store.messages.get(chatId) ?? [];
    chatMessages.push(message);
    this.store.messages.set(chatId, chatMessages);
    await this.saveStore();
    return chatMessages;
  }

  public async getMessages(chatId: string): Promise<ChatMessage[]> {
    return this.store.messages.get(chatId) ?? [];
  }

  public async createMessages(chatId: string): Promise<ChatMessage[]> {
    this.store.messages.set(chatId, [
      { role: ChatRole.Assistant, content: "Hello! How can I help you today?" },
    ]);
    await this.saveStore();
    return this.store.messages.get(chatId) ?? [];
  }

  private async loadStore(): Promise<void> {
    try {
      const data = await fs.readFile(FILE_PATH, "utf-8");
      const parsedData = JSON.parse(data) as {
        messages: Record<string, ChatMessage[]>;
      };
      this.store.messages = new Map(Object.entries(parsedData.messages));
    } catch (error) {
      if ((error as NodeJS.ErrnoException).code !== "ENOENT") {
        console.error("Error reading the database file:", error);
      }
    }
  }

  private async saveStore(): Promise<void> {
    try {
      const data = JSON.stringify(
        { messages: Object.fromEntries(this.store.messages) },
        null,
        2
      );
      await fs.writeFile(FILE_PATH, data, "utf-8");
    } catch (error) {
      console.error("Error writing to the database file:", error);
    }
  }
}
