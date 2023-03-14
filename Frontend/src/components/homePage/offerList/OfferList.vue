<script setup lang="ts">
import { offerExamples } from '@/data/examples/offerExamples'
import SingleOffer from './SingleOffer.vue'
</script>

<template>
  <div class="list">
    <SingleOffer v-for="offer in offerExamples" v-bind="offer" :key="offer.id" />
  </div>
</template>
<style lang="scss"></style>
