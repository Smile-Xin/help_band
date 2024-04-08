import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import dayjs from 'dayjs'
import './plugins/elementui'
import './plugins/antui'
import './plugins/http'
import './assets/css/style1.css'

import { Message } from 'element-ui';
import { MessageBox  } from 'element-ui'


Vue.filter('dateformat', function(indate, outdate) {
  return dayjs(indate).format(outdate)
})

Vue.prototype.$message = Message
Vue.prototype.$confirm = MessageBox
Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
