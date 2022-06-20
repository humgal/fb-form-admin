import axios from 'axios'
// 环境的切换
if (process.env.NODE_ENV === 'development') {
  axios.defaults.baseURL = '/api'
} else if (process.env.NODE_ENV === 'production') {
  axios.defaults.baseURL = 'http://114.113.127.115:9082/api'
}

// 请求拦截器
axios.interceptors.request.use(
  config => {
  // 每次发送请求之前判断是否存在token，如果存在，则统一在http请求的header都加上token
    const token = sessionStorage.getItem('token')
    token && (config.headers.Authorization = token)
    return config
  },
  error => {
    return Promise.error(error)
  })

axios.defaults.timeout = 60000

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'

// 响应拦截器
axios.interceptors.response.use(response => {
  if (response.status === 200) {
    return Promise.resolve(response)
  } else {
    return Promise.reject(response)
  }
}, error => {
  if (error.response.status) {
  // 对不同返回码对相应处理
    return Promise.reject(error.response)
  }
})

// get 请求
export function httpGet ({
  url,
  params = {}
}) {
  return new Promise((resolve, reject) => {
    axios.get(url, {
      params
    }).then((res) => {
      resolve(res.data)
    }).catch(err => {
      reject(err)
    })
  })
}
// delete 请求
export function httpDelete ({
  url,
  data = {}
}) {
  return new Promise((resolve) => {
    axios({
      url,
      method: 'delete',
      data
    }).then((res) => {
      resolve(res.data)
    })
  })
}

// post请求
export function httpPost ({
  url,
  data = {},
  params = {}
}) {
  return new Promise((resolve) => {
    axios({
      url,
      method: 'post',
      transformRequest: [function (data) {
        let ret = ''
        for (const it in data) {
          ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
        }
        return ret
      }],
      // 发送的数据
      data,
      // url参数
      params
    }).then(res => {
      resolve(res.data)
    })
  })
}

// put请求
export function httpPut ({
  url,
  data = {},
  params = {}
}) {
  return new Promise((resolve) => {
    axios({
      url,
      method: 'put',
      transformRequest: [function (data) {
        let ret = ''
        for (const it in data) {
          ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
        }
        return ret
      }],
      // 发送的数据
      data,
      // url参数
      params
    }).then(res => {
      resolve(res.data)
    })
  })
}
