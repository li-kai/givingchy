<template>
  <div>
  <h1>Crowdfunding the world</h1>
	<el-row :gutter="20">
    <el-col :xs="24" :sm="12" :md="8" :lg="6" v-for="proj in projects" :key="proj.id">
      <card :project="proj"></card>
    </el-col>
	</el-row>
  </div>
</template>

<script>
import Card from './Card.vue';
import { projects } from '../fixtures';
import axios from 'axios';

export default {
  name: 'root',
  components: {
    Card,
  },
  data() {
    return {
      projects,
    };
  },
  created() {
    axios
      .get('api/projects')
      .then((res) => {
        this.projects = res.data.sort((a, b) => a.id - b.id);
      })
      .catch((err) => console.error(err));
  },
};
</script>

<style scoped>

</style>
