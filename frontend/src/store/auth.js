import axios from "axios";
export default ({
  namespaced: true,
  state: {
    token: null,
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
    async signIn({ dispatch }, credentials) {
      let response = await axios.post("api/signin", credentials, { withCredentials: true });
      dispatch('attempt', response.data)
    },

    async attempt({ commit }, token) {
      commit('SET_TOKEN', token)

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
