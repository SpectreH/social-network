<template>
  <div class="d-flex align-items-center justify-content-between mb-3 flex-wrap">
    <div class="group-info d-flex align-items-center">
      <div class="me-3">
        <img class="rounded-circle img-fluid avatar-100" :src="group.picture" alt="">
      </div>
      <div class="info">
        <h4>{{ group.title }}</h4>
        <p class="mb-0"><i v-if="group.private" class="ri-lock-fill pe-2"></i>{{ group.private ? "Private Group" : "Public Group"}}. {{ group.totalFollowers }} followers</p>
      </div>
    </div>
    <div class="group-member d-flex align-items-center gap-2  mt-md-0 mt-2">
      <button type="submit" v-if="group.isFollowing" class="btn btn-primary mb-2">Invite</button>
      <button type="submit" v-if="!group.isFollowing && !group.isMyGroup && group.private" @click="request" class="btn btn-primary mb-2">Request to join</button> 
      <button type="submit" v-if="!group.isFollowing && !group.isMyGroup && !group.private" @click="follow" class="btn btn-primary mb-2">Follow</button> 
      <button type="submit" v-if="group.isFollowing && !group.isMyGroup" @click="unfollow" class="btn btn-primary mb-2">Unfollow</button>       
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "GroupHeader",
  props: {
    group: {type: Object}
  },
  methods: {
    async request() {
      let response = await axios.get('api/group/requesttofollow', { params: { id: this.group.id }, withCredentials: true } );
      this.parseResponse(response);
    },
    async follow() {
      let response = await axios.get('api/group/follow', { params: { id: this.group.id }, withCredentials: true } );
      this.parseResponse(response);

      if (response.data.ok === true) {
        this.$emit("follow");
      }
    },
    async unfollow() {
      let response = await axios.get('api/group/unfollow', { params: { id: this.group.id }, withCredentials: true } );
      this.parseResponse(response);

      if (response.data.ok === true) {
        this.$emit("unfollow");
      }      
    },
    parseResponse(response) {
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
    }    
  }
}
</script>

<style>

</style>