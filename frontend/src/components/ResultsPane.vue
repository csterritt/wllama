<template>
  <div>Results Pane</div>

  <div class="font-monocode">
    {{ store.results }}
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()
</script>
