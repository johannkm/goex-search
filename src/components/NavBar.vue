<template>
  <nav class="nav dark-theme">
    <div class="nav-left">
      <!-- <router-link to="/" exact class="nav-item title">
        Goex
      </router-link> -->
      <a href="/" class="nav-item title">
        Goex
      </a>
    </div>

    <span class="nav-toggle" :class="{ 'is-active' : toggled }" @click="toggle">
      <span></span>
      <span></span>
      <span></span>
    </span>

    <div class="nav-right nav-menu" :class="{ 'is-active' : toggled }">
      <a class="nav-item" href="https://github.com/johannkm/goex-search">
        <span class="icon">
          <i class="fa fa-github"></i>
        </span>
        Github
      </a>
    </div>

  </nav>
</template>

<script>
export default {
  name: 'nav-bar',
  data (){
    return{
      toggled: false
    }
  },
  methods: {
    toggle: function(){
      this.toggled = !this.toggled
    }
  }
}
</script>

<style scoped>

.nav {
  min-height: 42px!important;
}

.nav-item{
  /*padding-top: 0;
  padding-bottom: 0;*/
  color: #fff;
  font-size: 1rem;
  /*font-weight: 500;*/
  line-height: 1.2rem;
}
.nav-item:hover {
  color: #fff;
  z-index: 10!important;
}
.nav-item:active {
  color: #fff!important;
}

.title{
  text-transform: uppercase;
  font-weight: 600;
  letter-spacing: .2rem;
}
.nav-toggle span{
  background-color: #fff!important;
}
.nav-toggle{
  background-color: #transparent;

}
.nav-toggle:hover {
  background-color: #504455;
}

.nav-menu.is-active {
  background-color: #504455;
}
.nav-menu.nav-right {
  z-index: 10;
}

</style>
