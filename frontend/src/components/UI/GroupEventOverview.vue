<template>
  <div class="mb-3 d-flex align-items-center gap-3 link-anchor" data-bs-toggle="modal" data-bs-target="#eventModal-1">
    <img class="rounded-circle img-fluid avatar-40" src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/page-img/gi-1.jpg" alt="">
    <h6 class="mb-0">New Event</h6>
  </div>
  <GroupEventModal :modalId="'eventModal-1'" />
</template>

<script>
import GroupEventModal from "./GroupEventModal.vue"
export default {
  name: "GroupEventOverview",
  components: {
    GroupEventModal
  }
}
</script>

<style>

</style>