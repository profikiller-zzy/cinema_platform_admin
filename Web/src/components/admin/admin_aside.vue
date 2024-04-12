<template>
  <div class="admin_aside">
    <div class="admin_aside_header">
      <div class="admin_aside_logo">
        <img src="@/assets/logo.png">
      </div>
      <div class="admin_aside_title">管理员后台系统</div>
    </div>
    <div class="admin_aside_body">

    </div>
  </div>
</template>

<script>
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