<template>
  <div class="choice_box">
    <a-input class="input_box" v-model:value="TimePeriodRequest.movie_id" placeholder="请输入你想查询的电影ID" />
    <a-date-picker class="date_picker"
        v-model:value="TimePeriodRequest.start_time"
        :placeholder="TimePeriodRequest.start_time ? TimePeriodRequest.start_time.format('YYYY-MM-DD HH:mm') : '选择起始时间'"
        @change="handleStartChange"
    />
    <a-date-picker class="date_picker"
        v-model:value="TimePeriodRequest.end_time"
        :placeholder="TimePeriodRequest.end_time ? TimePeriodRequest.end_time.format('YYYY-MM-DD HH:mm') : '选择结束时间'"
        @change="handleEndChange"
    />
    <a-button type="primary" @click="getMovieBoxOffice">查询</a-button>
  </div>
  <div class="movie_box_office_chart" ref="chart"></div>
  <div class="total_box_office">
    <span v-if="data.visiblity">该电影在这段时间的总票房为 {{ data.totalBoxRevenue }} 元</span>
  </div>
</template>

<script setup>
import {reactive} from "vue";
import * as echarts from 'echarts';
import {movieBoxOfficeApi} from "@/api/movie_management_api.js";
import {message} from "ant-design-vue";
import {AdminInfoStore} from "@/stores/admin_info.js";

let TimePeriodRequest = reactive({
  movie_id: 0,
  start_time: "",
  end_time: "",
});

let data = reactive({
  totalBoxRevenue: 0,
  visiblity: false,
})

// handleStartChange 选择开始时间
function handleStartChange(value) {
  TimePeriodRequest.start_time = value;
}

// handleEndChange 选择结束时间
function handleEndChange(value) {
  TimePeriodRequest.end_time = value;
}

async function getMovieBoxOffice() {
  try {
    let store = AdminInfoStore()
    store.loadAdminInfo()
    const res = await movieBoxOfficeApi(TimePeriodRequest);
    if (res.code) {
      message.error(res.msg);
    } else {
      let dateList = res.data.date;
      let boxOfficePerDayList = res.data.box_office_per_day;
      let movieName = res.data.movie_name
      data.totalBoxRevenue = res.data.total_box_office;
      data.visiblity = true

      console.log(dateList)
      console.log(boxOfficePerDayList)
      // 绘制图表
      drawChart(dateList, boxOfficePerDayList, movieName);
    }
  } catch (error) {
    console.error("Failed to fetch chart data:", error);
    message.error("Failed to fetch chart data");
  }
}

// 绘制图表
function drawChart(dateList, boxOfficePerDayList, MovieName) {
  const chartContainer = document.querySelector('.movie_box_office_chart');
  const chart = echarts.init(chartContainer);

  console.log(dateList)
  console.log(boxOfficePerDayList)

  const option = {
    title: {
      text: MovieName + '票房查询结果'
    },
    tooltip: {
      trigger: 'axis'
    },
    xAxis: {
      type: 'category',
      data: dateList
    },
    yAxis: {
      type: 'value'
    },
    series: [{
      data: boxOfficePerDayList,
      type: 'line'
    }]
  };

  chart.setOption(option);
}
</script>

<style scoped lang="scss">
.choice_box {
  display: flex;
  justify-content: space-between;

  .input_box {
    width: 160px;
    flex: 1;
    margin-right: 10px;
  }
}

.movie_box_office_chart {
  width: 100%;
  height: 400px; /* 根据需要设置合适的高度 */
}
</style>
