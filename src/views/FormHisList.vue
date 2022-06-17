<template>
  <el-main >
    <ul  class="infinite-list" style="overflow: auto">
    <li v-for="(item ,index) in resdata" :key="index" @click='toFormHis(item.title,item.id)' class="infinite-list-item">
    <div class="hislist">
      <div> <h4>{{item.title}}</h4> </div>
          <div>{{item.uptime}}</div>
          <div>{{item.upuser}}</div>
    </div>
    <el-divider />
    </li>
    </ul>
  </el-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios'
export default defineComponent({
  name: 'FormHisList',
  data () {
    // 请求所有事项数据
    return {
      resdata: []
    }
  },
  methods: {
    getFormHisList () {
      axios.post('/api/form/hislist').then((res) => {
        console.log(res.data.Msg)
        this.resdata = JSON.parse(res.data.Msg)
      })
    },
    toFormHis (title : string, id: number) {
      this.$router.push({
        path: '/formHis',
        name: 'formHis',
        params: {
          title: title,
          id: id
        }
      })
    }
  },
  mounted () {
    this.getFormHisList()
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
