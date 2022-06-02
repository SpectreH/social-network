<template>
  <div class="iq-top-navbar">
    <div class="iq-navbar-custom">
      <nav class="navbar navbar-expand-lg navbar-light p-0">

        <div class="iq-navbar-logo d-flex justify-content-between">
					<router-link to="/">
						<a>
							<span>Social Network</span>
						</a>
					</router-link>
          <div class="iq-menu-bt align-self-center">
            <div class="wrapper-menu">
              <div class="main-circle">
								<i class="ri-menu-line"></i>
							</div>
            </div>
          </div>
        </div>

        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-label="Toggle navigation"
        >
          <i class="ri-menu-3-line"></i>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav ms-auto navbar-list gap-1">
            <li>
              <router-link to="/">
                <a to="/" class="d-flex align-items-center">
                  <i class="ri-home-line"></i>
                </a>
              </router-link>
            </li>

            <li class="nav-item dropdown">
              <nav-drop-down :menuAttr="requestDD" :requests="getRequests"></nav-drop-down>
            </li>

            <li class="nav-item dropdown">
              <nav-drop-down :menuAttr="notificaitonDD" ></nav-drop-down>
            </li>

            <li class="nav-item dropdown">
              <account-drop-down :accountAttr="accountDD"></account-drop-down>
            </li>
          </ul>
        </div>
      </nav>
    </div>
  </div>
</template>

<script>
import NavDropDown from "./NavDropDown.vue"
import AccountDropDown from "./AccountDropDown.vue"
import { mapGetters } from 'vuex'
export default {
  components: {
    NavDropDown,
    AccountDropDown 
  },
  name: "TopBar",
  data() {
    return {
      notificaitonDD: {
        id: "notification-drop",
        icon: "ri-notification-4-line",
        title: "All Notifications",
        elements: [
          {
            avatar: "https://png.pngtree.com/png-vector/20191103/ourlarge/pngtree-handsome-young-guy-avatar-cartoon-style-png-image_1947775.jpg",
            authorId: 1,
            author: "Test",
            sub: "New message",
            time: "28.03.12 12:23",
          }
        ],
      },
      requestDD: {
        id: "follow-drop",
        icon: "ri-group-line",
        title: "Requests",
      },
      accountDD: {
        id: "",
        name: ""
      }
    }
  },
  computed: {
    ...mapGetters({
      getRequests: 'requests',
    })
  },
  created() {
    const data = this.getUserData();

    this.accountDD.id =  data.id
    this.accountDD.name = `${data.firstName} ${data.lastName}` 
  },
  methods: {
    ...mapGetters({
      getUser: 'auth/user',
    }),
    getUserData() {
      return this.getUser()
    }
  }
};
</script>

<style>
</style>