<template>
  <div>
    <a-modal v-model:visible="data.AddModalVisible" title="添加用户" @ok="AddUser">
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
    <a-modal v-model:visible="data.EditModalVisible" title="编辑用户" @ok="EditUser">
      <a-form
        :model="EditFormState"
        name="basic"
        ref="formRef"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 20 }"
        autocomplete="off"
      >
        <a-form-item label="权限" name="role" :rules="[{ required: true, message: '请选择权限!' }]">
            <a-select
                    ref="select"
                    v-model:value="EditFormState.user_type"
                    style="width: 200px"
                    :options="roleOptions"
            >
            </a-select>
        </a-form-item>
        <a-form-item has-feedback label="年龄" name="age"
                     :rules="[{ required: false, message: '请输入年龄', trigger: 'blur' }]">
            <a-input v-model:value="EditFormState.age" placeholder="年龄"/>
        </a-form-item>
        <a-form-item has-feedback label="邮箱" name="user_name"
                   :rules="[{ required: false, message: '请输入邮箱!', trigger: 'blur' }]">
          <a-input v-model:value="EditFormState.email" placeholder="邮箱"/>
        </a-form-item>
        <a-form-item has-feedback label="电话号码" name="user_name"
                   :rules="[{ required: false, message: '请输入电话号码!', trigger: 'blur' }]">
          <a-input v-model:value="EditFormState.tel" placeholder="电话号码"/>
        </a-form-item>
      </a-form>
    </a-modal>
    <AdminTable
        @delete="userRemove"
        :columns="data.columns"
        base-url="/api/users"
        ref="adminTable"
        like-title="搜索用户名"
        is-delete is-add>
      <template #add>
        <a-button type="primary" @click="data.AddModalVisible = true">添加</a-button>
      </template>
      <template #edit="{record}">
        <a-button class="user_action update" @click="EditModel(record)" type="primary">编辑</a-button>
      </template>
      <template #cell="{ column, record }">
        <template v-if="column.key === 'avatar_url'">
          <img class="user_avatar" :src="record.avatar_url">
        </template>
      </template>
    </AdminTable>
  </div>
</template>

<script setup>
import AdminTable from "@/components/admin/admin_table.vue";
import {reactive, ref} from "vue";
import {userListApi, userCreateApi, userEditApi} from "@/api/user_management_api.js";
import {message} from "ant-design-vue";
import {AdminInfoStore} from "@/stores/admin_info.js";

// 对话框中规则验证
const formRef = ref({})
const adminTable = ref(null)
// 权限映射
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
// 密码校验，验证两次输入的密码是否一致
let validatePassword = async (_rule, value) => {
  if (value === "") {
    return Promise.reject("");
  } else if (value !== formState.password) {
    return Promise.reject("两次密码不一致!");
  } else {
    return Promise.resolve();
  }
}
// 用于置空formState
const _formState = {
  user_name: "",
  password: "",
  re_password: "",
  role: undefined,
  age: "",
  email: "",
  tel: "",
}
// 表单信息
const formState = reactive({
  user_name: "",
  password: "",
  re_password: "",
  role: undefined,
  age: "",
  email: "",
  tel: "",
})
const EditFormState = reactive({
  user_type: undefined,
  age: "",
  tel: "",
  email: "",
  user_id: 23
})
const filter = ref(undefined)
function OnFilter() {
  adminTable.value.ExportList({
    user_type: filter.value
  })
}

const data = reactive({
  AddModalVisible: false, // 对话框是否可见
  EditModalVisible: false,
  columns: [
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
})

function EditModel(record) {
  data.EditModalVisible = true
  EditFormState.user_id = record.id
  EditFormState.user_type = record.role_id
  EditFormState.age = record.age
  EditFormState.email = record.email
  EditFormState.tel = record.tel
}

function userRemove(user_id) {
  console.log(user_id)
}

// 创建用户
async function AddUser() {
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
    // await getUserList()
    // 置空
    Object.assign(formState, _formState)
    // // 清除验证规则
    // formRef.value.clearValidate()
    // GvbList.value.ExposeList()
  } catch (e) {
    console.log(e)
  }
}

// 编辑用户
async function EditUser() {
  let res = await userEditApi(EditFormState)
  if (res.code) {
    message.error(res.msg)
    return
  } else {
    message.success(res.msg)
    // getUserList()
  }
  data.EditModalVisible = false
}
// getUserList()

</script>

<style lang="scss">

</style>