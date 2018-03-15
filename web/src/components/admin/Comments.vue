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
      width="50">
    </el-table-column>
    <el-table-column
      prop="userId"
      label="User ID"
      width="70">
    </el-table-column>
    <el-table-column
      prop="content"
      label="Content">
    </el-table-column>
    <el-table-column
      prop="createdAt"
      label="Created At">
      <date-time
      slot-scope="scope"
      :datetime="scope.row.createdAt"></date-time>
    </el-table-column>
    <el-table-column
      label="Actions">
      <template slot-scope="scope">
        <el-button type="danger" size="small">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import axios from 'axios';
import DateTime from './DateTime';

export default {
  name: 'comments',
  components: {
    DateTime,
  },
  data() {
    return {
      rawComments: [],
    };
  },
  methods: {},
  created() {
    axios
      .get('/api/comments', this.credentials)
      .then((res) => {
        this.rawComments = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
  },
  computed: {
    comments() {
      return this.rawComments.slice()
        .sort((a, b) => a.id - b.id);
    },
  },
};
</script>

<style scoped>

</style>
