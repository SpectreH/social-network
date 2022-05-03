<template>
  <SignUp />
</template>

<script>
import { defineComponent } from 'vue';

import SignUp from '../components/SignUp.vue';

export default defineComponent({
  name: 'SignUpView',
  components: {
    SignUp
  }
});
</script>
