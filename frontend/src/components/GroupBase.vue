<template>
  <div v-if="group.id">
    <div class="col-lg-12">
      <GroupHeader @follow="follow" @unfollow="unfollow" @invite="invite" :group="group"/>
    </div>

    <div v-if="group.isFollowing"  class="row p-0">
      <div class="col-lg-8">
        <CreatePost groupPost :groupId="groupId"/>
        <router-link v-for="postData in showPosts" :key="postData.post.id" :to="'/post/' + postData.post.id" style="text-decoration: none; color: inherit;">
          <PostContent :postData="postData"/>
        </router-link>
      </div>

      <div class="col-lg-4">
        <GroupAbout :group="group"/>
        <GroupEventsFeed  :groupId="groupId" :events="showEvents"/>
      </div>
    </div>

    <div v-if="(group.private && !group.isFollowing) || !group.isFollowing" class="row p-0">
      <div class="col-lg-8 d-flex align-items-center">
        <h2 v-if="group.private" class="text-center"><i class="ri-lock-fill h2 me-3"></i>Group is private. Make a request to join.</h2>
        <h2 v-else class="text-center">Follow the group to see what is inside</h2>
      </div>

      <div class="col-lg-4">
        <GroupAbout :group="group"/>
      </div>
    </div>
  </div>
  <h2 v-if="group.ok === false" class="text-center mb-0">Group not found</h2>
</template>

<script>
import GroupEventsFeed from "./UI/GroupEventsFeed.vue"
import GroupAbout from "./UI/GroupAbout.vue"
import GroupHeader from "./UI/GroupHeader.vue"
import CreatePost from "./UI/CreatePost.vue"
import PostContent from "./UI/PostContent.vue"
import axios from "axios";
export default {
  name: "GroupBase",
  components: {
    GroupAbout,
    GroupHeader,
    GroupEventsFeed,
    CreatePost,
    PostContent
  },
  props: {
    groupId: {type: String, default: ""}
  },
  data() {
    return {
      group: {},
      postsToShow: 5
    }
  },
  watch: {
    groupId: {
      handler() {
        this.fetchGroupProfile()
      }
    }
  },
  computed: {
    showPosts() {
      if (!this.group.posts) {
        return []
      }

      return this.group.posts.slice(0, this.postsToShow) 
    },
    showEvents() {
      if (!this.group.events) {
        return []
      }

      return this.group.events
    }
  },
  async created() {
    this.fetchGroupProfile()
  },
  methods: {
    async fetchGroupProfile() {
      let response = await axios.get('api/group', { params: { id: this.groupId }, withCredentials: true } );
      this.group = response.data;
    },
    follow() {
      this.group.isFollowing = true;
    },
    unfollow() {
      this.group.isFollowing = false;
    },
    invite() {
      this.group.invite = false;
    }
  }
}
</script>

<style>

</style>