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
                                    disabled
                                  ></v-switch>
                              </v-col>
                              <v-col>
                                  <v-text-field type="date" :disabled="!state.useAsync"></v-text-field>
                              </v-col>
                              <v-col>
                                  <v-text-field type="time" :disabled="!state.useAsync"></v-text-field>
                              </v-col>
                          </v-row>
                      </v-form>
                  </v-card-text>
                  <v-card-actions>
                      <v-btn variant="elevated" color="success" v-on:click="send">送信</v-btn>
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
import { Space, DirectMessage } from "~/utils/model";

interface Props {
    useAsync: boolean
    spaces: Space[]
    selectedSpaces: Space[]
    directMessages: DirectMessage[]
    selectedDirectMessages: DirectMessage[]
    text: string
}

const state = reactive<Props>({
    useAsync: false,
    spaces: [],
    selectedSpaces: [],
    directMessages: [],
    selectedDirectMessages: [],
    text: ""
});

const chatController = new ChatController();

const send = () => {
    chatController.sendDirectMessages(state.selectedDirectMessages, state.text)
        .then((_: any) => {
            console.log("success");
        }).catch((err: Error) => {
            console.log(err)
        });

    chatController.sendMessages(state.selectedSpaces, state.text)
        .then((_: any) => {
            console.log("success");
        }).catch((err: Error) => {
            console.log(err);
        });
}

onMounted(() => {
    axios.get("/api/spaces/").then((res: { data: any; }) => {
        state.spaces = res.data;
    });

    axios.get("/api/members").then((res: { data: any; }) => {
        state.directMessages = camelcaseKeys(res.data);
    });
});
</script>