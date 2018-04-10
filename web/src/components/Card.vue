<template>
  <el-card class="card">
    <img alt="placeholder" :src="project.image" class="image">
    <div>
      <router-link :to="projectPath">
        <h3>{{project.title}}</h3>
      </router-link>
      <h4>{{project.description}}</h4>
      <span>{{timeLeft}}</span>
    </div>
    <el-row class="tags">
      <el-col :span="1" v-for="(tag, index) in project.tags" :key="index">
        <el-tag type="info">{{tag}}</el-tag>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
import distanceInWordsToNow from 'date-fns/distance_in_words_to_now';
import isFuture from 'date-fns/is_future';
export default {
  name: 'card',
  props: {
    project: {
      required: true,
    },
  },
  computed: {
    timeLeft() {
      const { endTime} = this.project;
      const ago = distanceInWordsToNow(this.project.endTime);
      if (isFuture(endTime)) {
        return `${ago} left`;
      }
      return `Ended ${ago} ago`;
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
.tags {
  margin: 0.5rem 0.5rem 0 0;
}
</style>
