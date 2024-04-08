import Vue from 'vue'
import axios from 'axios'

// axios请求地址
axios.defaults.baseURL = 'http://localhost:3001/api/v1'

Vue.prototype.$http = axios
axios.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem(
      "token"
    )}`;
    return config;
  });