<template>
  <div class="flex flex-col flex-grow min-w-full">
    <entry-pane></entry-pane>

    <toggle-pane></toggle-pane>

    <results-pane></results-pane>
  </div>
</template>

<script setup>
import EntryPane from './EntryPane.vue'
import ResultsPane from './ResultsPane.vue'
import TogglePane from './TogglePane.vue'
</script>
