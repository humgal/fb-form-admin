<template>
  <AppHeader/>
  <el-main >
    <ul  class="infinite-list" style="overflow: auto">
    <li v-for="i in 18" :key="i" class="infinite-list-item">
    <div class="hislist">
      <div> <h4>收集事项标题</h4> </div>
          <div>upuser</div>
          <div>uptime</div>
    </div>
    <el-divider />
    </li>
    </ul>
  </el-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import AppHeader from '@/components/AppHeader.vue'
export default defineComponent({
  name: 'FormHisList',
  components: {
    AppHeader
  }
})
</script>

<style >
.hislist {
  margin-left: 20px;
  text-align: left;
  width: 85%;

}

.infinite-list {

  padding: 0px;
  margin: 0;
  list-style: none;
}

</style>
