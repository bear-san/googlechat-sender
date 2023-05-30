<template>
  <v-app>
      <v-app-bar color="green-darken-2">
          <v-app-bar-title>
              Google Chat Sender
          </v-app-bar-title>
      </v-app-bar>
      <v-main>
          <v-container>
              <v-card>
                  <v-card-title>新規メッセージの作成</v-card-title>
                  <v-card-text>
                      <v-form>
                          <v-row>
                              <v-col>
                                  <v-combobox
                                      chips
                                      multiple
                                      :items="state.spaces"
                                      v-model="state.selectedSpaces"
                                      clearable
                                      id="send-target-space"
                                      label="送り先のスペース"
                                      :item-title="item => (item.displayName)"
                                  >
                                  </v-combobox>
                              </v-col>
                              <v-col>
                                  <v-combobox
                                      chips
                                      multiple
                                      :items="state.directMessages"
                                      v-model="state.selectedDirectMessages"
                                      clearable
                                      id="send-target-dm"
                                      label="ダイレクトメッセージ"
                                      :item-title="item => (item.displayName)"
                                  >
                                      <template v-slot:selection="{ attrs, item, select, selected }">
                                          <v-chip
                                              v-bind="attrs"
                                              :model-value="selected"
                                              closable
                                          >
                                              <strong>{{ item }}</strong>&nbsp;
                                              <span>(interest)</span>
                                          </v-chip>
                                      </template>
                                  </v-combobox>
                              </v-col>
                          </v-row>
                          <v-row>
                            <v-col>
                                <v-textarea
                                    label="テキスト"
                                    v-model="state.text"
                                ></v-textarea>
                            </v-col>
                          </v-row>
                          <v-row>
                              <v-col>
                                  <v-switch
                                    label="予約投稿"
                                    v-model="state.useAsync"
                                  ></v-switch>
                              </v-col>
                              <v-col>
                                  <v-text-field type="date" v-model="state.sendDate" :disabled="!state.useAsync"></v-text-field>
                              </v-col>
                              <v-col>
                                  <v-text-field type="time" v-model="state.sendTime" :disabled="!state.useAsync"></v-text-field>
                              </v-col>
                          </v-row>
                      </v-form>
                  </v-card-text>
                  <v-card-actions>
                      <v-btn variant="elevated" color="success" :disabled="state.processing" v-on:click="sendMessages" v-if="!state.useAsync">
                          送信
                      </v-btn>
                      <v-btn variant="elevated" color="success" :disabled="state.processing" v-on:click="scheduleMessages" v-if="state.useAsync">
                          予約送信
                      </v-btn>
                      <v-progress-circular v-show="state.processing" indeterminate style="margin-left: 10px;" model-value="20"></v-progress-circular>
                  </v-card-actions>
              </v-card>
          </v-container>
      </v-main>
  </v-app>
</template>

<script setup lang="ts">
import axios from 'axios';
import camelcaseKeys from 'camelcase-keys';
import { ChatController } from "~/utils/chat";
import { SpaceController } from "~/utils/spaceController";
import { Space, DirectMessage } from "~/utils/model";
import {Location} from "vscode-languageserver-types";

interface Props {
    useAsync: boolean
    spaces: Space[]
    selectedSpaces: Space[]
    directMessages: DirectMessage[]
    selectedDirectMessages: DirectMessage[]
    text: string,
    processing: boolean,
    sendDate: string,
    sendTime: string
}

const state = reactive<Props>({
    useAsync: false,
    spaces: [],
    selectedSpaces: [],
    directMessages: [],
    selectedDirectMessages: [],
    text: "",
    processing: false,
    sendDate: "",
    sendTime: ""
});

const chatController = new ChatController();
const spaceController = new SpaceController();

const sendMessages = async () => {
    state.processing = true;
    const dmResult = await chatController.sendDirectMessages(state.selectedDirectMessages, state.text);
    dmResult.forEach((v, i) => {
      if (v.status === "fulfilled") return;

      window.alert(`${state.selectedDirectMessages[i].displayName}へのメッセージ送信に失敗しました。`)
    });

    const spaceResult = await chatController.sendMessages(state.selectedSpaces, state.text);
    spaceResult.forEach((v, i) => {
      if (v.status === "fulfilled") return;

      window.alert(`スペース「${state.selectedSpaces[i].displayName}」へのメッセージ送信に失敗しました。`)
    });

    const successCount = dmResult.length + spaceResult.length;

    state.processing = false;
    state.selectedDirectMessages = [];
    state.selectedSpaces = [];
    state.text = "";

    window.alert(`${successCount}件のメッセージ（DM: ${dmResult.length}件、スペース: ${spaceResult.length}件）を送信しました！`);
}

const scheduleMessages = async () => {
    state.processing = true;

    if (state.sendDate == "" || state.sendTime == "") {
        window.alert("送信日時は必須です！");
        state.processing = false;
        return;
    }

    const d = new Date(Date.parse(`${state.sendDate} ${state.sendTime}`));

    const spaces = state.selectedSpaces.concat(await spaceController.findDirectMessages(state.selectedDirectMessages));
    const results = await chatController.scheduleMessages(spaces, state.text, d);

    state.processing = false;
    state.selectedDirectMessages = [];
    state.selectedSpaces = [];
    state.text = "";

    window.alert(`${results.length}件のメッセージ送信を予約しました！`);
}

onMounted(() => {
    axios.get("/api/auth/verify").catch(() => {
        window.location.assign("/api/auth/login");
    });

    axios.get("/api/spaces/").then((res: { data: any; }) => {
        state.spaces = res.data;
    });

    axios.get("/api/members").then((res: { data: any; }) => {
        state.directMessages = camelcaseKeys(res.data);
    });
});
</script>