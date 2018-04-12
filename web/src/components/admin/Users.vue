<template>
<div>
  <h1>Users</h1>
  <data-tables-server
    :data="users"
    :total="10000"
    :row-key="users.user_id"
    :load-data="fetchPage"
    :table-props="{ stripe: false, border: false }"
    :search-def="{ show: false }">
    <el-table-column
      prop="id"
      label="ID"
      width="50">
    </el-table-column>
    <el-table-column
      prop="username"
      label="Username">
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
  </data-tables-server>
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
  methods: {
    fetchPage({ page, pageSize }) {
      return axios
        .get(`/api/users?page=${page}&limit=${pageSize}`)
        .then((res) => {
          this.users = res.data;
        })
        .catch((err) => {
          this.$message(err);
        });
    },
  },
};
</script>

<style scoped>

</style>
