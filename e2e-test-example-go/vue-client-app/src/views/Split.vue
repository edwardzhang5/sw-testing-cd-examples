<template>
  <div class="home">
    <split-tip msg="Split the Tip"/>
  </div>
</template>

<script>
// @ is an alias to /src
import SplitTip from '@/components/SplitTip.vue';

export default {
  name: 'split',
  components: {
    SplitTip,
  },
};
</script>
