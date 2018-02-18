<template>
  <div>
    <h1>Crowdfunding the world</h1>
    <el-form
      ref="form"
      :model="project"
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
      <el-form-item label="Duration">
        <el-date-picker
          v-model="project.dateTimeRange"
          :picker-options="pickerOptions"
          type="datetimerange"
          format="d-M-yyyy h:mma"
          range-separator="To"
          start-placeholder="Start date"
          end-placeholder="End date">
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
export default {
  name: 'new',
  data() {
    return {
      project: {
        title: '',
        description: '',
        dateTimeRange: '',
        fundingRequired: '',
      },
      pickerOptions: {
        firstDayOfWeek: 1,
        disabledDate(date) {
          return date - Date.now() < 0;
        },
      },
    };
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
