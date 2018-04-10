<template>
<div>
  <h1>Users</h1>
  <el-table
    :data="users"
    :row-key="users.user_id"
    style="width: 100%">
    <el-table-column
      prop="id"
      label="ID"
      width="50">
    </el-table-column>
    <el-table-column
      prop="email"
      label="Email">
    </el-table-column>
    <el-table-column
      prop="isAdmin"
      label="Type">
      <template slot-scope="scope">
        <el-tag type="danger" v-if="scope.row.isAdmin">Admin</el-tag>
        <div v-else>User</div>
      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'users',
  data() {
    return {
      users: [],
    };
  },
  methods: {},
  created() {
    axios
      .get('/api/users', this.credentials)
      .then((res) => {
        this.users = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
  },
};
</script>

<style scoped>

</style>
