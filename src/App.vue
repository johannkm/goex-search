<template>
  <div id="app">

    <div class="dark-theme">
      <div class="container main-content">
        <nav-bar/>
      </div>
    </div>

    <div class="content-form">
      <div class="container main-content">
        <form class="search-form" v-on:submit.prevent="postSearch">
          <div class="field">
            <div class="columns is-gapless">
              <div class="column">
                <p class="control has-icon">
                  <input class="input" type="input" placeholder="Search (Optional)" v-model="term">
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
              <div class="column is-narrow">
                <p class="control">
                  <button type="submit" class="button is-primary" :class="{'is-loading': searching}">Go</button>

                </p>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>

    <div class="container main-content">

      <div class="alert-messages">
        <p class="start-help" v-if="begining">Search for restaurants in near a location. Goex will summarize all customer reviews.</p>
        <p v-if="noResponse">No matching results</p>
        <p v-if="noServer">Can't reach server.</p>
      </div>

      <table class="table is-striped yscroll">
        <loading-icon :active="searching" :isSmall="false"/>

        <tr v-for="b in response">
          <div class="table-row">
            <business-box :businessData="b"/>
            <hr>
          </div>
        </tr>
      </table>


    </div>
  </div>
</template>

<script>
import axios from 'axios'
// var debounce = require('lodash/debounce')
import BusinessBox from './Components/BusinessBox.vue'
import LoadingIcon from './Components/LoadingIcon.vue'
import NavBar from './Components/NavBar.vue'

export default {
  name: 'app',
  components: {
    BusinessBox,
    LoadingIcon,
    NavBar
  },
  data () {
    return {
      term: '',
      location: 'DC',
      response: [],
      noResponse: false,
      searching: false,
      begining: true,
      noServer: false
    }
  },
  methods: {
    postSearch: function() {
      this.begining = false
      this.noResponse = false
      this.searching = true
      this.noServer = false
      this.response = []
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
          if(response.data.businesses == null || response.data.businesses.length==0){
            vm.noResponse = true
            vm.response = []
          } else {
            vm.response = response.data.businesses.slice(0,2)
          }
          vm.searching = false
        })
        .catch(function(error){
          console.error(error)
          vm.searching = false
          vm.noServer = true
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
    max-width: 1000px;
    padding-right: 5px;
    padding-left: 5px;
  }

  .search-form {
    padding-top: 1rem;
    padding-bottom: 1rem;
  }
  .content-form {
    background-color: #F4F4F4
  }

  .yscroll {
    overflow-x: hidden;
    overflow-y: scroll;
  }

  hr {
    margin-bottom: 0;
    margin-top: 12px;
  }
  .table-row{
    margin-top: 20px;
  }

  tr:hover {
    background-color: transparent!important;
  }

  .dark-theme {
    background-color: #504455;
  }

  .alert-messages {
    padding-top: 1.2rem;
  }
  .start-help {
    color: #6F6F6F;
  }

</style>
