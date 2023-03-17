<script setup lang="ts">
import headerComponent from '@/pages/HeaderComponent.vue'
</script>

<template>
  <headerComponent />
  <div class="acces-wrapper">
    <div class="acces">
      <h2>Acces</h2>
      <p>Welcome to Jobguru!</p>
      <input type="email" placeholder="email" />
      <button class="accesBtn">Acces</button>
      <p class="pAcces">Or acces with:</p>
      <div class="login-methods">
        <button class="githubBtn">Github</button>
        <button class="googleBtn">Google</button>
      </div>
      <h3>First time here?</h3>
      <p>Our platform uses email based no pasword login to speed up the login process!</p>
    </div>
  </div>
</template>

<style>
body {
  background-color: #121213;
  color: #0fff50;
  font-family: 'Golos Text', sans-serif;
}

.acces {
  text-align: left;
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
  flex-direction: column;
  margin-top: 40px;
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
  background-color: transparent;
  border: #0fff50 1px solid;
  border-radius: 20px;
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
  background-color: transparent;
  border: #0fff50 1px solid;
  color: #0fff50;
  margin-top: 20px;
  font-size: 15px;
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
}
</style>
