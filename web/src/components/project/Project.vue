<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18" v-loading="isLoading">
      <el-row :gutter="20">
        <el-col :xs="24" :sm="16" :lg="10">
          <img :src="project.image" class="image">
        </el-col>
        <el-col :xs="24" :sm="8" :lg="14">
          <h1 class="proj-title">{{project.title}}</h1>
          <h2 class="proj-desc">{{project.description}}</h2>
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
          <div class="amount-raised">${{project.amountRaised}}</div>
          <div class="amount-required">raised of ${{project.amountRequired}} goal</div>
          <el-progress
            :percentage="fundingPercentage"
            :status="fundingStatus"
            class="amount-percentage">
          </el-progress>
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
          <el-slider :max="1000" v-model="fundingAmount" show-input class="amount-slider"></el-slider>
          <el-button type="primary" @click="submit()" :disabled="fundingAmount <= 0">
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
import Comments from './Comments.vue';
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
      const float = parseFloat(twoDecimalPlaces);
      return float <= 100 ? float : 100;
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
      this.project.amountRaised += this.fundingAmount;
      axios.post("/api/payments", {
        userId: 1,
        projectId: this.project.id,
        amount: this.fundingAmount,
      }).then(() => {
        this.fundingAmount = 0;
        // reset amount
        this.$notify({
          title: 'Success!',
          message: "You've backed a project",
          type: 'success',
        });
      }).catch((err) => {
        console.error(err);
      });
    },
  },
};
</script>

<style scoped>
.image {
  height: 100%;
  max-height: 100%;
}
.proj-title {
  margin: 0.2rem 0;
}
.proj-desc {
  margin: 0.1rem 0 0.2rem;
  color: #606266;
}
.amount-raised {
  font-size: 1.35em;
}
.amount-required {
  color: #606266;
}
.amount-percentage {
  display: flex;
  align-items: center;
  max-width: 16rem;
}
.amount-slider {
  max-width: 36rem;
}
.tags {
  margin: 0 0 1.25rem -0.05rem;
}
.tags-tag {
  margin: 0.25rem 0.25rem 0 0;
}
.keyline {
  margin-top: 2rem;
  border-top: 1px solid #e4e7ed;
}
</style>
