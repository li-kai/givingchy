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
        @keyup.enter.native="querySearch(searchTerm)"
        >
      </el-input>
    </el-col>
    <el-col class="categories" :xs="24">
      <el-tag 
        type="info" class="categories-tag"
        v-for="(cate, index) in categories" 
        :key="index" 
        @click.native="querySearch(cate.name)" 
        >
        {{cate.name}} ({{cate.projNum}})
      </el-tag>
    </el-col>
  </el-row>
	<el-row :gutter="20">
    <el-col :xs="24" :sm="12" :md="8" :lg="6" v-for="proj in projects" :key="proj.id">
      <card :project="proj"></card>
    </el-col>
	</el-row>
  <el-row type="flex" justify="space-around" v-if="shouldShowPagination">
    <el-pagination
      layout="prev, pager, next"
      :page-size="pageSize"
      :page-count="pageCount"
      :current-page.sync="pageNum"
      @current-change="queryPageNum"
      >
    </el-pagination>
  </el-row>
  </div>
</template>

<script>
import Card from './Card.vue';
import { projects } from '../fixtures';
import axios from 'axios';

const PAGE_NUM = 1;
const PAGE_SIZE = 12;
const PAGE_COUNT = 3;

function formatSearchQuery(searchTerm) {
  if (searchTerm === '') return searchTerm;
  return `&search=${searchTerm}`;
}

export default {
  name: 'root',
  components: {
    Card,
  },
  data() {
    return {
      projects,
      pageNum: PAGE_NUM,
      pageSize: PAGE_SIZE,
      pageCount: PAGE_COUNT,
      searchTerm: '',
      categories: [],
    };
  },
  created() {
    this.searchTerm = this.$route.query.search || '';
    this.pageNum = parseInt(this.$route.query.page, 10) || PAGE_NUM;
    this.fetchPage();
    this.fetchCategories();
  },
  beforeRouteUpdate(to, from, next) {
    this.searchTerm = to.query.search || '';
    this.pageNum = parseInt(to.query.page, 10) || PAGE_NUM;
    this.pageCount = PAGE_COUNT;
    this.fetchPage();
    next();
  },
  methods: {
    queryPageNum(pageNum) {
      this.$router.push({
        path: `/projects?page=${pageNum}${formatSearchQuery(this.searchTerm)}`,
      });
    },
    querySearch(searchTerm) {
      this.$router.push({ path: `/projects?page=${1}${formatSearchQuery(searchTerm)}` });
    },
    fetchPage() {
      axios
        .get(`api/projects?page=${this.pageNum}&limit=${this.pageSize}${formatSearchQuery(this.searchTerm)}`)
        .then((res) => {
          this.projects = res.data.sort((a, b) => a.id - b.id);
          if (this.pageNum >= this.pageCount && this.projects.length === PAGE_SIZE) {
            this.pageCount = this.pageNum + 1;
          }
        })
        .catch((err) => console.error(err));
    },
    fetchCategories() {
      axios
        .get('api/categories')
        .then((res) => {
          this.categories = res.data;
        })
        .catch((err) => console.error(err));
    },
  },
  computed: {
    shouldShowPagination() {
      return this.projects.length === this.pageSize || this.pageNum !== 1;
    },
  },
};
</script>

<style scoped>
.details {
  flex-wrap: wrap;
}

.search,
.categories {
  display: flex;
  align-items: center;
}

.search {
  margin: 1rem 0;
}

.categories {
  flex-wrap: wrap;
  margin-bottom: 1rem;
}

.categories-tag {
  cursor: pointer;
  margin: 0.125rem;
}
</style>
