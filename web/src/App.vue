<template>
	<div id="app">
		<header>
			<el-menu mode="horizontal" :router="true" :default-active="$route.path">
				<el-menu-item index="/">Home</el-menu-item>
				<el-menu-item index="/projects/new" v-if="isLoggedIn">New Project</el-menu-item>
        <li class="el-menu-item auth" @click="logout" v-if="isLoggedIn">Logout</li>
        <li class="el-menu-item auth" v-if="isLoggedIn">
          {{user.username}}
          <img class="auth-image" :src="user.image" alt="profile img">
        </li>
        <el-menu-item index="/admin" class="auth" v-if="isLoggedIn && user.isAdmin">
          Admin
        </el-menu-item>
				<el-menu-item index="/signup" class="auth" v-if="!isLoggedIn">Sign Up</el-menu-item>
				<el-menu-item index="/login" class="auth" v-if="!isLoggedIn">Login</el-menu-item>
			</el-menu>
		</header>
		<main class="main">
			<router-view></router-view>
		</main>
		<footer class="footer">
			CS2102 Project - Group 23
			<div>
				Members: Chen Yinfang, Li Kai, Peng Hanqiu, Tang Yifeng
			</div>
		</footer>
	</div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'app',
  computed: {
    ...mapGetters(['isLoggedIn', 'user']),
  },
  methods: {
    logout() {
      this.$store.dispatch('logoutUser');
    },
  },
};
</script>

<style>
html,
body {
  height: 100%;
  min-height: 100%;
  margin: 0;
}

#app {
  display: flex;
  flex-direction: column;
  height: 100%;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei',
    '微软雅黑', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 14px;
  line-height: 1.45;
  color: #2c3e50;
}

h1 {
  font-size: 24px;
}

h2 {
  font-size: 20px;
  font-weight: 500;
}

.el-menu--horizontal > .el-menu-item.auth {
  float: right;
}

.el-card > .el-card__body {
  padding: 0;
  height: 100%;
}

@keyframes placeHolderShimmer {
  0% {
    background-position: 100% 0;
  }
  100% {
    background-position: -100% 0;
  }
}

.image {
  max-height: 12.5rem;
  height: 165px;
  width: 100%;
  object-fit: cover;
  line-height: 165px;
  text-align: center;
  animation: 1.25s linear placeHolderShimmer infinite;
  background: #f6f7f8;
  background: linear-gradient(80deg, #eeeeee 8%, #dddddd 18%, #eeeeee 33%);
  background-size: 250% 100%;
}
</style>

<style scoped>
.main {
  flex: 1;
  padding: 20px;
}

.auth-image {
  height: 70%;
}

.footer {
  padding: 1rem;
}
</style>
