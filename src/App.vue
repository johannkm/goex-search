<template>
  <div id="app">
    <nav class="dark-theme">
      <div class="nav-left">
        <a class="nav-item" href="/">
          goex
        </a>
      </div>
    </nav>
    <div class="container main-content">

        <form class="search-form" v-on:submit.prevent="postSearch">
          <div class="field">
            <div class="columns is-gapless">
              <div class="column">
                <p class="control has-icon">
                  <input class="input" type="input" placeholder="Search" v-model="term">
                  <span class="icon is-small">
                    <i class="fa fa-search"></i>
                  </span>
                </p>
              </div>
              <div class="column">
                <p class="control has-icon">
                  <input class="input" type="input" placeholder="Place" v-model="location" required>
                  <span class="icon is-small">
                    <i class="fa fa-location-arrow"></i>
                  </span>
                </p>
              </div>
              <div class="column is-1">
                <p class="control">
                  <button type="submit" class="button is-primary" :class="{'is-loading': searching}">Go</button>

                </p>
              </div>
            </div>
          </div>
        </form>

        <p v-if="noResponse">No matching results</p>

        <table class="table is-striped yscroll">
          <tr v-for="b in response">
            <business-box :businessData="b"/>
            <hr>
          </tr>
        </table>


    </div>
  </div>
</template>

<script>
import axios from 'axios'
// var debounce = require('lodash/debounce')
import BusinessBox from './Components/BusinessBox.vue'

export default {
  name: 'app',
  components: {
    BusinessBox
  },
  data () {
    return {
      term: '',
      location: 'DC',
      response: [],
      noResponse: false,
      searching: false
    }
  },
  methods: {
    postSearch: function() {
      this.noResponse = false
      this.searching = true
      let location = this.location
      if(this.location == 'Current location'){
        if(navigator.geolocation){
          let pos = navigator.geolocation.getCurrentPosition((pos) => {
            location = pos.coords.latitude+','+pos.coords.longitude
          })
        }
      }
      console.log(location)
      var vm = this
      axios.post("http://localhost:8000/places",{ // TODO: remove for production
        term: vm.term,
        location: location
      })
        .then(function(response){
          console.log(response.data)
          vm.response = response.data.businesses.slice(0,2)
          if(vm.response == null || vm.response.length==0){
            vm.noResponse = true
          }
          vm.searching = false
        })
        .catch(function(error){
          console.error(error)
          vm.searching = false
        })
    }
  }
}
</script>

<style lang="sass">
  @import "~bulma"
</style>

<style scoped>

  .main-content {
    max-width: 1200px;
  }

  .search-form {
    padding-top: 1rem;
    padding-bottom: 1rem;
  }

  .dark-theme {
    background-color: #3F3244;
  }

  .nav-item{
    color: #fff;
  }
  .nav-item:hover {
    color: #fff;
  }

  .yscroll {
    overflow-x: hidden;
    overflow-y: scroll;
  }

  tr:hover {
    background-color: transparent!important;
  }

</style>
