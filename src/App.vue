<template>
  <div id="app">
    <div class="container">
      <div class="columns is-mobile">
          <div class="column is-8 is-offset-2">
            <input type="input" class="input" v-model="search" @input="postSearch" placeholder="Search"></input>
            <table class="table is-striped yscroll">
              <tr v-for="b in response">
                <article class="media">
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

export default {
  name: 'app',
  data () {
    return {
      search: '',
      response: []
    }
  },
  methods: {
    postSearch: debounce( function() {
      var vm = this
      axios.post("http://localhost:8000/places",{ // TODO: remove for production
        location: vm.search
      })
        .then(function(response){
          console.log(response.data)
          vm.response = response.data.businesses
        })
        .catch(function(error){
          console.error(error)
        })
    }, 500)
  }
}
</script>

<style lang="sass">
  @import "~bulma"
</style>

<style scoped>
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
