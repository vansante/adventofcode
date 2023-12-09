import run from "aocrunner"

const cardTypes1 = [
  "A",
  "K",
  "Q",
  "J",
  "T",
  "9",
  "8",
  "7",
  "6",
  "5",
  "4",
  "3",
  "2",
].reverse()

const cardTypes2 = [
  "A",
  "K",
  "Q",
  "T",
  "9",
  "8",
  "7",
  "6",
  "5",
  "4",
  "3",
  "2",
  "J",
].reverse()

interface Hand {
  cards: string
  bid: number
  strength: number
}

const parseInput = (rawInput: string): Array<Hand> => {
  return rawInput.split("\n").map((line: string): Hand => {
    return {
      cards: line.substring(0, 5),
      bid: parseInt(line.substring(6), 10),
      strength: -1,
    } as Hand
  })
}

const frequencies = (cards: string): Map<string, number> => {
  const freqs = new Map<string, number>()

  for (let i = 0; i < cards.length; i++) {
    if (freqs.has(cards[i])) {
      freqs.set(cards[i], (freqs.get(cards[i]) as number) + 1)
    } else {
      freqs.set(cards[i], 1)
    }
  }

  return freqs
}

const strength = (cards: string): number => {
  const freqs = frequencies(cards)

  switch (freqs.size) {
    case 1:
      return 10 // 5 of a kind
    case 2:
      if (freqs.get(cards[0]) == 2 || freqs.get(cards[0]) == 3) {
        return 8 // full house
      }
      return 9 // 4 of a kind
  }

  let result = null
  let pairCount = 0
  freqs.forEach((freq: number, card: string) => {
    if (freq === 3) {
      result = 7 // 3 of a kind
    }
    if (freq === 2) {
      pairCount++
    }
  })

  if (result) {
    return result
  }

  switch (pairCount) {
    case 2:
      return 6 // 2 pair
    case 1:
      return 5 // pair
  }

  return 4
}

const strongerCard = (
  a: string,
  b: string,
  cardTypes: Array<string>,
): number => {
  for (let i = 0; i < a.length; i++) {
    const aCard = cardTypes.indexOf(a[i])
    const bCard = cardTypes.indexOf(b[i])

    if (aCard > bCard) {
      return 1
    }
    if (bCard > aCard) {
      return -1
    }
  }

  console.error("Same cards!", a, b)
  return 0
}

const solve = (hands: Array<Hand>, cardTypes: Array<string>): number => {
  const sorted = hands.sort((a: Hand, b: Hand): number => {
    if (a.strength > b.strength) {
      return 1
    }
    if (b.strength > a.strength) {
      return -1
    }
    return strongerCard(a.cards, b.cards, cardTypes)
  })

  let total = 0
  sorted.forEach((h: Hand, idx: number) => {
    total += h.bid * (idx + 1)
  })

  return total
}

const part1 = (rawInput: string): number => {
  const hands = parseInput(rawInput)

  hands.forEach((h: Hand) => {
    h.strength = strength(h.cards)
  })

  return solve(hands, cardTypes1)
}

const jokerize = (cards: string): string => {
  if (cards.indexOf("J") === -1) {
    return cards
  }

  const freqs = frequencies(cards)

  let high = 0
  let card = ""
  freqs.forEach((f: number, c: string) => {
    if (f > high && c !== "J") {
      card = c
      high = f
    }
  })
  if (card === "") {
    card = "A"
  }

  return cards.replaceAll("J", card)
}

const part2 = (rawInput: string) => {
  const hands = parseInput(rawInput)

  hands.forEach((h: Hand) => {
    h.strength = strength(jokerize(h.cards))
  })

  return solve(hands, cardTypes2)
}

run({
  part1: {
    tests: [
      {
        input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
        expected: 6440,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
        expected: 5905,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
