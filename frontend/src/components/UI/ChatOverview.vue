<template>
  <a data-bs-toggle="pill" class="">
    <div class="d-flex align-items-center">
      <div class="avatar me-2">
        <img :src="chat.avatar" alt="chatuserimage" class="avatar-50 profile-pic rounded-circle">
      </div>
      <div class="chat-sidebar-name">
        <h6 class="mb-0">{{ chat.name }}</h6>
        <span>{{ chat.lastMessageContent }}</span>
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