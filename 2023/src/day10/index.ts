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
  x: number
  y: number
  connectors: Map<string, boolean>
  cardTiles: Map<string, Tile>
  pipeTiles: Map<string, Tile>
  distance: number
  isStart: boolean
  isReachable: boolean
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
        x,
        y,
        connectors: new Map<string, boolean>(),
        cardTiles: new Map<string, Tile>(),
        pipeTiles: new Map<string, Tile>(),
        distance: Number.MAX_SAFE_INTEGER,
        isStart: false,
        isReachable: false,
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
      if (x > 0) {
        const west = line[x - 1]
        tile.cardTiles.set("w", west)
        west.cardTiles.set("e", tile)
        if (tile.connectors.get("w") && west.connectors.get("e")) {
          tile.pipeTiles.set("w", west)
          west.pipeTiles.set("e", tile)
        }
      }
      if (y > 0) {
        const north = input.tiles[y - 1][x]
        tile.cardTiles.set("n", north)
        north.cardTiles.set("s", tile)
        if (tile.connectors.get("n") && north.connectors.get("s")) {
          tile.pipeTiles.set("n", north)
          north.pipeTiles.set("s", tile)
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

    const tile = t.pipeTiles.get(card)
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

const fixStartConnectors = (tile: Tile) => {
  for (const card of cardinals) {
    if (!tile.pipeTiles.get(card)) {
      tile.connectors.delete(card)
    }
  }
}

const contained = (tile: Tile): boolean => {
  if (tile.distance !== Number.MAX_SAFE_INTEGER) {
    return false
  }

  let pipeCount = 0
  let t = tile.cardTiles.get("e")

  while (t) {
    if (t.distance !== Number.MAX_SAFE_INTEGER && t?.connectors.get("n")) {
      pipeCount++
    }
    t = t.cardTiles.get("e")
  }

  return pipeCount % 2 === 1
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)

  walkCircle(input.start as Tile, "")
  fixStartConnectors(input.start as Tile)

  let count = 0
  for (const line of input.tiles) {
    for (const tile of line) {
      if (contained(tile)) {
        count++
      }
    }
  }

  return count
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
      {
        input: `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`,
        expected: 4,
      },
      {
        input: `..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`,
        expected: 4,
      },
      {
        input: `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`,
        expected: 8,
      },
      {
        input: `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`,
        expected: 10,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
