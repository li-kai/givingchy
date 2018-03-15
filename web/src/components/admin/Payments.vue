<template>
<div>
  <h1>Payments</h1>
  <el-table
    :data="payments"
    :row-key="payments.id"
    style="width: 100%">
    <el-table-column
      prop="id"
      label="ID"
      width="50">
    </el-table-column>
    <el-table-column
      prop="userId"
      label="User ID"
      width="70">
    </el-table-column>
    <el-table-column
      prop="amount"
      label="Paid">
      <div></div>
    </el-table-column>
    <el-table-column
      prop="paidAt"
      label="Paid At">
      <date-time
      slot-scope="scope"
      :datetime="scope.row.paidAt"></date-time>
    </el-table-column>
    <el-table-column
      label="Actions">
      <template slot-scope="scope">
        <el-button type="danger" size="small">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import axios from 'axios';
import DateTime from './DateTime';

export default {
  name: 'payments',
  components: {
    DateTime,
  },
  data() {
    return {
      payments: [],
    };
  },
  created() {
    axios
      .get('/api/payments', this.credentials)
      .then((res) => {
        this.payments = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
  },
};
</script>

<style scoped>

</style>
