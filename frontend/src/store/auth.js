import axios from "axios";
export default ({
  namespaced: true,
  state: {
    user: null
  },

  getters: {
  },

  mutations: {
    SET_TOKEN(state, token) {
      state.token = token
    },
    SET_USER(state, data) {
      state.user = data
    },
  },

  actions: {
    async authMe({ commit }) {
      try {
        let response = await axios.get('api/authme', { withCredentials: true })
        commit('SET_USER', response.data)
      } catch (e) {
        commit('SET_TOKEN', null)
        commit('SET_USER', null)
      }
    }
  }
})
