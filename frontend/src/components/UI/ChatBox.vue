<template>
  <div v-if="!chat.id" class="chat-start">
    <span class="iq-start-icon text-primary"><i class="ri-message-3-line"></i></span>
    <p id="chat-start" class="btn bg-white mt-3">Open chat to start conversation!</p>
  </div>
  <div v-else>

    <div class="chat-head">
      <header class="d-flex justify-content-between align-items-center bg-white pt-3  ps-3 pe-3 pb-3">
        <div class="d-flex align-items-center">
          <div class="sidebar-toggle">
            <i class="ri-menu-3-line"></i>
          </div>
          <div class="avatar chat-user-profile m-0 me-3">
            <img src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/06.jpg" alt="avatar" class="avatar-50 ">
            <span class="avatar-status"><i class="ri-checkbox-blank-circle-fill text-success"></i></span>
          </div>
          <h5 class="mb-0">{{chat.name}}</h5>
        </div>
      </header>
    </div>

    <div class="chat-content scroller">

    </div>

    <div class="chat-footer p-3 bg-white">
      <form class="d-flex align-items-center" action="#">
        <div class="chat-attagement d-flex">
          <a href="#"><i class="far fa-smile pe-3" aria-hidden="true"></i></a>
          <a href="#"><i class="fa fa-paperclip pe-3" aria-hidden="true"></i></a>
        </div>
        <input type="text" class="form-control me-3" placeholder="Type your message">
        <button type="submit" class="btn btn-primary d-flex align-items-center p-2"><i class="far fa-paper-plane" aria-hidden="true"></i><span class="d-none d-lg-block ms-1">Send</span></button>
      </form>
    </div>
    
  </div>
</template>

<script>
export default {
  name: "ChatBox",
  props: {
    chat: {type: Object, default: () => {
      return {}
    }}
  }
}
</script>

<style>

</style>