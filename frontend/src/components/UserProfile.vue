<template>
  <div class="col-sm-12">
    <div class="text-center" style="transform: translate(0, 60px); position: relative; z-index: 10;">
      <div class="profile-img">
        <img src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/11.png" alt="profile-img" class="avatar-130 img-fluid">
      </div>
      <div class="profile-detail">
        <h3>Bni Cyst</h3>
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
                          <p class="mb-0">690</p>
                        </li>
                        <li class="text-center ps-3">
                          <h6>Followers</h6>
                          <p class="mb-0">206</p>
                        </li>
                        <li class="text-center ps-3">
                          <h6>Following</h6>
                          <p class="mb-0">100</p>
                        </li>
                    </ul>
                  </div>
                  <div style="z-index: 100;">
                    <ul class="header-nav list-inline d-flex flex-wrap justify-end p-0 m-0">
                      <li><router-link to="/test"><i class="ri-settings-4-line"></i></router-link></li>
                    </ul>
                  </div>
              </div>
          </div>
        </div>
      </div>
    </div>

    <div class="card">
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

  <div class="col-sm-12">
    <div class="tab-content">
      <div class="tab-pane fade justify-content-center active show" id="timelineTab" role="tabpanel">
        <CreatePost/>
        <router-link to="/post/1" style="text-decoration: none; color: inherit;">
          <PostContent/>
        </router-link>
      </div>
      <div class="tab-pane fade" id="aboutTab" role="tabpanel">
        <AboutInfo/>
      </div>
      <div class="tab-pane fade" id="followersTab" role="tabpanel">
        <div class="card">
          <div class="card-body">
            <h2>Friends</h2>
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
                  <UserList :users="users"/>
                </div>
                <div class="tab-pane fade" id="followers" role="tabpanel">
                  <UserList :users="users" :toShow="'follower'"/>
                </div>
                <div class="tab-pane fade" id="following" role="tabpanel">
                  <UserList :users="users" :toShow="'following'"/>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import CreatePost from "./UI/CreatePost.vue"
import PostContent from "./UI/PostContent.vue"
import AboutInfo from "./UI/AboutInfo.vue"
import UserList from "./UI/UserList.vue"
export default {
  name: 'UserProfile',
  components: {
    CreatePost,
    PostContent,
    AboutInfo,
    UserList
  },
  data: () => ({
    users: [
      {
        fullName: "Denni Karin",
        image: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/19.jpg",
        followers: 12,
        type: "follower"
      },
      {
        fullName: "Mark Tests",
        image: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/19.jpg",
        followers: 5,
        type: "following"
      },
    ]

  })
}
</script>
