import { dbConfig } from "../config/mysql-config";
import { User } from "../models/User";
import { v4 as uuidv4 } from "uuid";
import { getJWTToken, verifyToken } from "../utils/authentication-utils";
import { createLogger } from "../utils/logger.utils";
var mysql = require('mysql2');

const logger = createLogger();

export function checkHealth() {
  logger.info('Authhentication Service health: :Live');
  return "Live";
}

export function createUser(userData) {
  logger.info('createUser service : Start');
  const userId = uuidv4();
  userData.userId = userId;
  var connection = mysql.createConnection(dbConfig);
  connection.connect();
  connection.query(`INSERT INTO User (
    userId,
    firstName,
    lastName,
    role,
    cartId,
    email,
    password) VALUES(?, ?, ?, ?, ?, ?, ?)`, [
        userData.userId,
        userData.firstName,
        userData.lastName,
        userData.role,
        null,
        userData.email,
        userData.password,
    ], (error, results) => {
        if (error) {
            throw Error(error); 
        }
    });
  logger.info('createUser service : Database insertion complete');
  connection.end();
  return userData;
}

export async function loginUser(user) {
  try {
    let token = '';
    const userId = user.userId;
    const password = user.password;
    const fetchedUser = await getUser(userId);
    if (fetchedUser.password === password) {
      token = getJWTToken({ userId, password });
    }
    return token;
  } catch(err) {
    throw Error('Error');
  }
}

export async function getUser(userId) {
  try {
    let fetchedUser;
    var connection = mysql.createConnection(dbConfig);
    connection.connect();
    const [rows, fields] = await connection.promise().query(`SELECT * FROM User WHERE userId="${userId}"`);
    fetchedUser = rows.length ? rows[0] : {};
    connection.end();
    return fetchedUser;
  } catch(err) {
    throw Error('Error');
  }
}

export async function verifySession(data) {
  const token = data.token;
  const verdict = await verifyToken(token);
  return verdict;
}