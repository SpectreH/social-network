<template>
  <div class="chat-sidebar-channel scroller mt-4 ps-2">
    <ul class="iq-chat-ui nav flex-column nav-pills">
      <h5 class="">Group Chats</h5>
      <li v-for="chat of chats.filter(ch => ch.type)" :key="chat.id" @click="callback(chat.id)" class="link-anchor">
        <ChatOverview :chat="chat" />
      </li>
      <p v-if="chats.filter(ch => ch.type).length === 0" class="text-center m-0 px-1">You don't have any group chats</p>
      <h5 class="mt-3">Direct Messages</h5>
      <li v-for="chat of chats.filter(ch => !ch.type)" :key="chat.id" @click="callback(chat.id)" class="link-anchor">
        <ChatOverview :chat="chat" />
      </li>
      <p v-if="chats.filter(ch => !ch.type).length === 0" class="text-center m-0 px-1">You don't have any direct messages</p>
    </ul>
  </div>
</template>

<script>
import ChatOverview from "./ChatOverview.vue"
export default {
  name: "ChatSideBar",
  props: {
    callback: {type: Function, default: () => {
      return "-1"
    }},
    chats: {type: Array}    
  },
  components: {
    ChatOverview
  },
  data: () => ({})
}
</script>