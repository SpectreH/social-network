<template>
  <div id="app">
    <side-bar></side-bar>
    <top-bar></top-bar>
    <div id="content-page" class="content-page">
      <div class="container">
        <div class="row">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopBar from "./components/UI/TopBar.vue"
import SideBar from "./components/UI/SideBar.vue"
export default {
  name: 'App',
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
