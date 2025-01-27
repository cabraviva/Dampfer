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

export async function setPermissionForUser(username: string, permission: UserListItem['permission']): Promise<boolean> {
    const req = await knorry('POST', '/api/users/set-permission', {
        username,
        permission
    }, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return (req.data as unknown) === true
}

export async function deleteUser(username: string): Promise<boolean> {
    const req = await knorry('POST', '/api/users/delete', {
        username
    }, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return (req.data as unknown) === true
}

export async function setPasswordForUser(username: string, password: string): Promise<boolean> {
    const req = await knorry('POST', '/api/users/set-password-sysadmin', {
        username,
        password
    }, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return (req.data as unknown) === true
}

export async function changeMyPassword(oldPassword: string, newPassword: string): Promise<true | string> {
    const req = await knorry('POST', '/api/me/change-password', {
        "old-password": oldPassword,
        "password": newPassword
    }, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return (req.data as unknown) === true ? true : req.data as unknown as string
}
