import run from "aocrunner"

interface Grid {
  g: Array<Array<string>>
}

const parseInput = (rawInput: string): Grid => {
  return {
    g: rawInput.split("\n").map((v: string): Array<string> => v.split("")),
  }
}

const rollRockNorth = (g: Grid, x: number, y: number) => {
  for (y--; y >= 0; y--) {
    if (g.g[y][x] === "#" || g.g[y][x] === "O") {
      break
    }

    g.g[y + 1][x] = "."
    g.g[y][x] = "O"
  }
}

const rollRockSouth = (g: Grid, x: number, y: number) => {
  for (y++; y < g.g.length; y++) {
    if (g.g[y][x] === "#" || g.g[y][x] === "O") {
      break
    }

    g.g[y - 1][x] = "."
    g.g[y][x] = "O"
  }
}

const rollRockWest = (g: Grid, x: number, y: number) => {
  for (x--; x >= 0; x--) {
    if (g.g[y][x] === "#" || g.g[y][x] === "O") {
      break
    }

    g.g[y][x + 1] = "."
    g.g[y][x] = "O"
  }
}

const rollRockEast = (g: Grid, x: number, y: number) => {
  for (x++; x < g.g[y].length; x++) {
    if (g.g[y][x] === "#" || g.g[y][x] === "O") {
      break
    }

    g.g[y][x - 1] = "."
    g.g[y][x] = "O"
  }
}

const rollAllNorth = (g: Grid) => {
  for (let y = 0; y < g.g.length; y++) {
    for (let x = 0; x < g.g[y].length; x++) {
      if (g.g[y][x] === "O") {
        rollRockNorth(g, x, y)
      }
    }
  }
}

const rollAllSouth = (g: Grid) => {
  for (let y = g.g.length - 1; y >= 0; y--) {
    for (let x = 0; x < g.g[y].length; x++) {
      if (g.g[y][x] === "O") {
        rollRockSouth(g, x, y)
      }
    }
  }
}

const rollAllEast = (g: Grid) => {
  for (let x = g.g[0].length - 1; x >= 0; x--) {
    for (let y = 0; y < g.g.length; y++) {
      if (g.g[y][x] === "O") {
        rollRockEast(g, x, y)
      }
    }
  }
}

const rollAllWest = (g: Grid) => {
  for (let x = 0; x < g.g[0].length; x++) {
    for (let y = 0; y < g.g.length; y++) {
      if (g.g[y][x] === "O") {
        rollRockWest(g, x, y)
      }
    }
  }
}

const totalLoad = (g: Grid): number => {
  let total = 0
  for (let y = 0; y < g.g.length; y++) {
    const load = g.g.length - y

    for (let x = 0; x < g.g[y].length; x++) {
      if (g.g[y][x] === "O") {
        total += load
      }
    }
  }
  return total
}

const gridString = (g: Grid): string => {
  let grid = ""
  for (let y = 0; y < g.g.length; y++) {
    let row = ""
    for (let x = 0; x < g.g[0].length; x++) {
      row += g.g[y][x]
    }
    grid += row + "\n"
  }
  return grid
}

const part1 = (rawInput: string): number => {
  const grid = parseInput(rawInput)

  rollAllNorth(grid)

  return totalLoad(grid)
}

const cycle = (g: Grid) => {
  rollAllNorth(g)
  rollAllWest(g)
  rollAllSouth(g)
  rollAllEast(g)
}

const cycles = (
  g: Grid,
  count: number,
  state: Set<string> | null = null,
): number => {
  for (let i = 0; i < count; i++) {
    cycle(g)

    if (state) {
      const str = gridString(g)
      if (state.has(str)) {
        return i + 1
      }
      state.add(str)
    }
  }
  return count
}

const part2 = (rawInput: string): number => {
  let grid = parseInput(rawInput)

  const count = 1_000_000_000

  let rep = [] as Array<number>
  let diff1 = -1
  let i = 0
  for (; i < count; i++) {
    i += cycles(grid, count, new Set<string>())
    rep.push(i)

    const len = rep.length
    if (len >= 3) {
      diff1 = rep[len - 1] - rep[len - 2]
      const diff2 = rep[len - 2] - rep[len - 3]
      if (diff1 === diff2) {
        console.log(`repeating repetition of ${diff1} detected at ${i}`)
        break
      }
    }
  }

  grid = parseInput(rawInput)
  const rounds = (count - i - 1) % (diff1 + i + 1)
  console.log(rounds)

  // let remainder = (count - i) % diff1
  // console.log(`${remainder} steps remaining`)

  cycles(grid, rounds)

  // > 98020
  // < 98876
  return totalLoad(grid)
}

run({
  part1: {
    tests: [
      {
        input: `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`,
        expected: 136,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`,
        expected: 64,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
