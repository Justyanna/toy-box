

import { z } from "astro/zod";
import { reference } from "astro:content";

export const roomSchema = z.object({
  id: z.number(),
  title: z.string(),
  date: z.date(),
  description: z.string(),
  game: reference('games'),
});

export type RoomModel = z.infer<typeof roomSchema>;