<template>
  <div class="chat-user">
    <a class="avatar m-0">
    <img :src="message.avatar" alt="avatar" class="avatar-35 ">
    </a>
    <span class="chat-time mt-1">{{message.time}}</span>
  </div>
  <div class="chat-detail">
    <div class="chat-message">
        <p>{{message.text}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "ChatMessage",
  props: {
    message: {type: Object, default: () => {
      return {}
    }}
  }
}
</script>

<style>

</style>