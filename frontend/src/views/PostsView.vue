<template>
  <div class="posts">
    <h1>Posts</h1>
    <post-list :posts="posts"></post-list>
  </div>
</template>
<script>
  import PostList from "../components/PostList.vue";
  export default {  
    name: "PostsView",  
    components: { PostList },
    data: function() {
      return {
        posts: [
          {
            id: 1,
            user: "mycurrentuser",
            date: "2021-01-19 15:30:30",
            post:"Today I'm feeling sooooo well."
          },
          {
            id: 2,
            user: "cshannon",
            date: "2021-01-19 15:25:20",
            post: "I think that I love giraffes"
          }
        ]
      };
    }
  };
</script>