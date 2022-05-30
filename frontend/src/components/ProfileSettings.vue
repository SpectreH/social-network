<template>
  <div class="col-lg-12">
    <div class="iq-edit-list-data">
      <div class="card">
        <div class="card-header d-flex justify-content-between">
          <div class="header-title">
            <h4 class="card-title">Personal Information Settings</h4>
          </div>
        </div>
        <div class="card-body">
          <form>
            <div class="form-group row align-items-center text-center">
              <div class="col-md-12">
                <div class="profile-img-edit">
                  <img
                    class="profile-pic rounded-circle"
                    :src="'http://localhost:4000/images/' + form.avatar"
                    alt="profile-pic"
                  />
                  <div class="p-image">
                    <label for="new-image" style="width: inherit; cursor: pointer;">
                      <i class="ri-pencil-line upload-button text-white"></i>
                    </label>
                    <input type="file" id="new-image" accept="image/png, image/gif, image/jpeg" ref="avatar-input" @change="updateAvatar" style="display: none;">
                  </div>
                </div>
              </div>
            </div>
            <div class="row align-items-center">

              <div class="form-group col-sm-6">
                <label class="form-label">First Name</label>
                <input
                  v-model="form.firstName"                
                  type="text"
                  class="form-control"
                  placeholder="First Name"
                  disabled/>
              </div>

              <div class="form-group col-sm-6">
                <label class="form-label">Last Name</label>
                <input        
                  v-model="form.lastName"       
                  type="text"
                  class="form-control"
                  placeholder="Last Name"
                  disabled/>
              </div>

              <div class="form-group col-sm-6">
                <label class="form-label">Date of Birth</label>
                <input           
                  v-model="form.birthDate"      
                  type="date"
                  class="form-control"
                  placeholder="Date of Birth"
                  disabled/>
              </div>

              <div class="form-group col-sm-6">
                <label class="form-label">Email</label>
                <input  
                  v-model="form.email"           
                  type="text"
                  class="form-control"
                  placeholder="Email"
                  disabled/>
              </div>

              <div class="form-group col-sm-6">
                <label class="form-label">Nickname</label>
                <input
                  minlength="4"    
                  v-model="inputs.nickname"      
                  type="text"
                  class="form-control"
                  placeholder="Nickname"/>
              </div>

              <div class="form-group">
                <label class="form-label">About me:</label>
                <textarea
                  v-model="inputs.aboutMe" 
                  class="form-control"
                  name="aboutMe"
                  rows="5"
                  style="line-height: 22px; resize: none"
                  placeholder="About me"></textarea>
              </div>

            </div>
            <button type="submit" class="btn btn-primary me-2" @click.prevent="updateData" :disabled="!submit">Update</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import axios from "axios";
export default {
  name: "ProfileSettings",
  data: () => ({
    form: {
      fistName: "",
      lastName: "",
      email: "",
      nickname: "",
      aboutMe: "",
      birthDate: "",
      avatar: "",
    },
    inputs: {
      aboutMe: "",
      nickname: ""
    },
    submit: false,
  }),
  watch: {
    inputs: {
      handler(newVal) {
        for (const [key, value] of Object.entries(newVal)) {
          if (this.form[key] != value) {
            console.log(this.form[key])
            this.submit = true;
            break;
          }

          this.submit = false;
        }
      },
      deep: true
    }
  },

  created() {
    this.form = this.getUser();
    this.inputs.aboutMe = this.form.aboutMe;
    this.inputs.nickname = this.form.nickname;
  },
  methods: {
    ...mapGetters({
      getUser: 'auth/user'
    }),
    ...mapMutations({
      updateUser: "auth/UPDATE_USER"
    }),

    async updateAvatar(e) {
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

      let response = await axios.post("api/profile/updateAvatar", {avatar: file}, { 
        headers: {
          "Content-Type": "multipart/form-data"
        },
      withCredentials: true,           
      });
      
      const data = response.data
      if (data.ok === false) {
        this.$toast.open({
          message: data.message,
          type: 'error',
        });
        return;
      } 

      this.updateUser({type: "avatar", data: response.data.data})
      this.$toast.open({
        message: "Profile picture successfully updated!",
        type: 'success',
      });
    },

    async updateData() {
      let response = await axios.post("api/profile/updateProfile", this.inputs, { 
        headers: {
          "Content-Type": "multipart/form-data"
        },
      withCredentials: true,           
      });

      if (response.data.ok === true) {
        this.$toast.open({
          message: response.data.message,
          type: 'success',
        });

        this.form.aboutMe = this.inputs.aboutMe
        this.form.nickname = this.inputs.nickname
        this.submit = false;

        this.updateUser({type: "nickname", data: this.inputs.nickname})
        this.updateUser({type: "aboutMe", data: this.inputs.aboutMe})
      }
    }
  }
};
</script>

<style>
</style>