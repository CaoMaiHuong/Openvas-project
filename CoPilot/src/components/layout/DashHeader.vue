<template>
  <header class="main-header">
    <span class="logo-mini">
      <a href="/"><img src="/static/img/81a6fc4f9c.svg" alt="Logo" class="img-responsive center-block logo"></a>
    </span>
    <!-- Header Navbar -->
    <nav class="navbar navbar-static-top" role="navigation"> 
      <!-- Sidebar toggle button-->
      <a href="javascript:;" class="sidebar-toggle" data-toggle="offcanvas" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>
      <span class="web-name">HỆ THỐNG QUẢN LÝ RỦI RO AN TOÀN THÔNG TIN</span>
      <!-- Navbar Right Menu -->
      <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">
          <!-- <messages-menu></messages-menu>
          <notifications-menu></notifications-menu> -->
          <!-- <tasks-menu></tasks-menu> -->
          <user-menu :user="user"></user-menu>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
import { mapState } from 'vuex'
// import MessagesMenu from './MessagesMenu'
// import NotificationsMenu from './NotificationsMenu'
// import TasksMenu from './TasksMenu'
import UserMenu from './UserMenu'

export default {
  name: 'DashHeader',
  components: {
    // MessagesMenu,
    // NotificationsMenu,
    // TasksMenu,
    UserMenu
  },
  props: ['user'],
  computed: {
    ...mapState([
      'userInfo'
    ])
  }
}
</script>
<style>
  .web-name{
    float: left;
    padding: 11px;
    color: #fff;
    font-size: 20px;
  }
</style>
