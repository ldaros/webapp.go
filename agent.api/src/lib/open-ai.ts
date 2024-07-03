import axios, { AxiosResponse } from "axios";

export type OpenAIConfig = {
  apiUrl: string;
  apiKey: string;
  model: string;
};

export type OpenAIRequest = {
  model: string;
  messages: {
    role: string;
    content: string;
  }[];
  temperature: number;
};

export type OpenAIResponse = {
  id: string;
  object: string;
  created: number;
  model: string;
  usage: {
    prompt_tokens: number;
    completion_tokens: number;
    total_tokens: number;
  };
  choices: {
    message: {
      role: string;
      content: string;
    };
    logprobs: any;
    finish_reason: string;
    index: number;
  }[];
};

export class OpenAI {
  private static instance: OpenAI;
  private config: OpenAIConfig;

  private constructor(config: OpenAIConfig) {
    this.config = config;
  }

  public static getInstance(config: OpenAIConfig): OpenAI {
    if (!OpenAI.instance) {
      if (!config.apiUrl) {
        throw new Error("apiUrl is required");
      }
      if (!config.apiKey) {
        throw new Error("apiKey is required");
      }
      if (!config.model) {
        throw new Error("model is required");
      }

      OpenAI.instance = new OpenAI(config);
    }
    return OpenAI.instance;
  }

  public async chat(
    messages: OpenAIRequest["messages"]
  ): Promise<OpenAIResponse> {
    try {
      const response: AxiosResponse<OpenAIResponse> = await axios.post(
        this.config.apiUrl,
        {
          model: this.config.model,
          messages,
          temperature: 0.5,
        },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${this.config.apiKey}`,
          },
        }
      );

      return response.data;
    } catch (error) {
      if (axios.isAxiosError(error) && error.response) {
        throw new Error(`Error: ${error.response.statusText}`);
      }

      const e = error as Error;
      throw new Error(`Error: ${e.message}`);
    }
  }
}
