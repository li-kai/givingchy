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
      width="70">
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
      label="StartTime">
      <template slot-scope="scope">
        <time class="time">
          <span>{{scope.row.startTime.weekday}}</span>
          <span>{{scope.row.startTime.date}}</span>
          <span>{{scope.row.startTime.time}}</span>
        </time>
      </template>
    </el-table-column>
    <el-table-column
      prop="endTime"
      label="EndTime">
      <template slot-scope="scope">
        <time class="time">
          <span>{{scope.row.endTime.weekday}}</span>
          <span>{{scope.row.endTime.date}}</span>
          <span>{{scope.row.endTime.time}}</span>
        </time>
      </template>
    </el-table-column>
    <el-table-column
      prop="category"
      label="Category">
    </el-table-column>
    <el-table-column
      prop="verified"
      label="Verified">
      <template slot-scope="scope">
        <el-button type="success" size="small"  v-if="scope.row.verified">Verified</el-button>
        <el-button type="danger" size="small" v-else>Unverified</el-button>
      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import axios from 'axios';

const formatter = new Intl.DateTimeFormat('en-US', {
  weekday: 'long',
  year: 'numeric',
  month: 'short',
  day: 'numeric',
  hour: 'numeric',
  minute: 'numeric',
  timeZoneName: 'short'
});
function formatDate(dateString) {
  const fullString = formatter.format(new Date(dateString));
  const [weekday, date, year, time] = fullString.split(', ');
  return {
    weekday,
    date: `${date}, ${year}`,
    time,
  }
}

export default {
  name: 'projects',
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
            startTime: formatDate(proj.startTime),
            endTime: formatDate(proj.endTime),
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
