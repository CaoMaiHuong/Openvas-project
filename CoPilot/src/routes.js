import DashView from './components/Dash.vue'
import LoginView from './components/Login.vue'
import SignupView from './components/Signup.vue'
import NotFoundView from './components/404.vue'
import Vue from 'vue'
import VeeValidate from 'vee-validate'
// Import Views - Dash
import DashboardView from './components/views/Dashboard.vue'
import TablesView from './components/views/Tables.vue'
import TasksView from './components/views/Tasks.vue'
import SettingView from './components/views/Setting.vue'
import AccessView from './components/views/Access.vue'
import ServerView from './components/views/Server.vue'
import ReposView from './components/views/Repos.vue'
import UsersView from './components/users/Users.vue'
// import UserCreateView from './components/users/Create.vue'
// import UserUpdateView from './components/users/Update.vue'
import TargetsView from './components/targets/Targets.vue'
// import TargetCreateView from './components/targets/Create.vue'
import TargetDetail from './components/targets/TargetDetail.vue'
import ScansView from './components/scans/Scans.vue'
import ReportByTask from './components/scans/ReportByTask.vue'
import Report from './components/scans/Report.vue'
import NvtsView from './components/databases/Nvt.vue'
import NvtDetail from './components/databases/NvtDetail.vue'
import CvesView from './components/databases/Cve.vue'
import CveDetail from './components/databases/CveDetail.vue'
import CpesView from './components/databases/Cpe.vue'
import CpeDetail from './components/databases/CpeDetail.vue'
import HostsView from './components/databases/Host.vue'
import HostDetail from './components/databases/HostDetail.vue'

Vue.use(VeeValidate, {
  inject: true,
  fieldsBagName: 'veeFields'
})
// Routes
const routes = [
  {
    path: '/login',
    component: LoginView
  },
  {
    path: '/signup',
    component: SignupView
  },
  {
    path: '/',
    component: DashView,
    children: [
      {
        path: 'dashboard',
        alias: '',
        component: DashboardView,
        name: 'Trang chủ',
        meta: {description: ''}
      }, {
        path: 'tables',
        component: TablesView,
        name: 'Tables',
        meta: {description: 'Simple and advance table in CoPilot'}
      }, {
        path: 'tasks',
        component: TasksView,
        name: 'Tasks',
        meta: {description: ''}
      }, {
        path: 'setting',
        component: SettingView,
        name: 'Settings',
        meta: {description: 'User settings page'}
      }, {
        path: 'access',
        component: AccessView,
        name: 'Access',
        meta: {description: 'Example of using maps'}
      }, {
        path: 'server',
        component: ServerView,
        name: 'Servers',
        meta: {description: 'List of our servers', requiresAuth: true}
      }, {
        path: 'repos',
        component: ReposView,
        name: 'Repository',
        meta: {description: 'List of popular javascript repos'}
      }, {
        path: 'users',
        component: UsersView,
        name: 'Người dùng',
        meta: {description: 'Danh sách người dùng'}
      }, {
        path: 'targets',
        component: TargetsView,
        name: 'Targets',
        meta: {description: 'Danh sách target'}
      }, {
        path: '/target/:id',
        component: TargetDetail,
        name: 'Target Detail',
        props: true,
        meta: {description: ''}
      }, {
        path: 'scans',
        component: ScansView,
        name: 'Scans',
        meta: {description: 'Danh sách task'}
      }, {
        path: '/reports/:id',
        component: ReportByTask,
        name: 'Danh sách báo cáo',
        props: true,
        meta: {description: ''}
      }, {
        path: '/report/:id',
        component: Report,
        name: 'Báo cáo',
        props: true,
        meta: {description: ''}
      }, {
        path: 'nvts',
        component: NvtsView,
        name: 'Nvts',
        meta: {description: 'Danh sách NVT'}
      }, {
        path: '/nvt/:id',
        component: NvtDetail,
        name: 'Nvt Detail',
        props: true,
        meta: {description: ''}
      }, {
        path: 'cves',
        component: CvesView,
        name: 'Cves',
        meta: {description: 'Danh sách CVE'}
      }, {
        path: '/cve/:name',
        component: CveDetail,
        name: 'Cve Detail',
        props: true,
        meta: {description: ''}
      }, {
        path: 'cpes',
        component: CpesView,
        name: 'Cpes',
        meta: {description: 'Danh sách CPE'}
      }, {
        path: '/cpe/:id',
        component: CpeDetail,
        name: 'Cpe Detail',
        props: true,
        meta: {description: ''}
      }, {
        path: 'hosts',
        component: HostsView,
        name: 'Hosts',
        meta: {description: 'Danh sách host'}
      }, {
        path: '/host/:id',
        component: HostDetail,
        name: 'Host Detail',
        props: true,
        meta: {description: ''}
      }
    ]
  }, {
    // not found handler
    path: '*',
    component: NotFoundView
  }
]

export default routes
