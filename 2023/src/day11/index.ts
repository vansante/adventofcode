import run from "aocrunner"

interface Coord {
  x: number
  y: number
  expansion: number
  galaxy: boolean
}

interface Galaxy {
  coord: Coord
}

interface Grid {
  orig: Array<string>
  exp: Array<Array<Coord>>
  galaxies: Array<Galaxy>
}

const parseInput = (rawInput: string, expFactor: number): Grid => {
  const g = {} as Grid

  g.orig = rawInput.split("\n")
  g.exp = []

  // expand rows
  for (let y = 0; y < g.orig.length; y++) {
    const line = g.orig[y].split("")
    let exp = 1
    if (line.filter((v: string) => v !== ".").length === 0) {
      exp = expFactor
    }

    g.exp.push(
      line.map((val: string, x: number): Coord => {
        return {
          y,
          x,
          galaxy: val === "#",
          expansion: exp,
        }
      }),
    )
  }

  // expand columns
  for (let x = 0; x < g.orig[0].length; x++) {
    let expanded = true
    for (let y = 0; y < g.orig.length; y++) {
      if (g.orig[y][x] !== ".") {
        expanded = false
      }
    }

    if (!expanded) {
      continue
    }

    for (let y = 0; y < g.exp.length; y++) {
      g.exp[y][x].expansion *= expFactor
    }
  }

  g.galaxies = []
  for (let y = 0; y < g.exp.length; y++) {
    for (let x = 0; x < g.exp[y].length; x++) {
      if (g.exp[y][x].galaxy) {
        g.galaxies.push({
          coord: g.exp[y][x],
        })
      }
    }
  }

  return g
}

const manhattanDist = (a: Coord, b: Coord): number => {
  return Math.abs(a.x - b.x) + Math.abs(a.y - b.y)
}

const distance = (g: Grid, a: Coord, b: Coord): number => {
  let distance = 0

  if (a.x <= b.x) {
    for (let x = a.x; x < b.x; x++) {
      distance += g.exp[a.y][x].expansion
    }
  } else {
    for (let x = b.x; x < a.x; x++) {
      distance += g.exp[a.y][x].expansion
    }
  }

  if (a.y <= b.y) {
    for (let y = a.y; y < b.y; y++) {
      distance += g.exp[y][b.x].expansion
    }
  } else {
    for (let y = b.y; y < a.y; y++) {
      distance += g.exp[y][b.x].expansion
    }
  }

  return distance
}

const distances = (grid: Grid): number => {
  let total = 0
  for (let i = 0; i < grid.galaxies.length; i++) {
    for (let j = i; j < grid.galaxies.length; j++) {
      total += distance(grid, grid.galaxies[i].coord, grid.galaxies[j].coord)
    }
  }

  return total
}

const part1 = (rawInput: string): number => {
  const grid = parseInput(rawInput, 2)

  return distances(grid)
}

const part2 = (rawInput: string) => {
  const grid = parseInput(rawInput, 1_000_000)

  return distances(grid)
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
        expected: 8410,
      },
    ],
    solution: (rawInput: string): number => {
      return distances(parseInput(rawInput, 100))
    },
  },
  trimTestInputs: true,
  onlyTests: false,
})
