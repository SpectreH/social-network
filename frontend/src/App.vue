<template>
  <div id="app">
    <auth-base v-if="false"></auth-base>
    <un-auth-base v-else></un-auth-base>
  </div>
</template>

<script>
import AuthBase from "./components/AuthBase.vue"
import UnAuthBase from "./components/UnAuthBase.vue"
export default {
  name: 'App',
  components: {
    AuthBase,
    UnAuthBase 
  }
}
</script>
