import { getVariable } from "../config/getVariables.js";
import { getUser } from "../services/authentication-service.js";
import jwt from 'jsonwebtoken';
export function getJWTToken(data) {
    const salt = getVariable("SALT");
    var token = jwt.sign(data, salt, { algorithm: getVariable("JWTALGO") });
    return token;
}

export async function verifyToken(token) {
    try {
        const salt = getVariable("SALT");
        let isValid = true;
        jwt.verify(token, salt, function (err, decoded) {
            if (err) {
                isValid = false;
            }
        });
        return isValid;
    } catch(err) {
        throw err;
    }
}

export async function getRole(token) {
    try {
        const salt = getVariable("SALT");
        let role = '';
        jwt.verify(token, salt, async function (err, decoded) {
            if (!err) {
                const userId = decoded.userId;
                const fetchedUser = await getUser(userId);
                if (fetchedUser) {
                    role = fetchedUser.role;
                }
            }
        });
        return role;
    } catch(err) {
        throw err;
    }
}

export async function getUserId(token) {
    try {
        const salt = getVariable("SALT");
        let userId = '';
        jwt.verify(token, salt, async function (err, decoded) {
            if (!err) {
                userId = decoded.userId;
            }
        });
        return userId;
    } catch(err) {
        throw err;
    }
}