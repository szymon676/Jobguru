<template>
  <div class="header">
    <div class="left">
      <h2>JOBGURU</h2>
      <img src="../assets/home.svg" alt="" />
    </div>
    <form>
      <input type="text" placeholder="e.g. Junior java developer" />
      <button><img src="../assets/magnifier.svg" alt="" /></button>
    </form>
  </div>
</template>

<script setup lang="ts"></script>

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
  color: #0fff50;
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
}

.header input {
  height: 10px;
  margin: 10px;
  border-radius: 15px;
  border: 1px solid #0fff50;
  padding: 10px;
  width: 100%;
  font-size: 16px;
  background-color: transparent;
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
  height: 50px;
  width: 50px;
  margin-right: 20px;
  outline: none;
  cursor: pointer;
}
</style>
