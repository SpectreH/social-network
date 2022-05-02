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
            <li class="active">
              <a href="../app/profile.html" class="">
                <i class="las la-user"></i><span>Profile</span>
              </a>
            </li>
            <li class="">
              <a href="../app/chat.html" class="">
                <i class="lab la-rocketchat"></i><span>Chat</span>
              </a>
            </li>
            <li class="">
              <a href="../app/group.html" class="">
                <i class="las la-users"></i><span>Group</span>
              </a>
            </li>
            <li class="">
              <a href="../dashboard/friend-list.html" class="">
                <i class="las la-user-friends"></i><span>Friend List</span>
              </a>
            </li>
            <li class="">
              <a href="../dashboard/notification.html" class="">
                <i class="las la-bell"></i><span>Notification</span>
              </a>
            </li>
            <li class="">
              <a href="../dashboard/friend-request.html" class="">
                <i class="las la-anchor"></i><span>Friend Request</span>
              </a>
            </li>
          </ul>
        </nav>
        <div class="p-5"></div>
      </div>
      <div class="scrollbar-track scrollbar-track-x" style="display: none">
        <div
          class="scrollbar-thumb scrollbar-thumb-x"
          style="width: 260px; transform: translate3d(0px, 0px, 0px)"
        ></div>
      </div>
      <div class="scrollbar-track scrollbar-track-y" style="display: block">
        <div
          class="scrollbar-thumb scrollbar-thumb-y"
          style="height: 667.305px; transform: translate3d(0px, 0px, 0px)"
        ></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "SideBar",
  props: {},
};
</script>
