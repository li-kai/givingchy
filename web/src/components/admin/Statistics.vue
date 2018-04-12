<template>
<div>
  <h1>Statistics</h1>
  <hr>
  <h2>User rankings</h2>
  <el-table
    :data="donors"
    style="width: 100%">
    <el-table-column
      prop="username"
      label="Username">
      <template slot-scope="scope">
        <div class="user">
          <img class="user-image" :src="scope.row.image" alt="user profile">
          <div>{{scope.row.username}}</div>
        </div>
      </template>
    </el-table-column>
    <el-table-column
      prop="totalDonation"
      label="Total amount donated">
    </el-table-column>
  </el-table>
  <hr>
  <h2>User payment statistics</h2>
  <el-table
    :data="userData"
    style="width: 100%">
    <el-table-column
      prop="percentile"
      label="Percentile">
    </el-table-column>
    <el-table-column
      prop="value"
      label="Value">
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'statistics',
  data() {
    return {
      userData: [],
      donors: [],
    };
  },
  created() {
    axios.get('/api/stats').then((res) => {
      const quartileKeys = ['Minimum', 'First Quartile', 'Median', 'Third Quartile', 'Maximum'];
      const quartileValues = res.data.find((stat) => stat.name === "users").values;
      this.userData = quartileKeys.map((key, i) => ({ percentile: key, value: quartileValues[i] }));
      this.donors = res.data.find((stat) => stat.name === "top_donors").values;
      console.log(this.donors);
    });
  },
};
</script>

<style scoped>
.user {
  display: flex;
  align-items: center;
}
.user-image {
  height: 70px;
  margin-right: 1rem;
}
</style>
