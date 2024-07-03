import express from "express";
import dotenv from "dotenv";
import { Logger } from "./lib/logger";
import { JsonDatabase } from "./lib/db";
import { OpenAI } from "./lib/open-ai";
import { ChatRole } from "./lib/chat";

dotenv.config();

const app = express();
const port = process.env.PORT || 3000;

const logger = Logger.getInstance();

// Middleware to parse JSON bodies
app.use(express.json());

app.get("/api/chat", async (req, res) => {
  try {
    const db = await JsonDatabase.getInstance();
    const id = req.query.id as string;

    if (!id) {
      res.status(400).send("id is required");
      return;
    }

    let chat = await db.getMessages(id);

    if (!chat) {
      chat = await db.createMessages(id);
    }

    res.send(chat);
  } catch (error) {
    const e = error as Error;
    logger.error(`Error in /api/chat GET: ${e.message}`);
    res.status(500).send("Internal server error");
  }
});

app.post("/api/chat", async (req, res) => {
  try {
    const db = await JsonDatabase.getInstance();
    const openAI = OpenAI.getInstance({
      apiUrl: process.env.OPENAI_API_URL!,
      apiKey: process.env.OPENAI_API_KEY!,
      model: process.env.OPENAI_MODEL!,
    });

    const id = req.query.id as string;
    const { message } = req.body;

    logger.debug(`Received request with id: ${id} and message: ${message}`);

    if (!id || !message) {
      res.status(400).send("id and message are required");
      return;
    }

    let chat = await db.getMessages(id);
    if (!chat) {
      chat = await db.createMessages(id);
    }

    chat = await db.addMessage(id, {
      role: ChatRole.User,
      content: message,
    });

    const response = await openAI.chat(
      chat.map((msg) => ({
        role: msg.role,
        content: msg.content,
      }))
    );

    const openAIResponse = response.choices[0].message;

    await db.addMessage(id, {
      role: ChatRole.Assistant,
      content: openAIResponse.content,
    });

    res.send(chat);
  } catch (error) {
    const e = error as Error;
    logger.error(`Error in /api/chat POST: ${e.message}`);
    res.status(500).send("Internal server error");
  }
});

app.listen(port, () => {
  logger.info(`Server is running on port ${port}`);
});
