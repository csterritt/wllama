<template>
  <div :class="store.entryPaneGrowClass">
    <div>Text Entry Pane</div>

    <div class="flex flex-col flex-grow">
      <div class="flex flex-row flex-grow">
        <div class="flex flex-col flex-grow">
          <textarea
            placeholder="Type here"
            class="input input-bordered flex-grow mr-4 mb-4"
            v-model="store.prompt"
            @keyup.enter="sendOut"
          >
          </textarea>
        </div>

        <button class="btn btn-primary" @click="sendOut">Send</button>
      </div>
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
