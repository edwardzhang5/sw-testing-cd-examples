<template>
  <div class="container">
    <div class="column is-two-thirds">
      <h1 class="title">{{ msg }}</h1>
      <p>Enter the Total Amount for the Bill and the Number of Guests.
        15% Gratuity will be added and the amount split between the guests
        as evenly as possible.<br><br></p>
      <div class="field">
        <label>Bill Total: <span class="has-text-success">${{ billTotal }}</span></label>
        <input class="input" id="cost" v-model="cost" placeholder="Enter the Bill Total $0.00">
      </div>
      <p>&nbsp;</p>
      <label>Number of Guests</label>
      <input class="input" id="guests" v-model="guests" placeholder="Enter the Bill Total $0.00">
      <h3 class="subtitle">Splitting Tip Between {{guests}} guest(s)</h3>
      <button class="button" id="submit" @click="onSubmit">Submit</button>
      <div v-show="splits">
        <hr>
        <ul>
          <li id='#splits' v-for="(split, index) in splits" :key="index">
            Guest {{index + 1}} - ${{split}}

          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SplitTip',
  props: {
    msg: String,
  },
  data() {
    return {
      guests: 0.0,
      cost: 0.0,
      splits: [],
      output: {},
    };
  },
  methods: {
    onSubmit() {
      axios({
        method: 'GET',
        url: 'http://localhost:9090/api/tip',
        params: {
          cost: this.cost,
          guests: this.guests,
        },
      })
        .then(res => {
          console.log(res.data);
          this.splits = res.data.splits;
        })
        .catch(err => console.log(err));
    },
  },
  computed: {
    billTotal() {
      let num = this.cost * 1.15;
      return num.toFixed(2);
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 30px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
