import {
    health,
    getProfile,
    createProfile,
    updateProfile,
    deleteProfile
  } from "../profile/profile-controller.js";
  import express from "express";
  
  export const profileRouter = express.Router();
  profileRouter.get("/profile/health", health);
  profileRouter.get("/profile/:username", getProfile);
  profileRouter.post("/profile/", createProfile);
  profileRouter.patch("/profile/:username", updateProfile);
  profileRouter.delete("/profile/:username", deleteProfile);