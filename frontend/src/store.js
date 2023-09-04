import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

import { PromptForResponse } from '../wailsjs/go/main/App.js'
import * as constants from './constants.js'

export const useStore = defineStore('store', () => {
  const paneState = ref(constants.EQUAL_PANE_STATE)
  const prompt = ref('')
  const results = ref('')

  const sendPrompt = async () => {
    if (prompt.value.trim() !== '') {
      results.value = await PromptForResponse(prompt.value.trim())
    }
  }

  const entryPaneGrowClass = computed(() => {
    if (paneState.value !== constants.MOSTLY_BOTTOM_PANE_STATE) {
      return 'flex-grow min-w-full'
    } else {
      return 'min-w-full'
    }
  })

  const resultsPaneGrowClass = computed(() => {
    if (paneState.value !== constants.MOSTLY_TOP_PANE_STATE) {
      return 'flex-grow min-w-full'
    } else {
      return 'min-w-full'
    }
  })

  const togglePaneState = () => {
    if (paneState.value === constants.EQUAL_PANE_STATE) {
      paneState.value = constants.MOSTLY_TOP_PANE_STATE
    } else if (paneState.value === constants.MOSTLY_TOP_PANE_STATE) {
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
  }
})
