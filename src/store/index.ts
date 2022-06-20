import Vuex from 'vuex'

const store = new Vuex.Store({
  state: {
    // 存储token
    Authorization: sessionStorage.getItem('Authorization') ? sessionStorage.getItem('Authorization') : ''
  },

  mutations: {
    // 修改token，并将token存入sessionStorage
    changeLogin (state, user) {
      state.Authorization = user.Authorization
      sessionStorage.setItem('Authorization', user.Authorization)
    }
  }
})

export default store
