<template>
  <el-main >
    <el-row>
    <el-select v-model="depvalue" placeholder="提报单位名称">
      <el-option
        v-for="item in depoptions"
        :key="item.value"
        :label="item.label"
        :value="item.value"
        />
      </el-select>

      <el-select v-model="timevalue" placeholder="一个月">
        <el-option
          v-for="item in timeoptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

    </el-row>

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
import { defineComponent, ref } from 'vue'
import axios from 'axios'

export default defineComponent({
  name: 'FormHisList',
  data () {
    // 请求所有事项数据
    return {
      resdata: [],
      depvalue: ref(''),
      depoptions: [
        {
          value: '1',
          label: '办公室'
        },
        {
          value: '2',
          label: '销售管理部'
        },
        {
          value: '3',
          label: '台商部'
        },
        {
          value: '4',
          label: '金融市场部'
        },
        {
          value: '5',
          label: '产品部'
        },
        {
          value: '6',
          label: '支付与数字银行部'
        },
        {
          value: '7',
          label: '信用卡事业部'
        },
        {
          value: '8',
          label: '理财事业部'
        },
        {
          value: '9',
          label: '授信管理部'
        },
        {
          value: '10',
          label: '风险管理部'
        },
        {
          value: '11',
          label: '法律合规部'
        },
        {
          value: '12',
          label: '信息科技部'
        },
        {
          value: '13',
          label: '运营管理部'
        },
        {
          value: '14',
          label: '计划财务部'
        },
        {
          value: '15',
          label: '人力资源部'
        },
        {
          value: '16',
          label: '营业部'
        },
        {
          value: '17',
          label: '上海业务管理中心'
        },
        {
          value: '18',
          label: '南京分行'
        },
        {
          value: '19',
          label: '苏州分行'
        },
        {
          value: '20',
          label: '深圳分行'
        },
        {
          value: '21',
          label: '广州分行'
        },
        {
          value: '22',
          label: '北京分行'
        },
        {
          value: '23',
          label: '天津分行'
        },
        {
          value: '24',
          label: '成都分行'
        },
        {
          value: '25',
          label: '重庆分行'
        },
        {
          value: '26',
          label: '武汉分行'
        },
        {
          value: '27',
          label: '西安分行'
        },
        {
          value: '28',
          label: '宁波分行'
        },
        {
          value: '29',
          label: '济南分行'
        },
        {
          value: '30',
          label: '南宁分行'
        }
      ],
      timevalue: ref(''),
      timeoptions: [
        {
          value: '1',
          label: '近一周'
        },
        {
          value: '2',
          label: '一个月'
        },
        {
          value: '3',
          label: '三个月'
        }
      ]
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
