import type { OfferExamplesType } from '@/data/examples/offerExamples'

export function getSkillsWithCounter(offers: OfferExamplesType) {
  const listWithoutCounting = [...new Set(offers.map((offer) => offer.skills).flat())]
  return listWithoutCounting.map((skill) => {
    const appearanceCounter = offers.reduce((acc, value) => {
      return value.skills.includes(skill) ? acc + 1 : acc
    }, 0)

    return { name: skill, counter: appearanceCounter }
  })
}
