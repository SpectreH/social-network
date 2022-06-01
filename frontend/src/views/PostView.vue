<template>
  <PostContent :postData="comppost" :preview="true"/>
</template>

<script>
import { defineComponent } from 'vue';
import axios from 'axios'
import PostContent from '../components/UI/PostContent.vue';

export default defineComponent({
  name: 'PostView',
  props: {
    postId: { type: String, default: "" }
  },
  data() {
    return {
      post: {},
    }
  },
  components: {
    PostContent
  },
  computed: {
    comppost() {
      return this.post
    }
  },
  async created() {
    let response = await axios.get("api/post", { params: { id: this.postId }, withCredentials: true} );
    this.post = response.data;
  }
});
</script>
