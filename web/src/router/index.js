import Vue from 'vue';
import Router from 'vue-router';
import Auth from '../auth';
import Root from '../components/Root';
import Project from '../components/project/Project';
import NewProject from '../components/project/New';
import Admin from '../components/auth/Admin';
import Login from '../components/auth/Login';
import SignUp from '../components/auth/SignUp';

Vue.use(Router);
const requiresAuth = true;

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Root,
      alias: '/projects',
    },
    {
      path: '/projects/new',
      component: NewProject,
      meta: { requiresAuth },
    },
    {
      path: '/projects/:id',
      component: Project,
      name: 'projects',
    },
    {
      path: '/login',
      component: Login,
    },
    {
      path: '/signup',
      component: SignUp,
    },
    {
      path: '/admin',
      component: Admin,
      meta: { requiresAuth },
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    if (!Auth.loggedIn()) {
      next({
        path: '/login',
        query: { redirect: to.fullPath },
      });
      return;
    }
  }
  next();
});

export default router;
