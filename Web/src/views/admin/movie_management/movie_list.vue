<template>
  <div>
    <a-modal v-model:visible="data.AddModalVisible" title="添加电影" @ok="AddMovie">
      <a-form
        :model="formState"
        name="basic"
        ref="formRef"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 20 }"
        autocomplete="off"
      >
         <template>
          <a-upload
            action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
            :multiple="false"
            :file-list="formState.pictureFile"
            @change="handleChange"
          >
            <a-button>
              <upload-outlined></upload-outlined>
              上传电影封面
            </a-button>
          </a-upload>
        </template>
        <a-form-item has-feedback label="电影名称" name="movie_name"
                 :rules="[{ required: true, message: '请输入电影名称!', trigger: 'blur' }]">
        <a-input v-model:value="formState.movie_name" placeholder="电影名"/>
        </a-form-item>
        <a-form-item has-feedback label="上映时间" name="release_date"
                   :rules="[{ required: true, message: '请输入上映时间!', trigger: 'blur' }]">
          <a-input-password v-model:value="formState.release_date" placeholder="上映时间"/>
        </a-form-item>
        <a-form-item has-feedback label="电影时长" name="play_time"
                     :rules="[{ required: true, message: '请输入电影时长', trigger: 'blur' }]">
            <a-input v-model:value="formState.play_time" placeholder="电影时长"/>
        </a-form-item>
        <a-form-item has-feedback label="导演" name="director"
                   :rules="[{ required: true, message: '请输入导演!', trigger: 'blur' }]">
          <a-input v-model:value="formState.director" placeholder="导演"/>
        </a-form-item>
        <a-form-item has-feedback label="演员" name="actors"
                   :rules="[{ required: true, message: '请输入演员!', trigger: 'blur' }]">
          <a-input v-model:value="formState.actors" placeholder="演员"/>
        </a-form-item>
      </a-form>
    </a-modal>
    <a-modal v-model:visible="data.EditModalVisible" title="编辑电影" @ok="EditMovie">
      <a-form
        :model="EditFormState"
        name="basic"
        ref="formRef"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 20 }"
        autocomplete="off"
      >
        <a-form-item has-feedback label="电影名称" name="movie_name"
                 :rules="[{ required: true, message: '请输入电影名称!', trigger: 'blur' }]">
        <a-input v-model:value="EditFormState.movie_name" placeholder="电影名"/>
        </a-form-item>
        <a-form-item has-feedback label="上映时间" name="release_date"
            :rules="[{ required: true, message: '请选择上映时间!', trigger: 'blur' }]">
            <a-date-picker
                v-model:value="EditFormState.release_date"
                :placeholder="EditFormState.release_date ? EditFormState.release_date.format('YYYY-MM-DD HH:mm') : '选择上映时间'"
                @change="handleDateChange"
            />
        </a-form-item>
        <a-form-item has-feedback label="电影时长" name="play_time"
                     :rules="[{ required: true, message: '请输入电影时长', trigger: 'blur' }]">
            <a-input v-model:value="EditFormState.play_time" placeholder="电影时长"/>
        </a-form-item>
        <a-form-item has-feedback label="导演" name="director"
                   :rules="[{ required: true, message: '请输入导演!', trigger: 'blur' }]">
          <a-input v-model:value="EditFormState.director" placeholder="导演"/>
        </a-form-item>
        <a-form-item has-feedback label="演员" name="actors"
                   :rules="[{ required: true, message: '请输入演员!', trigger: 'blur' }]">
          <a-input v-model:value="EditFormState.actors" placeholder="演员"/>
        </a-form-item>
      </a-form>
    </a-modal>
    <AdminTable
        @delete="movieRemove"
        :columns="data.columns"
        base-url="/api/movies"
        ref="adminTable"
        like-title="搜索电影名称"
        is-delete is-add is-edit>
      <template #add>
        <a-button type="primary" @click="data.AddModalVisible = true">添加电影</a-button>
      </template>
      <template #edit="{record}">
        <a-button class="user_action update" @click="EditModel(record)" type="primary">编辑</a-button>
      </template>
      <template #cell="{ column, record }">
        <template v-if="column.key === 'cover_picture_url'">
          <img class="cover_picture" :src="record.cover_picture_url">
        </template>
        <template v-if="column.key === 'play_time'">
          <span>{{ durationTransition(record.play_time) }}</span>
        </template>
        <template v-if="column.key === 'release_date'">
          <span>{{ dateTransition(record.release_date, 1) }}</span>
        </template>
      </template>
    </AdminTable>
  </div>
</template>

<script setup>
import dayjs from 'dayjs';
import AdminTable from "@/components/admin/admin_table.vue";
import {reactive, ref} from "vue";
import {movieCreateApi, movieEditApi} from "@/api/movie_management_api.js";
import cloneDeep from 'lodash/cloneDeep';
import {durationToNanoseconds, durationTransition} from "@/utils/durationTransition.js";
import {message} from "ant-design-vue";
import {dateTransition} from "../../../utils/dateTransition.js";

// 对话框中规则验证
const formRef = ref({})
const adminTable = ref(null)

// 用于置空formState
const _formState = {
  movie_name: "",
  release_date: "",
  play_time: "",
  director: "",
  actors: "",
  pictureFile: []
}
// formState 表单信息，用于添加电影信息
let formState = reactive({
  movie_name: "",
  release_date: "",
  play_time: "",
  director: "",
  actors: "",
  pictureFile: []
})

// EditFormState 表单信息，用于编辑电影信息
let EditFormState = reactive({
  movie_name: "",
  release_date: "",
  play_time: "",
  director: "",
  actors: "",
  movie_id: 0,
})

const data = reactive({
  AddModalVisible: false, // 对话框是否可见
  EditModalVisible: false,
  columns: [
    {title: '电影ID', dataIndex: 'id', key: 'id',},
    {title: '电影名', dataIndex: 'movie_name', key: 'movie_name',},
    {title: '上传时间', dataIndex: 'created_at', key: 'created_at',},
    {title: '电影封面', dataIndex: 'cover_picture_url', key: 'cover_picture_url',},
    {title: '上映时间', dataIndex: 'release_date', key: 'release_date',},
    {title: '电影时长', dataIndex: 'play_time', key: 'play_time',},
    {title: '导演', dataIndex: 'director', key: 'director',},
    {title: '演员', dataIndex: 'actors', key: 'actors',},
    {title: '操作', dataIndex: 'action', key: 'action',},
  ],
})

function handleChange(info) {
  let fileList = [...info.fileList]; // 拷贝一份文件列表

  // 限制文件上传数量
  fileList = fileList.slice(-1);

  // 将文件列表保存到组件状态中
  this.fileList = fileList;

  // 检查上传状态
  if (info.file.status === 'done') {
    this.$message.success(`${info.file.name} 文件上传成功`);
  } else if (info.file.status === 'error') {
    this.$message.error(`${info.file.name} 文件上传失败`);
  }
}

// 点击编辑按钮触发的事件
function EditModel(record) {
  data.EditModalVisible = true
  EditFormState.movie_id = record.id
  EditFormState.movie_name = record.movie_name
  EditFormState.release_date = dayjs(record.release_date)
  EditFormState.play_time = durationTransition(record.play_time)
  EditFormState.director = record.director
  EditFormState.actors = record.actors
  console.log(EditFormState)
}

// 编辑电影
async function EditMovie() {
  // EditFormStateRequest用于向后端请求，而EditFormState用于和前端的表单信息双向绑定
  // 这样不会打乱前端展示格式的同时，还能向后端发送满足后端请求格式的数据
  // 先对 EditFormState 进行深拷贝
  const EditFormStateRequest = cloneDeep(EditFormState);

  // 将 EditFormStateRequest.release_date 转换为时间字符串
  const formattedDate = dayjs(EditFormState.release_date).format('YYYY-MM-DDTHH:mm:ss[Z]');
  // 将格式化后的时间字符串重新赋值给 EditFormStateRequest.release_date
  EditFormStateRequest.release_date = formattedDate;

  // 将 EditFormState.play_time 转换为纳秒数
  const formatterDuration = durationToNanoseconds(EditFormState.play_time);
  EditFormStateRequest.play_time = formatterDuration;

  try {
    // 发送请求到后端
    let res = await movieEditApi(EditFormStateRequest);
    if (res.code) {
      // 处理错误消息
      message.error(res.msg);
    } else {
      // 处理成功消息
      message.success(res.msg);
      // 关闭编辑对话框
      data.EditModalVisible = false;
    }
  } catch (error) {
    // 处理请求失败的情况
    console.error("Error:", error);
    // 显示错误提示
    message.error("请求失败");
  }
}

function handleDateChange(value) {
    formState.release_date = value;
}

function movieRemove(user_id) {
  console.log(user_id)
}

// 创建用户
async function AddMovie() {
  try {
    // 主动触发验证
    await formRef.value.validate()
    let res = await movieCreateApi(formState)
    if (res.code) {
        message.error(res.msg)
        return
    }
    message.success(res.msg)
    data.AddModalVisible = false
    // 置空
    Object.assign(formState, _formState)
    // // 清除验证规则
  } catch (e) {
    console.log(e)
  }
}

</script>

<style lang="scss">
.cover_picture {
  width: 200px;
}

.cover_picture_upload {
  margin-bottom: 10px;
}

</style>