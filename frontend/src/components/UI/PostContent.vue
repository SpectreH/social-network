<template>
  <div class="col-sm-12">
    <div class="card card-block card-stretch card-height">
      <div class="card-body">

        <div class="user-post-data">
          <div class="d-flex justify-content-between">
            <div class="me-3">
              <router-link to="/user/1">
                <img
                  src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/01.jpg"
                  class="rounded-circle img-fluid"
                  alt=""
                />
              </router-link>
            </div>
            <div class="w-100">
              <div class="d-flex justify-content-between">
                <div class="">
                  <div class="d-flex gap-2">
                    <router-link to="/user/1">
                      <h5 class="mb-0 d-inline-block">Anna Sthesia</h5>
                    </router-link>
                    <span class="mb-0 d-inline-block">Added New Post</span>
                  </div>
                  <p class="mb-0 text-primary">Just Now</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="mt-3">
          <p>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi nulla
            dolor, ornare at commodo non, feugiat non nisi. Phasellus faucibus
            mollis pharetra. Proin blandit ac massa sed rhoncus...
            <router-link v-if="!preview" to="/post/1">Read more</router-link>
          </p>
        </div>

        <div class="user-post text-center">
          <img src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/page-img/p4.jpg" alt="post-image" class="img-fluid rounded">
        </div>

        <div class="comment-area mt-3">
          <div
            class="d-flex justify-content-between align-items-center flex-wrap"
          >
            <div class="like-block position-relative d-flex align-items-center">
              <div class="total-comment-block">
                <p class="m-0">Total comments: </p>
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
            <li class="mb-2">
              <PostComment/>
            </li>
          </ul>

          <div v class="text-center">
            <router-link v-if="!preview" to="/post/1">Read all comments</router-link>
            <button v-if="preview" type="button" class="btn btn-primary">Load more</button>
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
    preview: {type: Boolean, default: false}
  },

  methods: {

  }
};
</script>