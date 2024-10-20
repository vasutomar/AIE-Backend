
import {
    checkHealth,
    getDiscussionsForExam,
    createNewDiscussion,
    modifyDiscussion
} from "../services/discussion-service.js";
import { handleServerError, prepareServerResponse } from "../utils/common-utils.js";

export async function health(req, res) {
    const serverHealth = checkHealth();
    res.send({
        "health": serverHealth
    });
}

export async function getDiscussions(req, res) {
    try {
        const discussions = await getDiscussionsForExam(req.params.exam, req.query.count, req.query.page);
        res.send(prepareServerResponse(200, "Discussions Fetched", discussions));
    } catch (err) {
        handleServerError(err, res);
    }
}

export async function createDiscussion(req, res) {
    try {
        const body = req.body;
        const createdDiscussion = await createNewDiscussion(body);
        res.send(prepareServerResponse(201, "Discussion created", createdDiscussion));
    } catch (err) {
        handleServerError(err, res);
    }
}

export async function updateDiscussion(req, res) {
    try {
        const body = req.body;
        const updatedDiscussion = await modifyDiscussion(req.params.id, body);
        res.send(prepareServerResponse(204, "Discussion updated", updatedDiscussion));
    } catch (err) {
        handleServerError(err, res);
    }
}
