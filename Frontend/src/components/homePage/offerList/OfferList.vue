<script setup lang="ts">
import { offerExamples } from '@/data/examples/offerExamples'
</script>

<template>
  <div class="list">
    <div v-for="offer in offerExamples">
      {{ offer.salary }}
    </div>
  </div>
</template>
<style></style>
