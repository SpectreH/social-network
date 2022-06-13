<template>
  <div v-if="profile.id">
    <div class="col-sm-12">
      <div class="text-center" style="transform: translate(0, 60px); position: relative; z-index: 10;">
        <div>
          <img v-if="profile.isMyProfile" :src="getAvatar" alt="profile-img"  class="avatar-130 profile-pic rounded-circle">
          <img v-else :src="profile.avatar" alt="profile-img" class="avatar-130 profile-pic rounded-circle">
        </div>
        <div class="profile-detail">
          <h3> {{ profile.firstName }} {{ profile.lastName }} </h3>
        </div>
      </div>

      <div class="card">
        <div class="card-body profile-page p-0">
          <div class="profile-header">
            <div class="text-center">
                <div class="profile-info p-3 d-flex align-items-center justify-content-between position-relative">
                    <div class="social-info">
                      <ul class="social-data-block d-flex align-items-center justify-content-between list-inline p-0 m-0">
                          <li class="text-center ps-3">
                            <h6>Posts</h6>
                            <p class="mb-0">{{ profile.totalPosts }}</p>
                          </li>
                          <li class="text-center ps-3">
                            <h6>Followers</h6>
                            <p class="mb-0"> {{ profile.totalFollowers }} </p>
                          </li>
                          <li class="text-center ps-3">
                            <h6>Following</h6>
                            <p class="mb-0"> {{ profile.totalFollows }} </p>
                          </li>
                      </ul>
                    </div>
                    <div style="z-index: 100;">
                      <ul class="header-nav list-inline d-flex flex-wrap justify-end p-0 m-0">
                        <li v-if="profile.isMyProfile"><router-link to="/profile-settings"><i class="ri-settings-4-line"></i></router-link></li>
                        <a v-if="!profile.isMyProfile && !profile.private && !profile.following" @click="follow" class="me-3 btn btn-primary rounded">Follow</a>
                        <a v-if="!profile.isMyProfile && profile.following" @click="unfollow" class="me-3 btn btn-warning rounded">Unfollow</a>                        
                        <a v-if="!profile.isMyProfile && profile.private && !profile.following" @click="requestToFollow" class="me-3 btn btn-success rounded">Request To Follow</a>
                      </ul>
                    </div>
                </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="!profile.private || profile.isMyProfile || profile.following" class="card">
        <div class="card-body p-0">
          <div class="user-tabing">
            <ul class="nav nav-pills d-flex align-items-center justify-content-center profile-feed-items p-0 m-0">
              <li class="nav-item col-12 col-sm-3 p-0">
                <a class="nav-link active" href="#pills-timeline-tab" data-bs-toggle="pill" data-bs-target="#timelineTab" role="button">Timeline</a>
              </li>
              <li class="nav-item col-12 col-sm-3 p-0">
                <a class="nav-link" href="#pills-about-tab" data-bs-toggle="pill" data-bs-target="#aboutTab" role="button">About</a>
              </li>
              <li class="nav-item col-12 col-sm-3 p-0">
                <a class="nav-link" href="#pills-followers-tab" data-bs-toggle="pill" data-bs-target="#followersTab" role="button">Followers</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <div v-if="!profile.private || profile.isMyProfile || profile.following" class="col-sm-12">
      <div class="tab-content">
        <div class="tab-pane fade justify-content-center active show" id="timelineTab" role="tabpanel">
          <CreatePost v-if="profile.isMyProfile" />
          <router-link v-for="postData in showPosts" :key="postData.post.id" :to="'/post/' + postData.post.id" style="text-decoration: none; color: inherit;">
            <PostContent :postData="postData"/>
          </router-link>

          <div class="col-12 text-center mt-2 mb-2">
            <button v-if="postsToShow < profile.posts.length" type="button" class="btn btn-primary col-6"  @click="postsToShow += 5">Load more</button>    
          </div>
        </div>
        <div class="tab-pane fade" id="aboutTab" role="tabpanel">
          <AboutInfo :info="profile"/>
        </div>
        <div class="tab-pane fade" id="followersTab" role="tabpanel">
          <div class="card">
            <div class="card-body">
              <h2>Followers</h2>
              <div class="user-list-tab mt-2">
                <ul class="nav nav-pills d-flex align-items-center justify-content-left user-list-items p-0 mb-2">
                  <li>
                    <a class="nav-link active" data-bs-toggle="pill" href="#pill-all-users" data-bs-target="#all-users">All Users</a>
                  </li>
                  <li>
                    <a class="nav-link" data-bs-toggle="pill" href="#pill-followers" data-bs-target="#followers">Followers</a>
                  </li>
                  <li>
                    <a class="nav-link" data-bs-toggle="pill" href="#pill-following" data-bs-target="#following">Following</a>
                  </li>
                </ul>
                <div class="tab-content">
                  <div class="tab-pane fade active show" id="all-users" role="tabpanel">
                    <UserList :users="profile.followers" :isMe="profile.isMyProfile" />
                  </div>
                  <div class="tab-pane fade" id="followers" role="tabpanel">
                    <UserList :users="profile.followers" :toShow="'follower'" :isMe="profile.isMyProfile"/>
                  </div>
                  <div class="tab-pane fade" id="following" role="tabpanel">
                    <UserList :users="profile.followers" :toShow="'following'" :isMe="profile.isMyProfile"/>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-else>
      <h1 class="text-center mt-5">Account is Private</h1>
    </div>
  </div>

  <div v-else>
    <h1 class="text-center mt-5">Account is not found</h1>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import CreatePost from "./UI/CreatePost.vue"
import PostContent from "./UI/PostContent.vue"
import AboutInfo from "./UI/AboutInfo.vue"
import UserList from "./UI/UserList.vue"
import axios from "axios";
export default {
  name: 'UserProfile',
  components: {
    CreatePost,
    PostContent,
    AboutInfo,
    UserList
  },
  data: () => ({
    profile: {},
    postsToShow: 5
  }),
  props: {
    userId: { type: String, default: "" }
  },
  watch: {
    userId: {
      handler() {
        this.fetchUserProfile()
      }
    }
  },
  computed: {
    ...mapGetters({
      getAvatar: 'auth/avatar',
    }),     
    showPosts() {
     return this.profile.posts.slice(0, this.postsToShow) 
    }
  },
  created() {
    this.fetchUserProfile()
  },
  methods: {
    async fetchUserProfile() {
      let response = await axios.get('api/profile/fetchProfile', { params: { id: this.userId }, withCredentials: true } );
    
      if (response.data.ok === false) {
        console.log("Not found");
        return;
      }

      this.profile = response.data;
    },
    async follow() {
      let response = await axios.get('api/profile/follow', { params: { id: this.userId }, withCredentials: true } );
      this.parseResponse(response);

      if (response.data.ok === true) {
        this.profile.following = true;    
      }  
    },
    async unfollow() {
      let response = await axios.get('api/profile/unfollow', { params: { id: this.userId }, withCredentials: true } );
      this.parseResponse(response);

      if (response.data.ok === true) {
        this.profile.following = false;    
      }  
    },
    async requestToFollow() {
      let response = await axios.get('api/profile/requesttofollow', { params: { id: this.userId }, withCredentials: true } );
      this.parseResponse(response);

      if (response.data.ok === true) {
        this.$socket.send(JSON.stringify({
          dest: parseInt(this.userId),
          type: "followRequest" 
        }))
      }  
    },
    parseResponse(response) {
      if (response.data.ok === false) {
        this.$toast.open({
          message: response.data.message,
          type: 'error',
        });
        return;
      }

      this.$toast.open({
        message: response.data.message,
        type: 'success',
      });
    }
  }
}
</script>
