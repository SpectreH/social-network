<template>
  <div class="card-body p-0">
    <div v-if="users" class="row">
      <div v-for="user in mutableUsers.filter(filterFunc(toShow))" :key="user.id" class="col-md-6 col-lg-6 mb-3">
        <div class="iq-friendlist-block">
          <div class="d-flex align-items-center justify-content-between">
              <div class="d-flex align-items-center">
                <router-link :to="'/user/' + user.id">
                  <img :src="user.avatar" alt="profile-img" class="avatar-130 img-fluid">
                </router-link>
                <div class="friend-info ms-3">
                    <h5>{{user.firstName}} {{user.lastName}}</h5>
                    <p class="mb-0">{{user.totalFollowers}} followers</p>
                </div>
              </div>
              <div class="card-header-toolbar d-flex align-items-center">
                <div class="dropdown" v-if="isMe">
                  <div v-if="user.type === 'follower'">
                    <span class="dropdown-toggle btn btn-secondary me-2" id="followerButton" data-bs-toggle="dropdown" aria-expanded="false" role="button">
                      <i class="ri-check-line me-1 text-white"></i> Follower
                    </span>
                    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="followerButton" style="">
                      <a class="dropdown-item btn" @click="remove(user.id)">Remove</a>
                    </div>
                  </div>
                  <div v-else>
                    <span class="dropdown-toggle btn btn-secondary me-2" id="followingButton" data-bs-toggle="dropdown" aria-expanded="false" role="button">
                      <i class="ri-check-line me-1 text-white"></i> Following
                    </span>
                    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="followingButton" style="">
                      <a class="dropdown-item btn" @click="unfollow(user.id)">Unfollow</a>
                    </div>
                  </div>
                </div>
                <div v-else>
                  <div class="me-3" v-if="user.type === 'follower'">
                    Follower
                  </div>
                  <div class="me-3" v-else>
                    Following
                  </div>
                </div>
              </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      This user don't have any followers or follows
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "UserList",
  props: {
    toShow: {type: String, default: 'all'},
    users: {type: Array, default: () => { 
      return [] 
    }},
    isMe: {type: Boolean}
  },
  data() {
    return {
      mutableUsers: this.users
    }
  },
  methods: {
    filterFunc(toShow) {
      return function(element) {
        return element.type === toShow || toShow === "all";
      }
    },
    async remove(id) {
      let response = await axios.get('api/profile/removefollow', { params: { id: id }, withCredentials: true } );
      this.parseResponse(response, id);
    },
    async unfollow(id) {
      let response = await axios.get('api/profile/unfollow', { params: { id: id }, withCredentials: true } );
      this.parseResponse(response, id);
    },
    parseResponse(response, id) {
      if (response.data.ok === false) {
        this.$toast.open({
          message: response.data.message,
          type: 'error',
        });
        return;
      }

      this.$toast.open({
        message: response.data.message,
        type: 'success',
      });
      this.updateUsers(id);
    },     
    updateUsers(id) {
      for (let i = 0; i < this.mutableUsers.length; i++) {
        if (this.mutableUsers[i].id === id) {
          this.mutableUsers.splice(i, 1);
          break;
        }
      }
    }
  }
}
</script>

<style>

</style>