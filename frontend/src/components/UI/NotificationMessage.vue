<template>
  <div class="d-flex align-items-center">
    <div class="">
      <img :src="notification.avatar" class="avatar-40 rounded" alt="" />
    </div>
    <div class="ms-3 w-100">
      <h6 class="mb-0">{{ notification.author }}</h6>
      <div
        class="
          d-flex
          justify-content-between
          align-items-center
        "
      >
        <p class="mb-0">{{ notification.sub }}</p>
        <small class="float-right font-size-12">{{ notification.time }}</small>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "NotificationMessage",
  props: {
    notification: {type: Object, default: () => {
      return {}
    }}
  }
}
</script>

<style>

</style>