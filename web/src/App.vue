<template>
	<el-container id="app">
		<el-header>
			<el-menu mode="horizontal" :router="true" :default-active="$route.path">
				<el-menu-item index="/">Home</el-menu-item>
				<el-menu-item index="/projects/new">New Project</el-menu-item>
        <el-menu-item index="/admin" class="auth">Admin</el-menu-item>
				<el-menu-item index="/signup" class="auth">Sign Up</el-menu-item>
				<el-menu-item index="/login" class="auth">Login</el-menu-item>
			</el-menu>
		</el-header>
		<el-main>
			<router-view></router-view>
		</el-main>
	</el-container>
</template>

<script>
export default {
  name: 'app',
};
</script>

<style>
#app {
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei',
    '微软雅黑', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 14px;
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
</style>
