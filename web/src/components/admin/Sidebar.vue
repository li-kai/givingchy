<template>
  <el-row :gutter="20">
    <el-col :span="4">
    <el-aside width="100%" class="aside">
      <el-menu
        :router="true"
        :default-active="$route.path"
        class="el-menu-vertical-demo">
        <el-menu-item index="/admin/users">
          <i class="el-icon-mobile-phone"></i>
          <span>Users</span>
        </el-menu-item>
        <el-menu-item index="/admin/projects">
          <i class="el-icon-document"></i>
          <span>Projects</span>
        </el-menu-item>
        <el-menu-item index="/admin/payments">
          <i class="el-icon-goods"></i>
          <span>Payments</span>
        </el-menu-item>
        <el-menu-item index="/admin/comments">
          <i class="el-icon-message"></i>
          <span>Comments</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    </el-col>
    <el-col :span="20">
      <router-view></router-view>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'admin',
  computed: {
  },
};
</script>

<style scoped>
.aside {
  overflow: visible;
}
</style>
