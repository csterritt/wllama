<template>
  <div>Results Pane</div>

  <div class="font-monocode">
    <span v-if="store.waiting" class="italic"> Waiting for results... </span>

    <span v-else>
      {{ store.results }}
    </span>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()
</script>
