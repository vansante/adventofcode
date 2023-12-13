import run from "aocrunner"

interface Grid {
  g: Array<string>
}

const parseInput = (rawInput: string): Array<Grid> => {
  return rawInput.split("\n\n").map((grid: string): Grid => {
    return {
      g: grid.split("\n"),
    }
  })
}

const hasVerticalMirror = (g: Grid, x: number): boolean => {
  if (x === g.g[0].length - 1) {
    return false
  }

  const len = Math.min(x + 1, g.g[0].length - x - 1)
  if (len < 1) {
    return false
  }

  for (const line of g.g) {
    for (let char = 0; char < len; char++) {
      const a = line[x - char]
      const b = line[x + 1 + char]

      // console.log(x, len, char, a, "<>", b)
      if (a !== b) {
        return false
      }
    }
  }
  return true
}

const hasHorizontalMirror = (g: Grid, y: number): boolean => {
  if (y === g.g.length - 1) {
    return false
  }

  const len = Math.min(y + 1, g.g.length - y - 1)
  if (len < 1) {
    return false
  }

  for (let line = 0; line < len; line++) {
    const a = g.g[y - line]
    const b = g.g[y + 1 + line]

    // console.log(y, len, line, a, "<>", b)
    if (a !== b) {
      return false
    }
  }

  return true
}

const findMirrors = (g: Grid): [number, number] => {
  const horizontal = []
  for (let y = 0; y < g.g.length; y++) {
    if (hasHorizontalMirror(g, y)) {
      horizontal.push(y)
    }
  }
  if (horizontal.length > 1) {
    console.log("more than one horizontal detected", horizontal)
  }

  const vertical = []
  for (let x = 0; x < g.g[0].length; x++) {
    if (hasVerticalMirror(g, x)) {
      vertical.push(x)
    }
  }
  if (vertical.length > 1) {
    console.log("more than one vertical detected", vertical)
  }

  return [horizontal[0], vertical[0]]
}

const part1 = (rawInput: string) => {
  const grids = parseInput(rawInput)

  let total = 0
  for (const grid of grids) {
    const results = findMirrors(grid)
    if (results[0] !== undefined) {
      total += (results[0] + 1) * 100
    }
    if (results[1] !== undefined) {
      total += results[1] + 1
    }
  }

  // 43210 too low
  // === 43614
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
        input: `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
        expected: 405,
      },
      {
        input: `#....#.#.#....#.#
#....###.#....#.#
.##.#....#.##.#..
.###.###.#....#.#
.##.....##....##.
###.#.#..#.##.#..
##...#.#.######.#
##..##.####..####
#.#..###..#..#..#
.#####..#.#..#.#.
.#.##.#..#.##.#..`,
        expected: 12,
      },
      {
        input: `##..####..####.
##...##...####.
#.##.##.##.##.#
###.####.######
.####..####..##
.#.#...##.#..#.
#..........##..
....#..#.......
.##.#..#.##..##
#####..########
...#....#......`,
        expected: 12,
      },
      {
        input: `#...####.##.###
.#####.#...####
#.#.#..#.#..##.
#.#.#..#.#..##.
.#####.#...####
....####.##.###
###..###.###.#.
..##...#.#..#.#
#.#.#...###.###
.###..#...##..#
######.......#.
#...#####..##.#
#...#####..##.#
######.......#.
.###..#...##..#`,
        expected: 1200,
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
  onlyTests: false,
})
