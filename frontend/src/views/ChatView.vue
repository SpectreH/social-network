<template>
  <ChatBase :chatId="chatId" />
</template>

<script>
import ChatBase from "../components/ChatBase.vue"
export default {
  name: "ChatView",
  props: {
    chatId: { type: String, default: "0" }
  },
  components: {
    ChatBase
  }
}
</script>