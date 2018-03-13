<template>
  <el-card class="card">
    <img alt="placeholder" src="https://via.placeholder.com/300x150" class="image">
    <div>
      <router-link :to="projectPath">
        <h3>{{project.title}}</h3>
      </router-link>
      <h4>{{project.description}}</h4>
      <span>{{timeLeft}} left</span>
    </div>
    <el-row class="keywords">
      <el-col :span="1" v-for="(keyword, index) in project.keywords" :key="index">
        <el-tag type="info">{{keyword}}</el-tag>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
export default {
  name: 'card',
  props: {
    project: {
      required: true,
    },
  },
  computed: {
    timeLeft() {
      const secondsLeft = (new Date(this.project.endTime) - Date.now()) / 1000;
      const minutesLeft = secondsLeft / 60;
      if (minutesLeft < 60) return `${Math.round(minutesLeft)} minutes`;
      const hoursLeft = minutesLeft / 60;
      if (hoursLeft < 24) return `${Math.round(hoursLeft)} hours`;
      const daysLeft = hoursLeft / 24;
      return `${Math.round(daysLeft)} days`;
    },
    projectPath() {
      return { name: 'projects', params: { id: this.project.id } };
    },
  },
};
</script>

<style scoped>
.image {
  width: 100%;
  max-height: 12.5rem;
  object-fit: cover;
}
.card {
  height: 25rem;
  margin: 0 0 1.5rem;
}
.keywords {
  margin: 0.5rem 0.5rem 0 0;
}
</style>
