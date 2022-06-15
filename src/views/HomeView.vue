<template>
  <div class="home">
    <AppHeader/>
    <FormList></FormList>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import AppHeader from '@/components/AppHeader.vue' // @ is an alias to /src
import FormList from '@/views/FormList.vue'

export default defineComponent({
  name: 'HomeView',
  components: {
    AppHeader,
    FormList
  }
})
</script>
