<template>
  <nav class="nav dark-theme has-shadow">
    <div class="nav-left">
      <a class="nav-item title" href="/">
        Goex
      </a>
    </div>
    <div class="nav-right nav-menu">
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
  name: 'nav-bar'
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
.title{
  text-transform: uppercase;
  font-weight: 600;
  letter-spacing: .2rem;
}

.nav-item:hover {
  color: #fff;
}

</style>
