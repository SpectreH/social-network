<template>
  <div class="col-sm-12">
    <div id="post-modal-data" class="card card-block card-stretch card-height">
      <div class="card-header d-flex justify-content-between">
        <div class="header-title">
          <h4 class="card-title">Create Post</h4>
        </div>
      </div>
      <div class="card-body">
        <div class="d-flex align-items-center">
          <div class="user-img">
            <img
              :src="getAvatar"
              alt="userimg"
              class="avatar-60 profile-pic rounded-circle"
            />
          </div>
          <form
            class="post-text ms-3 w-100"
            action="javascript:void();"
          >
            <input
              type="text"
              @click="openModal"
              class="form-control rounded"
              placeholder="Write something here..."
              style="border: none"
            />
          </form>
        </div>
        <hr />
      </div>
      <NewPostModal :modalId="'postModal'" :groupPost="groupPost" :groupId="groupId" @closeModal="closeModal"/>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import bootstrapMin from 'bootstrap/dist/js/bootstrap.min';
import NewPostModal from "./NewPostModal.vue"
export default {
  name: "CreatePost",
  components: {
    NewPostModal
  },
  props: {
    groupPost: {type: Boolean, default: false},
    groupId: {type: String, default: ""}
  },
  data() {
    return {
      modal: null
    }
  },
  computed: {
    ...mapGetters({
      getAvatar: 'auth/avatar',
    }) 
  },
  mounted() {
    this.modal = new bootstrapMin.Modal(document.getElementById("postModal"));
  },
  methods: {
    openModal() {
      this.modal.show()
    },
    closeModal() {
      this.modal.hide()
    }
  }
};
</script>
