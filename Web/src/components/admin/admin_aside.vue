<template>
  <div class="admin_aside">
    <div class="admin_aside_header">
      <div class="admin_aside_logo">
        <img src="@/assets/logo.png">
      </div>
      <div class="admin_aside_title">管理员后台系统</div>
    </div>
    <div class="admin_aside_body">
       <a-menu
        v-model:selectedKeys="selectedKeys"
        style="width: 256px"
        mode="inline"
      >
        <a-menu-item key="admin_home">
          <template #icon>
            <i class="iconfont icon-home"></i>
          </template>
          管理员主页
        </a-menu-item>
        <a-sub-menu key="sub1">
          <template #icon>
            <i class="iconfont icon-yonghuguanli"></i>
          </template>
          <template #title>用户管理</template>
          <a-menu-item key="3">添加用户</a-menu-item>
          <a-menu-item key="3">用户列表</a-menu-item>
          <a-menu-item key="4">用户查询</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="sub2">
          <template #icon>
            <i class="iconfont icon-dianyingziyuan"></i>
          </template>
          <template #title>电影信息管理</template>
          <a-menu-item key="7">预告发布</a-menu-item>
          <a-menu-item key="8">电影列表</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </div>
  </div>
</template>

<script>
import "@/assets/css/iconfont.css"
export default {
  data() {
    return {
      menuItems: [
        { id: 1, label: 'Home', path: '/' },
        { id: 2, label: 'About', path: '/about' },
        { id: 3, label: 'Services', path: '/services' },
        { id: 4, label: 'Contact', path: '/contact' },
      ]
    };
  }
};
</script>

<style lang="scss">
.admin_aside {
  width: 240px;
  height: 100vh;
  background-color: #f0f0f0;

  .admin_aside_header {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    .admin_aside_logo {
      img {
        width: 180px;
        height: 126px;
      }

      &::after {
        content: "logo图片";
        display: flex;
        justify-content: center;
        align-content: center;
      }
    }

    .admin_aside_title {
      margin: 20px;
      font-size: 22px;

    }
  }
}


</style>