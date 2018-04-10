<template>
  <div>
  <el-row class="details" type="flex" justify="space-between">
    <el-col :xs="24" :md="8">
      <h1>Crowdfunding the world</h1>
    </el-col>
    <el-col class="search" :xs="24" :md="4">
      <el-input
        placeholder="Search projects"
        prefix-icon="el-icon-search"
        v-model="searchTerm"
        @keyup.enter.native="submitSearch"
        >
      </el-input>
    </el-col>
  </el-row>
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
      searchTerm: '',
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
        .get(`api/projects?page=${this.pageNum}&limit=${this.pageSize}&search=${this.searchTerm}`)
        .then((res) => {
          this.projects = res.data.sort((a, b) => a.id - b.id);
        })
        .catch((err) => console.error(err));
    },
    submitSearch() {
      this.fetchPage();
      this.searchTerm = '';
    }
  },
};
</script>

<style scoped>
.details {
  flex-wrap: wrap;
}

.search {
  display: flex;
  align-items: center;
  margin: 1rem 0;
}
</style>
