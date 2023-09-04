<template>
  <div class="flex-grow min-w-full">
    <div class="font-monocode">
      <div>Results Pane</div>

      <div class="flex flex-row">
        <div class="flex-grow">{{ store.results }}</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()
</script>
