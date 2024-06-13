import { defineCollection } from "astro:content";
import { gameSchema } from "../models/game";
import { roomSchema } from "../models/room";

const games = defineCollection({
  type: 'content',
  schema: gameSchema
});

const rooms = defineCollection({
  type: 'content',
  schema: roomSchema
});

export const collections = {rooms, games};