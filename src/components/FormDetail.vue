<template>
  <form-create :rule="rule" v-model:api="fApi" :option="options" />
</template>

<script lang="ts">

import { defineComponent } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'
export default defineComponent({
  name: 'FormDetail',
  data () {
    return {
      fApi: {},
      value: {},
      rule: [],
      options: {

        onSubmit: (formData: []) => {
          const open2 = () => {
            ElMessage({
              showClose: true,
              message: '提交成功',
              type: 'success'
            })
          }
          const open4 = () => {
            ElMessage({
              showClose: true,
              message: '提交失败',
              type: 'error'
            })
          }
          console.log(JSON.stringify(formData))
          axios({
            method: 'post',
            url: '/api/form/updatecontent',
            data: {
              id: this.$route.params.id,
              content: formData,
              title: this.$route.params.title
            },
            headers: {
              token: String(sessionStorage.getItem('token'))
            }

          }
          ).then((res) => {
            console.log(res)
            if (res.data.Code === 0) {
              open2()
            } else {
              open4()
            }
          }).catch((err) => {
            console.log(err)
          })
        },
        resetBtn: false
      }
    }
  },
  methods: {
    getFormDetail () {
      axios.post('/api/form/detail', {
        id: this.$route.params.id
      }).then((res) => {
        console.log(res.data.Msg)
        this.rule = JSON.parse(res.data.Msg)
      })
    }
  },
  beforeMount () {
    this.getFormDetail()
  }
})
</script>
