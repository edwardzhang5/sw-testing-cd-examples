<template>
  <div class="home">
    <div class="container">
      <div class="column is-two-thirds">
      <h1 id="bmi-title" class="title">Email Verifier</h1>
        <p></p>
        <div class="field">
          <label>Enter the email address to validate: <span class="has-text-info">{{ email }}</span></label>
          <input class="input" @focus="setSubmitted" id="email" v-model="email" placeholder="Enter Email Address to Verify">
        </div>
        <p>&nbsp;</p>
        <button class="button" id="submit" @click="onSubmit">Submit</button>
        <div v-if="validEmail">
          <hr>
          <p class="has-text-success">{{ email }} <span class="has-test-success"> is valid</span></p>
        </div>
        <div v-if="!validEmail && submitted">
          <hr>
          <p class="has-text-danger">{{ email }} <span class="has-test-danger"> is *NOT* valid</span></p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'email',
  data() {
    return {
      email: '',
      validEmail: false,
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
        url: 'http://localhost:9090/api/email',
        params: {
          email: this.email,
        },
      })
        .then(res => {
          console.log(res.data);
          this.validEmail = res.data.valid;
        })
        .catch(err => console.log(err));
    },
  },
};
</script>

<style>
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
