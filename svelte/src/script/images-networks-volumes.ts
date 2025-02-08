import knorry, { type JSONData, type JSONObject, type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
import { getCredentials } from './login'
import type { List } from 'flowbite-svelte'
defineKnorryOptions({
    easyMode: false
})

export async function listImages(): Promise<ListedImage[]> {
    const req = await knorry('GET', `/api/docker/image/list`, null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    let imagelist = req.data as unknown as any[]

    imagelist = imagelist.map(image => {
        image.CreatedAt = new Date(image.CreatedAt.replace(/[A-Z]/gi, ''))
        return image
    })

    return imagelist as ListedImage[]
}

interface ListedImage {
    Containers: string,
    CreatedAt: Date,
    CreatedSince: string,
    Digest: string,
    ID: string,
    Repository: string,
    SharedSize: string,
    Size: string,
    Tag: string,
    UniqueSize: string,
    VirtualSize: string
}