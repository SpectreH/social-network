import axios from "axios";
export default ({
  namespaced: true,
  state: {
    token: null,
    user: null,
  },

  getters: {
    authenticated(state) {
      return state.token && state.user;
    },

    id(state) {
      return state.user.id;
    },

    user(state) {
      return state.user;
    },
    
    avatar(state) {
      return state.user.avatar;
    }
  },

  mutations: {
    SET_TOKEN(state, token) {
      state.token = token;
    },
    SET_USER(state, data) {
      state.user = data;
    },
    UPDATE_USER(state, change) {
      state.user[change.type] = change.data;
    }
  },

  actions: {
    async authMe({ commit, state }, token) {
      if (token) {
        commit('SET_TOKEN', token);
      }

      if (!state.token) {
        return;
      }

      try {
        let response = await axios.get('api/authme', { withCredentials: true })
        commit('SET_USER', response.data);
      } catch (e) {
        commit('SET_TOKEN', null);
        commit('SET_USER', null);      
      }
    },

    async signOut({ commit }) {
      await axios.get('api/logout')
      commit('SET_TOKEN', null);
    }
  }
})
