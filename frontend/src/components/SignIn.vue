<template>
  <div class="sign-in-from">
    <h1 class="mb-0">Sign In</h1>
    <p>Enter your email address and password to access social network.</p>
    <form class="row mt-4" @submit.prevent="submit">
      <div class="col-12 form-group">
        <label class="form-label">Email</label>
        <input type="email" class="form-control mb-0" placeholder="Email" v-model="form.email" autocomplete/>
      </div>

      <div class="col-12 form-group">
        <label class="form-label">Password</label>
        <input type="password" class="form-control mb-0" placeholder="Password" v-model="form.password" autocomplete/>
      </div>

      <div class="d-inline-block w-100">
        <button type="submit" class="btn btn-primary float-end">
          Sign in
        </button>
      </div>

      <div class="sign-info">
        <span class="dark-color d-inline-block line-height-2">
          Don't have an account?
          <router-link to="/sign-up">Sign Up</router-link> 
        </span>
      </div>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex"
export default {
  name: 'SignIn',
  data: () => ({
    form: {
      email: "",
      password: ""
    }
  }),
  methods: {
    ...mapActions({
      signIn: 'auth/signIn'
    }),

    submit() {
      this.signIn(this.form)
    }
  }
}
</script>
