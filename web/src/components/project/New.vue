<template>
  <div>
    <h1>Crowdfunding the world</h1>
    <el-form
      ref="form"
      :model="project"
      :rules="rules"
      label-position="right" label-width="8rem"
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
      <el-form-item label="Funding required" prop="fundingRequired">
        <el-input-number
          v-model="project.fundingRequired"
          :min="1" :max="10000">
        </el-input-number>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submit()">
          Submit
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'new',
  data() {
    return {
      project: {
        title: '',
        description: '',
        endTime: '',
        fundingRequired: 1,
        userId: 1,
      },
      categories: [],
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
        fundingRequired: [
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
        console.log(res.data);
        this.categories = res.data.map((cat) => cat.name);
      })
      .catch((err) => console.error(err));
  },
  methods: {
    submit() {
      // TODO: sent to server side
    },
  },
};
</script>

<style scoped>

</style>
