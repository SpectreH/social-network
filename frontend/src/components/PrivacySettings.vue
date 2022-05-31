<template>
  <div class="col-lg-12">
    <div class="card">
      <div class="card-header d-flex justify-content-between">
        <div class="header-title">
          <h4 class="card-title">Privacy Setting</h4>
        </div>
      </div>
      <div class="card-body">
        <div class="acc-privacy">
          <div class="data-privacy">
            <h4 class="mb-2">Account Privacy</h4>
            <div class="form-check form-check-inline">
              <input
                v-model="accountPrivate"
                type="checkbox"
                class="form-check-input"
                id="acc-private"/>
              <label
                class="form-check-label privacy-status mb-2"
                for="acc-private"
                >Private Account</label>
            </div>
            <p>
              Lorem Ipsum is simply dummy text of the printing and typesetting
              industry. Lorem Ipsum has been the industry's standard dummy text
              ever since the 1500s, when an unknown printer took a galley of
              type and scrambled it to make a type specimen book
            </p>
          </div>
          <button type="submit" class="btn btn-primary me-2" @click.prevent="updatePrivate" :disabled="submit">Update</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import axios from "axios";
export default {
  data() {
    return {
      accountPrivate: false,
      submit: true
    }
  },

  name: "PrivacySettings",
  created() {
    const user = this.getUser();
    this.accountPrivate = user.private;
    this.submit = true;
  },
  watch: {
    accountPrivate: {
      handler() {
        this.submit = !this.submit;
      }
    }
  },
  methods: {
    ...mapGetters({
      getUser: 'auth/user'
    }),
    ...mapMutations({
      updateUser: "auth/UPDATE_USER"
    }),
    async updatePrivate() {
      let response = await axios.post("api/settings/updatePrivacy", { private: this.accountPrivate }, { 
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

        this.updateUser({type: "private", data: this.accountPrivate});
        this.submit = !this.submit;
      }
    }
  }
};
</script>

<style>
</style>