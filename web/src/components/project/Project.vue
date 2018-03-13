<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="16">
        <img src="https://via.placeholder.com/150x150" class="image">
      </el-col>
      <el-col :xs="24" :sm="8">
        <h1>{{project.title}}</h1>
        <h2>{{project.description}}</h2>
        <el-progress :percentage="fundingPercentage" :status="fundingStatus"></el-progress>
        <el-row class="keywords">
          <el-col :span="1" v-for="(keyword, index) in project.keywords" :key="index">
            <el-tag type="info">{{keyword}}</el-tag>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row>
      <el-col>
        <p>
          Started on {{ new Date(project.startDate).toLocaleString() }}
        </p>
        <p>
          Closes in {{ timeLeft }}
        </p>
        <el-slider :max="1000" v-model="fundingAmount" show-input></el-slider>
        <el-button type="primary" @click="submit()">
          Back project
        </el-button>
      </el-col>
    </el-row>
    </el-col>
  </el-row>
</template>

<script>
import projects from '../../fixtures';
export default {
  name: 'project',
  data() {
    return {
      project: projects[this.$route.params.id],
      fundingAmount: 0,
    };
  },
  computed: {
    timeLeft() {
      const secondsLeft = Date.now() - this.project.startDate + this.project.duration;
      const minutesLeft = secondsLeft / 60;
      if (minutesLeft < 60) return `${Math.round(minutesLeft)} minutes`;
      const hoursLeft = minutesLeft / 60;
      if (hoursLeft < 24) return `${Math.round(hoursLeft)} hours`;
      const daysLeft = hoursLeft / 24;
      return `${Math.round(daysLeft)} days`;
    },
    fundingPercentage() {
      const { fundingRaised, fundingRequired } = this.project;
      const twoDecimalPlaces = (fundingRaised / fundingRequired * 100).toFixed(2);
      return parseFloat(twoDecimalPlaces);
    },
    fundingStatus() {
      const secondsLeft = Date.now() - this.project.startDate + this.project.duration;
      if (secondsLeft < 0) return 'exception';
      return this.fundingPercentage >= 100 ? 'success' : '';
    },
  },
  methods: {
    submit() {
      this.$notify({
        title: 'Success!',
        message: "You've backed a project",
        type: 'success',
      });
    },
  },
};
</script>

<style scoped>
.image {
  width: 100%;
  object-fit: cover;
}
.keywords {
  margin: 0.5rem 0.5rem 0 0;
}
</style>
