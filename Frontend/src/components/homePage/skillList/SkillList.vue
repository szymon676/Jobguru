<script setup lang="ts">
import { useOfferStore } from '@/stores/offers/useOfferStore'
import { storeToRefs } from 'pinia'
const { getSkillList, toggleSelectedOffer } = useOfferStore()

const { selectedOffers } = storeToRefs(useOfferStore())
</script>

<template>
  <div class="listSkill">
    <button
      v-for="skill in getSkillList"
      @click="toggleSelectedOffer(skill.name)"
      :class="{ select: selectedOffers.includes(skill.name) }"
    >
      {{ skill.name }} ({{ skill.counter }})
    </button>
  </div>
</template>
<style lang="scss">
.listSkill {
  display: flex;
  flex-wrap: wrap;
  padding: 20px;
  justify-content: center;
  gap: 15px;
  button {
    font-size: 1.2rem;
    font-weight: 600;
    font-family: 'Roboto';
    padding: 10px;
    border: 2px solid white;
    background-color: transparent;
    color: white;
    cursor: pointer;
    &.select {
      background-color: green;
    }
  }
}
</style>
