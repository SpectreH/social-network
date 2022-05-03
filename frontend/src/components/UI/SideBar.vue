<template>
  <div class="iq-sidebar sidebar-default">
    <div
      id="sidebar-scrollbar"
      data-scrollbar="true"
      tabindex="-1"
      style="overflow: hidden; outline: none"
    >
    
      <div class="scroll-content">
        <nav class="iq-sidebar-menu">
          <ul id="iq-sidebar-toggle" class="iq-menu">
            <li v-for="element in routerList" :key="element.header" class="">
              <router-link :to="element.to">
                <i :class="element.icon"></i><span>{{element.header }}</span>
              </router-link>
            </li>
          </ul>
        </nav>
        <div class="p-5"></div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "SideBar",
  props: {},
  data() {
    return {
      routerList: [
        {
          header: "News Feed",
          icon: "las la-newspaper",
          to: "/"
        },
        {
          header: "Profile",
          icon: "las la-user",
          to: "/profile"
        },
        {
          header: "Chat",
          icon: "lab la-rocketchat",
          to: "/chat"
        },
        {
          header: "Groups",
          icon: "las la-users",
          to: "/groups"
        },
        {
          header: "Friend List",
          icon: "las la-user-friends",
          to: "/friend-list"
        },
        {
          header: "Friend Request",
          icon: "las la-plus",
          to: "/friend-request"
        }                                         
      ]
    }
  },
};
</script>
