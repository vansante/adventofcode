import run from "aocrunner"

interface Coord {
  x: number
  y: number
}

interface Galaxy {
  coord: Coord
}

interface Grid {
  orig: Array<string>
  exp: Array<string>
  galaxies: Array<Galaxy>
}

const parseInput = (rawInput: string): Grid => {
  const g = {} as Grid

  g.orig = rawInput.split("\n")
  g.exp = rawInput.split("\n")

  // expand rows
  for (let y = 0; y < g.exp.length; y++) {
    const line = g.exp[y]
    if (line.split("").filter((v: string) => v !== ".").length === 0) {
      g.exp.splice(y, 0, line)
      y++
    }
  }

  // expand columns
  for (let x = 0; x < g.exp[0].length; x++) {
    let rowDots = true
    for (let y = 0; y < g.exp.length; y++) {
      if (g.exp[y][x] !== ".") {
        rowDots = false
      }
    }

    if (!rowDots) {
      continue
    }

    for (let y = 0; y < g.exp.length; y++) {
      g.exp[y] = g.exp[y].slice(0, x) + "." + g.exp[y].slice(x)
    }
    x++
  }

  g.galaxies = []
  for (let y = 0; y < g.exp.length; y++) {
    for (let x = 0; x < g.exp[y].length; x++) {
      if (g.exp[y][x] === "#") {
        g.galaxies.push({
          coord: {
            x,
            y,
          },
        })
      }
    }
  }

  return g
}

const manhattanDist = (a: Coord, b: Coord): number => {
  return Math.abs(a.x - b.x) + Math.abs(a.y - b.y)
}

const part1 = (rawInput: string): number => {
  const grid = parseInput(rawInput)

  let total = 0
  for (let i = 0; i < grid.galaxies.length; i++) {
    for (let j = i; j < grid.galaxies.length; j++) {
      total += manhattanDist(grid.galaxies[i].coord, grid.galaxies[j].coord)
    }
  }

  return total
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)

  return
}

run({
  part1: {
    tests: [
      {
        input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
        expected: 374,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
        expected: 1030,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
