<template>

<span>
  <img :src="ratingImg"/>
  <small v-if="showTotal">{{total}} Reviews</small>
</span>

</template>

<script>
export default {
  name: "yelp-stars",
  props: {
    rating: {
      type: Number
    },
    total: {
      type: Number
    },
    showTotal: {
      type: Boolean,
      default: true
    }
  },
  computed: {
    ratingImg: function(){
      let a = (''+this.rating).charAt(0), b=''
      if ((''+this.rating).length>1) b='_half'
      return '/static/img/yelp_stars/small_'+a+b+'.png'
    }
  }
}
</script>
