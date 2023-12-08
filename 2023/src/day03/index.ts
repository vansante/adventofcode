import run from "aocrunner"

const parseInput = (rawInput: string): Array<string> => {
  return rawInput.split("\n")
}

const isDot = (str: string): boolean => {
  return str === "."
}

const isNumber = (str: string): boolean => {
  const code = str.charCodeAt(0)
  return code >= 48 && code <= 57
}

const isSymbol = (str: string): boolean => {
  return str.length > 0 && !isNumber(str) && !isDot(str)
}

const isGear = (str: string): boolean => {
  return str.length > 0 && str === "*"
}

interface Coord {
  x: number
  y: number
}

const surroundSymbols = (
  input: Array<string>,
  line: number,
  start: number,
  len: number,
  testFn: Function = isSymbol,
): Coord | null => {
  if (line > 0) {
    for (let i = start - 1; i < start + len + 1; i++) {
      if (
        i >= 0 &&
        i < input[line - 1].length &&
        testFn(input[line - 1].charAt(i))
      ) {
        return { x: line - 1, y: i }
      }
    }
  }

  if (start > 0 && testFn(input[line].charAt(start - 1))) {
    return { x: line, y: start - 1 }
  }

  if (
    start + len < input[line].length &&
    testFn(input[line].charAt(start + len))
  ) {
    return { x: line, y: start + len }
  }

  if (line < input.length - 1) {
    for (let i = start - 1; i < start + len + 1; i++) {
      if (
        i >= 0 &&
        i < input[line + 1].length &&
        testFn(input[line + 1].charAt(i))
      ) {
        return { x: line + 1, y: i }
      }
    }
  }

  return null
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput)

  let total = 0
  for (let i = 0; i < input.length; i++) {
    const line = input[i]

    let numStart = -1
    let numLen = 0
    for (let j = 0; j < line.length; j++) {
      const char = line.charAt(j)

      if (isNumber(char)) {
        if (numStart < 0) {
          numStart = j
          numLen = 0
        }
        numLen++
      } else if (numStart >= 0) {
        if (surroundSymbols(input, i, numStart, numLen)) {
          const numStr = line.substring(numStart, numStart + numLen)
          total += parseInt(numStr, 10)
        }

        numStart = -1
        numLen = 0
      }
    }

    if (numStart >= 0 && surroundSymbols(input, i, numStart, numLen)) {
      const numStr = line.substring(numStart, numStart + numLen)
      total += parseInt(numStr, 10)
    }

    numStart = -1
    numLen = 0
  }
  return total
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)

  const gears = new Map<string, Array<number>>()
  for (let i = 0; i < input.length; i++) {
    const line = input[i]

    let numStart = -1
    let numLen = 0
    for (let j = 0; j < line.length; j++) {
      const char = line.charAt(j)

      if (isNumber(char)) {
        if (numStart < 0) {
          numStart = j
          numLen = 0
        }
        numLen++
      } else if (numStart >= 0) {
        const coord = surroundSymbols(input, i, numStart, numLen, isGear)
        if (coord) {
          const numStr = line.substring(numStart, numStart + numLen)
          const num = parseInt(numStr, 10)
          if (gears.has(`${coord.x}_${coord.y}`)) {
            gears.get(`${coord.x}_${coord.y}`)?.push(num)
          } else {
            gears.set(`${coord.x}_${coord.y}`, [num])
          }
        }

        numStart = -1
        numLen = 0
      }
    }

    if (numStart >= 0) {
      const coord = surroundSymbols(input, i, numStart, numLen, isGear)
      if (coord) {
        const numStr = line.substring(numStart, numStart + numLen)
        const num = parseInt(numStr, 10)
        if (gears.has(`${coord.x}_${coord.y}`)) {
          gears.get(`${coord.x}_${coord.y}`)?.push(num)
        } else {
          gears.set(`${coord.x}_${coord.y}`, [num])
        }
      }
    }

    numStart = -1
    numLen = 0
  }

  let total = 0
  gears.forEach((val: Array<number>, key: string) => {
    if (val.length === 2) {
      total += val[0] * val[1]
    }
  })

  return total
}

run({
  part1: {
    tests: [
      {
        input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
        expected: 4361,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
        expected: 467835,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
