<template>
<div>
  <h1>Projects</h1>
  <data-tables-server
    :data="projects"
    :total="10000"
    :row-key="projects.id"
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
      prop="title"
      label="Title">
    </el-table-column>
    <el-table-column
      prop="description"
      label="Description">
    </el-table-column>
    <el-table-column
      prop="startTime"
      label="From">
      <date-time slot-scope="scope" :datetime="scope.row.startTime"></date-time>
    </el-table-column>
    <el-table-column
      prop="endTime"
      label="Until">
      <date-time slot-scope="scope" :datetime="scope.row.endTime"></date-time>
    </el-table-column>
    <el-table-column
      prop="category"
      label="Category">
    </el-table-column>
    <el-table-column
      prop="verified"
      label="Verified">
      <template slot-scope="scope">
        <el-button
          plain
          type="success"
          size="small"
          @click="verifyProject(scope.row)"
          v-if="scope.row.verified">
          Verified
        </el-button>
        <el-button
          plain
          type="danger"
          size="small"
          @click="verifyProject(scope.row)"
          v-else>
          Unverified
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
  name: 'projects',
  components: {
    DateTime,
  },
  data() {
    return {
      rawProjects: [],
    };
  },
  methods: {
    fetchPage({ page, pageSize }) {
      return axios
        .get(`/api/projects?page=${page}&limit=${pageSize}`)
        .then((res) => {
          this.rawProjects = res.data;
        })
        .catch((err) => {
          this.this.$notify.error({ title: 'Error', message: err.response.data.error });
          err;
        });
    },
    verifyProject(proj) {
      /* eslint-disable no-unused-vars */
      const {
        id,
        tags,
        startTime,
        amountRequiredString,
        amountRaised,
        amountRaisedString,
        ...relevantFields
      } = proj;
      /* eslint-enable no-unused-vars */
      const payload = {
        ...relevantFields,
        verified: !proj.verified,
      };
      axios
        .put(`/api/projects/${proj.id}`, payload)
        .then(() => {
          this.rawProjects = this.rawProjects.map(function(item) {
            return item.id === id ? { ...item, verified: !proj.verified } : item;
          });
        })
        .catch((err) => {
          this.$notify.error({ title: 'Error', message: err.response.data.error });
        });
    },
  },
  computed: {
    projects() {
      return this.rawProjects.map((proj) => {
        return {
          ...proj,
          amountRequiredString: `$${proj.amountRequired}`,
          amountRaisedString: `$${proj.amountRaised}`,
        };
      });
    },
  },
};
</script>

<style scoped>
.time {
  display: flex;
  flex-direction: column;
}
</style>
