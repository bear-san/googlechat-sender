import axios from "axios";
import { DirectMessage, Message, Space } from "~/utils/model";
import { SpaceController } from "~/utils/spaceController";

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
}