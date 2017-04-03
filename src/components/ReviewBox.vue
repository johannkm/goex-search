<template>

    <div class="columns">
      <div class="column is-1 no-pad">
        <small>{{postedDate}}</small>
      </div>

      <div class="column is-11 no-pad">

        <article class="media">

          <div class="media-content">
            <div class="content">
              <p>
                <strong>{{review.author_name}}</strong>
                <br>
                  <yelp-stars :rating="review.rating" :showTotal="false"></yelp-stars>
                <br>
                {{shortText}}
                <a class="read-more" @click="toggleExpanded" v-if="needsReadMore">
                  {{moreOrLess}}
                </a>
                <br>
              </p>
            </div>
          </div>

        </article>
      </div>
    </div>
</template>

<script>
import YelpStars from './YelpStars.vue'

export default{
  name: 'comment-box',
  components: {
    YelpStars
  },
  props: ['review'],
  methods : {
    secondsToDate: function(secs){
      var t = new Date(1970, 0, 1); // Epoch
      t.setSeconds(secs);
      return t;
    },
    toggleExpanded: function(){
      this.expanded = !this.expanded
    }
  },
  data: function(){
    return{
      expanded: false,
      monthNames: [
        'Jan','Feb','Mar','Apr','May','June','July','Aug','Sep','Oct','Nov','Dec'
      ]
    }
  },
  computed: {
    postedDate: function(){
      let date = this.secondsToDate(this.review.time)
      return this.monthNames[date.getMonth()]+' '+date.getDate()+', '+date.getFullYear()
    },

    splitText: function(){
      return this.review.text.trim().replace(/ +(?= )/g,'').split(' ')
    },
    needsReadMore: function(){
      return ( this.splitText.length > 35 )
    },
    shortText: function(){
      if( this.expanded || !this.needsReadMore ){
        return this.review.text
      }
      let str = ''
      for( let i in this.splitText.slice(0,35)){
        str += this.splitText[i]
        if(i!=34){
          str+=' '
        }
      }
      return str+'...'
    },
    moreOrLess: function(){
      return this.expanded? 'Less' : 'More'
    }
  }
}
</script>

<style scoped>
.no-pad{
  padding-top: 0!important;
  padding-bottom: 0!important;

}

.read-more{
  font-weight: 600;
  color: #000;
}

</style>
