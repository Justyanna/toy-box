import { getCollection } from "astro:content";

export async function fetchRoomEntries() {
    const roomEntries = await getCollection("rooms");
    return roomEntries.map((entry) => ({
        params: { slug: entry.slug },
        props: { entry },
    }));
}
