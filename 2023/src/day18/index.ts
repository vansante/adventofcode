import run from "aocrunner"

class Point {
  x: number
  y: number

  constructor(x: number, y: number) {
    this.x = x
    this.y = y
  }

  toString(): string {
    return `${this.x}_${this.y}`
  }

  equals(c: Point): boolean {
    return c.x === this.x && c.y === this.y
  }

  add(v: Point): Point {
    return new Point(this.x + v.x, this.y + v.y)
  }

  direction(c: Point): number {
    const v = new Point(c.x - this.x, c.y - this.y)
    for (let i = 0; i < directions.length; i++) {
      if (v.equals(directions[i])) {
        return i
      }
    }

    console.error(`${this} => ${c}: vector ${v} not found`)
    throw "vector not found"
  }
}

const directions = [
  new Point(0, -1),
  new Point(1, 0),
  new Point(0, 1),
  new Point(-1, 0),
]

interface Instruction {
  direction: string
  steps: number
  color: string
}

interface Input {
  instr: Array<Instruction>
  grid: Set<string>
}

const parseInput = (rawInput: string): Input => {
  const g = {} as Input
}

const part1 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  return 0
}

const part2 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  return 0
}

run({
  part1: {
    tests: [
      {
        input: `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`,
        expected: 0,
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
