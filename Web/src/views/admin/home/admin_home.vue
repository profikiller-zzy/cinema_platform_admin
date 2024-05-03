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
  <div class="home_data_chart_box"></div>
  <div class="home_data_chart_order"></div>
</template>

<script setup>
import { reactive, onMounted } from "vue";
import {dataStatisticsApi, dataStatisticsInSevenDaysApi} from "@/api/data_statistics_api.js";
import {message} from "ant-design-vue";
import * as echarts from 'echarts';

const iconList = reactive([
  "icon-xiaoshoue",
  "icon-shoupiao",
  "icon-changci-",
  "icon-yonghutongji",
  "icon-pingluntongji"
]);

const dataState = reactive({
  box_office_day: 0,
  box_office_week: 0,
  box_office_month: 0,
  total_users: 0,
  total_orders: 0
});

const data = reactive({
  data_preview_list: [
    { label: "本日票房统计", value: "" },
    { label: "本周票房统计", value: "" },
    { label: "本月票房统计", value: "" },
    { label: "用户总数", value: "" },
    { label: "订单总数", value: "" }
  ]
});

// 获取数据并更新dataState和data对象
async function getDataPreview() {
  try {
    const res = await dataStatisticsApi();
    if (res.code) {
      message.error(res.msg);
    } else {
      message.success("获取统计数据成功");
      const { box_office_day, box_office_week, box_office_month, total_users, total_orders } = res.data;
      data.data_preview_list[0].value = box_office_day.toFixed(2) + "元";
      data.data_preview_list[1].value = box_office_week.toFixed(2) + "元";
      data.data_preview_list[2].value = box_office_month.toFixed(2) + "元";
      data.data_preview_list[3].value = total_users;
      data.data_preview_list[4].value = total_orders;
    }
  } catch (error) {
    console.error("Failed to fetch data:", error);
    message.error("Failed to fetch data");
  }
}

// 渲染折线图
async function renderBoxChart() {
  const chartContainer = document.querySelector('.home_data_chart_box');
  const chart = echarts.init(chartContainer);

  // 获取最近七天票房数据
  try {
    const res = await dataStatisticsInSevenDaysApi(); // 假设这个是获取最近七天票房数据的接口函数
    if (res.code) {
      message.error(res.msg);
    } else {
      let dateList = res.data.date; // 假设接口返回的数据包括日期列表和票房数据列表
      let boxOfficeList = res.data.box_office_per_day

      const option = {
        title: {
          text: '最近七日单日票房统计'
        },
        tooltip: {
          trigger: 'axis'
        },
        xAxis: {
          type: 'category',
          data: dateList // 使用获取到的日期列表
        },
        yAxis: {
          type: 'value'
        },
        series: [{
          data: boxOfficeList, // 使用获取到的票房数据列表
          type: 'line'
        }]
      };

      chart.setOption(option);
    }
  } catch (error) {
    console.error("Failed to fetch chart data:", error);
    message.error("Failed to fetch chart data");
  }
}

async function renderOrderChart() {
  const chartContainer = document.querySelector('.home_data_chart_order');
  const chart = echarts.init(chartContainer);

  // 获取最近七天票房数据
  try {
    const res = await dataStatisticsInSevenDaysApi(); // 假设这个是获取最近七天票房数据的接口函数
    if (res.code) {
      message.error(res.msg);
    } else {
      let dateList = res.data.date; // 假设接口返回的数据包括日期列表和票房数据列表
      let ordersPerDayList = res.data.orders_per_day

      const option = {
        title: {
          text: '最近七日单日订单总数统计'
        },
        tooltip: {
          trigger: 'axis'
        },
        xAxis: {
          type: 'category',
          data: dateList // 使用获取到的日期列表
        },
        yAxis: {
          type: 'value'
        },
        series: [{
          data: ordersPerDayList, // 使用获取到的票房数据列表
          type: 'line'
        }]
      };

      chart.setOption(option);
    }
  } catch (error) {
    console.error("Failed to fetch chart data:", error);
    message.error("Failed to fetch chart data");
  }
}

// 在组件挂载后获取数据
onMounted(() => {
  getDataPreview();
  renderBoxChart();
  renderOrderChart();
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

.home_data_chart_box {
  width: 100%;
  height: 400px; /* 根据需要设置合适的高度 */
}

.home_data_chart_order {
  width: 100%;
  height: 400px; /* 根据需要设置合适的高度 */
}
</style>
