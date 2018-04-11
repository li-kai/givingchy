<template>
  <div>
    <h1>Crowdfunding the world</h1>
    <el-form
      ref="form"
      :model="project"
      :rules="rules"
      @submit.native.prevent="submit"
      label-position="right" label-width="10rem"
    >
      <el-form-item label="Title" prop="title">
        <el-input type="text"
          placeholder="Enter your title"
          v-model="project.title">
        </el-input>
      </el-form-item>
      <el-form-item label="Description" prop="description">
        <el-input
          type="textarea"
          :autosize="{ minRows: 2 }"
          placeholder="Enter your description"
          v-model="project.description">
        </el-input>
      </el-form-item>
      <el-form-item label="Deadline" prop="endTime">
        <el-date-picker
          v-model="project.endTime"
          type="datetime"
          placeholder="Select end date and time"
          :picker-options="pickerOptions">
        </el-date-picker>
      </el-form-item>
      <el-form-item label="Funding required" prop="amountRequired">
        <el-input-number
          v-model="project.amountRequired"
          :min="1" :max="10000">
        </el-input-number>
      </el-form-item>
      <el-form-item label="Category" prop="category">
        <el-select v-model="project.category" placeholder="Select category">
          <el-option
            v-for="item in categories"
            :key="item.name"
            :label="item.name"
            :value="item.name">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button native-type="submit" type="primary">
          Submit
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from 'axios';
import { categories } from '../../fixtures';

export default {
  name: 'new',
  data() {
    return {
      project: {
        title: '',
        description: '',
        endTime: '',
        amountRequired: 1,
        image: 'https://picsum.photos/300/150/?random' + Math.random(),
        category: '',
      },
      categories,
      rules: {
        title: [
          {
            required: true,
            message: 'Please enter your title',
            trigger: 'blur',
          },
          { min: 3, message: 'Must be at least 3 characters', trigger: 'blur' },
        ],
        description: [
          {
            required: true,
            message: 'Please enter your description',
            trigger: 'blur',
          },
          {
            min: 20,
            message: 'Please give a brief explanation of your project',
            trigger: 'blur',
          },
        ],
        endTime: [
          {
            required: true,
            message: 'Please enter when your project ends',
            trigger: 'blur',
          },
        ],
        amountRequired: [
          {
            required: true,
          },
        ],
        category: [
          {
            required: true,
          },
        ],
      },
      pickerOptions: {
        firstDayOfWeek: 1,
        disabledDate(date) {
          return date - Date.now() < 0;
        },
        shortcuts: [
          {
            text: 'Tomorrow',
            onClick(picker) {
              picker.$emit('pick', new Date());
            },
          },
          {
            text: 'Next week',
            onClick(picker) {
              const date = new Date();
              date.setTime(date.getTime() + 3600 * 1000 * 24 * 7);
              picker.$emit('pick', date);
            },
          },
        ],
      },
    };
  },
  created() {
    axios
      .get('/api/categories')
      .then((res) => {
        this.categories = res.data;
      })
      .catch((err) => console.error(err));
  },
  methods: {
    submit() {
      const project = { userId: this.$store.getters.user.userId, ...this.project };
      axios.post('/api/project', project).then((res) => {
        this.$router.push({ path: `/projects/${res.data}` });
      })
      .catch((err) => console.error(err));
    },
  },
};
</script>

<style scoped>

</style>
