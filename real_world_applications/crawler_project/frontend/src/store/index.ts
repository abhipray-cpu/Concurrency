// stores/counterStore.ts
import { defineStore } from 'pinia'
import actions from './actions'
import getters from './getters'
export const useCrawlerStore = defineStore('counter', {
  state: () => ({
    auth: false
  }),
  actions: actions,
  getters: getters
})
