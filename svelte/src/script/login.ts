import knorry, { type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
defineKnorryOptions({
    easyMode: false
})

export function credentialsSaved(): boolean {
    return typeof getCredentials() === 'string'
}

export function getCredentials(): string | null {
    return localStorage.getItem('credentials.jwt')
}

export function setCredentials(jwt: string): void {
    localStorage.setItem('credentials.jwt', jwt)
}

export function getSavedUsername(): string | null {
    return localStorage.getItem('credentials.user')
}

export function setSavedUsername(user: string): void {
    localStorage.setItem('credentials.user', user)
}

export async function isJWTValid(): Promise<boolean> {
    const req = await knorry('GET', '/api/endpoints', null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.status == 200
}