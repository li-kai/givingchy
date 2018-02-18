import Vue from "vue";
import Router from "vue-router";
import Root from "../components/Root.vue";
import NewProject from "../components/project/New.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/",
      component: Root,
      alias: "/projects"
    },
    {
      path: "/projects/new",
      component: NewProject
    }
  ]
});
