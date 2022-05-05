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
              <img class="profile-pic" src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/11.png" alt="profile-pic">
              <div class="p-image">
                <i class="ri-pencil-line upload-button text-white"></i>
                <input class="file-upload" type="file" accept="image/*">
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
        </div>
        <div class="modal-footer justify-content-center">
          <button type="button" class="col-4 btn btn-primary">Create</button>
        </div>
      </div>
    </div>
  </div>  
</template>

<script>
export default {
  name: "NewGroupModal",
  props: {
    modalId: {type: String, default: ""}
  },
  data: () => ({    
    newGroup: {
      title: "",
      currentShareSettings: 0,
      description: "",
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
  }),
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