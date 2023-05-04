<template>
  <div class="header">
    <div class="left">
      <h2>JOBGURU</h2>
      <img src="../assets/home.svg" alt="home icon" />
      <img src="../assets/job.svg" alt="bell icon" />
    </div>

    <form>
      <div id="dark-mode-switch">
        <input
          type="checkbox"
          id="switch"
          class="darkMode-checkbox"
          name="switch"
          @click.prevent="darkModeToggle"
        />
        <label for="switch" class="label"></label>
      </div>
      <input type="text" placeholder="e.g. Junior java developer" class="input" />
      <button><img src="../assets/magnifier.svg" alt="Magnifier icon" class="magnifier" /></button>
    </form>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'HeaderComponent',

  methods: {
    darkModeToggle() {
      const darkModeBtn = document.querySelector('.darkMode-checkbox')
      const body = document.querySelector('body')
      const header = document.querySelector('.header')
      const input = document.querySelector('.input')
      const homeIcon = document.querySelector('.left img')
      const bellIcon = document.querySelector('.left img:nth-child(2)')
      const magnifierIcon = document.querySelector('button:nth-child(3) img')
      const jobGuru = document.querySelector('.header h2')
      const magnifier = document.querySelector('.magnifier')

      darkModeBtn?.classList.toggle('DarkToggle')
      body?.classList.toggle('DarkToggle')
      header?.classList.toggle('DarkToggle')
      input?.classList.toggle('DarkToggle')
      homeIcon?.classList.toggle('DarkToggle')
      bellIcon?.classList.toggle('DarkToggle')
      magnifierIcon?.classList.toggle('DarkToggle')
      jobGuru?.classList.toggle('DarkToggle')

      this.$emit('dark-mode-toggle')
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
  max-height: 65px;
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
  height: 7px;
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

.label {
  width: 60px;
  height: 24px;
  position: relative;
  display: block;
  background: #525252;
  border-radius: 30px;
  box-shadow: inset 0px 5px 15px rgba(0, 0, 0, 0.4), inset 0px -5px 15px rgba(255, 255, 255, 0.3);
  cursor: pointer;
  transition: 0.3s;
}

.label:after {
  content: url('../assets/icon-moon-light.svg');
  width: 18px;
  height: 17px;
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  top: 3px;
  left: 7px;
  background: #ffffff;
  border-radius: 50%;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.2);
  transition: 0.3s;
}

#switch {
  display: none;
}

#switch:checked + .label {
  background: #ebebeb;
}

input:checked + label:after {
  left: 30px;
  transform: translateX(-100%);
  background: linear-gradient(180deg, #777, #3a3a3a);
}
</style>
