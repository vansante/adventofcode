import run from "aocrunner"

interface Coord {
  x: number
  y: number
}

const cardinals = ["n", "e", "s", "w"]

const directions = new Map<string, Coord>()
directions.set(cardinals[0], { x: 0, y: -1 })
directions.set(cardinals[1], { x: 1, y: 0 })
directions.set(cardinals[2], { x: 0, y: 1 })
directions.set(cardinals[3], { x: -1, y: 0 })

const opposite = (card: string): string => {
  const idx = (cardinals.indexOf(card) + 2) % cardinals.length
  return cardinals[idx]
}

interface Tile {
  connectors: Map<string, boolean>
  tiles: Map<string, Tile>
  distance: number
  isStart: boolean
}

interface Input {
  startX: number
  startY: number
  start: Tile | null
  tiles: Array<Array<Tile>>
}

const parseInput = (rawInput: string): Input => {
  const lines = rawInput.split("\n")

  const input = {
    startX: 0,
    startY: 0,
    start: null,
    tiles: [],
  } as Input

  for (let y = 0; y < lines.length; y++) {
    const line = lines[y]

    input.tiles.push([])
    for (let x = 0; x < line.length; x++) {
      const char = line[x]
      const t = {
        connectors: new Map<string, boolean>(),
        tiles: new Map<string, Tile>(),
        distance: Number.MAX_SAFE_INTEGER,
        isStart: false,
      } as Tile

      switch (char) {
        case "S":
          t.isStart = true
          t.distance = 0
          for (const card of cardinals) {
            t.connectors.set(card, true)
          }
          input.startY = y
          input.startX = x
          input.start = t
        case "|":
          t.connectors.set("n", true)
          t.connectors.set("s", true)
          break
        case "-":
          t.connectors.set("w", true)
          t.connectors.set("e", true)
          break
        case "L":
          t.connectors.set("n", true)
          t.connectors.set("e", true)
          break
        case "J":
          t.connectors.set("n", true)
          t.connectors.set("w", true)
          break
        case "7":
          t.connectors.set("s", true)
          t.connectors.set("w", true)
          break
        case "F":
          t.connectors.set("s", true)
          t.connectors.set("e", true)
          break
      }
      input.tiles[y][x] = t
    }
  }

  for (let y = 0; y < input.tiles.length; y++) {
    const line = input.tiles[y]

    for (let x = 0; x < line.length; x++) {
      const tile = line[x]
      if (x > 0 && tile.connectors.get("w")) {
        const west = line[x - 1]
        if (west.connectors.get("e")) {
          tile.tiles.set("w", west)
          west.tiles.set("e", tile)
        }
      }
      if (y > 0 && tile.connectors.get("n")) {
        const north = input.tiles[y - 1][x]
        if (north.connectors.get("s")) {
          tile.tiles.set("n", north)
          north.tiles.set("s", tile)
        }
      }
    }
  }

  return input
}

const connectingTile = (t: Tile, from: string): [string, Tile] => {
  for (const card of cardinals) {
    if (card === from) {
      continue
    }

    const tile = t.tiles.get(card)
    if (tile) {
      return [card, tile]
    }
  }

  console.error("no connecting tile found", t)
  throw "no connecting tile found"
}

const walkCircle = (t: Tile, card: string): string => {
  let dist = 0
  let tile = t
  let firstCard = ""
  card = opposite(card)

  while (!tile.isStart || dist === 0) {
    ;[card, tile] = connectingTile(tile, opposite(card))
    // console.log("step", card, tile)
    if (firstCard === "") {
      firstCard = card
    }

    dist++
    if (tile.distance > dist) {
      tile.distance = dist
    }
  }

  return firstCard
}

const part1 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  const card = walkCircle(input.start as Tile, "")
  walkCircle(input.start as Tile, card)

  // console.log(input.tiles)

  let dist = 0
  for (const line of input.tiles) {
    for (const tile of line) {
      if (tile.distance !== Number.MAX_SAFE_INTEGER && dist < tile.distance) {
        dist = tile.distance
      }
    }
  }

  return dist
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
  onlyTests: false,
})
