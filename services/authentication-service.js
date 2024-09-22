import { v4 as uuidv4 } from "uuid";
import { getVariable } from "../config/getVariables.js";
import { getJWTToken, verifyToken } from "../utils/authentication-utils.js";
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
  userData.userId = userId;

  let createdUser = {};
  let jwtToken;

  const connectedClient = await mongoClient.connect(uri);
  if (!connectedClient) throw new Error("Cannot connect to mongodb");

  try {
    const db = connectedClient.db(getVariable("DATABASE"));
    let users = db.collection('USERS');
    let storedUser = await users.findOne({ username: userData.username});
    if (!storedUser) {
      try {
        await users.insertOne(userData);
        jwtToken = getJWTToken({
          username: userData.username,
          firstName: userData.firstName,
          lastName: userData.lastName
        });

      } catch(err) {
        throw err;
      }
    } else {
      throw new Error("Username already taken");
    }
  } catch(err) {
    throw err;
  }
  return jwtToken;
}

export async function loginUser(user) {
  // try {
  //   let token = '';
  //   const userId = user.userId;
  //   const password = user.password;
  //   const fetchedUser = await getUser(userId);
  //   if (fetchedUser.password === password) {
  //     token = getJWTToken({ userId, password });
  //   }
  //   return token;
  // } catch(err) {
  //   throw Error('Error');
  // }
}

export async function getUser(userId) {
  // try {
  //   let fetchedUser;
  //   var connection = mysql.createConnection(dbConfig);
  //   connection.connect();
  //   const [rows, fields] = await connection.promise().query(`SELECT * FROM User WHERE userId="${userId}"`);
  //   fetchedUser = rows.length ? rows[0] : {};
  //   connection.end();
  //   return fetchedUser;
  // } catch(err) {
  //   throw Error('Error');
  // }
}

export async function verifySession(data) {
  const token = data.token;
  const verdict = await verifyToken(token);
  return verdict;
}
