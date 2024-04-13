<template>
  <!-- 数据预览部分 -->
  <div class="home_data_preview">
    <div class="home_data_preview_item" v-for="(item, index) in data.data_preview_list" :key="index">
      <div class="icon">
        <i :class="'iconfont ' + iconList[index]"></i>
      </div>
      <div class="text">
        <div class="data_title">{{ item.label }}</div>
        <div class="data_sum">{{ item.value }}</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from "vue";

const iconList = reactive([
  "icon-xiaoshoue",
  "icon-shoupiao",
  "icon-changci-",
  "icon-yonghutongji",
  "icon-pingluntongji"
]);

const data = reactive({
  data_preview_list: [
    {label: "今日销售额", value: 21},
    {label: "今日售票数量", value: 21},
    {label: "今日场次数量", value: 21},
    {label: "用户总数", value: 21},
    {label: "评论总数", value: 21}
  ]
});
</script>

<style lang="scss">
.home_data_preview {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 20px; /* 设置网格之间的间距 */
}

.home_data_preview_item {
  display: flex;
  background-color: white;
  border-radius: 5px;
  border: 1px solid #afffe9;
  padding: 20px;
}

.icon {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 60%;

  i {
     font-size: 32px; /* 调整图标的大小 */
  }
}


.text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 40%;
}

.data_title {
  font-size: 14px;
  font-weight: bold;
  margin-bottom: 6px;
}

.data_sum {
  font-size: 18px;
}
</style>
