import { offerExamples } from '@/data/examples/offerExamples'
import { defineStore } from 'pinia'
import { getSkillsWithCounter } from './helpers/getSkillsWithCounter'

export const useOfferStore = defineStore('offers', {
  state: () => {
    return { listOffer: offerExamples, selectedOffers: [] as Array<string> }
  },
  actions: {
    toggleSelectedOffer(nameOffer: string) {
      if (this.selectedOffers.includes(nameOffer)) {
        this.selectedOffers = this.selectedOffers.filter((offer) => offer !== nameOffer)
      } else {
        this.selectedOffers.push(nameOffer)
      }
    }
  },
  getters: {
    getSkillList: (state) => getSkillsWithCounter(state.listOffer)
  }
})
