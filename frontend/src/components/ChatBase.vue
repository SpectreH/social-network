<template>
  <div class="card">
    <div class="card-body chat-page p-0">
      <div class="chat-data-block">
        <div class="row">

          <div class="col-lg-3 chat-data-left scroller">
            <div class="chat-search pt-3 ps-3">
              <div class="d-flex align-items-center">
                <div class="chat-profile me-3">
                  <img src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/1.jpg" alt="chat-user" class="avatar-60 ">
                </div>
                <div class="chat-caption">
                  <h5 class="mb-0">Bni Jordan</h5>
                </div>
              </div>
            </div>
            <ChatSideBar :callback="changeChat"/>
          </div>
          
          <div class="col-lg-9 chat-data p-0 chat-data-right">
            <ChatBox :chat="chats[currentChat]" />
          </div>
          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ChatSideBar from "./UI/ChatSidebar.vue"
import ChatBox from "./UI/ChatBox.vue"
export default {
  name: "ChatBase",
  components: {
    ChatSideBar,
    ChatBox
  },
  props: {
    chatId: {type: String, default: "0"}
  },
  data: () => ({
    currentChat: null,
    chats: {
      "1":
      {
        id: "1",
        name: "Denni Karin",
        status: "online",
        messages: [
          {
            id: "1",
            from: "1",
            to: "2",
            time: "12:30",
            avatar: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/05.jpg",
            text: "Hey John, I am looking for the best admin template. Could you please help me to find it out? 🤔",
          }
        ],
      }
    }
  }),
  mounted() {
    this.currentChat = this.chatId;    
  },
  methods: {
    changeChat(chatId) {
      this.currentChat = chatId
    }
  }
}
</script>