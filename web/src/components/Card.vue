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
.card {
  display: flex;
  flex-direction: column;
  height: 25rem;
  margin: 0 0 1.5rem;
}

@keyframes placeHolderShimmer{
    0%{
        background-position: 100% 0
    }
    100%{
        background-position: -100% 0
    }
}

.image {
  max-height: 12.5rem;
  height: 150px;
  width: 100%;
  object-fit: cover;
  line-height: 150px;
  text-align: center;
  animation: 1.25s linear placeHolderShimmer infinite;
  background: #f6f7f8;
  background: linear-gradient(80deg, #eeeeee 8%, #dddddd 18%, #eeeeee 33%);
  background-size: 250% 100%;
}

.tags {
  margin: 0.5rem 0.5rem 0 0;
}
</style>
