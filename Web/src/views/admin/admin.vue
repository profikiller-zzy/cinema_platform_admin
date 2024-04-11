<template>
  <div class="admin">
    <aside></aside>

    <div class="admin_main">
      <header>
        <div class="header_left">
          <a-breadcrumb>
            <a-breadcrumb-item>后台首页</a-breadcrumb-item>
            <a-breadcrumb-item><a href="#">管理员中心</a></a-breadcrumb-item>
            <a-breadcrumb-item>用户管理</a-breadcrumb-item>
          </a-breadcrumb>
        </div>

        <div class="header_right">
          <div class="actions">
            <div class="icon_actions">
              <i class="fa fa-home"></i>
            </div>
            <div class="avatar">
              <img src="@/assets/img.png" alt="">
            </div>
            <div class="drop_menu">
              <a-dropdown>
                <a class="ant-dropdown-link" @click.prevent>
                  管理员菜单
                  <i class="fa-angle-down"></i>
                </a>
                <template #overlay>
                  <a-menu>
                    <a-menu-item>
                      <a href="javascript:;">1st menu item</a>
                    </a-menu-item>
                    <a-menu-item>
                      <a href="javascript:;">2nd menu item</a>
                    </a-menu-item>
                    <a-menu-item>
                      <a href="javascript:;">3rd menu item</a>
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
          </div>
        </div>
      </header>

      <div class="tabs">
        <i class="iconfont icon-yonghu"></i>
        <i class="fa fa-camera-retro fa-lg"></i>
      </div>

      <main>
        <div class="router-view">
          <router-view/>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
</script>

<style lang="scss">
.admin{
  width: 100%;
  display: flex;
  flex-direction: row;

  aside{
    width: 240px;
    height: 100vh;
    background-color: #2b3539;
  }

  .admin_main{
    width: calc(100% - 240px);

    header {
      padding: 0 20px;
      height: 60px;
      background-color: white;
      display: flex;
      justify-content: space-between;
      align-items: center; // 排列在当前列的中间位置

      .header_left {
      }

      .header_right {
        display: flex;
        align-content: center; /* 垂直居中 */

        .actions {
          display: flex;
          align-content: center;

          .icon_actions {
            margin-right: 20px;

            i {
              margin-left: 10px;
              cursor: pointer;
              font-size: 16px;
            }

            i:hover {
              color: #2184fc;
            }
          }

          .avatar {
            img {
              width: 40px;
              height: 40px;
              border-radius: 50%;
              margin-right: 20px;
            }

            .drop_menu {
              margin-left: 10px;
            }
          }
        }
      }
    }

    .tabs {
      height: 30px;
      border: 1px solid #f0eeee;
    }

      main {
        background-color: white;
        height: calc(100vh - 90px);
        padding: 20px;

        .router-view {
          background-color: white;
        }
      }

  }
}
</style>