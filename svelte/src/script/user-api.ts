import knorry, { type JSONData, type JSONObject, type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
import { getCredentials } from './login'
defineKnorryOptions({
    easyMode: false
})

export interface UserListItem {
    username: string,
    permission: 'insight' | 'admin' | 'system-admin'
}

export type UserList = UserListItem[]

export async function listUsers(): Promise<UserList> {
    const req = await knorry('GET', '/api/users/ls', null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.data as unknown as UserList
}

export interface CreateUserData {
    username: string,
    password: string,
    permission: UserListItem["permission"]
}

export async function createUser(userData: CreateUserData): Promise<boolean> {
    const req = await knorry('POST', '/api/users/create', userData, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return (req.data as unknown) === true
}