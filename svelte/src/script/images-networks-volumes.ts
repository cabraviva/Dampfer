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

export async function deleteImage(repository: string, tag: string): Promise<boolean> {
    const req = await knorry('GET', `/api/docker/image/rm?repository=${encodeURIComponent(repository)}&tag=${encodeURIComponent(tag)}`, null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.data as unknown as boolean
}

export async function inspectImage(id: string): Promise<InspectedImage> {
    const req = await knorry('POST', `/api/docker/image/inspect`, {
        id
    }, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    let imagelist = req.data as unknown as any[]

    return imagelist[0] as InspectedImage
}

export interface InspectedImage {
    "Architecture": "amd64" | "arm" | string,
    "Author": string,
    "Comment": "buildkit.dockerfile.v0" | string,
    "Config": {
        "ArgsEscaped": boolean,
        "AttachStderr": boolean,
        "AttachStdin": boolean,
        "AttachStdout": boolean,
        "Cmd": string[],
        "Domainname": string,
        "Entrypoint": unknown | null,
        "Env": string[]
        "Hostname": string,
        "Image": string,
        "Labels": null,
        "OnBuild": null,
        "OpenStdin": boolean,
        "StdinOnce": boolean,
        "Tty": boolean,
        "User": string,
        "Volumes": unknown | null,
        "WorkingDir": string
    },
    "ContainerConfig": {
        "AttachStderr": boolean,
        "AttachStdin": boolean,
        "AttachStdout": boolean,
        "Cmd": string[] | unknown | null,
        "Domainname": string,
        "Entrypoint": unknown | null,
        "Env": unknown | null,
        "Hostname": string,
        "Image": string,
        "Labels": unknown | null,
        "OnBuild": unknown | null,
        "OpenStdin": boolean,
        "StdinOnce": boolean,
        "Tty": boolean,
        "User": string,
        "Volumes": unknown,
        "WorkingDir": string
    },
    "Created": string,
    "DockerVersion": string,
    "GraphDriver"?: {
        "Data": unknown | null,
        "Name": string
    },
    "Id": string,
    "Metadata": {
        "LastTagTime": string,
        [key: string]: unknown
    },
    "Os": "linux" | string,
    "Parent": string,
    "RepoDigests": string
    "RepoTags": string
    "RootFS": unknown,
    "Size": number,
    "VirtualSize": number
}

export interface ListedImage {
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

export interface ListedImagesGrouped {
    [Repository: string]: ListedImagesGroupedItem
}

export interface ListedImagesGroupedItem {
    tags: string[],
    images: ListedImage[]
}