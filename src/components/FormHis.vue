<template>
  <form-create :rule="rule" v-model:api="fApi" :option="options" />
</template>

<script lang="ts">

import { defineComponent } from 'vue'
import axios from 'axios'

export default defineComponent({
  name: 'FormDetail',
  data () {
    return {
      fApi: {},
      value: {},

      options: {
        submitBtn: false,
        resetBtn: false
      },

      rule: []
    }
  },
  methods: {
    getFormDetail () {
      axios.post('/api/form/hisdetail', {
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
