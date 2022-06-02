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

      <div v-if="group.isFollowing" class="d-flex align-items-center justify-content-between">
        <button type="submit" class="btn btn-primary" style="border-top-right-radius: 0; border-bottom-right-radius: 0" @click="sendInvite">Invite</button>
        <SelectionDropDown additionalButton :selectionAttr="selectionDD"/>
      </div>

      <button type="submit" v-if="!group.isFollowing && !group.isMyGroup && group.private && !group.invite" @click="request" class="btn btn-success">Request to join</button> 
      <button type="submit" v-if="!group.isFollowing && !group.isMyGroup && group.invite" @click="acceptInvite" class="btn btn-success">Accept Invite</button>
      <button type="submit" v-if="!group.isFollowing && !group.isMyGroup && group.invite" @click="declineInvite" class="btn btn-danger">Decline Invite</button> 
      <button type="submit" v-if="!group.isFollowing && !group.isMyGroup && !group.private && !group.invite" @click="follow" class="btn btn-primary">Follow</button> 
      <button type="submit" v-if="group.isFollowing && !group.isMyGroup" @click="unfollow" class="btn btn-warning ">Unfollow</button>       
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import axios from "axios";
import SelectionDropDown from "./SelectionDropDown.vue"
export default {
  name: "GroupHeader",
  props: {
    group: {type: Object}
  },
  data() {
    return {
      selectionDD: {
        label: "",
        elements: []
      }      
    }
  },
  async created() {
    let response = await axios.get('api/followers', { withCredentials: true } );

    if (!response.data) {
      this.selectionDD.elements = [];
      return
    }

    this.selectionDD.elements = response.data;
  },
  components: {
    SelectionDropDown
  },
  methods: {
    ...mapGetters({
      getUserId: "auth/id"
    }),

    ...mapMutations({
      removeRequestById: "REMOVE_REQUEST_BY_ID"
    }),

    async request() {
      let response = await axios.get('api/group/requesttofollow', { params: { id: this.group.id }, withCredentials: true } );
      this.parseResponse(response);

      if (response.data.ok === true) {
        this.$socket.send(JSON.stringify({
          dest: this.group.creatorId,
          groupName: this.group.title,
          groupId: this.group.id,
          type: "groupFollowRequest" 
        }))
      }  
    },
    async acceptInvite() {
      let response = await axios.get('api/groupacceptrequest', { params: { id: this.getUserId(), groupId: this.group.id }, withCredentials: true } );
      this.parseResponse(response);
      if (response.data.ok === true) {
        this.$emit("invite");
        this.$emit("follow");        
        this.removeRequestById(this.group.id);   
      }
    },
    async declineInvite() {
      let response = await axios.get('api/groupdeclinerequest', { params: { id: this.getUserId(), groupId: this.group.id }, withCredentials: true } );
      this.parseResponse(response);
      if (response.data.ok === true) {
        this.$emit("invite");
        this.removeRequestById(this.group.id);   
      }
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
    async sendInvite() {
      let selected = this.selectionDD.elements.filter(e => {
        return e.selected
      })

      if (selected.length === 0) {
        this.$toast.open({
          message: 'You must select at least one folower!',
          type: 'error',
        });
        return;
      }

      let response = await axios.post("api/group/sendinvites", { id: this.group.id, invites: JSON.stringify(selected), }, { 
        withCredentials: true,
        headers: {
          "Content-Type": "multipart/form-data"
        } 
      });

      this.parseResponse(response);

      if (response.data.ok === true) {
        let sendedInvites = JSON.parse(response.data.data)

        if (sendedInvites !== null) {
          sendedInvites.forEach(e => {
            this.$socket.send(JSON.stringify({
              dest: e.id,
              groupName: this.group.title,
              groupId: this.group.id,
              type: "inviteRequest" 
            }))         
          })
        }

        for (let i = 0; i <  this.selectionDD.elements.length; i++) {
          this.selectionDD.elements[i].selected = false;
        }
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