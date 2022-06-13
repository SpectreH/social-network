<template>
  <div class="card">
    <div class="card-header d-flex justify-content-between">
      <div class="header-title">
        <h4 class="card-title">Group events</h4>
      </div>
    </div>
    <div class="card-body">
      <ul class="list-inline p-0 m-0">
        <li v-for="(event, index) in events" :key="'event'+index">
          <GroupEventOverview :event="event"/>
        </li>
        <li>
          <button type="submit" class="btn btn-primary d-block w-100" data-bs-toggle="modal" data-bs-target="#newEventModal"><i class="ri-add-line pe-2"></i>Create New Event</button>
          <NewEventModal :modalId="'newEventModal'" :groupId="groupId"/>
        </li>
      </ul>
    </div>

  </div>
</template>

<script>
import GroupEventOverview from "./GroupEventOverview.vue"
import NewEventModal from "./NewEventModal.vue"
export default {
  name: "GroupEventsFeed",
  components: {
    GroupEventOverview,
    NewEventModal
  },
  props: {
    groupId: {type: String, default: ""},
    events: {type: Array}
  }
}
</script>

<style>

</style>