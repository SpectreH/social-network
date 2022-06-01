<template>
  <h1 class="text-center">Logging out...</h1>
</template>

<script>
import { mapActions } from "vuex"
export default {
  async created() {
    await this.signOut();
  },
  methods: {
    ...mapActions({
      logOut: 'auth/signOut'
    }),
    async signOut() {
      await this.logOut()
      this.$disconnect();
      this.$router.push("/sign-in")
    }
  }
}
</script>