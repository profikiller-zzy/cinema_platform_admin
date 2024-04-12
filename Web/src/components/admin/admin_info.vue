<template>

  <div class="admin_info">
    <div class="avatar">
      <img src="@/assets/img.png" v-if="props.isAvatar">
    </div>
    <div class="drop_menu">
      <a-dropdown>
        <a class="ant-dropdown-link" @click.prevent>
          管理员菜单
          <i class="fa-angle-down"></i>
        </a>
        <template #overlay>
          <a-menu @click="onClick">
            <a-menu-item key="user">
              <a href="javascript:;">用户管理</a>
            </a-menu-item>
            <a-menu-item key="approval">
              <a href="javascript:;">审批管理</a>
            </a-menu-item>
            <a-menu-item key="movie">
              <a href="javascript:;">电影管理</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";

const router = useRouter()
const props = defineProps({
  isAvatar:{
    type:Boolean,
    default:false
  }
})
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

<style scoped lang="scss">
.admin_info {
  display: flex;
}

</style>