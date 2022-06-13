<template>
  <div class="iq-search-bar device-search d-flex justify-content-center mb-3 gap-5">
    <form action="#" class="searchbox">
      <a class="search-link h-100 d-flex align-items-center" href="#"><i class="ri-search-line"></i></a>
      <input v-model="search" type="text" class="text search-input" placeholder="Search group here...">
    </form>
    <button class="btn btn-primary d-block" @click="openModal">Create Group</button>
  </div>
  <NewGroupModal :modalId="'newGroupModal'" @closeModal="closeModal" />

  <div class="d-grid gap-3 d-grid-template-1fr-19">
    <div v-for="group of filteredGroups" :key="group.id" class="card">
      <GroupOverview :group="group" />
    </div>
  </div>
  <h4 v-if="filteredGroups.length === 0" class="text-center mb-0">Group not found</h4>
</template>

<script>
import GroupOverview from "./UI/GroupOverview.vue"
import NewGroupModal from "./UI/NewGroupModal.vue"
import bootstrapMin from 'bootstrap/dist/js/bootstrap.min';
import axios from "axios";
export default {
  name: 'NewsFeed',
  components: {
    GroupOverview,
    NewGroupModal
  },
  data: () => ({    
    search: "",
    modal: null,
    groups: []
  }),
  computed: {
    filteredGroups() {
      return this.groups.filter(group => {
        return group.title.toLowerCase().includes(this.search.toLowerCase())
      })
    }
  },
  async created() {
    let response = await axios.get("api/allgroups", { withCredentials: true} );

    if (response.data !== null) {
      this.groups = response.data
    }
  },
  mounted() {
    this.modal = new bootstrapMin.Modal(document.getElementById("newGroupModal"));
  },
  methods: {
    openModal() {
      this.modal.show()
    },
    closeModal() {
      this.modal.hide()
    }
  }  
}
</script>
