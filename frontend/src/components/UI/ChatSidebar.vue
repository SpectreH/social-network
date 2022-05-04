<template>
  <div class="chat-sidebar-channel scroller mt-4 ps-2">
    <ul class="iq-chat-ui nav flex-column nav-pills">
      <h5 class="">Group Chats</h5>
      <li v-for="chat of chatOverviews.filter(ch => ch.type === 'group')" :key="chat.id" @click="callback(chat.id)" class="link-anchor">
        <ChatOverview :chat="chat" />
      </li>
      <p v-if="chatOverviews.filter(ch => ch.type === 'group').length === 0" class="text-center m-0 px-1">You don't have any group chats</p>
      <h5 class="mt-3">Direct Messages</h5>
      <li v-for="chat of chatOverviews.filter(ch => ch.type === 'direct')" :key="chat.id" @click="callback(chat.id)" class="link-anchor">
        <ChatOverview :chat="chat" />
      </li>
      <p v-if="chatOverviews.filter(ch => ch.type === 'direct').length === 0" class="text-center m-0 px-1">You don't have any direct messages</p>
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
    }}    
  },
  components: {
    ChatOverview
  },
  data: () => ({
    chatOverviews: [
      {
        id: "1",
        avatar: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/07.jpg",
        status: "online",
        type: "group",
        name: "Denni Karin",
        lastMessageContent: "Test...",
        newMessageCounter: 5,
        lastMessageTime: "5 min"
      }
    ]
  })
}
</script>