<template>
  <side-bar></side-bar>
  <top-bar></top-bar>
  <div id="content-page" class="content-page">
    <div class="container">
      <div class="row justify-content-center">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import TopBar from "./UI/TopBar.vue"
import SideBar from "./UI/SideBar.vue"

export default {
  name: 'AuthBase',
  components: {
    TopBar,
    SideBar
  },
  data: () => ({
    
  }),
  mounted() {
    const pageMenu = document.querySelector(".wrapper-menu");

    pageMenu.addEventListener('click', () => {
      pageMenu.classList.toggle("open");
      document.body.classList.toggle("sidebar-main");
    })
  }
}
</script>
