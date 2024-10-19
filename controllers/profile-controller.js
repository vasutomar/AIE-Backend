
import {
    checkHealth,
    getUserProfile,
    createUserProfile,
    updateUserProfile,
    deleteUserProfile
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

export async function createProfile(req, res) {
    try {
        const body = req.body;
        const createdUserProfile = await createUserProfile(body);
        res.send(prepareServerResponse(201, "User profile created", createdUserProfile));
    } catch (err) {
        handleServerError(err, res);
    }
}

export async function deleteProfile(req, res) {
    try {
        const deletedUserProfile = await deleteUserProfile(req.params.username);
        res.send(prepareServerResponse(200, "User profile deleted", deletedUserProfile));
    } catch (err) {
        handleServerError(err, res);
    }
}