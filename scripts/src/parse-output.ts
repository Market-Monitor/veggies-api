import { v4 as uuidv4 } from "uuid";
import RESULT_DATA from "../output/result-copy.json";

async function main() {
  const parentIdList = new Map<string, string>();
  const idList = new Map<
    string,
    {
      id: string;
      parentId: string;
    }
  >();

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

          idList.set(catName, {
            id,
            parentId,
          });
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
            id: idList.get(catName)?.id,
            parentId: idList.get(catName)?.parentId,
          };
        });

        return {
          ...x,
          category: data.category,
          id: parentIdList.get(x.name),
        };
      });

      return data;
    });

    const nitem = {
      ...item,
      data: item.data.flatMap((item) => item.vegetables),
    };

    return nitem;
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
