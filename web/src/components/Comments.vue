<template>
  <section>
    <h3>Comments</h3>
    <ul v-if="comments.length > 0" class="comments">
      <li v-for="(comment, index) in comments" :key="index" class="comments-item">
        <p class="comments-p">{{comment.content}}</p>
        <div class="comments-timestamp">posted {{comment.timeAgo}}</div>
      </li>
    </ul>
    <p v-else class="comments-empty">
      No comments yet. Get the ball rolling!
    </p>
    <el-form ref="form" :model="form" label-position="top" class="form"
    >
      <el-form-item label="Leave a comment">
        <el-input
          type="textarea"
          :autosize="{ minRows: 2 }"
          v-model="form.text">
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="submit" :disabled="!form.text.length">
          Create
        </el-button>
      </el-form-item>
    </el-form>
  </section>
</template>

<script>
import axios from 'axios';
import distanceInWordsToNow from 'date-fns/distance_in_words_to_now';

export default {
  name: 'comments',
  props: {
    projectId: {
      required: true,
    },
  },
  data() {
    return {
      comments: [],
      form: {
        text: '',
      },
    };
  },
  created() {
    axios
      .get(`/api/projects/${this.projectId}/comments`)
      .then((res) => {
        this.comments = res.data.map((comment) => ({
          ...comment,
          timeAgo: distanceInWordsToNow(comment.createdAt, { addSuffix: true }),
        }));
      })
      .catch((err) => console.error(err));
  },
  methods: {
    submit() {
      this.comments.push({
        id: Math.floor(Math.random() * Number.MAX_SAFE_INTEGER),
        userId: 1,
        projectId: this.projectId,
        createdAt: Date.now(),
        timeAgo: distanceInWordsToNow(Date.now(), { addSuffix: true }),
        content: this.form.text,
      });
      axios
        .post('/api/comments', {
          userId: 1,
          projectId: this.projectId,
          content: this.form.text,
        })
        .then(() => {
          // clear text input
          this.form.text = '';
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
};
</script>

<style scoped>
.comments {
  list-style: none;
  margin: 0;
  padding: 0;
}
.comments-item {
  padding-bottom: 0.25rem;
}
.comments-p {
  margin-bottom: 0.15rem;
}
.comments-timestamp {
  color: #909399;
}
.comments-empty {
}
.form {
  margin-top: 1rem;
}
</style>
