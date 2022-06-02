<template>
  <a
    href="#"
    class="dropdown-toggle d-grid"
    data-bs-toggle="dropdown"
    aria-haspopup="true"
    aria-expanded="false"
    :id="menuAttr.id"
  >
    <i :class="menuAttr.icon"></i>
    <div v-if="requests.length !== 0" class="number-circle danger" style="font-size: 11px">{{ requests.length }}</div>
  </a>

  <div :class="menuAttr.id === 'follow-drop' ? 'sub-drop-large' : '' " class="sub-drop dropdown-menu" :aria-labelledby="menuAttr.id">
    <div class="card shadow-none m-0">
      <div class="card-header d-flex justify-content-between bg-primary">
        <div class="header-title bg-primary">
          <h5 class="mb-0 text-white">{{ menuAttr.title }}</h5>
        </div>
        <small class="badge bg-light text-dark">{{ requests.length }}</small>
      </div>
      <div class="card-body p-0">
        <div v-if="menuAttr.id === 'notification-drop'">
          <div v-for="(request, index) in requests" :key="index" href="#" class="iq-sub-card">
            <NotificationMessage :notification="request" />
          </div>
          <div v-if="requests.length === 0">
            <p class="text-center mb-0">You don't have any notification</p>     
          </div>
        </div>

        <div v-else-if="menuAttr.id === 'follow-drop'">
          <div v-for="(request, index) in requests" :key="index.id" class="iq-follow-request">
            <FollowRequest :request="request" :index="index"/>
          </div>
          <div v-if="requests.length === 0">
            <p class="text-center mb-0">You don't have any request</p>     
          </div>
        </div>
      </div>
    </div> 
  </div>
</template>

<script>
import FollowRequest from "./FollowRequest.vue"
import NotificationMessage from "./NotificationMessage.vue"
export default {
  name: "NavDropDown",
  props: {
    menuAttr: {type: Object, default: () => { 
      return {
        id: "",
        icon: "",
        title: ""
      } 
    }},
    requests: {type: Array, default: () => {
      return []
    }}
  },
  components: {
    FollowRequest,
    NotificationMessage
  }
}
</script>

<style>
.number-circle {
  font: 32px Arial, sans-serif;
  position: absolute;

  width: 2em;
  height: 2em;
  box-sizing: initial;
  
  background-color: red;
  border: 0.1em solid #666;
  color: white;
  text-align: center;
  border-radius: 50%;    

  line-height: 2em;
  box-sizing: content-box;
  
  margin-top: 7px;
  margin-left: 13px;
}
</style>