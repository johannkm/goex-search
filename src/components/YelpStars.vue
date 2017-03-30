<template>

<div>
  <img :src="ratingImg"/>
  <small>{{total}} Reviews</small>
</div>

</template>

<script>
export default {
  name: "yelp-stars",
  props: ['rating','total'],
  computed: {
    ratingImg: function(){
      let a = (''+this.rating).charAt(0), b=''
      if ((''+this.rating).length>1) b='_half'
      return '/static/img/yelp_stars/small_'+a+b+'.png'
    }
  }
}
</script>
