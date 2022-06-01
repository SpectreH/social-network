<template>
  <div class="card">
    <div class="card-header d-flex justify-content-between">
        <div class="header-title">
          <h4 class="card-title">Follow Requests</h4>
        </div>
    </div>
    <div class="card-body">
        <ul class="request-list list-inline m-0 p-0">
          <li v-for="(request, index) of getRequests" :key="index" class="d-flex align-items-center flex-wrap">
            <FollowRequest :request="request" />
          </li>
          <li v-if="!getRequests" class="d-block text-center mb-0 pb-0">
            <p class="me-3 mb-0">You don't have any request</p>
          </li>
        </ul>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import FollowRequest from "./UI/FollowRequest.vue" 
export default {
  name: "RequestsFeed",
  components: {
    FollowRequest
  },
  computed: {
    ...mapGetters({
      getRequests: 'requests',
    })
  },
}
</script>

<style>

</style>