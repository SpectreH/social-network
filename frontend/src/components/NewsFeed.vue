<template>
  <div class="col-lg-8 row m-0 p-0">
    <CreatePost/>
    <router-link to="/post/1" style="text-decoration: none; color: inherit;">
      <PostContent/>
    </router-link>
  </div>
</template>

<script>
import CreatePost from "./UI/CreatePost.vue"
import PostContent from "./UI/PostContent.vue"
export default {
  name: 'NewsFeed',
  components: {
    CreatePost,
    PostContent
  },
  data: () => ({})
}
</script>
