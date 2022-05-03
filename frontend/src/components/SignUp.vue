<template>
  <div class="sign-in-from">
    <h1 class="mb-0">Sign Up</h1>
    <p>Fill your personal information to access social network.</p>
    <form class="row mt-4">
      <div v-for="input in inputs" :key="input.id" class="col-6 form-group" :class="input.required ? 'required' : _">
        <label class="form-label" :for="input.id"
          >{{ input.label }}</label
        >

        <input
          v-if="input.id !== 'aboutMe'"
          :type="input.type" 
          class="form-control mb-0"
          :id="input.id"
          :placeholder="input.label"
          :required="input.required"
          :accept="input.type === 'file' ? 'image/png, image/gif, image/jpeg' : _"
          :minlength="input.minLenght ? input.minLenght : _"
        />

        <textarea 
          v-else 
          :type="input.type" 
          class="form-control mb-0" 
          :id="input.id" 
          :placeholder="input.label" 
          :required="input.required" 
          style="resize: none; height: 60px"/>
      </div>

      <div class="d-inline-block w-100">
        <div class="form-check d-inline-block mt-2 pt-1">
          <input type="checkbox" class="form-check-input" id="customCheck1" />
          <label class="form-check-label" for="customCheck1">
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
export default {
  name: "SignUp",
  data: () => ({
    inputs: [
      {
        label: "First Name",
        id: "firstName",
        type: "text",
        required: true
      },
      {
        label: "Last Name",
        id: "lastName",
        type: "text",
        required: true
      },
      {
        label: "Email",
        id: "email",
        type: "email",
        required: true
      },
      {
        label: "Date of Birth",
        id: "birthDate",
        type: "date",
        required: true
      },
      {
        label: "Nickname",
        id: "nickname",
        type: "text",     
        minLenght: 4,   
        required: false
      },
      {
        label: "Avatar",
        id: "avatar",
        type: "file",
        required: false
      },
      {
        label: "About me",
        id: "aboutMe",
        type: "text",
        required: false 
      }, 
      {
        label: "Password",
        id: "password",
        type: "password",
        minLenght: 8,
        required: true 
      },      
    ]
  }),
};
</script>

<style scoped>
.form-group.required .form-label:after {
  content:"*";
  color:red;
}
</style>