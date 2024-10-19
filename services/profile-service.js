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

export async function createUserProfile(body) {
  logger.info("createProfile service : Start");
  let createdProfile;
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("createProfile service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let profileCollection = db.collection("PROFILE");
    try {
      createdProfile = await profileCollection.insertOne(body);
      logger.info("createProfile service : Profile created");
    } catch (err) {
      logger.info("createProfile service : Cannot create profile");
      throw err;
    }
  } catch (err) {
    logger.info(
      "createProfile service : DB Initialization or connection issues"
    );
    throw err;
  }
  logger.info("createProfile service : API Completed");
  return createdProfile;
}

export async function deleteUserProfile(username) {
  logger.info("deleteProfile service : Start");
  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("deleteProfile service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let profileCollection = db.collection("PROFILE");
    try {
      await profileCollection.deleteOne({
        username,
      });
      logger.info("deleteProfile service : Profile deleted");
    } catch (err) {
      logger.info("deleteProfile service : Cannot delete profile");
      throw err;
    }
  } catch (err) {
    logger.info(
      "deleteProfile service : DB Initialization or connection issues"
    );
    throw err;
  }
  logger.info("deleteProfile service : API Completed");
  return true;
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
      updatedProfile = await profileCollection.updateOne({ username }, body);
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
