<template>
  <div class="busi-box">
    <div class="image is-64x64" align="left">
      <div class="thumb" :style="{ backgroundImage: 'url(' + businessData.image_url + ')' }"></div>
    </div>
    <strong>{{businessData.name}}</strong>
    <span class="tag">
      {{businessData.categories[0].title}}
    </span>
    <br>

    <yelp-stars :rating="businessData.rating" :total="businessData.review_count"></yelp-stars>

    <span class="price">
      <span v-for="x in businessData.price">
        $
      </span>
    </span>

    <div>
      <loading-icon :active="summaryLoading" :isSmall="true"/>
      <strong class="keyword">{{title}}</strong>
      {{summary}}
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import YelpStars from './YelpStars.vue'
import LoadingIcon from './LoadingIcon.vue'

export default{
  name: 'business-box',
  components: {
    YelpStars,
    LoadingIcon
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
      this.title = ''
      this.summary = ''
      this.summaryLoading = true
      var vm = this
      axios.post("http://localhost:8000/summary",{ // TODO: remove for production
        name: data.name,
        latitude: data.coordinates.latitude,
        longitude: data.coordinates.longitude
      })
        .then(function(response){
          console.log(response.data)
          vm.title = response.data.keyword
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

.image {
  float: left;
  margin-right: 9px;
  margin-bottom: 5px;
}

.price {
  padding-left: 10px;
}

.keyword{
  background-color: #E7FAF8;
}

/*article {
  margin-bottom: 5px;
}*/

</style>
