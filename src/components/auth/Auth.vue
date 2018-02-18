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
          label-position="right" label-width="6rem"
        >
          <el-form-item label="Username" prop="username">
            <el-input type="text"
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
            <el-button type="primary" @click="submit()">
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
export default {
  name: 'auth',
  props: {
    isLogin: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      // We need to initialize the component with any
      // properties that will be used in it
      credentials: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          {
            required: true,
            message: 'Please enter your username',
            trigger: 'blur'
          },
          { min: 3, message: 'Must be at least 3 characters', trigger: 'blur' }
        ],
        password: [
          {
            required: true,
            message: 'Please enter your password',
            trigger: 'blur'
          },
          {
            min: 8,
            max: 100,
            message: 'Must be between 8 and 100 characters',
            trigger: 'blur'
          }
        ]
      },
      error: ''
    };
  },
  methods: {
    submit() {
      // auth.login(this, this.credentials, 'secretquote');
    }
  }
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
