<template>
  <a data-bs-toggle="pill" class="">
    <div class="d-flex align-items-center">
      <div class="avatar me-2">
        <img :src="chat.avatar" alt="chatuserimage" class="avatar-50 ">
        <span class="avatar-status"><i class="ri-checkbox-blank-circle-fill" :class="chat.status === 'online' ? 'text-success' : 'text-gray'" ></i></span>
      </div>
      <div class="chat-sidebar-name">
        <h6 class="mb-0">{{ chat.name }}</h6>
        <span>{{ chat.lastMessageContent }}</span>
      </div>
      <div class="chat-meta float-right text-center mt-2 me-1">
        <div class="chat-msg-counter bg-primary text-white" v-if="chat.newMessageCounter !== 0">{{ chat.newMessageCounter }}</div>
        <span v-if="chat.lastMessageContent !== 0" class="text-nowrap">{{ chat.lastMessageTime }}</span>
      </div>
    </div>
  </a>
</template>

<script>
export default {
  name: "ChatOverview",
  props: {
    chat: {type: Object, default: () => { return {} }}
  }
}
</script>

<style>

</style>