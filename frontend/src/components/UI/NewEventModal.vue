<template>
  <div class="modal fade" :id="modalId" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true" style="overflow: inherit">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">Create Event</h5>
          <button
            type="button"
            class="btn btn-secondary"
            data-bs-dismiss="modal">
            <i class="ri-close-line"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="d-flex justify-content-center">
            <div class="profile-img-edit">
              <img class="profile-pic rounded-circle" :src="!groupAvatarCreated ? groupAvatarDefault : groupAvatarCreated" alt="profile-pic">
              <div class="p-image">
                <label for="new-image" style="width: inherit; cursor: pointer;">
                  <i class="ri-pencil-line upload-button text-white"></i>
                </label>
                <input type="file" id="new-image" accept="image/png, image/gif, image/jpeg" ref="avatar-input" @change="updatePicture" style="display: none;">
              </div>
            </div>
          </div>
          <hr>
          
          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Title:</h6>
            <input type="text" class="form-control w-75" v-model="newEvent.title" placeholder="Event title">
          </div>
          <hr>

          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Date:</h6>
            <input type="date" class="form-control w-75" v-model="newEvent.date" placeholder="Event date">
          </div>
          <hr>

          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Description:</h6>
            <textarea id="descriptionTextArea" type="text" class="form-control w-75" v-model="newEvent.description" placeholder="Event description" style="resize: none; height: 60px"> </textarea>
          </div>
        </div>
        <div class="modal-footer justify-content-center">
          <button type="button" class="col-4 btn btn-primary" @click="submit">Create</button>
        </div>
      </div>
    </div>
  </div>  
</template>

<script>
import axios from "axios"
export default {
  name: "NewEventModal",
  props: {
    modalId: {type: String, default: ""},
    groupId: {type: String, default: ""},
  },
  data: () => ({    
    groupAvatarDefault: "http://localhost:4000/images/default_avatar.png",
    groupAvatarCreated: "", 
    newEvent: {
      eventAvatar: null,
      title: "",
      description: "",
      date: null,
    },
  }),
  methods: {
    async submit() {
      if (!this.newEvent.title) {
        this.$toast.open({
          message: 'Please fill event title!',
          type: 'error',
        });
        return;
      }

      if (!this.newEvent.description) {
        this.$toast.open({
          message: 'Please fill event description!',
          type: 'error',
        });
        return;
      }

      this.newEvent.groupId = this.groupId;

      let response = await axios.post("api/group/newevent", this.newEvent, { 
        withCredentials: true,
        headers: {
          "Content-Type": "multipart/form-data"
        } 
      });

      if (response.data.ok === true) {
        this.$toast.open({
          message: response.data.message,
          type: 'success',
        });

        this.$socket.send(JSON.stringify({
          eventName: this.newEvent.title,
          eventDate: this.newEvent.date,
          eventId: parseInt(response.data.data),
          groupId: parseInt(this.groupId),
          type: "newEvent" 
        }))    

        this.$router.go(this.$router.currentRoute)
      }
    },
    updatePicture(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        return;
      }

      const file = files[0]

      const [extension] = file.type.split("/")
      if ((!(extension == "image"))) {
        this.$toast.open({
          message: 'Only images allowed to upload!',
          type: 'error',
        });
        e.target.value = null;
        this.newEvent.eventAvatar = null;
        this.groupAvatarCreated = "";       
        return
      }

      if (file.size > 2048000) { // 2 MB
        this.$toast.open({
          message: 'Image size must be less than 2 MB!',
          type: 'error',
        });
        e.target.value = null;
        this.newEvent.eventAvatar = null;
        this.groupAvatarCreated = "";        
        return
      }

      this.newEvent.eventAvatar = file;
      this.groupAvatarCreated = URL.createObjectURL(this.newEvent.eventAvatar);
    }
  },
  mounted() {
    const textarea = document.getElementById("descriptionTextArea");
    const limit = 120; // height limit

    textarea.oninput = function() {
      textarea.style.height = "";
      textarea.style.height = Math.min(textarea.scrollHeight, limit) + "px";
    };
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