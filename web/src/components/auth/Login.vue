<template>
  <auth :isLogin="true"></auth>
</template>

<script>
import Auth from './Auth';
export default {
  name: 'login',
  components: { Auth }
};
</script>
