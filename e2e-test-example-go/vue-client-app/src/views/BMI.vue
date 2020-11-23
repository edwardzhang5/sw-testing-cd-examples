<template>
  <div class="home">
    <div class="container">
      <div class="column is-two-thirds">
      <h1 id="bmi-title" class="title">Body Mass Index</h1>
        <p></p>
        <div class="field">
          <label>Enter Height in Feet</label>
          <input class="input" @focus="setSubmitted" id="hFeet" v-model="hFeet" placeholder="Height in Feet">
          <label>Enter Height in Inches</label>
          <input class="input" @focus="setSubmitted" id="hInches" v-model="hInches" placeholder="Height in Inches">
          <label>Enter Weight</label>
          <input class="input" @focus="setSubmitted" id="weight" v-model="weight" placeholder="Weight">
        </div>
        <p>&nbsp;</p>
        <button class="button" id="submit" @click="onSubmit">Submit</button>
        <div v-if="message">
          <hr>
          <p class="has-text-info">{{ message }} - BMI Result: {{ bmi }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'bmi',
  data() {
    return {
      hInches: 0.0,
      hFeet: 0.0,
      weight: 0.0,
      message: '',
      bmi: 0.0,
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
        url: 'http://localhost:9090/api/bmi',
        params: {
          hInches: this.hInches,
          hFeet: this.hFeet,
          weight: this.weight,
        },
      })
        .then(res => {
          console.log(res.data);
          this.message = res.data.message;
          this.bmi = res.data.BMI;
        })
        .catch(err => console.log(err));
    },
  },
};
</script>
