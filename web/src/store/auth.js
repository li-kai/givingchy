// initial state
const state = {
  user: {},
  token: '',
};

// getters
const getters = {
  isLoggedIn: (state) => !!state.token,
  user: (state) => state.user,
};

// actions
const actions = {
  loginUser({ commit }, payload) {
    commit('loginUser', payload);
  },
  logoutUser({ commit }) {
    commit('logoutUser')
  },
};

// mutations
const mutations = {
  loginUser(state, {
    token,
    ...user
  }) {
    state.user = user;
    state.token = token;
  },
  logoutUser(state) {
    state.user = {};
    state.token = '';
  }
};

export default {
  state,
  getters,
  actions,
  mutations,
};
