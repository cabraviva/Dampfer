import knorry, { type JSONData, type JSONObject, type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
import { getCredentials } from './login'
defineKnorryOptions({
    easyMode: false
})

export interface UserInfo {
    username: string,
    permission: 'insight' | 'admin' | 'system-admin'
    insight: boolean
    admin: boolean
    systemAdmin: boolean
}

export async function whoami(): Promise<UserInfo> {
    const req = await knorry('GET', '/api/whoami', null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.data as unknown as UserInfo
}

export async function pwChangeCount(): Promise<number> {
    const req = await knorry('GET', '/api/me/how-often-was-pw-changed', null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return Number.isInteger(req.data as unknown as number) ? (req.data as unknown as number) : 0
}
