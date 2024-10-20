import { getVariable } from "../config/getVariables.js";
import { createLogger } from "../utils/logger-utils.js";
import mongo from "mongodb";

const logger = createLogger();
const mongoClient = mongo.MongoClient;
const uri = getVariable("MONGODBURI");

export function checkHealth() {
  logger.info("Profile Service health: :Live");
  return "Live";
}

export async function getUserProfile(username) {
  logger.info("getProfile service : Start");
  let fetchedProfile;
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("getProfile service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let profileCollection = db.collection("PROFILE");
    fetchedProfile = await profileCollection.findOne({
      username,
    });
    if (fetchedProfile) {
      try {
        logger.info("getProfile service : Got Profile");
      } catch (err) {
        throw err;
      }
    } else {
      logger.info("getProfile service : No Profile found");
      throw new Error("No Profile found");
    }
  } catch (err) {
    throw err;
  }
  logger.info("getProfile service : API Completed");
  return fetchedProfile;
}

export async function updateUserProfile(username, body) {
  logger.info("updateProfile service : Start");
  let updatedProfile;
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("updateProfile service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let profileCollection = db.collection("PROFILE");
    try {
      updatedProfile = profileCollection.updateOne(
        { username },
        {
          $set: body,
        },
        { upset: false }
      );
      logger.info("updateProfile service : Profile updated");
    } catch (err) {
      logger.info("updateProfile service : Cannot update profile");
      throw err;
    }
  } catch (err) {
    logger.info(
      "updateProfile service : DB Initialization or connection issues"
    );
    throw err;
  }
  logger.info("updateProfile service : API Completed");
  return updatedProfile;
}
