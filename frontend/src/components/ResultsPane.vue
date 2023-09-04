<template>
  <div :class="store.resultsPaneGrowClass">
    <div>Results Pane</div>

    <div class="flex flex-row flex-grow">
      <div class="h-full w-full font-monocode overflow-auto mr-4">
        {{ store.results }}
      </div>

      <button class="btn btn-primary" @click="copyText">Copy</button>
    </div>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()

const copyText = () => {
  console.log(`Entered copyText`)
}
</script>
