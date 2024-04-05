/* eslint-disable @typescript-eslint/no-explicit-any */
import prisma from '$lib/prisma.js';
import { redirect, error } from '@sveltejs/kit';
import { customAlphabet } from 'nanoid';

export const actions = {
    async create(event) {
        if(event.request.method == "POST") {
            try {    
                const data = await event.request.formData();
                const url = data.get("url")?.toString();
    
                if (!url) {
                    throw error(400, "URL is null");
                }
    
                const urlPattern = /^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$/;
                if (!urlPattern.test(url)) {
                    throw error(400, "Invalid URL format");
                }
    

                const existingURL = await prisma.registry.findUnique({
                    where: {
                        url
                    }
                })

                if(existingURL) {
                    return redirect(307, `/slug/${existingURL.slug}`)
                }

                const alphabet = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
                const randomizedText = customAlphabet(alphabet);
    
                const slug = randomizedText(10).toString();
    
                await prisma.registry.create({
                    data: {
                        slug,
                        url,
                    }
                })
    
                return redirect(307, `/slug/${slug}`);
            } catch (err: any) {
                error(err.title)
            }
        }
    }
};
