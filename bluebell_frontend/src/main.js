import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from './service/api'

Vue.prototype.$axios = axios;
Vue.config.productionTip = false


router.beforeEach((to, from, next) => {
  console.log(to);
  console.log(from);
  if (to.meta.requireAuth) { // 判断该路由是否需要登录权限
    if (localStorage.getItem("loginResult")) { //判断本地是否存在access_token
      next();
    } else {
      if (to.path === '/login') {
        next();
      } else {
        next({
          path: '/login'
        })
      }
    }
  }
  else {
    next();
  }
  /*如果本地 存在 token 则 不允许直接跳转到 登录页面*/
  if (to.fullPath == "/login") {
    if (localStorage.getItem("loginResult")) {
      next({
        path: from.fullPath
      });
    } else {
      next();
    }
  }
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
