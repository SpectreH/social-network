<template>
  <div class="sign-in-from">
    <h1 class="mb-0">Sign In</h1>
    <p>Enter your email address and password to access social network.</p>
    <form class="row mt-4">
      <div v-for="input in inputs" :key="input.id" class="col-12 form-group" :class="input.required ? 'required' : _">
        <label class="form-label" :for="input.id"
          >{{ input.label }}</label
        >

        <input
          :type="input.type" 
          class="form-control mb-0"
          :id="input.id"
          :placeholder="input.label"
          :required="input.required"
        />
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

export default {
  name: 'SignIn',
  data: () => ({
    inputs: [
      {
        label: "Email",
        id: "email",
        type: "email",
        required: true
      },
      {
        label: "Password",
        id: "password",
        type: "password",
        required: true
      }
    ]

  }),
}
</script>
