<template>
  <div class="mb-3 d-flex align-items-center gap-3 link-anchor" data-bs-toggle="modal" :data-bs-target='"#eventModal-" + event.id'>
    <img class="rounded-circle img-fluid avatar-40" :src="event.picture">
    <h6 class="mb-0">{{ event.title }}</h6>
  </div>
  <GroupEventModal :modalId="`eventModal-${event.id}`" :event="event" />
</template>

<script>
import GroupEventModal from "./GroupEventModal.vue"
export default {
  name: "GroupEventOverview",
  components: {
    GroupEventModal
  },
  props: {
    event: {type: Object}
  }
}
</script>

<style>

</style>