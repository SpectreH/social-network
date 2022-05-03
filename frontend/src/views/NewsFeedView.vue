<template>
  <NewsFeed />
</template>

<script>
import { defineComponent } from 'vue';

// Components
import NewsFeed from '../components/NewsFeed.vue';

export default defineComponent({
  name: 'HomeView',

  components: {
    NewsFeed,
  },
});
</script>
