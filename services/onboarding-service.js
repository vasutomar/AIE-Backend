import { getVariable } from "../config/getVariables.js";
import { createLogger } from "../utils/logger-utils.js";
import mongo from "mongodb";

const logger = createLogger();
const mongoClient = mongo.MongoClient;
const uri = getVariable("MONGODBURI");

export function checkHealth() {
  logger.info("Authentication Service health: :Live");
  return "Live";
}

export async function getQuestions() {
  logger.info("getQuestions service : Start");
  let fetchedQuestions;  
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("getQuestions service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let onboardingCollection = db.collection("ONBOARDING");
    let questionDocument = await onboardingCollection.findOne({});
    if (questionDocument) {
      try {
        logger.info("getQuestions service : Got questions");
        fetchedQuestions = questionDocument.questions;
      } catch (err) {
        throw err;
      }
    } else {
      logger.info("getQuestions service : No questions found");
      throw new Error("No questions found");
    }
  } catch (err) {
    throw err;
  }
  logger.info("getQuestions service : API Completed");
  return fetchedQuestions;
}