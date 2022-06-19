<template>
  <el-menu
    :default-active="1"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
  >
    <el-menu-item index="0">LOGO</el-menu-item>
      <el-menu-item index="1"><router-link to="/formlist">事项填写</router-link> </el-menu-item>
    <el-menu-item index="2"><router-link to="/formhislist">事项记录</router-link> </el-menu-item>
    <div class="flex-grow" />

    <el-menu-item index="3" ><a href="http://localhost:8088">表单编辑器</a></el-menu-item>

    <el-sub-menu index="4">
      <template #title>系统管理</template>
      <el-menu-item index="4-1" >用户管理</el-menu-item>
      <el-menu-item index="4-2" >部门管理</el-menu-item>
     <el-menu-item index="4-2" >表单授权</el-menu-item>
      <el-menu-item index="4-4">退出登录</el-menu-item>

    </el-sub-menu>
  </el-menu>
</template>

<script lang="ts" >
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'AppHeader'
})

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>
