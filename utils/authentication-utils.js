import { getVariable } from "../config/getVariables.js";
import jwt from 'jsonwebtoken';
import crypto from 'crypto';

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

export async function getUsername(token) {
    try {
        const salt = getVariable("SALT");
        let username = '';
        jwt.verify(token, salt, async function (err, decoded) {
            if (!err) {
                username = decoded.username;
            }
        });
        return username;
    } catch(err) {
        throw err;
    }
}

export function hashPassword(password) {
    const secret = getVariable('PASSALGO');  
    const hash = crypto.createHmac('sha256', secret)  
                       .update(password)  
                       .digest('hex'); 
    return hash; 
    //Dehashing algorithm
    /*
        var mykey = crypto.createDecipher('aes-128-cbc', 'mypassword');
        var mystr = mykey.update('34feb914c099df25794bf9ccb85bea72', 'hex', 'utf8')
        mystr += mykey.final('utf8');
    */
}
