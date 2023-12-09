import run from "aocrunner"

const parseInput = (rawInput: string): Array<Array<number>> => {
  return rawInput.split("\n").map((line: string): Array<number> => {
    return line.split(" ").map((val: string) => parseInt(val.trim(), 10))
  })
}

const getSums = (vals: Array<number>): Array<number> => {
  const sums = [] as Array<number>

  for (let i = 1; i < vals.length; i++) {
    sums.push(vals[i] - vals[i - 1])
  }
  return sums
}

const isZeroes = (vals: Array<number>): boolean => {
  return vals.filter((val: number): boolean => val !== 0).length === 0
}

const nextValue = (vals: Array<number>, delta: number): number => {
  return vals[vals.length - 1] + delta
}

const makeStack = (vals: Array<number>): Array<Array<number>> => {
  let cur = vals
  const stack = [vals]
  while (true) {
    cur = getSums(cur)

    stack.push(cur)
    if (isZeroes(cur)) {
      break
    }
  }
  return stack
}

const findNext = (vals: Array<number>): number => {
  const stack = makeStack(vals)

  // add the extra zero to the end
  let newVal = 0
  for (let i = stack.length - 2; i >= 0; i--) {
    newVal = nextValue(stack[i], newVal)
  }
  return newVal
}

const part1 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  let total = 0
  for (const vals of input) {
    total += findNext(vals)
  }
  return total
}

const previousValue = (vals: Array<number>, delta: number): number => {
  return vals[0] - delta
}

const findPrevious = (vals: Array<number>): number => {
  const stack = makeStack(vals)

  // add the extra zero to the beginning
  let newVal = 0
  for (let i = stack.length - 2; i >= 0; i--) {
    newVal = previousValue(stack[i], newVal)
  }
  return newVal
}

const part2 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  let total = 0
  for (const vals of input) {
    total += findPrevious(vals)
  }
  return total
}

run({
  part1: {
    tests: [
      {
        input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`,
        expected: 114,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`,
        expected: 2,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
