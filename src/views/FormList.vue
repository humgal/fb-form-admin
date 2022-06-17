<template>

  <el-main>

    <div class="center-card">
      <el-card shadow="hover" class="form-item" v-for="(item, index) in resdata" :key="index"
        @click='toFormDetail(item.title,item.id)'>
        <div>
          <h4>{{ item.title }}</h4>
        </div>
        <div>{{ item.upuser }}</div>
        <div>{{ item.uptime }}</div>
      </el-card>
    </div>

  </el-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios'
export default defineComponent({
  name: 'FormList',
  data () {
    // 请求所有事项数据
    return {
      resdata: []
    }
  },
  methods: {
    getFormList () {
      axios.post('/api/form/list').then((res) => {
        console.log(res.data.Msg)
        this.resdata = JSON.parse(res.data.Msg)
      })
    },
    toFormDetail (title: string, id: number) {
      this.$router.push({
        path: '/formDetail',
        name: 'formDetail',
        params: {
          id: id,
          title: title
        }
      })
    }
  },
  mounted () {
    this.getFormList()
  }
})
</script>

<style >
.center-card {
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
