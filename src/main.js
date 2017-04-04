import Vue from 'vue'
import App from './App.vue'
import SearchForm from './pages/SearchForm.vue'
import VueScrollTo from 'vue-scrollto'
import Router from 'vue-router'

Vue.use(VueScrollTo)
Vue.use(Router)

const routes = [
  { name: 'search', path: '/', component: SearchForm }
]

const router = new Router({
  routes
})

const app = new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
