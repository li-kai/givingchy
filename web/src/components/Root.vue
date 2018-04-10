<template>
  <div>
  <h1>Crowdfunding the world</h1>
	<el-row :gutter="20">
    <el-col :xs="24" :sm="12" :md="8" :lg="6" v-for="proj in projects" :key="proj.id">
      <card :project="proj"></card>
    </el-col>
	</el-row>
  <el-row type="flex" justify="space-around">
    <el-pagination
      layout="prev, pager, next"
      :page-size="pageSize"
      :page-count="pageCount"
      :current-page.sync="pageNum"
      @current-change="handleCurrentChange"
      >
    </el-pagination>
  </el-row>
  <footer>
    CS2102 Project - Group 23
    <div>
      Members: Chen Yinfang, Li Kai, Peng Hanqiu, Tang Yifeng
    </div>
  </footer>
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
      pageNum: 1,
      pageSize: 12,
      pageCount: 5,
    };
  },
  created() {
    this.fetchPage();
  },
  methods: {
    handleCurrentChange(val) {
      this.pageNum = val;
      if (val >= this.pageCount) {
        this.pageCount += 5;
      }
      this.fetchPage();
    },
    fetchPage() {
      axios
        .get(`api/projects?page=${this.pageNum}&limit=${this.pageSize}`)
        .then((res) => {
          this.projects = res.data.sort((a, b) => a.id - b.id);
        })
        .catch((err) => console.error(err));
    },
  },
};
</script>

<style scoped>

</style>
