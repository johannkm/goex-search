<template>
  <nav class="nav dark-theme">
    <div class="nav-left">
      <a class="nav-item title" href="#">
        Goex
      </a>
    </div>

    <span class="nav-toggle" :class="{ 'is-active' : toggled }" @click="toggle">
      <span></span>
      <span></span>
      <span></span>
    </span>

    <div class="nav-right nav-menu" :class="{ 'is-active' : toggled }">
      <a class="nav-item" href="/">
        <span class="icon">
          <i class="fa fa-github"></i>
        </span>
        View on Github
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
  padding: 0;
  color: #fff;
  font-size: 1rem;
  font-weight: 500;
  line-height: 1.2rem;
}
.nav-item:hover {
  color: #fff;
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

.nav-menu {
  background-color: #504455
}

</style>
