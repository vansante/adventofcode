import run from "aocrunner"

interface Tile {
  c: Coord
  loss: number
}

class Coord {
  x: number
  y: number

  constructor(x: number, y: number) {
    this.x = x
    this.y = y
  }

  toString(): string {
    return `${this.x}_${this.y}`
  }

  equals(c: Coord): boolean {
    return c.x === this.x && c.y === this.y
  }

  add(v: Coord): Coord {
    return new Coord(this.x + v.x, this.y + v.y)
  }

  direction(c: Coord): number {
    const v = new Coord(c.x - this.x, c.y - this.y)
    for (let i = 0; i < directions.length; i++) {
      if (v.equals(directions[i])) {
        return i
      }
    }

    console.error(`${this} => ${c}: vector ${v} not found`)
    throw "vector not found"
  }
}

const fromString = (str: string): Coord => {
  const coords = str.split("_")

  return new Coord(parseInt(coords[0], 10), parseInt(coords[1], 10))
}

const directions = [
  new Coord(0, -1),
  new Coord(1, 0),
  new Coord(0, 1),
  new Coord(-1, 0),
]

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
          c: new Coord(x, y),
          loss: parseInt(v, 10),
        } as Tile
      }),
    )
  })

  return g
}

const getLoss = (
  g: Grid,
  x: number,
  y: number,
  defaultVal: number = Number.MAX_SAFE_INTEGER,
): number => {
  if (y < 0 || y >= g.tiles.length) {
    return defaultVal
  }
  if (x < 0 || x >= g.tiles[0].length) {
    return defaultVal
  }

  return g.tiles[y][x].loss
}

const forbiddenDirection = (a: Coord, b: Coord, c: Coord): number => {
  const dir1 = a.direction(b)
  const dir2 = b.direction(c)

  if (dir1 === dir2) {
    return dir1
  }
  return -1
}

const dijkstra = (g: Grid, start: Coord, end: Coord): number => {
  const visited = new Set<string>()
  const distance = new Map<string, number>()
  const previous = new Map<string, Coord>()
  const queue = [] as Array<Coord>

  distance.set(start.toString(), 0)
  queue.push(start)

  while (queue.length) {
    const current = queue.shift() as Coord
    if (current.equals(end)) {
      break
    }
    visited.add(current.toString())

    let forbiddenDir = -1
    const prevCoord = previous.get(current.toString())
    if (prevCoord) {
      const prevPrevCoord = previous.get(prevCoord.toString())

      if (prevPrevCoord) {
        forbiddenDir = forbiddenDirection(prevPrevCoord, prevCoord, current)
        console.log(
          "FORBIDDEN",
          forbiddenDir,
          prevPrevCoord,
          prevCoord,
          current,
        )
      }
    }

    for (let i = 0; i < directions.length; i++) {
      const dir = directions[i]
      if (i === forbiddenDir) {
        continue
      }

      const neighbor = current.add(dir)
      if (neighbor.equals(current)) {
        continue
      }
      if (visited.has(neighbor.toString())) {
        continue
      }

      const loss = getLoss(g, neighbor.x, neighbor.y)
      if (loss >= Number.MAX_SAFE_INTEGER) {
        continue
      }

      const currentDist = distance.has(neighbor.toString())
        ? (distance.get(neighbor.toString()) as number)
        : Number.MAX_SAFE_INTEGER
      const shortest = (distance.get(current.toString()) as number) + loss

      if (shortest >= currentDist) {
        continue
      }

      queue.push(neighbor)
      distance.set(neighbor.toString(), shortest)
      previous.set(neighbor.toString(), current)
    }
  }

  return distance.get(end.toString()) || Number.MAX_SAFE_INTEGER
}

const part1 = (rawInput: string): number => {
  const grid = parseInput(rawInput)

  return dijkstra(
    grid,
    new Coord(0, 0),
    new Coord(grid.tiles[0].length - 1, grid.tiles.length - 1),
  )
}

const part2 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  return 0
}

run({
  part1: {
    tests: [
      {
        input: `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`,
        expected: 102,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      // {
      //   input: ``,
      //   expected: 0,
      // },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: true,
})
