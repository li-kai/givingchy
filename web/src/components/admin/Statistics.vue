<template>
<div>
  <h1>Statistics</h1>
  <data-tables-server
    :data="comments"
    :total="10000"
    :row-key="comments.id"
    :load-data="fetchPage"
    :table-props="{ stripe: false, border: false }"
    :search-def="{ show: false }">
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
      slot-scope="date"
      :datetime="date.row.createdAt"></date-time>
    </el-table-column>
    <el-table-column
      label="Actions">
      <template slot-scope="actions">
        <el-button
          type="danger"
          size="small"
          @click="deleteComment(actions.row.id)">
          Delete
          </el-button>
      </template>
    </el-table-column>
  </data-tables-server>
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
      comments: [],
    };
  },
  methods: {
    fetchPage({ page, pageSize }) {
      return axios
        .get(`/api/comments?page=${page}&limit=${pageSize}`)
        .then((res) => {
          this.comments = res.data;
        })
        .catch((err) => {
          this.$notify.error({ title: 'Error', message: err.response.data.error });
        });
    },
    deleteComment(id) {
      axios
        .delete(`/api/comments/${id}`)
        .then((res) => {
          this.comments = this.comments.filter((comment) => comment.id !== id);
        })
        .catch((err) => {
          this.$notify.error({ title: 'Error', message: err.response.data.error });
        });
    },
  },
};
</script>

<style scoped>

</style>
