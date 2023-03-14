<script setup lang="ts">
import { defineProps } from 'vue'
import type { SingleOfferType } from '@/data/examples/offerExamples'

const { company, title, salary, skills, description } = defineProps<SingleOfferType>()
</script>

<template>
  <div class="list">
    <div class="singleOffer">
      <h2>{{ title }}</h2>
      <div>{{ company }}</div>
      <p>{{ description }}</p>
      <div class="salary">
        Salary: <span>{{ salary }}</span>
      </div>
      <div class="skills">
        <div v-for="skill in skills" class="skill">{{ skill }}</div>
      </div>
      <button>See more</button>
    </div>
  </div>
</template>
<style lang="scss">
.singleOffer {
  text-align: left;
  width: 100%;
  padding: 15px;
  background-color: white;
  border: 3px solid rgb(163, 163, 163);
  border-radius: 3px;
  max-width: 920px;
  margin: 0 auto 20px;
  box-shadow: 6px -6px 0px -1px rgb(0, 163, 46);
  font-family: 'Roboto', sans-serif;
  transition: all 0.2s ease-in-out;
  cursor: pointer;
  &:hover {
    box-shadow: -6px -6px 0px -1px rgb(0, 163, 46);
  }
  h2 {
    margin: 0 0 5px;
  }
  .skills {
    margin-top: 5px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    padding: 0;
    font-weight: 900;
    font-size: 1.2rem;
    .skill {
      padding-right: 15px;
    }
  }
  .salary {
    margin-top: 20px;
    span {
      color: rgb(0, 163, 46);
      font-weight: 900;
      font-size: 1.2rem;
    }
  }
  button {
    padding: 10px;
    background-color: rgb(0, 163, 46);
    font-weight: 600;
    font-size: 1.2rem;
    color: white;
    margin-top: 20px;
    border: none;
    border-radius: 3px;
  }
}
</style>
