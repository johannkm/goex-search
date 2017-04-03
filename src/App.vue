<template>
  <div id="app">

    <div class="content-except-footer">
      
      <div class="dark-theme">
        <div class="container main-content">
          <nav-bar/>
        </div>
      </div>

      <router-view></router-view>

    </div>

    <footer class="footer">
      <small>
        <p>
          Made with data from
          <a href="https://www.yelp.com" class="bottom-links">
            <span class="icon is-small icon-cred">
              <i class="fa fa-yelp"></i>
            </span>
            Yelp
          </a>
          &nbspand
          <a href="https://developers.google.com/places/" class="bottom-links">
            <span class="icon is-small icon-cred">
              <i class="fa fa-map-marker"></i>
            </span>
            Google Places
          </a>
        </p>
      </small>
    </footer>

  </div>
</template>

<script>
import NavBar from './Components/NavBar.vue'

export default {
  name: 'app',
  components: {
    NavBar
  }
}
</script>

<style lang="sass">
  @import "~bulma"
</style>

<style>
html{
  height: 100%;
}
body {
  min-height: 100vh;
  flex-direction: column;
  /*flex-shrink: 0;*/
}
#app {
  height: 100vh;
  flex-direction: column;
  display: flex;
}
.content-except-footer {
  flex: 1 0 auto;
  padding: var(--space) var(--space) 0;
  width: 100%;
}
.content-except-footer::after {
  content: '\00a0'; /* &nbsp; */
  display: block;
  margin-top: var(--space);
  height: 0px;
  visibility: hidden;
}
@media (--break-lg) {
  .content-except-footer {
    padding-top: var(--space-lg);
  }
  .content-except-footer::after {
    margin-top: var(--space-lg);
  }
}

.content-except-footer--full {
  padding: 0;
}
.content-except-footer--full::after {
  content: none;
}

</style>

<style scoped>

  .main-content {
    max-width: 1000px;
    padding-right: 5px;
    padding-left: 5px;
  }

  .dark-theme {
    background-color: #504455;
  }

  .footer {
    font-size: 13px;
    flex: none;
    color: #6F6F6F;
    padding: 0.5rem;
    background-color: transparent;
    text-align: center;
  }
  .bottom-links{
    color: #2E2E2E
  }

</style>
