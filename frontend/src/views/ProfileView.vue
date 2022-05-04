<template>
  <UserProfile/>
</template>

<script>
import { defineComponent } from 'vue';

import UserProfile from '../components/UserProfile.vue';

export default defineComponent({
  name: 'ProfileView',
  props: {
    userId: { type: String, default: "" }
  },
  components: {
    UserProfile
  }
});
</script>
