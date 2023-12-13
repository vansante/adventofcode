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

const hasVerticalMirror = (g: Grid, xCheck: number): boolean => {
  if (xCheck === g.g[0].length - 1) {
    return false
  }

  const len = Math.min(xCheck + 1, g.g[0].length - xCheck - 1)
  if (len < 1) {
    return false
  }

  for (let y = 0; y < g.g.length; y++) {
    for (let x = 0; x < len; x++) {
      const a = g.g[y][xCheck - x]
      const b = g.g[y][xCheck + 1 + x]

      if (a !== b) {
        return false
      }
    }
  }
  return true
}

const hasHorizontalMirror = (g: Grid, yCheck: number): boolean => {
  if (yCheck === g.g.length - 1) {
    return false
  }

  const len = Math.min(yCheck + 1, g.g.length - yCheck - 1)
  if (len < 1) {
    return false
  }

  for (let y = 0; y < len; y++) {
    for (let x = 0; x < g.g[y].length; x++) {
      const a = g.g[yCheck - y][x]
      const b = g.g[yCheck + 1 + y][x]

      if (a !== b) {
        return false
      }
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
        expected: 400,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
