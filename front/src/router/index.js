import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from "../views/Login.vue";
import Home from '@/views/Home.vue';
import Admin from "@/views/Admin.vue";
import SignUp from '@/views/SignUp.vue';

const TaskList = () => import('@/components/task/TaskList.vue')
const TaskDetail = () => import('@/components/task/TaskDetail.vue')
const AddTask = () => import('@/components/task/AddTask.vue')
const AcceptedTask = () => import('@/components/task/AcceptedTask.vue')
const PublishedTask = () => import('@/components/task/PublishedTask')
const FinishedTask = () => import('@/components/task/FinishedTask')
const Community = () => import('@/components/Community')
const ArticleDetail = () => import('@/components/article/Details')

const AdminIndex = () => import('@/components/admin/Index')
const Profile = () => import('@/components/user/Profile')
const AdminArticleList = () => import('@/components/article/ArticleList')
const AdminAddArticle = () => import('@/components/article/AddArticle')

const letter = () => import('@/components/letter/Letter')
const Chat = () => import('@/components/letter/Chat')

const L7 = () => import('@/components/L7/L7')
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: Home,
    children: [
      { path: '/index', name: "Index", component: TaskList, meta: { title: '欢迎来到HelpBank' } },
      { path: '/', name: 'Home', component: TaskList, meta: { title: '欢迎来到HelpBank' } },
      { path: '/task/detail/:taskId', name: "Detail", component: TaskDetail, meta: { title: '任务详情' }, props: true },
      { path: '/task/add', component: AddTask, meta: { title: '发布任务' } },
      { path: '/task/edit/:id', component: AddTask, meta: { title: '发布任务' }, props: true },
      { path: '/task/publishedTask', component: PublishedTask, meta: { title: '任务列表' } },
      { path: '/task/acceptedTask', component: AcceptedTask, meta: { title: '任务列表' } },
      { path: '/task/finishedTask', component: FinishedTask, meta: { title: '任务列表' } },
      { path: '/community', component: Community, meta: { title: '社区' } },
      { path: '/article/detail/:id', component: ArticleDetail, meta: { title: '文章详情' }, props: true },
      {
        path: 'letter', name: "Letter", component: letter, children: [
          { path: 'chat/:userId', name: 'Chat', component: Chat, meta: { title: '私信' } }
        ], meta: { title: '私信' }, props: true
      },
      { path: '/L7', name: "L7", component: L7, meta: { title: 'L7' } },
      // { path: '/letter', name: "Letter", component: letter, meta: { title: '私信' }, props: true },
    ],
    meta: { title: '首页' }
  },
  {
    path: '/admin',
    component: Admin,
    name: 'Admin',
    children: [
      { path: '/', name: "AdminIndex", component: AdminIndex, meta: { title: '后台管理' } },
      { path: 'profile', name: "Profile", component: Profile, meta: { title: '个人中心' } },
      { path: 'artList', name: "AdminArticleList", component: AdminArticleList, meta: { title: '文章列表' } },
      { path: 'addArticle/:id', name: "AdminAddArticle", component: AdminAddArticle, meta: { title: '发布文章' }, props: true },
      { path: 'addArticle', name: "AdminAddArticle", component: AdminAddArticle, meta: { title: '发布文章' } },
    ],
    meta: { title: '后台管理' }
  },
  // {
  //   path: '/letter',
  //   component: letter,
  //   name: 'Letter',
  //   children: [
  //     { path: '/', name: "Letter", component: letter, meta: { title: '私信' } },
  //   ],
  //   meta: { title: '私信' }
  // },
  {
    path: '/signUp',
    component: SignUp,
    name: 'SignUp',
    meta: { title: '注册' }
  },
  {
    path: '/login',
    component: Login,
    name: 'Login',
    meta: { title: '登录' }
  },

]

const router = new VueRouter({
  // mode: 'history',
  // base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem("token");
  if (to.meta.title) {
    document.title = to.meta.title ? to.meta.title : '加载中'
  }
  if (to.name === "Login" || to.name === "Index" || to.name === "Home" || to.name === "Detail" || to.name === "SignUp") {
    return next();
  } else {
    if (!token) {
      next({ name: "Login" });
    } else {
      next();
    }
  }
});

// 设置页面标题
// router.beforeEach((to, from, next) => {
//   if (to.meta.title) {
//     document.title = to.meta.title ? to.meta.title : '加载中'
//   }
//   next()
// })

export default router
