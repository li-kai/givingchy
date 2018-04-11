<template>
  <div>
    <el-row type="flex" justify="center">
      <el-col class="gap" :xs="24" :sm="10">
        <div class="alert alert-danger" v-if="error">
          <p>{{ error }}</p>
        </div>
        <el-form
          ref="form"
          :model="credentials" :rules="rules"
          @submit.native.prevent="submit"
          label-position="right" label-width="6rem"
        >
          <el-form-item label="Email" prop="email">
            <el-input type="text"
              placeholder="Enter your email"
              v-model="credentials.email">
            </el-input>
          </el-form-item>
          <el-form-item v-if="!isLogin" label="Username" prop="username">
            <el-input type="username"
              placeholder="Enter your username"
              v-model="credentials.username">
            </el-input>
          </el-form-item>
          <el-form-item label="Password" prop="password">
            <el-input type="password"
              placeholder="Enter your password"
              v-model="credentials.password">
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button native-type="submit" type="primary">
              {{ isLogin ? 'Log in' : 'Sign up'}}
            </el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col class="aside" :xs="24" :sm="10">
        {{isLogin ? "New user?" : "Already have an account?"}}
        <router-link :to="isLogin ? '/signup' : '/login'">
          {{isLogin ? "Sign up" : "Log in"}}
        </router-link>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'auth',
  props: {
    isLogin: {
      type: Boolean,
      required: true,
    },
  },
  data() {
    return {
      // We need to initialize the component with any
      // properties that will be used in it
      credentials: {
        email: '',
        username: '',
        image: 'https://www.gravatar.com/avatar/',
        password: '',
      },
      rules: {
        email: [
          {
            required: true,
            message: 'Please enter your email',
            trigger: 'blur',
          },
          { min: 3, message: 'Must be at least 3 characters', trigger: 'blur' },
        ],
        username: [
          {
            required: !this.isLogin,
            message: 'Please enter your username',
            trigger: 'blur',
          },
          { min: 3, message: 'Must be at least 3 characters', trigger: 'blur' },
        ],
        password: [
          {
            required: true,
            message: 'Please enter your password',
            trigger: 'blur',
          },
          {
            min: 8,
            max: 100,
            message: 'Must be between 8 and 100 characters',
            trigger: 'blur',
          },
        ],
      },
      error: '',
    };
  },
  methods: {
    submit() {
      const endPoint = this.isLogin ? '/api/auth' : 'api/user';
      const body = this.isLogin
        ? {
            email: this.credentials.email,
            password: this.credentials.password,
          }
        : this.credentials;
      axios
        .post(endPoint, body)
        .then((res) => {
          this.$store.dispatch('loginUser', res.data);
          this.$router.push(this.$route.query.redirect || '/');
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
};
</script>

<style scoped>
.gap {
  border-bottom: 1px solid #e4e7ed;
}
.aside {
  margin: 0.3rem;
}
</style>
