<template>
  <div class="modal fade" :id="modalId" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true" style="overflow: inherit">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">{{ event.title }}</h5>
          <button
            type="button"
            class="btn btn-secondary"
            data-bs-dismiss="modal"
          >
            <i class="ri-close-line"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="d-flex justify-content-center">
            <div class="profile-img-edit">
              <img class="profile-pic rounded-circle" :src="event.picture" alt="profile-pic">
            </div>
          </div>
          <hr>

          <div class="d-flex gap-2 mb-2">
            <h6 class="">Date:</h6>
            <h6 class="">{{ event.date }}</h6>
          </div>
          <hr>

          <div class="d-grid align-items-center gap-2 mb-2">
            <h6 class="">Description:</h6>
            <h6 class="">{{ event.description }}</h6>
          </div>

          <hr>          
          <div class="d-flex justify-content-between">
            <div class="d-flex gap-2">
              <h6 class="">Will participate:</h6>
              <h6 class="">{{ countAccepters }}</h6>
            </div>
            <div class="d-flex gap-2">
              <h6 class="">Refused:</h6>
              <h6 class="">{{ countDecliners }}</h6>
            </div>
          </div>
        </div>
        <div v-if="!checkParticipater" class="modal-footer justify-content-center">
          <button type="button" class="col-4 btn btn-primary" @click="accept">Will participate</button>
          <button type="button" class="col-4 btn btn-primary" @click="refuse">Refuse</button>
        </div>
        <div v-else class="modal-footer justify-content-center">
          <h4 v-if="atteningEvent" >You are participating this event!</h4>
          <h4 v-else>You aren't participating this event!</h4>
        </div> 
      </div>
    </div>
  </div>  
</template>

<script>
import { mapGetters } from 'vuex'
import axios from "axios";
export default {
  name: "GroupEventModal",
  props: {
    modalId: {type: String, default: ""},
    event: {type: Object}
  },
  data: () => ({    
    participants: []
  }),
  created() {
    if (this.event.participants !== null) {
      this.participants = this.event.participants;
    }
  },
  computed: {
    countAccepters() {
      return this.participants.filter(p => p.willAttend === true && p.eventId === this.event.id).length
    },
    countDecliners() {
      return this.participants.filter(p => p.willAttend === false && p.eventId === this.event.id).length
    },
    checkParticipater() {
      return this.participants.filter(p => p.participantId === this.getId() && p.eventId === this.event.id).length != 0
    },
    atteningEvent() {
      return this.participants.filter(p => p.participantId === this.getId() && p.willAttend === true && p.eventId === this.event.id).length != 0
    }    
  },
  methods: {
    ...mapGetters({
      getId: 'auth/id',
    }),
    async accept() {
      let response = await axios.get('api/event/accept', { params: { id: this.event.id }, withCredentials: true } );
      this.parseResponse(response);
      if (response.data.ok === true) {
        this.participants.push(
          {
            eventId: this.event.id,
            participantId: this.getId(),
            willAttend: true
          })
      }
    },
    async refuse() {
      let response = await axios.get('api/event/decline', { params: { id: this.event.id }, withCredentials: true } );
      this.parseResponse(response);
      if (response.data.ok === true) {
        this.participants.push(
          {
            eventId: this.event.id,
            participantId: this.getId(),
            willAttend: false
          })
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

<style scoped>
textarea {
  resize: none;
  overflow: hidden;
  min-height: 30px;
  max-height: 120px;
}
</style>