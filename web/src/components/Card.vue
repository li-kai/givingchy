<template>
  <el-card class="card">
    <router-link class="card-link" :to="projectPath">
      <img alt="placeholder" :src="project.image" class="image">
      <div class="card-content">
          <h3>{{project.title}}</h3>
        <h4 class="card-desc">{{project.description}}</h4>
        <span class="card-time">{{timeLeft}}</span>
      </div>
      <div class="tags">
        <el-tag 
          class="tags-tag" 
          type="info" 
          size="small" 
          color="#FFF"
          v-for="(tag, index) in project.tags" 
          :key="index">
          {{tag}}
        </el-tag>
      </div>
    </router-link>
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
  height: 25rem;
  margin: 0 0 1.5rem;
  padding: 0;
}

.card-link {
  display: flex;
  flex-direction: column;
  height: 100%;
  color: inherit;
  text-decoration: none;
}

.card-content {
  flex: 1;
  padding: 0 1rem;
}

.card-desc {
  margin: 0.25rem 0;
}

.card-time {
  color: #606266;
}

.tags {
  align-self: bottom;
  margin: 1rem;
}

.tags-tag {
  margin: 0.25rem 0.25rem 0 0;
}
</style>
