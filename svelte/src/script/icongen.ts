import knorry, { type JSONData, type JSONObject, type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
import { getCredentials } from './login'
defineKnorryOptions({
    easyMode: false
})

export async function searchIcons(query: string): Promise<string[]> {
    const req = await knorry('GET', `/api/icongen/search?q=${encodeURIComponent(query)}`, null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.data as unknown as string[]
}