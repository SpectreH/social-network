<template>
  <div class="iq-search-bar device-search d-flex justify-content-center mb-3">
    <form action="#" class="searchbox">
      <a class="search-link h-100 d-flex align-items-center" href="#"><i class="ri-search-line"></i></a>
      <input v-model="search" type="text" class="text search-input" placeholder="Search group here...">
    </form>
  </div>
  <div class="d-grid gap-3 d-grid-template-1fr-19">
    <div v-for="group of filteredGroups" :key="group.id" class="card">
      <GroupOverview :group="group" />
    </div>
  </div>
  <h4 v-if="filteredGroups.length === 0" class="text-center mb-0">Group not found</h4>
</template>

<script>
import GroupOverview from "./UI/GroupOverview.vue"
export default {
  name: 'NewsFeed',
  components: {
    GroupOverview
  },
  data: () => ({
    search: "",
    groups: [
      {
        id: "1",
        owner: "1",
        type: "public",
        title: "For Designers",
        image: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/page-img/gi-1.jpg",
        posts: 100,
        members: 100,
      },
      {
        id: "2",
        owner: "2",
        type: "public",
        title: "Test",
        image: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/page-img/gi-1.jpg",
        posts: 5,
        members: 5,
      },
    ]
  }),
  computed: {
    filteredGroups() {
      return this.groups.filter(group => {
        return group.title.toLowerCase().includes(this.search.toLowerCase())
      })
    }
  }
}
</script>
