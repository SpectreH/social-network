<template>
  <div class="card">
    <div class="card-header d-flex justify-content-between">
        <div class="header-title">
          <h4 class="card-title">Follow Requests</h4>
        </div>
    </div>
    <div class="card-body">
        <ul class="request-list list-inline m-0 p-0">
          <li v-for="request of requests" :key="request.id" class="d-flex align-items-center flex-wrap">
            <FollowRequest :request="request" />
          </li>
          <li v-if="requests.length === 0" class="d-block text-center mb-0 pb-0">
            <p class="me-3 mb-0">You don't have any request</p>
          </li>
        </ul>
    </div>
  </div>
</template>

<script>
import FollowRequest from "./UI/FollowRequest.vue" 
export default {
  name: "RequestsFeed",
  components: {
    FollowRequest
  },
  data: () => ({
    requests: [
      {
        id: 0,
        avatar: "https://png.pngtree.com/png-vector/20191103/ourlarge/pngtree-handsome-young-guy-avatar-cartoon-style-png-image_1947775.jpg",
        author: "Jaques Amole",
        sub: "Wants to join your group",
      },
      {
        id: 1,
        avatar: "https://png.pngtree.com/png-vector/20191103/ourlarge/pngtree-handsome-young-guy-avatar-cartoon-style-png-image_1947775.jpg",
        author: "Denni Amole",
        sub: "Wants be your follower",
      },
    ]
  })
}
</script>

<style>

</style>