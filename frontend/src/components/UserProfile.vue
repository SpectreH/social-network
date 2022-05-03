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
                  <a class="nav-link active" href="#pills-timeline-tab" data-bs-toggle="pill" data-bs-target="#timeline" role="button">Timeline</a>
                </li>
                <li class="nav-item col-12 col-sm-3 p-0">
                  <a class="nav-link" href="#pills-about-tab" data-bs-toggle="pill" data-bs-target="#about" role="button">About</a>
                </li>
                <li class="nav-item col-12 col-sm-3 p-0">
                  <a class="nav-link" href="#pills-friends-tab" data-bs-toggle="pill" data-bs-target="#friends" role="button">Followers</a>
                </li>
                <li class="nav-item col-12 col-sm-3 p-0">
                  <a class="nav-link" href="#pills-photos-tab" data-bs-toggle="pill" data-bs-target="#photos" role="button">Groups</a>
                </li>
            </ul>
          </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UserProfile',
  components: {
  },
  data: () => ({})
}
</script>
