<template>
  <div class="admin">
    <Admin_aside></Admin_aside>

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
            <admin_info></admin_info>
          </div>
        </div>
      </header>

      <div class="tabs">

      </div>

      <main>
        <div class="router_view">
          <router-view></router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";
import Admin_aside from "@/components/admin/admin_aside.vue";
import Admin_info from "@/components/admin/admin_info.vue";

const route = useRouter()

function onClick({key}) {
  if (key === "user"){
    router.push({
      name:"login"
    })
  }
  else if (key === "approval") {
    console.log("approval")
  }
}
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

        .router_view {
          background-color: white;
        }
      }

  }
}
</style>