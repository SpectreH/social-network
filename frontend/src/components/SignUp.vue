<template>
  <div class="sign-in-from">
    <h1 class="mb-0">Sign Up</h1>
    <p>Fill your personal information to access social network.</p>
    <form class="row mt-4" @submit.prevent="submit">
      <div class="col-6 form-group required">
        <label for="firstName" class="form-label">First Name</label>
        <input v-model="form.firstName" id="firstName" placeholder="First Name" type="text" class="form-control mb-0" required/>
        <div class="invalid-feedback" id="firstName-feedback"></div>
      </div>

      <div class="col-6 form-group required">
        <label for="lastName" class="form-label">Last Name</label>
        <input v-model="form.lastName" id="lastName" placeholder="Last Name" type="text" class="form-control mb-0" required/>
        <div class="invalid-feedback" id="lastName-feedback"></div>
      </div>

      <div class="col-6 form-group required">
        <label for="email" class="form-label">Email</label>
        <input v-model="form.email" id="email" placeholder="Email" type="email" class="form-control mb-0" required/>
        <div class="invalid-feedback" id="email-feedback"></div>
      </div>

      <div class="col-6 form-group required">
        <label for="birthDate" class="form-label">Date of Birth</label>
        <input v-model="form.birthDate" id="birthDate" placeholder="Date of Birth" type="date" class="form-control mb-0" required/>
        <div class="invalid-feedback" id="birthDate-feedback"></div>
      </div>

      <div class="col-6 form-group">
        <label for="nickname" class="form-label">Nickname</label>
        <input v-model="form.nickname" id="nickname" placeholder="Nickname" type="text" class="form-control mb-0" minlength="4"/>
        <div class="invalid-feedback" id="nickname-feedback"></div>
      </div>

      <div class="col-6 form-group">
        <label for="avatar" class="form-label">Avatar</label>
        <input @change="setAvatar" id="avatar" placeholder="Avatar" type="file" accept="image/png, image/gif, image/jpeg" class="form-control mb-0"/>
        <div class="invalid-feedback" id="avatar-feedback"></div>
      </div>

      <div class="col-6 form-group">
        <label for="aboutMe" class="form-label">About me</label>
        <textarea v-model="form.aboutMe" type="text" id="aboutMe" placeholder="About me" class="form-control mb-0" style="resize: none; height: 60px"/>
      </div>

      <div class="col-6 form-group required">
        <label for="password" class="form-label">Password</label>
        <input v-model="form.password" id="password" placeholder="Password" type="password" class="form-control mb-0" autocomplete="false" required/>
        <div class="invalid-feedback" id="password-feedback"></div>
      </div>

      <div class="d-inline-block w-100">
        <div class="form-group form-check d-inline-block mt-2 pt-1 mb-0 required">
          <input type="checkbox" class="form-check-input" id="customCheck1" required/>
          <label class="form-label" for="customCheck1">
            I accept 
          <a href="#">Terms and Conditions</a>
          </label>
        </div>
        <button type="submit" class="btn btn-primary float-end">Sign Up</button>
      </div>

      <div class="sign-info">
        <span class="dark-color d-inline-block line-height-2">
          Already Have Account ?
          <router-link to="/sign-in">Log In</router-link> 
        </span>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "SignUp",
  data: () => ({
    validation: {
      input: null,
      message: null
    },
    form: {
      firstName: "",
      lastName: "",
      email: "",
      birthDate: null,
      nickname: "",
      avatar: null,
      aboutMe: "",
      password: "",
    }
  }),
  methods: {
    async submit() {
      try {
        let response = await axios.post("api/signup", this.form, { 
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data"
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
        
        this.$router.push("sign-in");
      } catch (e) {
        console.log("Something went wrong")
      }
    },
    setAvatar(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        return;
      }

      const file = files[0]

      const [extension] = file.type.split("/")
      if ((!(extension == "image"))) {
        this.$toast.open({
          message: 'Only images allowed to upload!',
          type: 'error',
        });
        return
      }

      if (file.size > 2048000) { // 2 MB
        this.$toast.open({
          message: 'Image size must be less than 2 MB!',
          type: 'error',
        });
        return
      }

      this.form.avatar = file;
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

};
</script>

<style scoped>
.form-group.required .form-label:after {
  content:"*";
  color:red;
}
</style>