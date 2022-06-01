<template>
  <div class="modal fade" :id="modalId" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true" style="overflow: inherit">
    <div class="modal-dialog">
      <div class="modal-content">
        
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">Create Post</h5>
          <button
            type="button"
            class="btn btn-secondary"
            data-bs-dismiss="modal"
          >
            <i class="ri-close-line"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="d-flex align-items-center">
            <div class="user-img">
              <img
                src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/1.jpg"
                alt="userimg"
                class="avatar-60 rounded-circle img-fluid"
              />
            </div>
            <form class="post-text ms-3 w-100" action="javascript:void();">
              <textarea
                type="text"
                id="newPostTextArea"
                class="form-control rounded"
                placeholder="Write something here..."
                style="border: none;"
                v-model="postContent"
              />
            </form>
          </div>
          
          <hr />
          <div class="image-option">
            <div class="d-flex align-items-center justify-content-between">
              <div class="col-6 d-flex align-items-center">
                <h6>Choose the image to upload</h6>
              </div>
              <div class="col-6">
                <input
                  class="form-control mb-0"
                  accept="image/png, image/gif, image/jpeg"
                  type="file"
                  @change="setImage"
                />
              </div>
            </div>
          </div>

          <div v-if="!groupPost">
            <hr />
            <div class="share-option">
              <div class="d-flex align-items-center justify-content-between">
                <div class="d-flex align-items-center">
                  <h6>Choose the privacy of the post</h6>
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
                      <a v-for="(setting, index) in shareSettings" :key="setting.type" class="dropdown-item p-3" href="#" @click="currentShareSettings = index">
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
          </div>

          <div v-if="currentShareSettings === 2">
            <hr />
            <div class="share-option">
              <div class="d-flex align-items-center justify-content-between">
                <div class="d-flex align-items-center">
                  <h6>Choose the followers to give access</h6>
                </div>
                <SelectionDropDown :selectionAttr="selectionDD"/>
              </div>
            </div>
          </div>
        </div>

        <div class="modal-footer justify-content-center">
          <button type="button" class="col-4 btn btn-primary" @click="submit">Post</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import SelectionDropDown from "./SelectionDropDown.vue"
export default {
  name: "NewPostModal",
  props: {
    groupPost: {type: Boolean, default: false},
    modalId: {type: String, default: ""},
  },
  components: {
    SelectionDropDown
  },
  data() {
    return {
      postContent: "",
      postImage: null,
      currentShareSettings: 0,
      shareSettings: [
        {
          type: "Public",
          sub: "Anyone of network",
          icon: "ri-save-line",
        },
        {
          type: "Followers",
          sub: "Only your followers",
          icon: "ri-close-circle-line",
        },
        {
          type: "Only for...",
          sub: "Shows only to some followers",
          icon: "ri-user-unfollow-line",
        },
      ],
      selectionDD: {
        label: "Followers",
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
  computed: {
    setShareSettingType() {
      return this.shareSettings[this.currentShareSettings].type
    }
  },
  mounted() {
    const textarea = document.getElementById("newPostTextArea");
    const limit = 120; // height limit

    textarea.oninput = function() {
      textarea.style.height = "";
      textarea.style.height = Math.min(textarea.scrollHeight, limit) + "px";
    };
  },
  methods: {
    async submit() {
      let form = {
        postContent: this.postContent,
        postImage: this.postImage,
        postShare: this.currentShareSettings,
        followers: JSON.stringify(this.selectionDD.elements)
      };

      if (!this.postContent) {
        this.$toast.open({
          message: 'Please fill post content!',
          type: 'error',
        });
        return;
      }

      if (this.currentShareSettings === 2) {
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

        form.followers = JSON.stringify(selected);
      }

      let response = await axios.post("api/post/new", form, { 
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
      }
    },
    setImage(e) {
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
        this.postImage = null;    
        return
      }

      if (file.size > 2048000) { // 2 MB
        this.$toast.open({
          message: 'Image size must be less than 2 MB!',
          type: 'error',
        });
        e.target.value=null;
        this.postImage = null;        
        return
      }

      this.postImage = file;
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