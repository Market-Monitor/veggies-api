import { createOpenAI } from "@ai-sdk/openai";

export const openai = createOpenAI({
  compatibility: "strict",
  apiKey: process.env.OPENAI_API_KEY!, // OPENAI_API_KEY || DEEPSEEK_API_KEY
  // baseURL: "https://api.deepseek.com",
  // fetch: (input, init) => {
  //   if (init) {
  //     init.body = JSON.stringify({
  //       ...JSON.parse(init.body as string),
  //       ...{ response_format: { type: "json_object" } },
  //     });
  //   }

  //   return fetch(input, init);
  // },
});
