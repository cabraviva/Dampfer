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
    const req = await knorry('GET', '/api/jwt-valid-check', null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.status == 200
}

/**
 * Tries to login with given credentials and returns whether the login attempt was successful, a bold text and a message
 */
export async function login(username: string, password: string): Promise<[boolean, string, string]> {
    const req = await knorry('POST', '/login', {
        username,
        password
    }) as KnorryResponseObj

    if (req.status == 403) {
        // Login failed
        return [false, 'Login failed:', 'Invalid credentials\nTip: Default user is admin with password "admin"']
    } else if (req.status == 200) {
        // Login successful
        const jwt = req.data as string
        setCredentials(jwt)
        if (!isJWTValid()) {
            return [false, 'Login failed:', 'Attempt was successful but jwt seems to be invalid. Please see log for details!']
        }

        return [true, 'Login successful', '']
    } else {
        // Internal Server Error
        const errMsg = req.data
        return [false, `Login failed:`, `Internal sever error, please see logs for details: ${errMsg}`]
    }
}