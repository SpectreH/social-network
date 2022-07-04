<template>
  <div class="sign-in-from">
    <h1 class="mb-0">Sign In</h1>
    <p>Enter your email address and password to access social network.</p>
    <form class="row mt-4" @submit.prevent="submit">
      <div class="col-12 form-group">
        <label class="form-label">Email</label>
        <input type="email" class="form-control mb-0" id="email" placeholder="Email" v-model="form.email" autocomplete/>
        <div class="invalid-feedback" id="email-feedback"></div>
      </div>

      <div class="col-12 form-group">
        <label class="form-label">Password</label>
        <input type="password" class="form-control mb-0" id="password" placeholder="Password" v-model="form.password" autocomplete/>
        <div class="invalid-feedback" id="password-feedback"></div>
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
import axios from 'axios'
import { mapActions } from "vuex"
export default {
  name: 'SignIn',
  data: () => ({
    validation: {
      input: null,
      message: null
    },    
    form: {
      email: "",
      password: ""
    }
  }),
  methods: {
    ...mapActions({
      authMe: 'auth/authMe'
    }),

    async submit() {
      try {
        let response = await axios.post("api/signin", this.form, {
          withCredentials: true,
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          }         
        });

        const data = response.data
        if (data.ok === false) {
          this.validation = {
            input: data.input,
            message: data.message
          }
          return;
        } 
        
        await this.authMe(data.token); 
        
        this.$connect("ws://localhost:4000/api/socket");
        this.$router.push("/")
      } catch (e) {
        console.log("Something went wrong")
      }
    }
  },
  watch: {
    validation: {
      handler(val, oldVal){
        if (val.input) {
          console.log(val.input)
          document.querySelector(`#${val.input}`).classList.toggle("is-invalid");
          document.querySelector(`#${val.input}-feedback`).textContent = val.message;
        }

        if (oldVal.input) {
          document.querySelector(`#${oldVal.input}`).classList.toggle("is-invalid");
        }
      },
      deep: true
    }
  }
}
</script>
