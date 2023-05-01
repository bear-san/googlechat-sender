import { Space, DirectMessage } from "~/utils/model";
import axios from "axios";
import camelcaseKeys from 'camelcase-keys';

export class SpaceController {
    findDirectMessages = async (users: DirectMessage[]) => {
        return await Promise.all(users.map(async (u) => {
            return await this.findDirectMessage(u);
        }))
    }

    findDirectMessage = (u: DirectMessage) => {
        return new Promise<Space>(async (resolve, reject) => {
            axios.get(`/api/members/${u.googleUserId}/space`).then((res) => {
                const data: Space = res.data;
                resolve(data);
            }).catch((err) => {
                reject(err);
            });
        });
    }
}