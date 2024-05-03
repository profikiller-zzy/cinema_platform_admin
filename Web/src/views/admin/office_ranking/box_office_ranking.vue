<template>
  <div class="total-box">
    <span id="year-box">最近一周电影票房总榜前100</span>
  </div>
  <div class="rank-title">
    <ul class="row">
      <li class="col0">排名</li>
      <li class="col1">片名</li>
      <li class="col2">最近一周票房(元)</li>
    </ul>
  </div>
  <div class="rank-list">
    <div v-for="(item, index) in data.rank_list" :key="index">
      <ul class="row">
        <li class="col0">{{ item.rank }}</li>
        <li class="col1">
            <p class="first-line">{{ item.movie_name }}</p>
            <p class="second-line">上映时间: {{ item.release_date }}</p>
        </li>
        <li class="col2">{{ item.box_office.toFixed(2) }}</li>
      </ul>
      <!-- 分界线 -->
      <div class="divider" v-if="index < data.rank_list.length - 1"></div>
    </div>
  </div>
</template>

<script setup>
import {reactive} from "vue";
import {movieRankApi} from "@/api/movie_management_api.js";
import {AdminInfoStore} from "@/stores/admin_info.js";

const data = reactive({
  rank_list: [],
})

async function getRankList() {
  let store = AdminInfoStore()
  store.loadAdminInfo()
  let res = await movieRankApi()
  data.rank_list = res.data
}
getRankList()
</script>

<style scoped lang="scss">
.total-box {
  color: #333;
  font-size: 16px;
  line-height: 1.4;
  padding: .24rem;
  background-color: #ff7f7f;
  text-align: center;
}

.rank-title {
  width: 100%;
  background: #ff7f7f;
  color: #fff;

  .row {
    padding: 0 .24rem;
    font-size: 16px;
    display: flex;

    .col0, .col1, .col2 {
      flex: 1; /* 平均分配剩余空间给三列 */
    }

    .col1 {
      display: flex;
      flex-direction: column;
      justify-content: center; /* 垂直居中排列 */

      .first-line {
        font-size: 14px; /* 调整第一行字体大小 */
      }

      .second-line {
        font-size: 12px; /* 调整第二行字体大小 */
        color: #999; /* 使用灰色字体颜色 */
      }
    }
  }
}

.rank-list {
  width: 100%;
  background: #40609F;
  color: #fff;
  padding: 0.24rem;

  .row {
    padding: 0 .24rem;
    font-size: 14px; /* 调整字体大小 */
    display: flex;

    .col0, .col1, .col2 {
      flex: 1; /* 平均分配剩余空间给三列 */
      padding: 0.2rem; /* 添加适当的内边距 */
    }
  }

  /* 分界线样式 */
  .divider {
    width: 100%;
    height: 1px;
    background-color: #999; /* 分界线颜色 */
    margin: 0.5rem 0; /* 控制分界线上下间距 */
  }
}
</style>