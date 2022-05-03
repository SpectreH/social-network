<template>
  <div class="col-lg-8 row m-0 p-0">
    <CreatePost/>
  </div>
</template>

<script>
import CreatePost from "./UI/CreatePost.vue"
export default {
  name: 'NewsFeed',
  components: {
    CreatePost
  },
  data: () => ({})
}
</script>
