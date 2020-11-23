import Vue from 'vue';
import Router from 'vue-router';
import Split from './views/Split.vue';
import Retire from './views/Retire.vue';
import Email from './views/Email.vue';
import Distance from './views/Distance.vue';
import BMI from './views/BMI.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'split',
      component: Split,
    },
    {
      path: '/bmi',
      name: 'bmw',
      component: BMI,
    },
    {
      path: '/retire',
      name: 'retire',
      component: Retire,
    },
    {
      path: '/email',
      name: 'email',
      component: Email,
    },
    {
      path: '/distance',
      name: 'distance',
      component: Distance,
    },
  ],
});
