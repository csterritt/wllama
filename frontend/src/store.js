import { ref } from 'vue'
import { defineStore } from 'pinia'

import { PromptForResponse } from '../wailsjs/go/main/App.js'

export const useStore = defineStore('store', () => {
  const prompt = ref('')
  const results = ref('')

  const sendPrompt = async () => {
    if (prompt.value.trim() !== '') {
      results.value = await PromptForResponse(prompt.value.trim())
    }
  }

  return { prompt, results, sendPrompt }
})
