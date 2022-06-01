<template>
  <div v-if="postData.ok === false" >
    <h3 class="text-center">Post Not Founded!</h3>
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
                    <span class="mb-0 d-inline-block">Added New Post</span>
                  </div>
                  <p class="mb-0 text-primary">{{ postData.post.createdAt }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="mt-3">
          <p>
            {{postData.post.content}}
            <router-link v-if="!preview" to="/post/1">Read more</router-link>
          </p>
        </div>

        <div class="user-post text-center col-6 m-auto">
          <img :src="postData.post.picture" alt="post-image" class="img-fluid rounded">
        </div>

        <div class="comment-area mt-3">
          <div
            class="d-flex justify-content-between align-items-center flex-wrap"
          >
            <div class="like-block position-relative d-flex align-items-center">
              <div class="total-comment-block">
                <p class="m-0">Total comments: {{ commentsLength }} </p>
              </div>
            </div>
          </div>

          <hr />
          
          <div v-if="preview">
            <form class="comment-text d-flex align-items-center mt-3" action="javascript:void(0);">
              <input
                type="text"
                class="form-control rounded"
                placeholder="Enter Your Comment"
              />
              <div class="comment-attagement d-flex">
                <a class="link-anchor">
                  <i class="ri-camera-line me-3"></i>
                </a>
              </div>
            </form>
            <hr />
          </div>
      
          <ul class="post-comments list-inline p-0 m-0">
            <li class="mb-2" v-for="(comment,index) of postData.post.comments" :key="index">
              <PostComment />
            </li>
          </ul>

          <div v-if="postData.post.comments" class="text-center">
            <router-link v-if="!preview" to="/post/1">Read all comments</router-link>
            <button v-if="preview" type="button" class="btn btn-primary">Load more</button>
          </div>
          <div v-else>
            <h3 class="text-center">No comments yet</h3>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script>
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
  computed: {
    commentsLength() {
      if (!this.postData.comments) {
        return 0
      } 

      return this.postData.comments.length
    }
  }
};
</script>