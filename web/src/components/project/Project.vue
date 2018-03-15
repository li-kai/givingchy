<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18" v-loading="isLoading">
      <el-row :gutter="20">
        <el-col :xs="24" :sm="16" :lg="10">
          <img src="https://via.placeholder.com/150x150" class="image">
        </el-col>
        <el-col :xs="24" :sm="8" :lg="14">
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
            Started on {{ new Date(project.startTime).toLocaleString() }}
          </p>
          <p>
            {{ timeLeft }}
          </p>
          <el-slider :max="1000" v-model="fundingAmount" show-input></el-slider>
          <el-button type="primary" @click="submit()">
            Back project
          </el-button>
        </el-col>
      </el-row>
      <el-row v-if="!isLoading" class="keyline">
        <el-col>
          <comments :project-id="project.id"/>
        </el-col>
      </el-row>
    </el-col>
  </el-row>
</template>

<script>
import Comments from '../Comments.vue';
import axios from 'axios';
import distanceInWordsToNow from 'date-fns/distance_in_words_to_now';
import isFuture from 'date-fns/is_future';
import { projects } from '../../fixtures';

export default {
  name: 'project',
  components: {
    Comments,
  },
  data() {
    return {
      project: projects[0],
      fundingAmount: 0,
      isLoading: true,
    };
  },
  computed: {
    timeLeft() {
      const { endTime } = this.project;
      const ago = distanceInWordsToNow(this.project.endTime);
      if (isFuture(endTime)) {
        return `Closes in ${ago}`;
      }
      return "Project has ended";
    },
    fundingPercentage() {
      const { amountRaised, amountRequired } = this.project;
      const twoDecimalPlaces = (amountRaised / amountRequired * 100).toFixed(2);
      return parseFloat(twoDecimalPlaces);
    },
    fundingStatus() {
      const secondsLeft = Date.now() - this.project.endTime;
      if (secondsLeft < 0) return 'exception';
      return this.fundingPercentage >= 100 ? 'success' : '';
    },
  },
  created() {
    axios
      .get(`/api/projects/${this.$route.params.id}`)
      .then((res) => {
        this.project = res.data;
        this.isLoading = false;
      })
      .catch((err) => console.error(err));
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
.keyline {
  margin-top: 2rem;
  border-top: 1px solid #e4e7ed;
}
</style>
