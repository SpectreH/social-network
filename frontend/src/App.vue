<template>
  <div>
    <auth-base v-if="authenticated"></auth-base>
    <un-auth-base v-else></un-auth-base>

    <div class="page-alert" style="display: none; width: fit-content; justify-content: center;" type="Error" id="page-alert"
      onclick="ClearAlertMessage(this)">
      <div class="alert" id="alert">
        <div class="alert-message" name="alert" id="alert-message">
          <h3 class="alert-text" name="alert-text" id="alert-text"></h3>
        </div>
      </div>
      <div class="page-alert-close" data-closable>
        <svg width="20px" height="20px" viewBox="0 0 24 24">
          <path
            d="M13.41 12l4.3-4.29a1 1 0 1 0-1.42-1.42L12 10.59l-4.29-4.3a1 1 0 0 0-1.42 1.42l4.3 4.29-4.3 4.29a1 1 0 0 0 0 1.42 1 1 0 0 0 1.42 0l4.29-4.3 4.29 4.3a1 1 0 0 0 1.42 0 1 1 0 0 0 0-1.42z" />
        </svg>
      </div>
    </div>

  </div>
</template>

<script>
import AuthBase from "./components/AuthBase.vue"
import UnAuthBase from "./components/UnAuthBase.vue"
import { mapGetters } from 'vuex'
export default {
  name: 'App',
  components: {
    AuthBase,
    UnAuthBase
  },
  computed: {
    ...mapGetters({
      authenticated: 'auth/authenticated',
    })
  },
  created() {
    if (this.authenticated) {
      this.$connect("ws://127.0.0.1:4000/api/socket");
    }
  }
}
</script>
