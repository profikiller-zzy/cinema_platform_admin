<template>
  <div class="container">
    <!-- search_box 搜索框 主要用于模糊匹配 -->
    <div class="search_box">
      <a-input-search
        v-model:value="value"
        placeholder="请输入你想搜索的用户"
        enter-button
        @search="onSearch"
        style="width: 300px"
      />
    </div>
    <!-- actions 主要是一些定义行为的按钮 -->
    <div class="actions">
      <a-button type="primary">添加用户</a-button>
      <a-button type="danger" @click="userDelete" v-if="data.selectedRowKeys.length">删除用户</a-button>
    </div>
    <!-- tables 用于展示用户数据的列表 -->
    <div class="tables">
      <a-table
          :columns="data.columns"
          :data-source="data.list"
          :pagination="false"
          :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange}"
          rowKey="id"
      />
    </div>
    <!-- pages 用于展示分页 -->
    <div class="pages">
       <a-pagination
          v-model:current="page.pageNum"
          v-model:page-size="page.pageSize"
          :showSizeChanger="false"
          :total="page.total"
          :show-total="total => `共 ${total} 条`"
       />
    </div>
  </div>
</template>

<script setup>
import {reactive} from "vue";

const page = reactive({
  pageNum: 1,
  pageSize: 10,
  total: 50
})

const data = reactive({
  columns:[
    {title: '用户ID', dataIndex: 'id', key: 'id',},
    {title: '用户名', dataIndex: 'user_name', key: 'user_name',},
    {title: '年龄', dataIndex: 'age', key: 'age',},
    {title: '电话号码', dataIndex: 'tel', key: 'tel',},
    {title: '邮箱', dataIndex: 'email', key: 'email',},
    {title: '用户类型', dataIndex: 'user_type', key: 'user_type',},
    {title: '注册时间', dataIndex: 'created_at', key: 'created_at',},
  ],
  list:[
    {
      id: 22,
      created_at: "2024-03-15T21:30:51+08:00",
      updated_at: "2024-03-15T21:30:51+08:00",
      user_name: "profikiller_admin",
      age: "18",
      tel: "136****0045",
      email: "e****@gmail.com",
      user_type: "管理员"
    }
  ],
  selectedRowKeys: [],
})

function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
}

function userDelete(selectedKeys) {
  console.log(data.selectedRowKeys)
}
</script>

<style lang="scss">
.container {
  background-color: white;

  .search_box {
    padding: 10px;
    border-bottom: 1px solid #e8e6ff;
  }

  .actions {
    padding: 10px;

    .ant-btn {
      margin-left: 10px;
    }
  }

  .tables {
    padding: 10px;
  }

  .pages {
    display: flex;
    justify-content: center;
    padding: 10px;
  }
}
</style>