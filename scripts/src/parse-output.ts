import { v4 as uuidv4 } from "uuid";
import RESULT_DATA from "../output/result-copy.json";

async function main() {
  const parentIdList = new Map<string, string>();
  const idList = new Map<string, string>();

  for (const item of RESULT_DATA) {
    for (const data of item.data) {
      for (const x of data.vegetables) {
        const parentId = uuidv4();

        // Set only parent id if not already set
        if (!parentIdList.get(x.name)) {
          parentIdList.set(x.name, parentId);
        }

        for (const category of x.data) {
          const id = uuidv4();

          const catName = `${x.name} ${category.name}`.trim();

          // If id already exists, skip
          if (idList.get(catName)) {
            continue;
          }

          idList.set(catName, id);
        }
      }
    }
  }

  const parsed = RESULT_DATA.map((item) => {
    item.data = item.data.map((data) => {
      data.vegetables = data.vegetables.map((x) => {
        x.data = x.data.map((category) => {
          const catName = `${x.name} ${category.name}`.trim();

          return {
            ...category,
            id: idList.get(catName),
          };
        });

        return {
          ...x,
          id: parentIdList.get(x.name),
        };
      });

      return data;
    });

    return item;
  });

  await Bun.write(
    "./output/parent-ids.json",
    JSON.stringify(Object.fromEntries(parentIdList.entries()), null, 2)
  );
  await Bun.write(
    "./output/classes-ids.json",
    JSON.stringify(Object.fromEntries(idList.entries()), null, 2)
  );

  await Bun.write("./output/PARSED.json", JSON.stringify(parsed, null, 2));
}

main();
