import axios from "axios";
import {DirectMessage, Message, Space} from "~/utils/model";

export class ChatController {
    sendDirectMessages = async (targets: DirectMessage[], text: string) => {
        return await Promise.all(targets.map(async (target) => {
            return await this.sendDirectMessage(target, text);
        }));
    }

    sendDirectMessage = (target: DirectMessage, text: string) => {
        return new Promise<Message>((resolve, reject) => {
            axios.post(`/api/members/${target.googleUserId}/messages`, {
                text: text
            }).then((res) => {
                const data: Message = res.data;
                resolve(data);
            }).catch((e) => {
                reject(e);
            })
        })
    }

    sendMessages = async (targets: Space[], text: string) => {
        return await Promise.all(targets.map(async (t) => {
            return await this.sendMessage(t, text);
        }));
    }

    sendMessage = (target: Space, text: string) => {
        return new Promise<Message>((resolve, reject) => {
            axios.post(`/api/${target.name}/messages`, {
                text: text
            }).then((res) => {
                const data: Message = res.data;
                resolve(data);
            }).catch((e) => {
                reject(e);
            })
        })
    }

    scheduleMessages = async (targets: Space[], text: string, sendTime: Date) => {
        return Promise.all(targets.map(async (t) => {
            return this.scheduleMessage(t, text, sendTime);
        }));
    }

    scheduleMessage = (target: Space, text: string, sendTime: Date) => {
        return new Promise((resolve, reject) => {
            const t = sendTime.getTime() / 1000;

            axios.post(`/api/schedules`, {
                space: target,
                message: {
                    text: text,
                },
                send_at: t,
            }).then((res) => {
                resolve(res);
            }).catch((e: Error) => {
                reject(e);
            });
        })
    }
}