import { z } from "astro/zod";

export const gameSchema = z.object({
    id: z.number(),
    layout: z.string(),
    title: z.string(),
    publishDate: z.date(),
    description: z.string(),
});

export type GameModel = z.infer<typeof gameSchema>;