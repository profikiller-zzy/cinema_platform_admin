<template>
    <div class="login_background">
        <div class="login_container">
            <div class="login_title">
                管理员登录
            </div>
            <div class="login_input">
                <a-input
                        v-model:value="admin.user_name"
                        placeholder="管理员用户名"
                >
                    <template #prefix>
                        <i class="iconfont icon-guanliyuandenglu"></i>
                    </template>
                </a-input>
                <a-input-password
                        v-model:value="admin.password"
                        placeholder="请输入密码"
                >
                    <template #prefix>
                        <i class="iconfont icon-mima"></i>
                    </template>
                </a-input-password>
            </div>
            <div class="login_button">
                <a-button type="primary" @click="Login">登录</a-button>
            </div>
        </div>
    </div>
</template>

<script setup>
import {reactive} from "vue";
import {message} from 'ant-design-vue'
import {useRouter} from "vue-router";
import {useRoute} from "vue-router";
import {adminLogin} from "@/api/login_api.js";
import {parseToken} from "@/utils/jwt.js";
import {AdminInfoStore} from "@/stores/admin_info.js";


const router = useRouter()
const route = useRoute()
const admin = reactive({
    user_name: "",
    password: "",
})

async function Login() {
  if (admin.user_name.trim() === "") {
      message.error('请输入用户名')
      return
  }
  if (admin.password.trim() === "") {
      message.error('请输入密码')
      return
  }

  let res = await adminLogin(admin)
  if (res.code) {
        message.error(res.msg)
        return
    }
  message.success("登录成功")

  // res.data就是用户的ID
  let store = AdminInfoStore()
  let adminInfo = parseToken(res.data)
  adminInfo.token = res.data
  store.setAdminInfo(adminInfo)

  // 成功后跳转
  const redirect_url = route.query.redirect_url
  if (redirect_url === undefined) {
    // 跳转至home页面
      setTimeout(() => {
          router.push({name: "adminHome"})
      }, 1000)
      return
  }
  setTimeout(() => {
      router.push({path: redirect_url})
  }, 1000)

}
</script>


<style lang="scss">
.login_background {
  height: 100vh;
  width: 100%;
  background-image: url("../assets/background/background.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;

  .login_container {
    position: fixed;
    top: 0;
    right: 0;
    width: 350px;
    height: 100vh;
    background-color: rgba(255, 255, 255, 0.7);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .login_title {
      font-size: 25px;
      font-weight: 600;
      color: #5955ff;
      margin-bottom: 15px;
    }

    .login_input {
      width: 75%;

      .ant-input {
        margin-left: 10px;
      }

      .ant-input-affix-wrapper {
        margin-bottom: 10px;
      }
    }

    .login_button {
      width: 75%;
      margin-bottom: 10px;

      .ant-btn {
        width: 100%;
      }
    }

  }
}

</style>