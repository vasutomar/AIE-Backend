import { getVariable } from "../config/getVariables.js";
import { createLogger } from "../utils/logger-utils.js";
import mongo, { ObjectId } from "mongodb";

const logger = createLogger();
const mongoClient = mongo.MongoClient;
const uri = getVariable("MONGODBURI");

export function checkHealth() {
  logger.info("Discussion Service health: :Live");
  return "Live";
}

export async function getDiscussionsForExam(exam, count, page) {
  logger.info("getDiscussion service : Start");
  let fetchedDiscussions;
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("getDiscussion service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let discussionCollection = db.collection("DISCUSSIONS");
    try {
      fetchedDiscussions = discussionCollection
        .find({ exam })
        .skip((page - 1) * count)
        .limit(parseInt(count, 10))
        .toArray(function(err, result) {
            if (err) throw err;
            console.log(result);
            fetchedDiscussions = result;
        });
    } catch (err) {
      logger.error("getDiscussion service : Error while getting discussions");
      throw err;
    }
  } catch (err) {
    logger.error("getDiscussion service : Error while establishing DB Connection");
    throw err;
  }
  logger.info("getDiscussion service : API Completed");
  return fetchedDiscussions;
}

export async function createNewDiscussion(body) {
  logger.info("createDiscussion service : Start");
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("createDiscussion service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let discussionCollection = db.collection("DISCUSSIONS");
    try {
      await discussionCollection.insertOne(body);
      logger.info("createDiscussion service : Discussion created");
    } catch (err) {
      logger.info("createDiscussion service : Cannot create discussion");
      throw err;
    }
  } catch (err) {
    logger.error(
      "createDiscussion service : DB Initialization or connection issues"
    );
    throw err;
  }
  logger.info("createDiscussion service : API Completed");
  return true;
}

export async function modifyDiscussion(id, body) {
  logger.info("modifyDiscussion service : Start");
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("modifyDiscussion service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let discussionCollection = db.collection("DISCUSSIONS");
    try {
      updatedProfile = discussionCollection.updateOne(
        { "_id": ObjectId(id) },
        {
          $set: body,
        },
        { upset: false }
      );
      logger.info("modifyDiscussion service : Discussion updated");
    } catch (err) {
      logger.info("modifyDiscussion service : Cannot update discussion");
      throw err;
    }
  } catch (err) {
    logger.info(
      "modifyDiscussion service : DB Initialization or connection issues"
    );
    throw err;
  }
  logger.info("modifyDiscussion service : API Completed");
  return true;
}
