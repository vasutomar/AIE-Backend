
import {
    checkHealth,
    createUser,
    loginUser,
    verifySession
} from "../services/authentication-service.js";
import { handleServerError, prepareServerResponse } from "../utils/common-utils.js";

export async function health(req, res) {
    const serverHealth = checkHealth();
    res.send({
        "health": serverHealth
    });
}

export async function signup(req, res) {
    try {
        const body = req.body;
        const createdUser = await createUser(body);
        res.send(prepareServerResponse(201, "User created", createdUser));
    } catch (err) {
        handleServerError(err, res);
    }
}

export async function signin(req, res) {
    try {
        const body = req.body;
        const token = await loginUser(body.username, body.password);
        if (token) {
            res.send(prepareServerResponse(200, "User logged in!", token));
        } else {
            res.send(prepareServerResponse(204, "Invalid email or password", token));
        }
    } catch (err) {
        handleServerError(err, res);
    }
}

export async function verify(req, res) {
    try {
        const body = req.body;
        const verdict = await verifySession(body);
        if (verdict) {
            res.send(prepareServerResponse(200, "Valid session", null));
        } else {
            res.send(prepareServerResponse(401, "Invalid session", null));
        }
    } catch (err) {
        handleServerError(err, res);
    }
}