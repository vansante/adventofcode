import run from "aocrunner"

const parseInput = (rawInput: string): Array<string> => {
  return rawInput.split("\n")
}

const translation1 = new Map<string, number>()
translation1.set("1", 1)
translation1.set("2", 2)
translation1.set("3", 3)
translation1.set("4", 4)
translation1.set("5", 5)
translation1.set("6", 6)
translation1.set("7", 7)
translation1.set("8", 8)
translation1.set("9", 9)

const translation2 = new Map<string, number>()
translation1.forEach((val: number, key: string) => {
  translation2.set(key, val)
})
translation2.set("one", 1)
translation2.set("two", 2)
translation2.set("three", 3)
translation2.set("four", 4)
translation2.set("five", 5)
translation2.set("six", 6)
translation2.set("seven", 7)
translation2.set("eight", 8)
translation2.set("nine", 9)

const solve = (rawInput: string, trans: Map<string, number>): number => {
  const input = parseInput(rawInput)

  let total = 0
  for (const line of input) {
    const first = findFirstNumber(line, trans)
    if (Number.isNaN(first)) {
      console.error("No first number found", line, first)
    }

    const last = findLastNumber(line, trans)
    if (Number.isNaN(last)) {
      console.error("No last number found", line, last)
    }

    const lineNum = first * 10 + last
    total += lineNum
  }

  return total
}

const findFirstNumber = (line: string, trans: Map<string, number>): number => {
  let idx = Number.MAX_SAFE_INTEGER
  let num = NaN
  trans.forEach((val: number, key: string) => {
    let numIdx = line.indexOf(key)
    if (numIdx >= 0 && numIdx < idx) {
      idx = numIdx
      num = val
    }
  })
  return num
}

const findLastNumber = (line: string, trans: Map<string, number>): number => {
  let idx = Number.MIN_SAFE_INTEGER
  let num = NaN
  trans.forEach((val: number, key: string) => {
    let numIdx = line.lastIndexOf(key)
    if (numIdx >= 0 && numIdx > idx) {
      idx = numIdx
      num = val
    }
  })
  return num
}

const part1 = (rawInput: string): number => {
  return solve(rawInput, translation1)
}

const part2 = (rawInput: string): number => {
  return solve(rawInput, translation2)
}

run({
  part1: {
    tests: [
      {
        input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
        expected: 142,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
        expected: 281,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
