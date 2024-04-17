<template>
  <div class="container">

    <a-modal v-model:visible="data.AddModalVisible" title="添加用户" @ok="handleOk">
            <a-form
              :model="formState"
              name="basic"
              ref="formRef"
              :label-col="{ span: 4 }"
              :wrapper-col="{ span: 20 }"
              autocomplete="off"
            >
              <a-form-item has-feedback label="用户名" name="user_name"
                         :rules="[{ required: true, message: '请输入用户名!', trigger: 'blur' }]">
                <a-input v-model:value="formState.user_name" placeholder="用户名"/>
              </a-form-item>
              <a-form-item has-feedback label="密码" name="password"
                         :rules="[{ required: true, message: '请输入密码!', trigger: 'blur' }]">
                <a-input-password v-model:value="formState.password" placeholder="密码"/>
              </a-form-item>
              <a-form-item has-feedback label="确认密码" name="re_password"
                         :rules="[{ required: true, message: '请再次输入密码!' },
                         {validator: validatePassword, trigger: 'change'}]">
                <a-input-password v-model:value="formState.re_password" placeholder="确认密码"/>
              </a-form-item>
              <a-form-item label="权限" name="role" :rules="[{ required: true, message: '请选择权限!' }]">
                  <a-select
                          ref="select"
                          v-model:value="formState.role"
                          style="width: 200px"
                          :options="roleOptions"
                  >
                  </a-select>
              </a-form-item>
              <a-form-item has-feedback label="年龄" name="age"
                           :rules="[{ required: false, message: '请输入年龄', trigger: 'blur' }]">
                  <a-input v-model:value="formState.age" placeholder="年龄"/>
              </a-form-item>
              <a-form-item has-feedback label="邮箱" name="user_name"
                         :rules="[{ required: false, message: '请输入邮箱!', trigger: 'blur' }]">
                <a-input v-model:value="formState.email" placeholder="邮箱"/>
              </a-form-item>
              <a-form-item has-feedback label="电话号码" name="user_name"
                         :rules="[{ required: false, message: '请输入电话号码!', trigger: 'blur' }]">
                <a-input v-model:value="formState.tel" placeholder="电话号码"/>
              </a-form-item>
            </a-form>
        </a-modal>

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
      <a-button type="primary" @click="data.AddModalVisible = true">添加用户</a-button>
      <a-popconfirm
              title="你确定要批量删除这些用户么？"
              ok-text="删除"
              cancel-text="取消"
              @confirm="usersRemove"
            >
              <a-button type="danger" v-if="data.selectedRowKeys.length">删除用户</a-button>
      </a-popconfirm>
    </div>
    <!-- tables 用于展示用户数据的列表 -->
    <div class="tables">
      <a-table
          :columns="data.columns"
          :data-source="data.list"
          :pagination="false"
          :total="data.count"
          :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange}"
          rowKey="id">
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'avatar_url'">
            <img class="user_avatar" :src="record.avatar_url">
          </template>
          <template v-if="column.key === 'created_at'">
            <span>{{ dateTransition(record.created_at) }}</span>
          </template>
          <template v-if="column.key === 'action'">
            <a-button class="user_action update" type="primary">编辑</a-button>
            <a-popconfirm
              title="是否确认删除？"
              ok-text="删除"
              cancel-text="取消"
              @confirm="userRemove(record.id)"
            >
              <a-button class="user_action delete" type="danger">删除</a-button>
            </a-popconfirm>
          </template>
        </template>
      </a-table>
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
import {reactive, ref} from "vue";
import {dateTransition} from "@/utils/dateTransition.js";
import {userListApi, userCreateApi, userRemoveApi} from "@/api/user_management_api.js";
import {message} from "ant-design-vue";
import {AdminInfoStore} from "@/stores/admin_info.js";

// 对话框中规则验证
const formRef = ref({})
const idList = reactive({
    id_list: []
})

const page = reactive({
  page_num: 1,
  page_size: 5,
})

// 用于置空
const _formState = {
  user_name: "",
  password: "",
  re_password: "",
  role: undefined,
  age: "",
  email: "",
  tel: "",
}

const formState = reactive({
  user_name: "",
  password: "",
  re_password: "",
  role: undefined,
  age: "",
  email: "",
  tel: "",
})

const EditformState = reactive({
  user_id: "",
  role: ""
})

const roleOptions = [
    {
        value: 1,
        label: "系统管理员"
    },
    {
        value: 2,
        label: "影院用户"
    },
    {
        value: 3,
        label: "普通用户"
    },
]

const data = reactive({
  columns:[
    {title: '用户ID', dataIndex: 'id', key: 'id',},
    {title: '用户名', dataIndex: 'user_name', key: 'user_name',},
    {title: '头像', dataIndex: 'avatar_url', key: 'avatar_url',},
    {title: '年龄', dataIndex: 'age', key: 'age',},
    {title: '电话号码', dataIndex: 'tel', key: 'tel',},
    {title: '邮箱', dataIndex: 'email', key: 'email',},
    {title: '用户类型', dataIndex: 'user_type', key: 'user_type',},
    {title: '注册时间', dataIndex: 'created_at', key: 'created_at',},
    {title: '操作', dataIndex: 'action', key: 'action',},
  ],
  list:[
    {
      id: 22,
      created_at: "2024-03-15T21:30:51+08:00",
      updated_at: "2024-03-15T21:30:51+08:00",
      user_name: "profikiller_admin",
      avatar_url: "http://sbxaiblxv.hn-bkt.clouddn.com/cinema/2024-04-14_19:54:55.876_头像4.jpg",
      age: "18",
      tel: "136****0045",
      email: "e****@gmail.com",
      user_type: "管理员"
    }
  ],
  selectedRowKeys: [],
  count :0,
  AddModalVisible: false, // 对话框是否可见
})

function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
}

async function handleOk() {
  try {
    // 主动触发验证
    await formRef.value.validate()
    // console.log(formState)
    let res = await userCreateApi(formState)
    if (res.code) {
        message.error(res.msg)
        return
    }
    message.success(res.msg)
    data.AddModalVisible = false
    getUserList()
    // 置空
    Object.assign(formState, _formState)
    // // 清除验证规则
    // formRef.value.clearValidate()
    // GvbList.value.ExposeList()
  } catch (e) {
    console.log(e)
  }
}

// 密码校验
let validatePassword = async (_rule, value) => {
  if (value === "") {
    return Promise.reject("");
  } else if (value !== formState.password) {
      return Promise.reject("两次密码不一致!");
  } else {
      return Promise.resolve();
  }
}

async function getUserList() {
  // 在这里获取用户列表并且将用户数据存入data.list中
  let store = AdminInfoStore()
  store.loadAdminInfo()
  let res = await userListApi(page)
  data.list = res.data.data_list
  data.count = res.data.count
}
getUserList()

function pageChange() {
  getUserList()
}

async function userRemove(user_id) {
  try {
    let res = await userRemoveApi([user_id]);
    if (res.code === 0) {
      const failedDeletes = res.data.filter(item => !item.is_success);
      if (failedDeletes.length > 0) {
        // 输出失败删除的用户信息及原因
        // 显示错误提示
        message.error("用户删除失败");
        failedDeletes.forEach(item => {
          message.error(`用户 ${item.user_id} 删除失败: ${item.msg}`);
        });
      } else {
        // 所有用户删除成功的情况
        message.success(`用户删除成功`);
      }
    } else {
      // 显示错误提示
      message.error(res.msg || "删除用户失败");
    }
    await getUserList();
  } catch (error) {
    // 处理请求失败的情况
    console.error("Error:", error);
    // 显示错误提示
    message.error("请求失败");
  }
}

async function usersRemove() {
  try {
    const res = await userRemoveApi(data.selectedRowKeys);
    if (res.code === 0) {
      const failedDeletes = res.data.filter(item => !item.is_success);
      if (failedDeletes.length > 0) {
        // 输出失败删除的用户信息及原因
        // 显示错误提示
        message.error("部分用户删除失败");
        failedDeletes.forEach(item => {
          message.error(`用户 ${item.user_id} 删除失败: ${item.msg}`);
        });
      } else {
        // 所有用户删除成功的情况
        message.success("所有用户删除成功");
      }
    } else {
      // 显示错误提示
      message.error(res.msg || "删除用户失败");
    }
    await getUserList();
  } catch (error) {
    // 处理请求失败的情况
    console.error("Error:", error);
    // 显示错误提示
    message.error("请求失败");
  }
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