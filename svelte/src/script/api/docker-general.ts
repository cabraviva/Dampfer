import knorry, { type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
import { getCredentials } from '../login'
defineKnorryOptions({
    easyMode: false
})

interface DockerReadyStatus {
  "Ready": boolean,
  "Installed": boolean,
  "DaemonRunning": boolean,
  "ComposeInstalled": boolean,
  "ComposeVersion": 'v1' | 'v2' | 'NOT_FOUND',
  "Msg": string
}

export async function dockerReady(): Promise<DockerReadyStatus> {
    const req = await knorry('GET', '/api/docker/ready', null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return (req.data as unknown) as DockerReadyStatus
}