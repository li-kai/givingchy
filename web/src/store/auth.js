// initial state
const state = {
  user: JSON.parse(sessionStorage.getItem("user")) || {},
  token: sessionStorage.getItem("token") || "",
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
    sessionStorage.setItem("user", JSON.stringify(user));
    sessionStorage.setItem("token", token);
    state.user = user;
    state.token = token;
  },
  logoutUser(state) {
    sessionStorage.clear();
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
