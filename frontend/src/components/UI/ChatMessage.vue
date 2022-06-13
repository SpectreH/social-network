<template>
  <div class="chat-user">
    <a class="avatar m-0">
      <img :src="avatar" alt="avatar" class="avatar-35 profile-pic rounded-circle">
    </a>
    <span class="mt-1 d-block">{{ message.authorName }}</span>
  </div>
  <div class="chat-detail">
    <div class="chat-message">
      <p class="pb-3">{{message.text}}</p>
      <span class="chat-time mt-2" style="display: contents">{{ time }}</span>
    </div>
  </div>
</template>

<script>
import getTime from '@/main'
export default {
  name: "ChatMessage",
  props: {
    message: {type: Object, default: () => {
      return {}
    }},
    avatar: {type: String}
  },
  computed: {
    time() {
      return getTime(this.message.time)
    }
  }
}
</script>

<style>

</style>