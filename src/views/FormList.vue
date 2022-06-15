<template>
<AppHeader/>
  <el-main >

    <div class="center-card" >
    <el-card shadow="hover" class="form-item" v-for="n in 20" :key="n">
          <div> <h4>收集事项标题</h4> </div>
          <div>upuser</div>
          <div>uptime</div>
        </el-card>
    </div>

  </el-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import AppHeader from '@/components/AppHeader.vue'
export default defineComponent({
  name: 'FormList',
  components: {
    AppHeader
  }
})
</script>

<style >
.center-card {
  margin-top: 30px;
  text-align: left;
  width: 100%;
  display: flex;
  flex-wrap: wrap;

}
.center-card .form-item {
  margin-top: 30px;
  margin-left: 15px;
  display: flex;
  height: 150px;
  width: 200px;
}

.infinite-list {

  padding: 0px;
  margin: 0;
  list-style: none;
}

</style>
