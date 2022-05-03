<template>
  <PostContent :preview="true"/>
</template>

<script>
import { defineComponent } from 'vue';

import PostContent from '../components/UI/PostContent.vue';

export default defineComponent({
  name: 'PostView',
  props: {
    postId: { type: String, default: "" }
  },
  components: {
    PostContent
  }
});
</script>
