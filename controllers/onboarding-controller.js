
import {
    checkHealth,
    getQuestions
} from "../services/onboarding-service.js";
import { verifyToken } from "../utils/authentication-utils.js";
import { handleServerError, prepareServerResponse } from "../utils/common-utils.js";

export async function health(req, res) {
    const serverHealth = checkHealth();
    res.send({
        "health": serverHealth
    });
}

export async function questions(req, res) {
    try {
        const headers = req.headers;
        const isValidSession = await verifyToken(headers.authorization);
        if (isValidSession) {
            const questions = await getQuestions();
            res.send(prepareServerResponse(200, "Questions fetced", questions));
        } else {
            res.send(prepareServerResponse(419, "Session expired"));
        }
    } catch (err) {
        handleServerError(err, res);
    }
}
