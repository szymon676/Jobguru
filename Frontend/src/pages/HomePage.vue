<script lang="ts">
import HeaderComponent from '@/components/HeaderComponent.vue'
import { defineComponent } from 'vue'

export default defineComponent({
  components: {
    headerComponent: HeaderComponent
  },
  methods: {
    handleDarkModeToggle() {
      const googleBtn = document.querySelector('.googleBtn')
      const ghBtn = document.querySelector('.githubBtn')
      const accesBtn = document.querySelector('.accesBtn')
      const emailInput = document.querySelector('.email')

      googleBtn?.classList.toggle('DarkToggle')
      ghBtn?.classList.toggle('DarkToggle')
      accesBtn?.classList.toggle('DarkToggle')
      emailInput?.classList.toggle('DarkToggle')
    }
  }
})
</script>

<template>
  <div>
    <headerComponent @darkModeToggle="handleDarkModeToggle" />
    <div class="acces-wrapper">
      <div class="acces">
        <h2>Acces</h2>
        <p>Welcome to Jobguru!</p>
        <input type="email" placeholder="email" class="email" />
        <button class="accesBtn">Acces</button>
        <p class="pAcces">Or acces with:</p>
        <div class="login-methods">
          <button class="githubBtn">
            <img src="../assets/github.svg" alt="github logo" /> Github
          </button>
          <button class="googleBtn">
            <img src="../assets/google.svg" alt="Google logo" /> Google
          </button>
        </div>
        <h3>First time here?</h3>
        <p>Our platform uses email based no pasword login to speed up the login process!</p>
      </div>
    </div>
  </div>
</template>

<style>
body {
  background-color: #121212;
  color: #fff;
  font-family: 'Golos Text', sans-serif;
  overflow-x: hidden;
}

.acces {
  text-align: left;
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
  flex-direction: column;
  margin-top: 40px;
  padding: 10px;
}

.acces-wrapper {
  display: flex;
  justify-content: center;
}

.acces h2 {
  font-size: 40px;
  margin-bottom: 10px;
}

.acces input {
  background-color: #1f1f1f;
  border: #1f1f1f 1px solid;
  border-radius: 20px;
  color: #a0a0a0;
  height: 20px;
  width: 300px;
  padding: 10px;
  font-size: 17px;
}

.accesBtn,
.githubBtn,
.googleBtn {
  padding: 10px 20px;
  border-radius: 20px;
  background-color: #1f1f1f;
  border: #1f1f1f 1px solid;
  color: #b6b6b6;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  font-size: 15px;
  cursor: pointer;
}

.googleBtn img,
.githubBtn img {
  height: 25px;
  margin-right: 5px;
}

.acces p {
  margin-bottom: 13px;
}

.githubBtn {
  margin-right: 20px;
}

.acces .pAcces {
  margin-top: 10px;
  margin-bottom: -3px;
  font-size: 20px;
}

.login-methods {
  margin-bottom: 10px;
  display: flex;
}

.DarkToggle {
  background-color: #fff !important;
  color: #121213 !important;
}
</style>
