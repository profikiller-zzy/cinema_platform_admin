<template>
    <div class="login_background">
        <div class="login_container">
            <div class="login_title">
                用户登录
            </div>
            <div class="login_input">
                <a-input
                        v-model:value="user.user_name"
                        placeholder="管理员用户名"
                >
                    <template #prefix>
                        <i class="iconfont icon-guanliyuandenglu"></i>
                    </template>
                </a-input>
                <a-input-password
                        v-model:value="user.password"
                        placeholder="请输入密码"
                >
                    <template #prefix>
                        <i class="iconfont icon-mima"></i>
                    </template>
                </a-input-password>
            </div>
            <div class="login_button">
                <a-button type="primary" @click="emailLogin">登录</a-button>
            </div>
        </div>
    </div>
</template>

<script setup>
import {reactive} from "vue";
import {message} from 'ant-design-vue'
import {useRouter} from "vue-router";
import {useRoute} from "vue-router";

const router = useRouter()
const route = useRoute()
const user = reactive({
    user_name: "",
    password: "",
})

async function emailLogin() {
    if (user.user_name.trim() === "") {
        message.error('请输入用户名')
        return
    }
    if (user.password.trim() === "") {
        message.error('请输入密码')
        return
    }
    console.log("登录成功")

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


    .login_other {
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;

      &::before {
        width: 30%;
        height: 1px;
        background-color: #252121;
        display: inline-block;
        content: "";
        margin-right: 5px;
      }

      &::after {
        width: 30%;
        height: 1px;
        background-color: #252121;
        display: inline-block;
        content: "";
        margin-left: 5px;
      }
    }

    .login_icons {
      .login_qq {
        margin-top: 10px;
        height: 40px;
        width: 40px;
        cursor: pointer;
      }
    }

  }
}

</style>