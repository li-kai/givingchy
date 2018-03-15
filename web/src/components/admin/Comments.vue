<template>
<div>
  <h1>Comments</h1>
  <el-table
    :data="comments"
    :row-key="comments.id"
    style="width: 100%">
    <el-table-column
      prop="id"
      label="ID"
      width="70">
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
        <div v-else>Comment</div>
      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'comments',
  data() {
    return {
      comments: [],
    };
  },
  methods: {},
  created() {
    axios
      .get('/api/comments', this.credentials)
      .then((res) => {
        console.log(res.data[0].isAdmin);
        this.comments = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
  },
};
</script>

<style scoped>

</style>
