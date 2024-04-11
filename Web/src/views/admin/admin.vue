<template>
  <div class="admin">
    <aside>
    </aside>

    <div class="admin_main">
      <header></header>

      <div class="tabs"></div>
        <div>
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
    width: calc(100vh - 240px);

    header {
      padding: 0 20px;
      height: 60px;
      background-color: white;
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