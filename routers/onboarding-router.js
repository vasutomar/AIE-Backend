import {
    health,
    questions
  } from "../controllers/onboarding-controller.js";
  import express from "express";
  
  export const onboardingRouter = express.Router();
  onboardingRouter.get("/onboarding/health", health);
  onboardingRouter.get("/onboarding/questions", questions);