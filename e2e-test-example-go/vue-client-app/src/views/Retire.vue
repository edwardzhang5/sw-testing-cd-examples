<template>
  <div class="home">
    <div class="container">
      <div class="column is-two-thirds">
        <h1 id="bmi-title" class="title">Retirement Planner</h1>
        <p></p>
        <div class="field">
          <label>Enter Your Age: {{ age }}</label>
          <input class="input" id="age" v-model="age"
                 placeholder="Age">
          <label>Enter Your Current Salary: <span>{{ salary | money }}</span></label>
          <input class="input" id="salary" v-model="salary"
                 placeholder="Salary">
          <label>Enter Your Target Amount Saved: <span>{{ saved }}%</span></label>
          <input class="input" id="saved" v-model="saved"
                 placeholder="Saved">
          <label>Enter Your Retirement Goal Amount: <span>{{ goal | money }}</span></label>
          <input class="input" id="goal" v-model="goal"
                 placeholder="Goal">
        </div>
        <p>&nbsp;</p>
        <button class="button" id="submit" @click="onSubmit">Submit</button>
        <div v-if="reached">
          <hr>
          <h3 class="has-text-info"> Congrats...Goal of {{ goal | money }} reached at {{ retirementAge }} years old! </h3>
        </div>
        <div v-else-if="!reached && submitted">
          <hr>
          <h3 class="has-text-danger"> Goal of {{ goal | money }} *not* reached, sorry. </h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import formatMoney from 'accounting-js/lib/formatMoney.js';

export default {
  name: 'retire',
  data() {
    return {
      age: 0,
      salary: 0,
      saved: 0,
      goal: 0,
      message: '',
      reached: false,
      submitted: false,
      retirementAge: 0,
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
        url: 'http://localhost:9090/api/retire',
        params: {
          age: this.age,
          salary: this.salary,
          saved: this.saved,
          goal: this.goal,
        },
      })
        .then(res => {
          console.log(res.data);
          this.message = res.data.message;
          this.reached = res.data.reached;
          this.retirementAge = res.data.retAge;
        })
        .catch(err => console.log(err));
    },
  },
  filters: {
    money(value) {
      return formatMoney(value);
    },
  },
};
</script>
