import Vue from "vue";
import Router from "vue-router";
import Root from "../components/Root";
import Project from "../components/project/Project";
import NewProject from "../components/project/New";
import Login from "../components/auth/Login";
import SignUp from "../components/auth/SignUp";

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
    },
    {
      path: "/projects/:id",
      component: Project,
      name: "projects"
    },
    {
      path: "/login",
      component: Login
    },
    {
      path: "/signup",
      component: SignUp
    }
  ]
});
