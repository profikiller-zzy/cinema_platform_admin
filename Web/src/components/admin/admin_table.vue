<template>
  <div class="container">
    <!-- search_box 搜索框 主要用于模糊匹配 -->
    <div class="search_box">
      <a-input-search
          v-model:value="page.key"
          :placeholder="props.likeTitle"
          enter-button
          @search="onSearch"
          style="width: 300px"
      />

      <slot name="filters">

      </slot>

      <div class="user_list_refresh">
        <a-button @click="Refresh"><i class="iconfont icon-shuaxin"></i></a-button>
      </div>
    </div>
    <!-- actions 主要是一些定义行为的按钮 -->
    <div class="actions">
      <slot name="add">
        <a-button type="primary" @click="data.AddModalVisible = true" v-if="props.isAdd">添加</a-button>
      </slot>

      <slot name="batchRemove">
        <a-popconfirm
            title="你确定要批量删除么？"
            ok-text="删除"
            cancel-text="取消"
            @confirm="batchRemove"
        >
          <a-button type="danger" v-if="props.isDelete && data.selectedRowKeys.length">批量删除</a-button>
        </a-popconfirm>
      </slot>
    </div>
    <!-- tables 用于展示用户数据的列表 -->
    <div class="tables">
      <a-spin :spinning="data.spinning" tip="加载中" :delay="300">
        <a-table
            :columns="props.columns"
            :data-source="data.list"
            :pagination="false"
            :total="data.count"
            :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange}"
            rowKey="id">
          <template #bodyCell="{ column, record }">
            <slot name="cell" v-bind="{ column, record }">
              <template v-if="column.key === 'created_at'">
                <span>{{ dateTransition(record.created_at, 0) }}</span>
              </template>
              <template v-if="column.key === 'action'">
                <slot name="edit" v-bind="{column, record}">
                  <a-button type="primary" v-if="isEdit">编辑</a-button>
                </slot>

                <slot name="delete" v-bind="{column, record}">
                  <a-popconfirm
                      title="是否确认删除？"
                      ok-text="删除"
                      cancel-text="取消"
                      @confirm="sigleRemove(record.id)"
                      v-if="props.isDelete"
                  >
                    <a-button class="user_action delete" type="danger">删除</a-button>
                  </a-popconfirm>
                </slot>
              </template>
            </slot>
          </template>
        </a-table>
      </a-spin>
      </div>
    <!-- pages 用于展示分页 -->
    <div class="pages">
      <a-pagination
          v-model:current="page.page_num"
          v-model:page-size="page.page_size"
          @change="pageChange"
          :showSizeChanger="false"
          :total="data.count"
          :show-total="total => `共 ${total} 条`"
      />
    </div>
  </div>
</template>

<script setup>
import {baseApiList, baseRemoveApi} from "@/api/base_api.js";
import {reactive} from "vue";
import {dateTransition} from "@/utils/dateTransition.js";
import {message} from "ant-design-vue";
import {AdminInfoStore} from "@/stores/admin_info.js";

const emits = defineEmits(["delete"])
const props = defineProps({
  columns: {
    type: Array
  },
  baseUrl: {
    type: String
  },
  isAdd: {
    type: Boolean,
    default: false
  },
  isEdit: {
    type: Boolean,
    default: false
  },
  isDelete: {
    type: Boolean,
    default: false
  },
  filters: {
    type: Array,
  },
  // 用于模糊搜索的名字
  likeTitle: {
    type: String,
    default: "模糊搜索",
  }
})


const page = reactive({
  page_num: 1,
  page_size: 5,
  key: "",
})

const data = reactive({
  list: [],
  selectedRowKeys: [], // 复选框选择的ID列表
  count: 0,
  spinning: true, // 默认是在加载中
})

// 选择复选框
function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
}

async function getDataList(params) {
  // 在这里获取用户列表并且将用户数据存入data.list中
  data.spinning = true
  let store = AdminInfoStore()
  store.loadAdminInfo()
  let res = await baseApiList(props.baseUrl, params)
  data.list = res.data.data_list
  data.count = res.data.count
  data.spinning = false
}
getDataList(page)

function pageChange() {
  getDataList(page)
}

async function batchRemove() {
  try {
    let res = await baseRemoveApi(props.baseUrl, data.selectedRowKeys);
    if (res.code === 0) {
      const failedDeletes = res.data.filter(item => !item.is_success);
      if (failedDeletes.length > 0) {
        // 输出失败删除的用户信息及原因
        // 显示错误提示
        message.error("删除失败");
        failedDeletes.forEach(item => {
          message.error(item.msg);
        });
      } else {
        // 所有用户删除成功的情况
        message.success(`删除成功`);
      }
    } else {
      // 显示错误提示
      message.error(res.msg);
    }
    getDataList(page);
    emits("delete", data.selectedRowKeys)
    data.selectedRowKeys = []
  } catch (error) {
    // 处理请求失败的情况
    console.error("Error:", error);
    // 显示错误提示
    message.error("请求失败");
  }
}

async function sigleRemove(id) {
  try {
    let res = await baseRemoveApi(props.baseUrl, [id]);
    if (res.code === 0) {
      res.data.forEach(item => {
        if (item.is_success) {
          // 打印成功消息
          message.success(item.msg);
        } else {
          // 打印失败消息
          message.error(item.msg);
        }
      });
    } else {
      // 显示错误提示
      message.error(res.msg);
    }
    getDataList(page);
    emits("delete", data.selectedRowKeys);
    data.selectedRowKeys = [];
  } catch (error) {
    // 处理请求失败的情况
    console.error("Error:", error);
    // 显示错误提示
    message.error("请求失败");
  }
}

// 刷新用户列表页面
function Refresh() {
  message.success("刷新成功")
  getDataList(page)
}

function onSearch() {
  page.key = page.key.trim()
  page.page_num = 1
  getDataList(page)
}

function ExportList(params) {
  page.page_num = 1
  Object.assign(page, params)
  getDataList(page)
}

defineExpose({
  ExportList,
  Refresh
})

</script>

<style lang="scss">
.container {
  background-color: white;

  .search_box {
    padding: 10px;
    border-bottom: 1px solid #e8e6ff;
    position: relative;

    .user_list_select {
      margin-left: 10px;
    }

    .user_list_refresh {
      position: absolute;
      right: 10px;
      top: 10px;
    }
  }

  .actions {
    padding: 10px;

    .ant-btn {
      margin-left: 10px;
    }
  }

  .tables {
    padding: 0 10px 10px 10px;

    .user_avatar {
      width: 50px;
      height: 50px;
      border-radius: 50%;
    }

    .user_action.update {
      margin-right: 10px;
    }
  }

  .pages {
    display: flex;
    justify-content: center;
    padding: 10px;
  }
}
</style>