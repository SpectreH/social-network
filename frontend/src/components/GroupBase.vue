<template>
  <div class="col-lg-12">
    <GroupHeader/>
  </div>

  <div class="row p-0">
    <div class="col-lg-8">
      <CreatePost groupPost/>
      <router-link to="/post/1" style="text-decoration: none; color: inherit;">
        <PostContent/>
      </router-link>
    </div>

    <div class="col-lg-4">
      <GroupAbout/>
      <GroupEventsFeed/>
    </div>
  </div>

  <div class="row p-0">
    <div class="col-lg-8">
      <h2 class="text-center"><i class="ri-lock-fill h2 me-3"></i>Group is private. Make a request to join.</h2>
    </div>

    <div class="col-lg-4">
      <GroupAbout/>
    </div>
  </div>

</template>

<script>
import GroupEventsFeed from "./UI/GroupEventsFeed.vue"
import GroupAbout from "./UI/GroupAbout.vue"
import GroupHeader from "./UI/GroupHeader.vue"
import CreatePost from "./UI/CreatePost.vue"
import PostContent from "./UI/PostContent.vue"
export default {
  name: "GroupBase",
  components: {
    GroupAbout,
    GroupHeader,
    GroupEventsFeed,
    CreatePost,
    PostContent
  }
}
</script>

<style>

</style>