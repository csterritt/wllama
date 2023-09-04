<template>
  <div class="flex flex-col flex-grow min-w-full">
    <entry-pane></entry-pane>

    <div class="flex flex-row">
      <div class="flex-grow divider">
        <button class="btn btn-secondary" @click="store.togglePaneState">
          T
        </button>
      </div>
    </div>
    <results-pane></results-pane>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

import EntryPane from './EntryPane.vue'
import ResultsPane from './ResultsPane.vue'

const store = useStore()
</script>
