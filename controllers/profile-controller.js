
import {
    checkHealth,
    getUserProfile,
    updateUserProfile,
} from "../services/profile-service.js";
import { handleServerError, prepareServerResponse } from "../utils/common-utils.js";

export async function health(req, res) {
    const serverHealth = checkHealth();
    res.send({
        "health": serverHealth
    });
}

export async function getProfile(req, res) {
    try {
        const userProfile = await getUserProfile(req.params.username);
        res.send(prepareServerResponse(201, "Profile Fetched", userProfile));
    } catch (err) {
        handleServerError(err, res);
    }
}

export async function updateProfile(req, res) {
    try {
        const body = req.body;
        const updatedUserProfile = await updateUserProfile(req.params.username, body);
        res.send(prepareServerResponse(204, "User profile updated", updatedUserProfile));
    } catch (err) {
        handleServerError(err, res);
    }
}
