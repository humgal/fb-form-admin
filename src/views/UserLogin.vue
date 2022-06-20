<template>
<div class="ms-login">
  <el-row>
    <el-col :span="8"><div class="grid-content ep-bg-purple" /></el-col>
    <el-col :span="8"><div class="grid-content ep-bg-purple-light" />
    <div class="ms-title">用户登录</div>
      <el-form>
        <el-form-item label="用户名">
          <el-input v-model="form.name" placeholder="用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码&nbsp&nbsp&nbsp">
          <el-input v-model="form.password" placeholder="密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="login-btn" @click="onSubmit" >登录</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="8"><div class="grid-content ep-bg-purple" /></el-col>
  </el-row>
</div>
</template>

<script lang="ts">
import axios from 'axios'
import { defineComponent } from 'vue'
import { ElMessage } from 'element-plus'
const form = {
  name: '',
  password: ''
}

export default defineComponent({
  name: 'UserLogin',
  data () {
    return {
      form: form
    }
  },
  methods: {
    onSubmit () {
      console.log(this.form)
      const open2 = () => {
        ElMessage({
          showClose: true,
          message: '登录成功',
          type: 'success'
        })
      }
      const open4 = () => {
        ElMessage({
          showClose: true,
          message: '登录失败',
          type: 'error'
        })
      }
      axios.post('/api/login', { name: form.name, password: form.password }).then((res) => {
        console.log(res)
        if (res.data.Code === 1) {
          open2()
          sessionStorage.setItem('token', res.data.Msg)
          this.$router.push('/')
        } else {
          open4()
        }
      })
    }
  }
})
</script>

<style scoped>

.ms-login {
  margin-top: 200px;
}
.ms-title {
    width: 100%;
    text-align: center;
    font-size: 20px;
    margin-bottom: 30px;
}

.login-btn {
    text-align: center;
        width: 100%;
    height: 36px;
}
</style>
