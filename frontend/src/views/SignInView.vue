<template>
  <SignIn />
</template>

<script>
import { defineComponent } from 'vue';

import SignIn from '../components/SignIn.vue';

export default defineComponent({
  name: 'SignInView',
  components: {
    SignIn
  }
});
</script>
