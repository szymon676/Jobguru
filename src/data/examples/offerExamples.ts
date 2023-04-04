export type SingleOfferType = {
  id: number
  title: string
  company: string
  skills: Array<string>
  salary: string
  description: string
}

export type OfferExamplesType = Array<SingleOfferType>

export const offerExamples: OfferExamplesType = [
  {
    id: 1,
    title: 'go backend developer',
    company: 'januszex123',
    skills: ['go', 'postgresql', 'redis'],
    salary: '30000',
    description: 'backend developer needed'
  },
  {
    id: 2,
    title: 'typescript developer',
    company: 'OGDEVS123',
    skills: ['js', 'ts', 'mongodb'],
    salary: '20000',
    description: 'typescript developer asap'
  },
  {
    id: 3,
    title: 'javascript developer',
    company: 'OGDEVS123',
    skills: ['js', 'mongodb'],
    salary: '20000',
    description: 'javascript developer asap'
  },
  {
    id: 4,
    title: 'rust developer',
    company: 'OGDEVS123',
    skills: ['rust', 'mongodb'],
    salary: '25000',
    description: 'rust developer asap'
  }
]
