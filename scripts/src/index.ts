import { generateObject } from "ai";
import data from "./data/DATA.json";
import { openai } from "./openai";

async function main() {
  const filteredData = data.filter((item) =>
    item.content?.includes("𝐁𝐀𝐏𝐓𝐂 𝐌𝐎𝐍𝐈𝐓𝐎𝐑𝐄𝐃 𝐖𝐇𝐎𝐋𝐄𝐒𝐀𝐋𝐄 𝐏𝐑𝐈𝐂𝐄𝐒")
  );

  const sysPrompt = `
Parse the DATASET and return the following JSON object:

<JSON>
{
  "data": [
    {
      "category": "SOLID | SARI-SARI",
      "vegetables": [
        {
          "name": "[vegetable name]",
          "price": [] // price range: [lowest, highest]
        },
        // or
        {
          "name": "[vegetable name]",
          "categories": [
            {
              "name": "[category name]",
              "price": [] // price range: [lowest, highest]
            }
          ]
        }
      ]
    }
  ],
  "date": "[date]"
}
</JSON>
 
<DATA-SET>
[[data-set]]
</DATA-SET>
    `.trim();

  let parsedOutput: Array<any> = [];

  for (const item of filteredData) {
    const result = await generateObject({
      model: openai("gpt-4o"),
      prompt: sysPrompt.replace("[[data-set]]", item.content ?? ""),
      output: "no-schema",
    });

    console.log("[i] ...done");

    parsedOutput.push(result.object);
  }

  // Save parsed output
  await Bun.write(
    "./output/result.json",
    JSON.stringify(parsedOutput, null, 2)
  );
}

main()
  .then(() => {
    console.log("[i] DONE");
  })
  .catch((err) => {
    console.error("[e] ERROR", err);
  });
