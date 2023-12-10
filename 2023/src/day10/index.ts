import run from "aocrunner"

interface Coord {
  x: number
  y: number
}

const cardinals = ['n', 'e', 's', 'w']

const directions = new Map<string, Coord>()
directions.set(cardinals[0], {x: 0, y: -1})
directions.set(cardinals[1], {x: 1, y: 0})
directions.set(cardinals[2], {x: 0, y: 1})
directions.set(cardinals[3], {x: -1, y: 0})

interface Tile {
  connectors: Map<string, boolean>
  tiles: Map<string, Tile>

  isStart: boolean
}

interface Input {
  startX: number
  startY: number
  tiles: Array<Array<Tile>>
}

const parseInput = (rawInput: string): Input => {
  const lines = rawInput.split("\n")

  const input = {
    startX: 0,
    startY: 0,
    tiles: [],
  } as Input

  for (let y = 0; y < lines.length; y++) {
    const line = lines[y]

    input.tiles.push([])
    for (let x = 0; x < line.length; x++) {
      const char = line[x]
      const t = {
        connectors: new Map<string, boolean>(),
        isStart: false,
      } as Tile

      switch (char) {
        case "S":
          t.isStart = true
          for (const card of cardinals) {
            t.connectors.set(card, true)
          }
          input.startY = y
          input.startX = x
        case "|":
          t.connectors.set('n', true)
          t.connectors.set('s', true)
          break
        case "-":
          t.connectors.set('w', true)
          t.connectors.set('e', true)
          break
        case "L":
          t.connectors.set('n', true)
          t.connectors.set('e', true)
          break
        case "J":
          t.connectors.set('n', true)
          t.connectors.set('w', true)
          break
        case "7":
          t.connectors.set('s', true)
          t.connectors.set('w', true)
          break
        case "F":
          t.connectors.set('s', true)
          t.connectors.set('e', true)
          break
      }
      input.tiles[y][x] = t
    }
  }

  for (let y = 0; y < input.tiles.length; y++) {
    const line = input.tiles[y]

    for (let x = 0; x < line.length; x++) {
      if (x > 0) {
        line[x].tiles.set('w', line[x-1])
        line[x-1].tiles.set('e', line[x])
      }
      if (y > 0) {
        input.tiles[y][x].tiles.set('n', input.tiles[y-1][x])
        input.tiles[y-1][x].tiles.set('s', input.tiles[y][x])
      }
    }
  }

  return input
}

const connectingTile = (t: Tile, cardinal: string): Record<boolean, Tile> {
  
}

const part1 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  console.log(input.tiles)

  let total = 0

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
        input: `.....
.S-7.
.|.|.
.L-J.
.....`,
        expected: 4,
      },
      {
        input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`,
        expected: 8,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      // {
      //   input: ``,
      //   expected: "",
      // },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: true,
})
