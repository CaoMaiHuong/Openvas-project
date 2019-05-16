<template>
  <ul class="sidebar-menu">
    <router-link tag="li" class="pageLink" to="/">
      <a>
        <i class="fa fa-desktop"></i>
        <span class="page">{{ $t('dashboardMsg') }}</span>
      </a>
    </router-link>
    <router-link tag="li" class="pageLink" to="/targets">
      <a>
        <i class="fa fa-bullseye"></i>
        <span class="page">Target</span>
      </a>
    </router-link>
    <router-link tag="li" class="pageLink" to="/hosts">
      <a>
        <i class="fa fa-globe"></i>
        <span class="page">Host</span>
      </a>
    </router-link>
    <router-link tag="li" class="pageLink" to="/tasks">
      <a>
        <i class="fa fa-area-chart"></i>
        <span class="page">Tác vụ</span>
      </a>
    </router-link>
    <!-- <router-link tag="li" class="pageLink" to="/vulnerabilities">
      <a>
        <i class="fa fa-bug"></i>
        <span class="page">Vulnerabilities</span>
      </a>
    </router-link> -->
    <router-link tag="li" class="pageLink" to="/users">
      <a>
        <i class="fa fa-users"></i>
        <span class="page">{{ $t('usermanaMsg') }}</span>
      </a>
    </router-link>
    <li class="treeview">
      <a href="#">
        <i class="fa fa-folder-o"></i>
        <span class="treeview-title">Tập dữ liệu mẫu</span>
        <span class="pull-right-container pull-right">
          <i class="fa fa-angle-left fa-fw"></i>
        </span>
      </a>
      <ul class="treeview-menu database-menu">
        <li>
          <router-link tag="li" class="pageLink" to="/cves">
          <a href="#">
            <i class="fa fa-circle-o"></i> CVEs
          </a>
          </router-link>
        </li>
        <li>
          <router-link tag="li" class="pageLink" to="/cpes">
          <a href="#">
            <i class="fa fa-circle-o"></i> CPEs
          </a>
          </router-link>
        </li>
        
        <li>
          <router-link tag="li" class="pageLink" to="/nvts">
          <a href="#">
            <i class="fa fa-circle-o"></i> NVTs
          </a>
          </router-link>
        </li>
      </ul>
    </li>
    <!-- <li class="header">PAGES</li>
    <router-link tag="li" class="pageLink" to="/login">
      <a>
        <i class="fa fa-circle-o text-yellow"></i>
        <span class="page"> Login</span>
      </a>
    </router-link> -->
  </ul>
</template>
<script>
export default {
  name: 'SidebarMenu'
}
</script>
<style>
/* override default */
.sidebar-menu > li > a {
  padding: 12px 15px 12px 15px;
}

.sidebar-menu li.active > a > .fa-angle-left,
.sidebar-menu li.active > a > .pull-right-container > .fa-angle-left {
  animation-name: rotate;
  animation-duration: 0.2s;
  animation-fill-mode: forwards;
} 
.database-menu li{
  padding: 2.5px 5px 2.5px 10px;
}
.database-menu li a i {
  margin-right: 5px;
}
.database-menu li a:hover{
  color: aliceblue;
}
.treeview-title {
  z-index: 1;
}

@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(-90deg);
  }
}
</style>
