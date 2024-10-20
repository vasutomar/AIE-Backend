import { v4 as uuidv4 } from "uuid";
import { getVariable } from "../config/getVariables.js";
import {
  getJWTToken,
  hashPassword,
  verifyToken,
  getUsername,
} from "../utils/authentication-utils.js";
import { createLogger } from "../utils/logger-utils.js";
import mongo from "mongodb";

const logger = createLogger();
const mongoClient = mongo.MongoClient;
const uri = getVariable("MONGODBURI");

export function checkHealth() {
  logger.info("Authhentication Service health: :Live");
  return "Live";
}

export async function createUser(userData) {
  logger.info("createUser service : Start");

  const userId = uuidv4();
  const { username, firstName, lastName, password } = userData;
  userData.userId = userId;
  let jwtToken;

  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("createUser service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let usersCollection = db.collection("USERS");
    let profileCollection = db.collection("PROFILE");
    let storedUser = await usersCollection.findOne({ username });
    if (!storedUser) {
      try {
        logger.info("createUser service : Creating user");
        const hashedPassword = hashPassword(password);
        userData.password = hashedPassword;
        const { phone, email, ...userInfo } = userData;
        await usersCollection.insertOne(userInfo);
        logger.info("createUser service : User created");
        await profileCollection.insertOne({
          username: userData.username,
          phone,
          email,
          exams: [],
          salt: ''
        });
        logger.info("createUser service : User profile created");
        jwtToken = getJWTToken({
          username,
          firstName,
          lastName,
        });
      } catch (err) {
        throw err;
      }
    } else {
      logger.info("createUser service : Username already present");
      throw new Error("Username already taken");
    }
  } catch (err) {
    throw err;
  }
  logger.info("createUser service : User signup completed");
  return jwtToken;
}

export async function loginUser(username, password) {
  const hashedPassword = hashPassword(password);

  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    logger.info("loginUser service : MongoDB Connection established");
    const db = connectedClient.db(getVariable("DATABASE"));
    let users = db.collection("USERS");
    let storedUser = await users.findOne({ username });

    var jwtToken = "";

    if (!storedUser) {
      logger.error("loginUser service : User not found");
      throw new Error("User not found, Incorrect username");
    } else {
      try {
        logger.info("loginUser service : User found");
        let authenticatedUser = await users.findOne({
          username,
          password: hashedPassword,
        });
        if (authenticatedUser) {
          jwtToken = getJWTToken({
            username,
            firstName: authenticatedUser.firstName,
            lastName: authenticatedUser.lastName,
          });
          logger.info("loginUser service : Valid User");
        } else {
          throw new Error("loginUser service : Invalid password");
        }
      } catch (err) {
        throw err;
      }
    }
  } catch (err) {
    throw err;
  }
  logger.info("loginUser service : User signin completed");
  return jwtToken;
}

export async function verifySession(data) {
  const token = data.token;
  const verdict = await verifyToken(token);
  return verdict;
}
