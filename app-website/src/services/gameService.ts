import { getCollection } from "astro:content";

export async function fetchGamesEntries() {
    const gamesEntries = await getCollection("games");
    return gamesEntries.map((entry) => ({
        params: { slug: entry.slug },
        props: { entry },
    }));
}