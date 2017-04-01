<template>
  <article class="resp-row media">
    <figure class="media-left">
      <p class="image is-64x64">
        <img class="thumb" :src="businessData.image_url">
      </p>
    </figure>
    <div class="media-content">
      <div class="content">
        <p>
          <strong>{{businessData.name}}</strong>
          <span class="tag">
            {{businessData.categories[0].title}}
          </span>
          <br>
        </p>
        <p>
          <yelp-stars :rating="businessData.rating" :total="businessData.review_count"></yelp-stars>
        </p>
      </div>
      <div class="content" :class="{'is-loading, button': summaryLoading}">
          <h3>{{title}}</h3>
          <p>{{summary}}</p>
      </div>
    </div>
  </article>
</template>

<script>
import axios from 'axios'
import YelpStars from './YelpStars.vue'

export default{
  name: 'business-box',
  components: {
    YelpStars
  },
  props: ['businessData'],
  data: function() {
    return {
      summaryLoading: 'false',
      title: '',
      summary: ''
    }
  },
  methods: {
    getSummary: function(data){
      this.summaryLoading = true
      var vm = this
      axios.post("http://localhost:8000/summary",{ // TODO: remove for production
        name: data.name,
        latitude: data.coordinates.latitude,
        longitude: data.coordinates.longitude
      })
        .then(function(response){
          console.log(response.data)
          vm.summary = response.data.text
          vm.summaryLoading = false
        })
        .catch(function(error){
          console.error(error)
          vm.summaryLoading = false
        })
    }
  },
  watch: {
    businessData: function(newVal, oldVal) {
      this.getSummary(newVal)
    }
  },
  mounted: function(){
      this.getSummary(this.businessData)
  }
}
</script>

<style scoped>

.thumb {
  display: inline-block;
  width: 64px;
  height: 64px;
  background-position: center center;
  background-size: cover;
}

.resp-row {
  margin: 5px;
}

</style>
