import Vue from 'vue';
import Router from 'vue-router';
import Auth from '../auth';
import Root from '../components/Root';
import Project from '../components/project/Project';
import NewProject from '../components/project/NewProject';
import AdminSidebar from '../components/admin/Sidebar';
import UsersAdmin from '../components/admin/Users';
import ProjectsAdmin from '../components/admin/Projects';
import PaymentsAdmin from '../components/admin/Payments';
import CommentsAdmin from '../components/admin/Comments';
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
      component: AdminSidebar,
      redirect: '/admin/users',
      meta: { requiresAuth },
      children: [
        {
          path: 'users',
          component: UsersAdmin,
          meta: { requiresAuth },
        },
        {
          path: 'projects',
          component: ProjectsAdmin,
          meta: { requiresAuth },
        },
        {
          path: 'payments',
          component: PaymentsAdmin,
          meta: { requiresAuth },
        },
        {
          path: 'comments',
          component: CommentsAdmin,
          meta: { requiresAuth },
        },
      ],
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
