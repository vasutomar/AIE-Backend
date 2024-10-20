import {
  health,
  getDiscussions,
  createDiscussion,
  updateDiscussion,
} from "../controllers/discussion-controller.js";
import express from "express";

export const discussionRouter = express.Router();
discussionRouter.get("/discussion/health", health);
discussionRouter.get("/discussion/:exam", getDiscussions);
discussionRouter.post("/discussion/", createDiscussion);
discussionRouter.patch("/discussion/:id", updateDiscussion);
