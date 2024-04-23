<template>
  <div>
    <AdminTable
        @delete="userRemove"
        :columns="data.columns"
        base-url="/api/cinemas"
        ref="adminTable"
        like-title="搜索电影院名称">
    </AdminTable>
  </div>
</template>

<script setup>
import AdminTable from "@/components/admin/admin_table.vue";
import {reactive, ref} from "vue";
import { userCreateApi, userEditApi} from "@/api/user_management_api.js";
import {message} from "ant-design-vue";
import {AdminInfoStore} from "@/stores/admin_info.js";

// 对话框中规则验证
const formRef = ref({})
const adminTable = ref(null)

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
    {title: '影院ID', dataIndex: 'id', key: 'id',},
    {title: '注册时间', dataIndex: 'created_at', key: 'created_at',},
    {title: '电影院名称', dataIndex: 'name', key: 'name',},
    {title: '地址信息', dataIndex: 'address_information', key: 'address_information',},
    {title: '联系电话', dataIndex: 'tel_number', key: 'tel_number',},
    {title: '对应影院用户ID', dataIndex: 'cinema_user_id', key: 'cinema_user_id',},
  ],
})

function EditModel(record) {
  data.EditModalVisible = true
  EditFormState.user_id = record.id
  EditFormState.user_type = record.user_type
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