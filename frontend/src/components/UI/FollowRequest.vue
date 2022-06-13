<template>
  <div class="col-12 iq-follow-request">
    <div
      class="
        iq-sub-card iq-sub-card-big
        d-flex
        align-items-center
        justify-content-between
      "
    >
      <div class="col-6 d-flex align-items-center">
        <img :src="request.avatar" class="avatar-40 rounded" alt="" />
        <div class="ms-3">
          <h6 class="mb-0">{{ request.author }}</h6>
          <p class="mb-0">{{ request.sub }}</p>
        </div>
      </div>
      <div class="col-6 d-flex align-items-center justify-content-end">
        <a
          @click="accept"
          class="me-3 btn btn-primary rounded"
          >Confirm</a
        >
        <a
          @click="decline"
          class="me-3 btn btn-secondary rounded"
          >Decline</a
        >
      </div>
    </div>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import axios from "axios";
export default {
  name: "FollowRequest",
  props: { 
    request: {type: Object, default: () => {
      return {}
    }},
    index: {type: Number}
  },
  methods: {
    ...mapMutations({
      removeRequest: "REMOVE_REQUEST"
    }),

    async accept() {
      if (this.request.type == "followRequest") {
        var response = await axios.get('api/acceptrequest', { params: { id: this.request.authorId, type: this.request.type }, withCredentials: true } );
      } else if (this.request.type == "groupFollowRequest") {
        response = await axios.get('api/groupacceptrequest', { params: { id: this.request.authorId, groupId: this.request.groupId, type: this.request.type }, withCredentials: true } );
      } else if (this.request.type == "inviteRequest") {
        response = await axios.get('api/groupacceptrequest', { params: { id: this.request.dest, groupId: this.request.groupId, type: this.request.type }, withCredentials: true } );
      } else if (this.request.type == "newEvent") {
        response = await axios.get('api/event/accept', { params: { id: this.request.eventId }, withCredentials: true } );
      }

      this.parseResponse(response);
    },
    async decline() {
      if (this.request.type == "followRequest") {
        var response = await axios.get('api/declinerequest', { params: { id: this.request.authorId, type: this.request.type }, withCredentials: true } );
      } else if (this.request.type == "groupFollowRequest") {
        response = await axios.get('api/groupdeclinerequest', { params: { id: this.request.authorId, groupId: this.request.groupId, type: this.request.type }, withCredentials: true } );
      } else if (this.request.type == "inviteRequest") {
        response = await axios.get('api/groupdeclinerequest', { params: { id: this.request.dest, groupId: this.request.groupId, type: this.request.type }, withCredentials: true } );
      } else if (this.request.type == "newEvent") {
        response = await axios.get('api/event/decline', { params: { id: this.request.eventId }, withCredentials: true } );
      }

      this.parseResponse(response);
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
      
      this.removeRequest(this.index);      
    }
  }
}
</script>

<style>

</style>