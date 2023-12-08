import run from "aocrunner"

interface Card {
  id: number
  wins: Set<number>
  nums: Array<number>
  count: number
}

const parseInput = (rawInput: string): Array<Card> => {
  const cards = []

  const lines = rawInput.split("\n")
  for (const line of lines) {
    const card = {
      id: NaN,
      wins: new Set(),
      nums: [],
      count: 1
    } as Card
    
    card.id = parseInt(line.substring(5, line.indexOf(':')))

    const nums = line.substring(line.indexOf(':') + 2).split(' | ')
    card.wins = new Set(
      nums[0].split(' ').map((val: string): number => {
        return parseInt(val, 10)
      }).filter((val: number) => !Number.isNaN(val))
    )
    card.nums = nums[1].split(' ').map((val: string): number => {
      return parseInt(val, 10)
    }).filter((val: number) => !Number.isNaN(val))

    cards.push(card)
  }

  return cards
}

const part1 = (rawInput: string): number => {
  const cards = parseInput(rawInput)

  let total = 0
  for (const card of cards) {
    let score = 0
    for (const num of card.nums) {
      if (card.wins.has(num)) {
        if (score === 0) {
          score = 1
        } else {
          score *= 2
        }
      }
    }

    total += score
  }

  return total
}

const part2 = (rawInput: string): number => {
  const cards = parseInput(rawInput)

  for (let i = 0; i < cards.length; i++) {
    const card = cards[i]

    let score = 0
    for (const num of card.nums) {
      if (card.wins.has(num)) {
        score++
      }
    }

    for (let j = i + 1; j < i + score + 1; j++) {
      cards[j].count += card.count
    }
  }

  let total = 0
  for (const card of cards) {
    total += card.count
  }

  return total
}

run({
  part1: {
    tests: [
      {
        input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
        expected: 13,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
        expected: 30,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
