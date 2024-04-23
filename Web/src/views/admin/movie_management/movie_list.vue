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
        <a-form-item has-feedback label="电影名称" name="movie_name"
                 :rules="[{ required: true, message: '请输入电影名称!', trigger: 'blur' }]">
        <a-input v-model:value="formState.movie_name" placeholder="电影名"/>
        </a-form-item>
        <a-form-item has-feedback label="上映时间" name="release_date"
                   :rules="[{ required: true, message: '请输入上映时间!', trigger: 'blur' }]">
          <a-date-picker
              v-model:value="formState.release_date"
              :placeholder="formState.release_date ? formState.release_date.format('YYYY-MM-DD HH:mm') : '选择上映时间'"
              @change="handleDateChange"
          />
        </a-form-item>
        <a-form-item has-feedback label="电影封面" name="coverPicture"
                 :rules="[{ required: true, message: '请输入电影名称!', trigger: 'blur' }]">
          <a-upload
            v-model:file-list="data.fileList"
            :custom-request="customRequest"
            list-type="picture-card"
            :headers="{ token: adminInfoStore.adminInfo.token }"
            name="coverPicture"
          >
            <div>上传电影封面</div>
          </a-upload>
        </a-form-item>
        <a-form-item has-feedback label="电影时长" name="play_time"
                     :rules="[{ required: true, message: '请选择电影时长', trigger: 'blur' }]">
          <div>
            <a-select v-model:value="film_length.hours" style="width: 80px;" @change="updatePlayTime">
              <a-select-option v-for="hour in hourOptions" :key="hour" :value="hour">{{ hour }} 小时</a-select-option>
            </a-select>
            <a-select v-model:value="film_length.minutes" style="width: 80px;" @change="updatePlayTime">
              <a-select-option v-for="minute in minuteOptions" :key="minute" :value="minute">{{ minute }} 分钟</a-select-option>
            </a-select>
            <a-select v-model:value="film_length.seconds" style="width: 80px;" @change="updatePlayTime">
              <a-select-option v-for="second in secondOptions" :key="second" :value="second">{{ second }} 秒</a-select-option>
            </a-select>
          </div>
            <span>{{ formState.play_time }}</span>
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
import {movieCreateApi, movieEditApi, movieCoverUploadApi} from "@/api/movie_management_api.js";
import cloneDeep from 'lodash/cloneDeep';
import {durationToNanoseconds, durationTransition} from "@/utils/durationTransition.js";
import {message} from "ant-design-vue";
import {dateTransition} from "../../../utils/dateTransition.js";
import {AdminInfoStore} from "@/stores/admin_info.js";
// 对话框中规则验证
const formRef = ref({})
const adminTable = ref(null)

const adminInfoStore = AdminInfoStore()

// 用于置空formState
const _formState = {
  movie_name: "",
  release_date: "",
  play_time: "",
  director: "",
  actors: "",
  cover_picture_id: 0,
  cover_picture_url: ""
}
// formState 表单信息，用于添加电影信息
let formState = reactive({
  movie_name: "",
  cover_picture_id: 0,
  cover_picture_url: "",
  release_date: "",
  play_time: "",
  director: "",
  actors: "",
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
    {title: '电影封面', dataIndex: 'coverPicture', key: 'coverPicture',},
    {title: '电影时长', dataIndex: 'play_time', key: 'play_time',},
    {title: '导演', dataIndex: 'director', key: 'director',},
    {title: '演员', dataIndex: 'actors', key: 'actors',},
    {title: '操作', dataIndex: 'action', key: 'action',},
  ],
  fileList: []
})

const film_length = reactive({
  hours: 0,
  minutes: 0,
  seconds: 0,
})

// 用于用户自选电影时长
const hourOptions = [...Array(24).keys()]; // 0-23 小时
const minuteOptions = [...Array(60).keys()]; // 0-59 分钟
const secondOptions = [...Array(60).keys()]; // 0-59 秒
function formattedDuration(hours, minutes, seconds) {
  return `${hours}小时${minutes}分钟${seconds}秒`;
}
function updatePlayTime() {
  formState.play_time = formattedDuration(film_length.hours, film_length.minutes, film_length.seconds)
  console.log(film_length)
  console.log(formState.play_time)
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

async function customRequest(file) {
  formState.cover_picture_id = 0
  formState.cover_picture_url = ""
  try {
    // 创建 FormData 对象
    const formData = new FormData();
    // 获取文件对象
    const fileObj = file.file;
    // 将文件添加到 FormData 对象中，使用 coverPicture 作为参数名
    formData.append('coverPicture', fileObj);

    let res = await movieCoverUploadApi(formData);
    // 处理响应
    if(res.code) {
      message.error(res.msg)
    } else {
      message.success(res.msg)
      formState.cover_picture_id = res.data.picture_id
      formState.cover_picture_url = res.data.file_path
    }
  } catch (error) {
    // 处理上传失败的情况
    message.error('上传失败，请重试');
  }
}

// AddMovie 添加电影
async function AddMovie() {
  // 先对 EditFormState 进行深拷贝
  const AddFormStateRequest = cloneDeep(formState);

  // 将 EditFormStateRequest.release_date 转换为时间字符串
  const formattedDate = dayjs(formState.release_date).format('YYYY-MM-DDTHH:mm:ss[Z]');
  // 将格式化后的时间字符串重新赋值给 EditFormStateRequest.release_date
  AddFormStateRequest.release_date = formattedDate;

  // 将 EditFormState.play_time 转换为纳秒数
  const formatterDuration = durationToNanoseconds(formState.play_time);
  AddFormStateRequest.play_time = formatterDuration;

  try {
        // 发送请求到后端
    let res = await movieCreateApi(AddFormStateRequest);
    if (res.code) {
      // 处理错误消息
      message.error(res.msg);
      data.AddModalVisible = false;
    } else {
      // 处理成功消息
      message.success(res.msg);
      // 关闭编辑对话框
      data.AddModalVisible = false;
    }
  } catch (error) {
    // 处理请求失败的情况
    console.error("Error:", error);
    // 显示错误提示
    message.error("请求失败");
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