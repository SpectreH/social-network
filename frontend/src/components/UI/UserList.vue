<template>
  <div class="card-body p-0">
    <div class="row">
      <div v-for="user in users.filter(filterFunc(toShow))" :key="user.id" class="col-md-6 col-lg-6 mb-3">
        <div class="iq-friendlist-block">
          <div class="d-flex align-items-center justify-content-between">
              <div class="d-flex align-items-center">
                <a href="#">
                <img :src="user.image" alt="profile-img" class="img-fluid">
                </a>
                <div class="friend-info ms-3">
                    <h5>{{user.fullName}}</h5>
                    <p class="mb-0">{{user.followers}} followers</p>
                </div>
              </div>
              <div class="card-header-toolbar d-flex align-items-center">
                <div class="dropdown">
                  <div v-if="user.type === 'follower'">
                    <span class="dropdown-toggle btn btn-secondary me-2" id="followerButton" data-bs-toggle="dropdown" aria-expanded="false" role="button">
                      <i class="ri-check-line me-1 text-white"></i> Follower
                    </span>
                    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="followerButton" style="">
                      <a class="dropdown-item" href="#">Remove</a>
                    </div>
                  </div>
                  <div v-else>
                    <span class="dropdown-toggle btn btn-secondary me-2" id="followingButton" data-bs-toggle="dropdown" aria-expanded="false" role="button">
                      <i class="ri-check-line me-1 text-white"></i> Following
                    </span>
                    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="followingButton" style="">
                      <a class="dropdown-item" href="#">Unfollow</a>
                    </div>
                  </div>
                </div>
              </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "UserList",
  props: {
    toShow: {type: String, default: 'all'},
    users: {type: Array, default: () => { 
      return [] 
    }}
  },
  methods: {
    filterFunc(toShow) {
      return function(element) {
        return element.type === toShow || toShow === "all";
      }
    }
  }
}
</script>

<style>

</style>