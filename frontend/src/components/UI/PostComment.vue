
<template>
  <div class="d-flex">
    <div class="user-img" style="min-width: 35px">
      <router-link :to="'/user/' + comment.authorId">
        <img
          :src="comment.avatar"
          alt="userimg"
          class="avatar-35 profile-pic rounded-circle img-fluid"
          width="35px"
          height="35px"
        />
      </router-link>  
    </div>
    <div class="comment-data-block ms-3">
      <router-link :to="'/user/' + comment.authorId">
        <h6 class="d-inline-flex">{{ comment.firstName }} {{ comment.lastName }}</h6>
      </router-link>
      <p class="mb-0" style="font-size: 14px">{{ comment.content }}</p>
      <img v-if="comment.picture" :src="comment.picture" alt="" style="max-width: 300px; max-height: 300px">
      <div class="d-flex flex-wrap align-items-center comment-activity">
        <span class="text-primary"> {{ timeFormat }} </span>
      </div>
    </div>
  </div>
  <hr>
</template>

<script>
import getTime from '@/main'
export default {
  name: "PostComment",
  props: {
    comment: {type: Object}
  },
  computed: {
    timeFormat() {
      return getTime(this.comment.createdAt)
    }
  }
}
</script>