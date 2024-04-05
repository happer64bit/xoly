import prisma from '$lib/prisma'
import { error } from '@sveltejs/kit'

export async function load({ params, request }) {
    if(request.method == 'GET') {
        const { slug }= params
    
        try {
            const searchedSlug = await prisma.registry.findUnique({
                where: {
                    slug
                }
            })
    
            if(searchedSlug) {
                return searchedSlug
            } else {
                return error(404)
            }
            
        } catch (err) {
            return error(500)
        }
    }
    
}