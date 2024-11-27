import { mongoClient } from "./mongodb";

import VeggiesClasses from "../output/classes-ids.json";
import Veggies from "../output/parent-ids.json";
import HistoryPrices from "../output/PARSED.json";

async function main() {
  const db = mongoClient.db("MarketMonitor");

  // Collections
  const veggies = db.collection("Veggies");
  const veggiesClasses = db.collection("VeggiesClasses");
  const historyPrices = db.collection("HistoryPrices");

  // Insert veggies
  for (const [name, id] of Object.entries(Veggies)) {
    await veggies.insertOne({ name, id });
  }

  // Insert veggies classes
  for (const [name, id] of Object.entries(VeggiesClasses)) {
    await veggiesClasses.insertOne({ name, id });
  }

  // Insert history prices
  for (const item of HistoryPrices) {
    await historyPrices.insertOne(item);
  }

  await mongoClient.close();
}

main();
