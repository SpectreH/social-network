import { createStore } from 'vuex'
import auth from './auth/auth'
import { app } from '@/main'

export default createStore({
  state: {
    socket: {
      // Connection Status
      isConnected: false,
      // Message content
      message: "",
      // Reconnect error
      reconnectError: false,
      // Heartbeat message sending time
      heartBeatInterval: 50000,
      // Heartbeat timer
      heartBeatTimer: 0,

      requests: []
    }
  },
  mutations: {
    // Connection open
    SOCKET_ONOPEN(state, event) {
      app.config.globalProperties.$socket = event.currentTarget;
      console.log("The line is connected: " + new Date());
      state.socket.isConnected = true;
      // When the connection is successful, start sending heartbeat messages regularly to avoid being disconnected by the server
      state.socket.heartBeatTimer = setInterval(() => {
        const message = "heart_beat";
        state.socket.isConnected &&
          app.config.globalProperties.$socket.send(message);
      }, state.socket.heartBeatInterval);
    },
    // Connection closed
    SOCKET_ONCLOSE(state, event) {
      state.socket.isConnected = false;
      // Stop the heartbeat message when the connection is closed
      clearInterval(state.socket.heartBeatTimer);
      state.socket.heartBeatTimer = 0;
      console.log("The line is disconnected: " + new Date());
      console.log(event);
    },
    // An error occurred
    SOCKET_ONERROR(state, event) {
      console.error(state, event);
    },
    // Receive the message sent by the server
    SOCKET_ONMESSAGE(state, message) {
      const socketMessage = JSON.parse(message.data)

      switch (socketMessage.type) {
        case "followRequest":
          state.socket.requests.push({
            avatar: "http://localhost:4000/images/" + socketMessage.avatar,
            authorId: socketMessage.from,
            author: socketMessage.fromName,
            sub: "Wants to be your follower",
          })
          break;
      }

      state.socket.message = message;
    },
    // Auto reconnect
    SOCKET_RECONNECT(state, count) {
      console.info("Message system reconnecting...", state, count);
    },
    // Reconnect error
    SOCKET_RECONNECT_ERROR(state) {
      state.socket.reconnectError = true;
    }
  },
  getters: {
    requests(state) {
      console.log(state.socket.requests);

      return state.socket.requests;
    },
  },
  actions: {
  },
  modules: {
    auth
  }
})
