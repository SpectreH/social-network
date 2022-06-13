<template>
  <div v-if="!chat.id" class="chat-start">
    <span class="iq-start-icon text-primary"><i class="ri-message-3-line"></i></span>
    <p id="chat-start" class="btn bg-white mt-3">Open chat to start conversation!</p>
  </div>
  <div v-else>

    <div class="chat-head">
      <header class="d-flex justify-content-between align-items-center bg-white pt-3  ps-3 pe-3 pb-3">
        <div class="d-flex align-items-center">
          <div class="sidebar-toggle">
            <i class="ri-menu-3-line"></i>
          </div>
          <div class="avatar chat-user-profile m-0 me-3">
            <img :src="chat.avatar" alt="avatar" class="avatar-50 profile-pic rounded-circle">
          </div>
          <h5 class="mb-0">{{chat.name}}</h5>
        </div>
      </header>
    </div>

    <div class="chat-content scroller">
      <div v-for="(message, index) of chat.messages" :key="index" :class="message.author !== getId() ? 'chat-left' : 'other-user'" class="chat d-flex mb-2">
        <ChatMessage :message="message" :avatar="message.author !== getId() ? message.profilePicture : getAvatar()" />
      </div>
    </div>

    <div class="chat-footer p-3 bg-white">
      <form class="d-flex align-items-center" action="#">
        <div class="chat-attagement d-flex">
          <span class="" id="followerButton" data-bs-toggle="dropdown" aria-expanded="false" role="button">
            <i class="las la-smile pe-3" aria-hidden="true"></i>
          </span>
          <div class="dropdown-menu dropdown-menu-start" aria-labelledby="followerButton" style="">
            <Picker :data="emojiIndex" :showSearch="false" :showCategories="false" :showPreview="false" :showSkinTones="false" @select="showEmoji" />
          </div>
        </div>
        <input type="text" v-model="inputMessage"  class="form-control me-3" placeholder="Type your message">
        <button type="submit" class="btn btn-primary d-flex align-items-center p-2" @click.prevent="submit"><i class="lab la-telegram-plane" aria-hidden="true"></i><span class="d-none d-lg-block ms-1">Send</span></button>
      </form>
    </div>
    
  </div>
</template>

<script>
import axios from "axios";
import { mapGetters } from 'vuex'
import ChatMessage from "./ChatMessage.vue";
import data from "emoji-mart-vue-fast/data/all.json";
import { Picker, EmojiIndex } from "emoji-mart-vue-fast/src";
let emojiIndex = new EmojiIndex(data);

export default {
  name: "ChatBox",
  props: {
    chat: {type: Object, default: () => {
      return {}
    }}
  },
  components: {
    Picker,
    ChatMessage
  },
  data() {
    return {
      inputMessage: "",
      emojiIndex: emojiIndex,
    };
  },
  methods: {
    ...mapGetters({
      getId: 'auth/id',
      getAvatar: "auth/avatar"
    }),
    showEmoji(emoji) {
      this.inputMessage += emoji.native;
    },
    async submit() {
      let response = await axios.post('api/chat/submit', { id: this.chat.id, message: this.inputMessage }, { 
        headers: {
          "Content-Type": "multipart/form-data"
        },
        withCredentials: true,           
      });

      if (response.data.ok === true) {
        this.$emit('newMessage', this.inputMessage);
        this.inputMessage = "";    
      }
    }
  }  
}
</script>

<style>

</style>