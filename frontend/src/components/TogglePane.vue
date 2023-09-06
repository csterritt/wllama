<template>
  <div class="divider">
    <button class="btn btn-square btn-accent" @click="store.togglePaneState">
      <img
        alt="toggle panes"
        class="w-8 h-8"
        src="../assets/icons/arrow-up-341-svgrepo-com.svg"
      />
    </button>
  </div>
</template>

<script setup>
import { useStore } from '../store.js'

const store = useStore()
</script>
