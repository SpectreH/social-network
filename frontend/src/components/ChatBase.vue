<template>
  <div class="card">
    <div class="card-body chat-page p-0">
      <div class="chat-data-block">
        <div class="row">

          <div class="col-lg-3 chat-data-left scroller">
            <ChatSideBar :chats="chatOverviews"  :callback="changeChat"/>
          </div>
          
          <div class="col-lg-9 chat-data p-0 chat-data-right">
            <ChatBox :chat="chats[currentChat]" @newMessage="newMessage" />
          </div>
          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import axios from "axios";
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
      // "1":
      // {
      //   id: "1",
      //   name: "Denni Karin",
      //   type: "direct",
      //   avatar: "https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/05.jpg",
      //   messages: [
      //     {
      //       author: 1,
      //       time: "12:30",
      //       text: "Hello 🤔",
      //     },
      //     {
      //       author: 2,
      //       time: "12:30",
      //       text: "Hello 🤔",
      //     }
      //   ],
      // }
    }
  }),
  computed: {
    chatOverviews() {
      let res = [];
      for (const [key, chat] of Object.entries(this.chats)) {
        key;
        res.push({
          id: chat.id,
          name: chat.name,
          type: chat.type,
          avatar: chat.avatar,
          lastMessageContent: chat.messages.length == 0 ? "" : chat.messages[chat.messages.length - 1].text.slice(0, 10) + "..."
        })
      }

      return res
    }
  },
  async created() {
    let response = await axios.get("api/allchats", { withCredentials: true} );
    if (response.data) {
      let res = {}

      response.data.forEach(chat => {
        res[chat.id] = chat

        if (chat.messages === null) {
          res[chat.id].messages = [];
        }
      });

      this.chats = res;
    }
  },
  mounted() {
    this.currentChat = this.chatId;    
  },
  methods: {
    ...mapGetters({
      getId: 'auth/id'
    }),
    changeChat(chatId) {
      this.currentChat = chatId
    },
    newMessage(message) {
      this.chats[this.currentChat].messages.push({
        author: this.getId(),
        time: new Date(Date.now()).toISOString(),
        text: message,
      })
    }
  }
}
</script>