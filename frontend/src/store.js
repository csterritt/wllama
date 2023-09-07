import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

import { PromptForResponse } from '../wailsjs/go/main/App.js'
import * as constants from './constants.js'

export const useStore = defineStore('store', () => {
  const paneState = ref(constants.EQUAL_PANE_STATE)
  const prompt = ref('')
  const results = ref('')
  const waiting = ref(false)

  const sendPrompt = async () => {
    if (prompt.value.trim() !== '') {
      waiting.value = true
      results.value = await PromptForResponse(prompt.value.trim())
      waiting.value = false
    }
  }

  const entryPaneGrowClass = computed(() => {
    if (paneState.value !== constants.MOSTLY_BOTTOM_PANE_STATE) {
      return constants.EXPAND_CLASSES
    } else {
      return constants.NO_EXPAND_CLASSES
    }
  })

  const resultsPaneGrowClass = computed(() => {
    if (paneState.value !== constants.MOSTLY_TOP_PANE_STATE) {
      return constants.EXPAND_CLASSES
    } else {
      return constants.NO_EXPAND_CLASSES
    }
  })

  const togglePaneState = () => {
    if (paneState.value === constants.EQUAL_PANE_STATE) {
      paneState.value = constants.MOSTLY_BOTTOM_PANE_STATE
    } else if (paneState.value === constants.MOSTLY_BOTTOM_PANE_STATE) {
      paneState.value = constants.EQUAL_PANE_STATE
    }
  }

  return {
    entryPaneGrowClass,
    paneState,
    prompt,
    results,
    resultsPaneGrowClass,
    sendPrompt,
    togglePaneState,
    waiting,
  }
})
