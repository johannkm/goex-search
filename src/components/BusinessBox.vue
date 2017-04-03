<template>
  <div class="busi-box">

    <a class="main-link" :href="businessData.url">
      <div class="image is-64x64" align="left">
        <div class="thumb" :style="{ backgroundImage: 'url(' + businessData.image_url + ')' }"></div>
      </div>
      <strong>{{businessData.name}}</strong>
    </a>
    <span class="tag">
      {{businessData.categories[0].title}}
    </span>
    <span v-if="businessData.is_closed" class="tag is-danger">
      Closed
    </span>
    <br>

    <yelp-stars :rating="businessData.rating" :total="businessData.review_count"></yelp-stars>

    <span class="price">
      <span v-for="x in businessData.price">
        $
      </span>
    </span>

    <div class="reviewSummary">
      <loading-icon :active="summaryLoading" :isSmall="true"/>
      <strong class="keyword" :style="{ 'background-color': sentimentColor }">{{title}}</strong>
      {{summary}}
    </div>

    <span class="float-right">
      <small>
        <a class="review-button" @click="toggleReviews">
          Show Reviews
          <span class="icon">
            <i class="fa" :class="{'fa-chevron-down': !reviewsExpanded, 'fa-chevron-up': reviewsExpanded}"></i>
          </span>
        </a>
      </small>
    </span>

    <div v-if="reviewsExpanded">
      <hr>
      <review-box class="reviews" v-for="b in reviews" :review="b" />
    </div>

  </div>
</template>

<script>
import axios from 'axios'
import YelpStars from './YelpStars.vue'
import LoadingIcon from './LoadingIcon.vue'
import ReviewBox from './ReviewBox.vue'

export default{
  name: 'business-box',
  components: {
    YelpStars,
    LoadingIcon,
    ReviewBox
  },
  props: ['businessData','key'],
  data: function() {
    return {
      summaryLoading: 'false',
      title: '',
      summary: '',
      titleSentiment: '',
      reviews: [],
      reviewsExpanded: false
    }
  },
  computed: {
    sentimentColor: function(){
      if (this.titleSentiment>0.1){
        return '#E7FAF8'
      }
      if( this.titleSentiment<0.1){
        return '#F6D5A9'
      }
      return '#F4F4F4'
    }
  },
  methods: {
    getSummary: function(data){
      this.title = ''
      this.summary = ''
      this.reviews = []
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
          vm.titleSentiment = response.data.keyword_sentiment
          vm.summaryLoading = false
          vm.reviews = response.data.google_place_review.reviews
        })
        .catch(function(error){
          console.error(error)
          vm.summaryLoading = false
        })
    },
    toggleReviews: function(){
      this.reviewsExpanded = !this.reviewsExpanded
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
  white-space: nowrap;
}

.keyword{
  text-transform: capitalize;
}

.reviewSummary {
  clear: both;
}

.main-link:hover{
  text-decoration: underline!important;
}

hr {
  margin-bottom: 0;
  margin-top: 12px;
}

.reviews{
  margin-top: 1em;
  margin-bottom: 1em;
}

.review-button{
  color: #000!important;
  margin-top: 1em;
}

.icon{
  color: #A3A3A3;
  font-size: 11px;
}

/*article {
  margin-bottom: 5px;
}*/

</style>
