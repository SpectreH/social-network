<template>
  <div class="modal fade" :id="modalId" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true" style="overflow: inherit">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">Create Group</h5>
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
            <input type="text" class="form-control w-75" v-model="newGroup.title" placeholder="Group title">
          </div>
          <hr>

          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Description:</h6>
            <textarea id="descriptionTextArea" type="text" class="form-control w-75" v-model="newGroup.description" placeholder="Group description" style="resize: none; height: 60px"> </textarea>
          </div>
          <hr>

          <div class="share-option">
            <div class="d-flex align-items-center justify-content-between">
              <div class="d-flex align-items-center">
                <h6>Choose the privacy of the group</h6>
              </div>
              <div class="card-post-toolbar">
                <div class="dropdown">
                  <span
                    class="dropdown-toggle"
                    data-bs-toggle="dropdown"
                    aria-haspopup="true"
                    aria-expanded="false"
                    role="button"
                  >
                    <span class="btn btn-primary">{{setShareSettingType}}</span>
                  </span>
                  <div class="dropdown-menu mt-2 m-0 p-0">
                    <a v-for="(setting, index) in shareSettings" :key="setting.type" class="dropdown-item p-3" href="#" @click="newGroup.currentShareSettings = index">
                      <div class="d-flex align-items-top">
                        <i class="h4" :class="setting.icon"></i>
                        <div class="data ms-2">
                          <h6>{{setting.type}}</h6>
                          <p class="mb-0">{{setting.sub}}</p>
                        </div>
                      </div>
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <hr>

          <div class="share-option">
            <div class="d-flex align-items-center justify-content-between">
              <div class="d-flex align-items-center">
                <h6>Choose the followers to invite</h6>
              </div>
              <SelectionDropDown :selectionAttr="selectionDD"/>
            </div>
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
import SelectionDropDown from "./SelectionDropDown.vue"
export default {
  name: "NewGroupModal",
  props: {
    modalId: {type: String, default: ""}
  },
  components: {
    SelectionDropDown
  },
  data: () => ({   
    groupAvatarDefault: "http://localhost:4000/images/default_avatar.png",
    groupAvatarCreated: "", 
    newGroup: {
      groupAvatar: null,
      title: "",
      currentShareSettings: 0,
      description: "",
      invites: [],
    },
    shareSettings: [
      {
        type: "Public",
        sub: "Anyone of network can join",
        icon: "ri-save-line",
      },
      {
        type: "Private",
        sub: "Join only with request or invite",
        icon: "ri-lock-line",
      }
    ],
    selectionDD: {
      label: "Followers",
      elements: []
    }
  }),
  async created() {
    let response = await axios.get('api/followers', { withCredentials: true } );

    if (!response.data) {
      this.selectionDD.elements = [];
      return
    }

    this.selectionDD.elements = response.data;
  },
  computed: {
    setShareSettingType() {
      return this.shareSettings[this.newGroup.currentShareSettings].type
    }
  },
  mounted() {
    const textarea = document.getElementById("descriptionTextArea");
    const limit = 120; // height limit

    textarea.oninput = function() {
      textarea.style.height = "";
      textarea.style.height = Math.min(textarea.scrollHeight, limit) + "px";
    };
  },
  methods: {
    async submit() {
      if (!this.newGroup.title) {
        this.$toast.open({
          message: 'Please fill group title!',
          type: 'error',
        });
        return;
      }

      if (!this.newGroup.description) {
        this.$toast.open({
          message: 'Please fill group description!',
          type: 'error',
        });
        return;
      }

      let invites = this.selectionDD.elements.filter(e => {
        return e.selected
      })

      this.newGroup.invites = JSON.stringify(invites);

      let response = await axios.post("api/group/new", this.newGroup, { 
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

        invites.forEach(e => {
          this.$socket.send(JSON.stringify({
            dest: e.id,
            groupName: this.newGroup.title,
            groupId: parseInt(response.data.data),
            type: "inviteRequest" 
          }))         
        })

        this.$emit('closeModal')
        this.$router.push("/group/" + response.data.data)
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
        this.newGroup.groupAvatar = null;
        this.groupAvatarCreated = "";       
        return
      }

      if (file.size > 2048000) { // 2 MB
        this.$toast.open({
          message: 'Image size must be less than 2 MB!',
          type: 'error',
        });
        e.target.value = null;
        this.newGroup.groupAvatar = null;
        this.groupAvatarCreated = "";        
        return
      }

      this.newGroup.groupAvatar = file;
      this.groupAvatarCreated = URL.createObjectURL(this.newGroup.groupAvatar);
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