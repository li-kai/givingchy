<template>
<div>
  <h1>Projects</h1>
  <el-table
    :data="projects"
    :row-key="projects.id"
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
        <el-button plain type="success" size="small"  v-if="scope.row.verified">Verified</el-button>
        <el-button plain type="danger" size="small" v-else>Unverified</el-button>
      </template>
    </el-table-column>
  </el-table>
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
  methods: {},
  created() {
    axios
      .get('/api/projects', this.credentials)
      .then((res) => {
        this.rawProjects = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
  },
  computed: {
    projects() {
      return this.rawProjects
        .map((proj) => {
          return {
            ...proj,
            amountRequiredString: `$${proj.amountRequired}`,
            amountRaisedString: `$${proj.amountRaised}`,
          };
        })
        .sort((a, b) => a.id - b.id);
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
