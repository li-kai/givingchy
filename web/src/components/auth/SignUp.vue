<template>
  <auth :isLogin="false"></auth>
</template>

<script>
import Auth from './Auth';
export default {
  name: 'signup',
  components: { Auth }
};
</script>
