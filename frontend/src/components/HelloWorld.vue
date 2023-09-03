<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input
        id="name"
        v-model="data.name"
        autocomplete="off"
        class="input"
        type="text"
        @keyup.enter="greet"
      />
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>

<script setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'

const data = reactive({
  name: '',
  resultText: 'Please enter your name below 👇',
})

function greet() {
  Greet(data.name).then((result) => {
    data.resultText = result
  })
}
</script>
