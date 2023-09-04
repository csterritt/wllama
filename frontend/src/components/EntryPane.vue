<template>
  <div class="flex-grow min-w-full">
    <div>Text Entry Pane</div>

    <div class="flex flex-row">
      <input
        type="text"
        placeholder="Type here"
        class="input input-bordered flex-grow mr-4"
        v-model="store.prompt"
        @keyup.enter="sendOut"
      />

      <button class="btn btn-primary" @click="sendOut">Send</button>
    </div>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()

const sendOut = () => {
  store.sendPrompt()
}
</script>
