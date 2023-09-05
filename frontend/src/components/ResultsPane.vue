<template>
  <div :class="store.resultsPaneGrowClass">
    <div>Results Pane</div>

    <div class="flex flex-row flex-grow">
      <div class="h-full w-full font-monocode overflow-auto mr-4">
        {{ store.results }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()
</script>
