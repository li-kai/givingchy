<template>
<div>
  <h1>Payments</h1>
  <data-tables-server
    :data="payments"
    :total="10000"
    :row-key="payments.id"
    :load-data="fetchPage"
    :table-props="{ stripe: false, border: false }"
    :search-def="{ show: false }">
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
      slot-scope="date"
      :datetime="date.row.paidAt"></date-time>
    </el-table-column>
    <el-table-column
      label="Actions">
      <template slot-scope="actions">
        <el-button
          type="danger"
          size="small"
          @click="deletePayment(actions.row.id)">
          Delete
        </el-button>
      </template>
    </el-table-column>
  </data-tables-server>
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
  methods: {
    fetchPage({ page, pageSize }) {
      return axios
        .get(`/api/payments?page=${page}&limit=${pageSize}`)
        .then((res) => {
          this.payments = res.data;
        })
        .catch((err) => {
          this.$notify.error({ title: 'Error', message: err.response.data.error });
        });
    },
    deletePayment(id) {
      axios
        .delete(`/api/payments/${id}`)
        .then(() => {
          this.payments = this.payments.filter((payment) => payment.id !== id);
        })
        .catch((err) => {
          this.$notify.error({ title: 'Error', message: err.response.data.error });
        });
    },
  },
};
</script>

<style scoped>

</style>
