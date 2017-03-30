<template>
  <div id="app">
    <nav class="dark-theme">
      <div class="nav-left">
        <a class="nav-item" href="/">
          goex
        </a>
      </div>
    </nav>
    <div class="main-content container">
      <div class="columns">
          <div class="column is-8 is-offset-2">

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


            <!-- <input type="input" class="input" v-model="search" @input="postSearch" placeholder="Search"></input> -->

            <p v-if="noResponse">No matching results</p>

            <table class="table is-striped yscroll">
              <tr v-for="b in response">
                <article class="resp-row media">
                  <figure class="media-left">
                    <p class="image is-64x64">
                      <img class="thumb" :src="b.image_url">
                    </p>
                  </figure>
                  <div class="media-content">
                    <div class="content">
                      <p>
                        <strong>{{b.name}}</strong> <small>{{b.categories[0].title}}</small>
                        <br>
                      </p>
                      <p>
                        <yelp-stars :rating="b.rating" :total="b.review_count"></yelp-stars>
                      </p>
                    </div>
                  </div>
                </article>
              </tr>
            </table>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
var debounce = require('lodash/debounce')

import YelpStars from './components/YelpStars.vue'

export default {
  name: 'app',
  components: {
    YelpStars
  },
  data () {
    return {
      term: '',
      location: 'Current location',
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
          vm.response = response.data.businesses
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

  .resp-row {
    margin: 5px;
  }

  .yscroll {
    overflow-x: hidden;
    overflow-y: scroll;
  }
  .thumb {
    display: inline-block;
    width: 64px;
    height: 64px;
    background-position: center center;
    background-size: cover;
  }
</style>
