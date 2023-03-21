<script setup lang="ts">
import { offerExamples } from '@/data/examples/offerExamples'
import SkillList from '../skillList/SkillList.vue'
import SingleOffer from './SingleOffer.vue'
</script>

<template>
  <SkillList />
  <div class="list">
    <SingleOffer v-for="offer in offerExamples" v-bind="offer" :key="offer.id" />
  </div>
</template>
<style lang="scss">
.list {
  color: black;
}
</style>
