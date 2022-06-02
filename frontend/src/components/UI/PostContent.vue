<template>
  <div v-if="postData.ok === false" >
    <h3 class="text-center">{{ postData.message }}</h3>
  </div>

  <div v-if="postData.post" class="col-sm-12">
    <div class="card card-block card-stretch card-height">
      <div class="card-body">

        <div class="user-post-data">
          <div class="d-flex justify-content-between">
            <div class="me-3">
              <router-link to="/user/1">
                <img
                  :src="postData.post.avatar"
                  class="rounded-circle avatar-50"
                  alt=""
                />
              </router-link>
            </div>
            <div class="w-100">
              <div class="d-flex justify-content-between">
                <div class="">
                  <div class="d-flex gap-2">
                    <router-link :to="'/user/' + postData.post.authorId">
                      <h5 class="mb-0 d-inline-block">{{postData.post.firstName}} {{postData.post.lastName}}</h5>
                    </router-link>
                    <span class="mb-0 d-inline-block">Added New Post
                      <span v-if="postData.post.groupTitle">
                        In
                        <router-link :to="'/group/' + postData.post.groupId">{{ postData.post.groupTitle }}</router-link>
                      </span> 
                    </span>
                  </div>
                  <p class="mb-0 text-primary">{{ postData.post.createdAt }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="mt-3 mb-2">
          <p v-if="!preview">
            {{postData.post.content.substring(0, 150) + "....."}}
          </p>
          <p v-else>
            {{ postData.post.content }}
          </p>
        </div>

        <div v-if="postData.post.picture" class="user-post text-center col-6 m-auto">
          <img :src="postData.post.picture" alt="post-image" class="img-fluid rounded">
        </div>

        <div class="comment-area">
          <div v-if="preview">
            <form class="comment-text d-flex align-items-center mt-3" action="javascript:void(0);">
              <div class="input-group">
                <input type="text" class="form-control" v-model="comment" placeholder="Enter Your Comment" style="border-top-right-radius: 0; border-bottom-right-radius: 0">
                <div class="input-group-append">
                  <label for="new-image" class="btn btn-outline-secondary" style="border-top-left-radius: 0; border-bottom-left-radius: 0">
                    <div class="d-flex">
                      <i class="ri-camera-line"></i><div v-if="commentImage" class="mx-1">{{ commentImage.name }}</div>
                    </div>
                  </label>
                  <input
                    class="form-control mb-0"
                    id="new-image"
                    accept="image/png, image/gif, image/jpeg"
                    type="file"
                    @change="setImage"
                    style="display: none"
                  />
                  <button class="btn btn-outline-secondary mx-2" type="button" @click="submit">Submit</button>
                </div>
              </div>
            </form>
          </div>
          <hr />

          <p v-if="preview" class="m-0 mb-3">Total comments: {{ commentsLength }} </p>

          <ul class="post-comments list-inline p-0 m-0">
            <li class="mb-2" v-for="(commentary,index) of comments" :key="index">
              <PostComment :comment="commentary" />
            </li>
          </ul>

          <div v-if="comments != 0" class="text-center">
            <router-link v-if="!preview" to="/post/1">Read all comments</router-link>
            <button v-if="preview && showComments < commentsLength" type="button" class="btn btn-primary"  @click="showComments += 5">Load more</button>
          </div>
          <div v-else>
            <h5 class="text-center">No comments yet</h5>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import PostComment from "./PostComment.vue"
export default {
  name: "PostContent",
  components: {
    PostComment
  },
  props: {
    preview: {type: Boolean, default: false},
    postData: {type: Object},
  },
  data() {
    return {
      comment: "",
      showComments: this.preview ? 5 : 1,
      commentImage: null
    }
  },
  computed: {
    commentsLength() {
      if (!this.postData.comments) {
        return 0
      } 

      return this.postData.comments.length
    },
    comments() {
      if (!this.postData.comments) {
        return []
      }
      return this.postData.comments.slice(0, this.showComments);
    }
  },
  methods: {
    async submit() {
      let form = {
        commentContent: this.comment,
        commentImage: this.commentImage,
        postId: this.postData.post.id,
      };

      if (!this.comment) {
        this.$toast.open({
          message: "Comment can't be empty!",
          type: 'error',
        });
        return;
      }

      let response = await axios.post("api/post/comment", form, { 
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

        this.$router.go(0)
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
        this.commentImage = null;    
        return
      }

      if (file.size > 2048000) { // 2 MB
        this.$toast.open({
          message: 'Image size must be less than 2 MB!',
          type: 'error',
        });
        e.target.value=null;
        this.commentImage = null;        
        return
      }

      this.commentImage = file;
    }    
  }
};
</script>