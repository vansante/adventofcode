import run from "aocrunner"

interface Coord {
  x: number
  y: number
}

const North = 0
const East = 1
const South = 2
const West = 3

const directions = [
  { x: 0, y: -1 },
  { x: 1, y: 0 },
  { x: 0, y: 1 },
  { x: -1, y: 0 },
] as Array<Coord>

interface Tile {
  x: number
  y: number
  type: string
  beams: number
  visited: Set<number>
}

interface Grid {
  tiles: Array<Array<Tile>>
}

const parseInput = (rawInput: string): Grid => {
  const g = {
    tiles: [] as Array<Array<Tile>>,
  }

  rawInput.split("\n").forEach((v: string, y: number) => {
    g.tiles.push(
      v.split("").map((v: string, x: number): Tile => {
        return {
          x,
          y,
          type: v,
          beams: 0,
          visited: new Set<number>(),
        } as Tile
      }),
    )
  })

  return g
}

const addVector = (c: Coord, v: Coord): Coord => {
  return {
    x: c.x + v.x,
    y: c.y + v.y,
  }
}

const inBounds = (c: Coord, g: Grid): boolean => {
  if (c.x < 0 || c.y < 0) {
    return false
  }
  if (c.x >= g.tiles[0].length || c.y >= g.tiles.length) {
    return false
  }
  return true
}

const resetGrid = (g: Grid) => {
  g.tiles.forEach((row: Array<Tile>) => {
    row.forEach((t: Tile) => {
      t.beams = 0
      t.visited = new Set<number>()
    })
  })
}

const trackBeam = (g: Grid, c: Coord, dir: number) => {
  while (inBounds(c, g)) {
    const t = g.tiles[c.y][c.x]
    if (t.visited.has(dir)) {
      return
    }
    t.visited.add(dir)

    t.beams++

    switch (t.type) {
      case ".":
        break
      case "|":
        if (dir === West || dir === East) {
          trackBeam(g, addVector(c, directions[North]), North)
          trackBeam(g, addVector(c, directions[South]), South)
          return // let the 2 subroutines take over
        }
        break
      case "-":
        if (dir === North || dir === South) {
          trackBeam(g, addVector(c, directions[East]), East)
          trackBeam(g, addVector(c, directions[West]), West)
          return // let the 2 subroutines take over
        }
        break
      case "/":
        switch (dir) {
          case North:
            dir = East
            break
          case East:
            dir = North
            break
          case South:
            dir = West
            break
          case West:
            dir = South
            break
        }
        break
      case "\\":
        switch (dir) {
          case North:
            dir = West
            break
          case East:
            dir = South
            break
          case South:
            dir = East
            break
          case West:
            dir = North
            break
        }
        break
      default:
        throw `unknown tile type ${t.type}`
    }

    c = addVector(c, directions[dir])
  }
}

const countEnergized = (g: Grid): number => {
  let total = 0
  g.tiles.forEach((row: Array<Tile>) => {
    row.forEach((t: Tile) => {
      if (t.beams > 0) {
        total++
      }
    })
  })

  return total
}

const part1 = (rawInput: string): number => {
  const grid = parseInput(rawInput)

  trackBeam(grid, { x: 0, y: 0 }, East)

  return countEnergized(grid)
}

const part2 = (rawInput: string): number => {
  const grid = parseInput(rawInput)

  let mostEnergized = 0
  for (let y = 0; y < grid.tiles.length; y++) {
    trackBeam(grid, { x: 0, y: y }, East)
    const energizedWest = countEnergized(grid)
    resetGrid(grid)

    if (energizedWest > mostEnergized) {
      mostEnergized = energizedWest
    }

    trackBeam(grid, { x: grid.tiles[0].length - 1, y: y }, West)
    const energizedEast = countEnergized(grid)
    resetGrid(grid)

    if (energizedEast > mostEnergized) {
      mostEnergized = energizedEast
    }
  }

  for (let x = 0; x < grid.tiles[0].length; x++) {
    trackBeam(grid, { x: x, y: 0 }, South)
    const energizedNorth = countEnergized(grid)
    resetGrid(grid)

    if (energizedNorth > mostEnergized) {
      mostEnergized = energizedNorth
    }

    trackBeam(grid, { x: x, y: grid.tiles.length - 1 }, South)
    const energizedSouth = countEnergized(grid)
    resetGrid(grid)

    if (energizedSouth > mostEnergized) {
      mostEnergized = energizedSouth
    }
  }

  return mostEnergized
}

run({
  part1: {
    tests: [
      {
        input: `.|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|....`,
        expected: 46,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `.|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|....`,
        expected: 51,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
