<template>
  <div class="home_data_preview">
    <div class="home_data_preview_item">
      <div class="icon">
        <i class="iconfont icon-tianjiayonghu"></i>
      </div>
      <div class="text">
        <div class="data_title">今日票房总数</div>
        <div class="data_sum">20</div>
      </div>
    </div>

  </div>
</template>

<script setup>
import {reactive} from "vue";
const iconList = reactive({

})

const data = reactive({
  data_preview_list:[
    {
      label: "今日销售额",
      value: 21,
    },
    {
      label: "今日售票数量",
      value: 21,
    },
    {
      label: "今日场次数量",
      value: 21,
    },
    {
      label: "用户总数",
      value: 21,
    },
    {
      label: "评论总数",
      value: 21,
    },
  ]
})
</script>

<style scoped lang="scss">
.home_data_preview {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  gap: 20px; /* 项目之间的水平间距 */

  .home_data_preview_item {
    display: flex;
    align-items: center;
    background-color: white;
    border-radius: 5px;
    border: 1px solid #afffe9;
    padding: 20px;

    .icon {
      margin-right: 20px;
      font-size: 40px;
      color: #1890ff; /* 图标颜色，根据需要调整 */
    }

    .text {
      flex: 1;
      display: flex;
      flex-direction: column;
      justify-content: center; /* 垂直居中 */
      font-size: 14px;
      text-align: center; /* 水平居中 */

      .title {
        font-weight: bold;
        margin-bottom: 6px;

        .data_title {
          color: white;
        }

        .data_sum {
          font-size: 18px;
        }
      }

      .value {
        font-size: 20px;
      }
    }
  }
}
</style>