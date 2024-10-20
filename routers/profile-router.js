import {
  health,
  getProfile,
  updateProfile,
} from "../controllers/profile-controller.js";
import express from "express";

export const profileRouter = express.Router();
profileRouter.get("/profile/health", health);
profileRouter.get("/profile/:username", getProfile);
profileRouter.patch("/profile/:username", updateProfile);
