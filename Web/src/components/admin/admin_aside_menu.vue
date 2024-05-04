<template>
  <a-menu
    v-model:selectedKeys="selectedKeys"
    style="width: 256px"
    mode="inline"
    @click="goto"
  >
    <!-- 使用v-for循环渲染menuList中的数据 -->
    <template v-for="menu in data.menuList" :key="menu.name">
      <!-- 判断是否为子菜单 -->
      <a-sub-menu v-if="menu.children && menu.children.length > 0" :key="menu.name">
        <!-- 渲染子菜单的图标和标题 -->
        <template #title>
          <i :class="menu.icon"></i>
          <span>{{ menu.title }}</span>
        </template>
        <!-- 循环渲染子菜单中的菜单项 -->
        <a-menu-item v-for="child in menu.children" :key="child.name">
          <template #icon><i :class="child.icon"></i></template>
          <span>{{ child.title }}</span>
        </a-menu-item>
      </a-sub-menu>
      <!-- 如果不是子菜单，则直接渲染菜单项 -->
      <a-menu-item :key="menu.name" v-else>
        <template #icon><i :class="menu.icon"></i></template>
        <span>{{ menu.title }}</span>
      </a-menu-item>
    </template>
  </a-menu>
</template>

<script setup>

import "@/assets/css/iconfont.css"
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const selectedKeys = ref(['1'])
const data = reactive({
  menuList:[
    {
      id: 1,
      icon: "iconfont icon-home", // 图标名称
      title: "管理员首页",          // 菜单标题
      name: "admin_home",               // 要跳转的路由的名称
      children: [
        {
          id: 11,
          icon: "iconfont icon-jia",
          title: "管理员首页",
          name: "admin_home",
        },
        {
          id: 8,
          icon: "iconfont icon-TOP",
          title: "电影票房排行榜",
          name: "box_office_ranking",
        },
        {
          id: 9,
          icon: "iconfont icon-TOP",
          title: "影院票房排行榜",
          name: "cinema_box_office_ranking",
        }
      ]
    },
    {
      id: 2,
      icon: "iconfont icon-yonghuguanli", // 图标名称
      title: "用户管理",          // 菜单标题
      name: "",                  // 用户管理不会跳转路由，所以这里直接为空
      children: [
        {
          id: 3,
          icon: "iconfont icon-yonghuliebiao",
          title: "用户列表",
          name: "user_list",
        },
      ]
    },
    {
      id: 4,
      icon: "iconfont icon-dianyingziyuan", // 图标名称
      title: "电影管理",          // 菜单标题
      name: "",                  // 用户管理不会跳转路由，所以这里直接为空
      children: [
        {
          id: 5,
          icon: "iconfont icon-liebiao",
          title: "电影列表",
          name: "movie_list",
        },
        {
          id: 10,
          icon: "iconfont icon-p",
          title: "电影票房信息查看",
          name: "movie_box_office",
        },
      ]
    },
    {
      id: 6,
      icon: "iconfont icon-yingyuanguanli", // 图标名称
      title: "影院管理",          // 菜单标题
      name: "",                  // 用户管理不会跳转路由，所以这里直接为空
      children: [
        {
          id: 7,
          icon: "iconfont icon-liebiao",
          title: "影院列表",
          name: "cinema_list",
        },
      ]
    },
  ]
})

function goto(event) {
  router.push({
    name: event.key
  })
  console.log(event)
}

</script>


<style scoped lang="scss">

</style>