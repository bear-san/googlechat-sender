export type Space = {
    name: string,
    type: string,
    displayName: string,
    spaceType: string
}

export type DirectMessage = {
    employeeNumber: string,
    email: string,
    googleUserId: string,
    displayName: string
}

export type Message = {
    text: string
}