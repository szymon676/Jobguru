<template>
  <div class="header">
    <div class="left">
      <h2>JOBGURU</h2>
      <img src="../assets/home.svg" alt="home icon" />
      <img src="../assets/job.svg" alt="bell icon" />
    </div>

    <form>
      <button class="darkModeBtn" @click.prevent="darkModeToggle">
        <img src="../assets/icon-moon-light.svg" alt="a moon icon" class="darkModeImg" />
      </button>
      <input type="text" placeholder="e.g. Junior java developer" />
      <button><img src="../assets/magnifier.svg" alt="Magnifier icon" /></button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'

export default defineComponent({
  name: 'HeaderComponent',

  methods: {
    darkModeToggle() {
      const darkModeBtn = document.querySelector('.darkModeBtn')
      const darkModeImg = document.querySelector('.darkModeImg')
      const body = document.querySelector('body')
      const header = document.querySelector('.header')
      const input = document.querySelector('input')
      const homeIcon = document.querySelector('.left img')
      const bellIcon = document.querySelector('.left img:nth-child(2)')
      const magnifierIcon = document.querySelector('button:nth-child(3) img')
      const jobGuru = document.querySelector('.header h2')

      darkModeBtn?.addEventListener('click', () => {
        darkModeBtn.classList.toggle('DarkToggle')
        darkModeImg?.classList.toggle('DarkToggle')
        body?.classList.toggle('DarkToggle')
        header?.classList.toggle('DarkToggle')
        input?.classList.toggle('DarkToggle')
        homeIcon?.classList.toggle('DarkToggle')
        bellIcon?.classList.toggle('DarkToggle')
        magnifierIcon?.classList.toggle('DarkToggle')
        jobGuru?.classList.toggle('DarkToggle')
      })
    }
  }
})
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Golos+Text:wght@700&display=swap');

.header {
  background-color: #121213;
  width: 100vw;
  font-family: 'Golos Text', sans-serif;
  display: flex;
  justify-content: space-between;
  padding-left: 10px;
  padding-right: 10px;
  border-bottom: #3a3a3c 1px solid;
  max-height: 50px;
}

.header h2 {
  color: #fff;
  text-align: left;
  padding: 10px;
}

.header .left {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-left: 10px;
}

.header .left img {
  margin-left: 10px;
  cursor: pointer;
  height: 25px;
}

.header input {
  height: 10px;
  margin: 10px;
  border-radius: 15px;
  border: 1px solid #1f1f1f;
  padding: 10px;
  width: 100%;
  font-size: 16px;
  background-color: #1f1f1f;
  outline: none;
  color: #fff;
}

.header form {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.header button {
  background-color: transparent;
  border: none;
  border-radius: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  width: 50px;
  margin-right: 20px;
  outline: none;
  cursor: pointer;
}

.darkModeBtn {
  background-color: #1f1f1f;
  border: none;
  border-radius: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  width: 50px;
  margin-right: 20px;
  outline: none;
  cursor: pointer;
}

.DarkToggle {
  background-color: #fff !important;
  color: #121213 !important;
}
</style>
