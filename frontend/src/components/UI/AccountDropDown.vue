<template>
  <a
    href="#"
    class="d-flex align-items-center dropdown-toggle"
    data-bs-toggle="dropdown"
    aria-haspopup="true"
    aria-expanded="false"
    :id="accountAttr.id"
  >
    <img :src="'http://localhost:4000/images/' + getAvatar()" class="img-fluid rounded-circle me-3" alt="user" />
    <div class="caption">
      <h6 class="mb-0 line-height">{{ accountAttr.name }}</h6>
    </div>
  </a>
  <div
    class="sub-drop dropdown-menu caption-menu"
    :aria-labelledby="accountAttr.id"
  >
    <div class="card shadow-none m-0">
      <div class="card-header bg-primary">
        <div class="header-title">
          <h5 class="mb-0 text-white">Hello, {{ accountAttr.name }}!</h5>
        </div>
      </div>
      <div class="card-body p-0">
        <router-link class="iq-sub-card iq-bg-primary-hover" :to="`/user/` + accountAttr.id">
          <div class="d-flex align-items-center">
            <div class="rounded card-icon bg-soft-primary">
              <i class="ri-file-user-line"></i>
            </div>
            <div class="ms-3">
              <h6 class="mb-0">My Profile</h6>
              <p class="mb-0 font-size-12">
                View personal profile details.
              </p>
            </div>
          </div>
        </router-link>
        
        <router-link class="iq-sub-card iq-bg-primary-hover" to="/profile-settings">
          <div class="d-flex align-items-center">
            <div class="rounded card-icon bg-soft-warning">
              <i class="ri-profile-line"></i>
            </div>
            <div class="ms-3">
              <h6 class="mb-0">Edit Profile</h6>
              <p class="mb-0 font-size-12">
                Modify your personal details.
              </p>
            </div>
          </div>
        </router-link>

        <router-link class="link-anchor iq-sub-card iq-bg-primary-hover" to="/privacy-settings">
          <div class="d-flex align-items-center">
            <div class="rounded card-icon bg-soft-danger">
              <i class="ri-lock-line"></i>
            </div>
            <div class="ms-3">
              <h6 class="mb-0">Privacy Settings</h6>
              <p class="mb-0 font-size-12">
                Control your privacy parameters.
              </p>
            </div>
          </div>
        </router-link>

        <div class="d-inline-block w-100 text-center p-3">
          <a class="btn btn-primary iq-sign-btn" role="button" @click="logout">
            Sign out<i class="ri-login-box-line ms-2"></i>
          </a>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: "AccountDropDown",
  props: {
    accountAttr: {type: Object, default: () => { 
      return {
        id: "",
        name: "",
      } 
    }}
  },
  methods: {
    ...mapGetters({
      getAvatar: 'auth/avatar',
    }),
    logout() {
      this.$router.push("/logout");
    }
  }
}
</script>