<template>
  <div class="col-lg-8 row m-0 p-0">
    <CreatePost/>
    <router-link v-for="postData in showPosts" :key="postData.post.id" :to="'/post/' + postData.post.id" style="text-decoration: none; color: inherit;">
      <PostContent :postData="postData"/>
    </router-link>
    <div class="col-12 text-center mt-2 mb-2">
      <button v-if="postsToShow < posts.length" type="button" class="btn btn-primary col-6"  @click="postsToShow += 5">Load more</button>    
    </div>
  </div>
</template>

<script>
import CreatePost from "./UI/CreatePost.vue"
import PostContent from "./UI/PostContent.vue"
import axios from "axios";
export default {
  name: 'NewsFeed',
  components: {
    CreatePost,
    PostContent
  },
  data: () => ({
    posts: [],
    postsToShow: 5
  }),
  computed: {
    showPosts() {
     return this.posts.slice(0, this.postsToShow) 
    }
  },
  async created() {
    let response = await axios.get("api/allposts", { withCredentials: true} );
    this.posts = response.data
  }
}
</script>
