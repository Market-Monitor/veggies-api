import { MongoClient } from "mongodb";

const uri = process.env.MONGODB_URI!;

export const mongoClient = new MongoClient(uri);
