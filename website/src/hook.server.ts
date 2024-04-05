import { DB_URL } from "$env/static/private";
import mongoose from "mongoose";

/** @type {import('@sveltejs/kit').Handle} */
export async function handle() {
    mongoose.connect(DB_URL)
}