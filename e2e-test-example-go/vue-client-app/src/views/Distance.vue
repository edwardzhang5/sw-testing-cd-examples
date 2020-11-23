<template>
  <div class="home">
    <div class="container">
      <div class="column is-two-thirds">
        <h1 id="bmi-title" class="title">Distance Formula</h1>
        <p></p>
        <div class="field">
          <label>Enter x1, x2, y1, y2</label>
          <input class="input"  id="x1" v-model="x1"
                 placeholder="x1">
          <input class="input" id="x2" v-model="x2"
                 placeholder="x2">
          <input class="input" id="y1" v-model="y1"
                 placeholder="y1">
          <input class="input" id="y2" v-model="y2"
                 placeholder="y2">
        </div>
        <p>&nbsp;</p>
        <button class="button" id="submit" @click="onSubmit">Submit</button>
        <div v-if="result">
          <hr>
          <h3 class="has-text-info"> {{ message }} </h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'distance',
  data() {
    return {
      x1: 0.0,
      x2: 0.0,
      y1: 0.0,
      y2: 0.0,
      result: 0.0,
      message: '',
      submitted: false,
    };
  },
  methods: {
    setSubmitted() {
      this.submitted = false;
    },
    onSubmit() {
      this.submitted = true;
      axios({
        method: 'GET',
        url: 'http://localhost:9090/api/distance',
        params: {
          x1: this.x1,
          x2: this.x2,
          y1: this.y1,
          y2: this.y2,
        },
      })
        .then(res => {
          console.log(res.data);
          this.message = res.data.message;
          this.result = res.data.result;
        })
        .catch(err => console.log(err));
    },
  },
};
</script>
