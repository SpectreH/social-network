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
    chats: {}
  }),
  watch: {
    "$store.state.message": {
      handler(newMessage) {
        this.chats[newMessage.chatId].messages.push(newMessage)
      },
      deep: true
    }
  },
  computed: {
    chatOverviews() {
      let res = [];
      for (const [key, chat] of Object.entries(this.chats)) {
        key;
        res.push({
          id: chat.id,
          name: chat.name,
          type: chat.isGroupChat,
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
      getId: 'auth/id',
    }),
    changeChat(chatId) {
      this.currentChat = chatId;
    },
    newMessage(message) {
      this.$socket.send(JSON.stringify({
        isGroupChat: this.chats[this.currentChat].isGroupChat,
        chatId: this.currentChat,
        authorId: this.getId(),
        sub: message,
        dest: this.chats[this.currentChat].destId,
        type: "newMessage" 
      }));

      this.addMessage(this.getId(), message);
    },
    addMessage(author, message) {
      this.chats[this.currentChat].messages.push({
        author: author,
        time: new Date(Date.now()).toISOString(),
        text: message,
      });
    },
  }
}
</script>