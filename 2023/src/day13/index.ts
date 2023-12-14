import run from "aocrunner"

interface Grid {
  g: Array<Array<string>>
}

const parseInput = (rawInput: string): Array<Grid> => {
  return rawInput.split("\n\n").map((grid: string): Grid => {
    return {
      g: grid.split("\n").map((v: string) => {
        return v.split("")
      }),
    }
  })
}

const hasVerticalMirror = (
  g: Grid,
  xCheck: number,
): [boolean, undefined | [number, number]] => {
  if (xCheck === g.g[0].length - 1) {
    return [false, undefined]
  }

  const len = Math.min(xCheck + 1, g.g[0].length - xCheck - 1)
  if (len < 1) {
    return [false, undefined]
  }

  let mistakes = 0
  let coord = [0, 0] as [number, number]
  for (let y = 0; y < g.g.length; y++) {
    for (let x = 0; x < len; x++) {
      const a = g.g[y][xCheck - x]
      const b = g.g[y][xCheck + 1 + x]

      if (a !== b) {
        mistakes++
        coord = [x, y]
        if (mistakes > 1) {
          return [false, undefined]
        }
      }
    }
  }

  return [mistakes === 0, coord]
}

const hasHorizontalMirror = (
  g: Grid,
  yCheck: number,
): [boolean, undefined | [number, number]] => {
  if (yCheck === g.g.length - 1) {
    return [false, undefined]
  }

  const len = Math.min(yCheck + 1, g.g.length - yCheck - 1)
  if (len < 1) {
    return [false, undefined]
  }

  let mistakes = 0
  let coord = [0, 0] as [number, number]
  for (let y = 0; y < len; y++) {
    for (let x = 0; x < g.g[y].length; x++) {
      const a = g.g[yCheck - y][x]
      const b = g.g[yCheck + 1 + y][x]

      if (a !== b) {
        mistakes++
        coord = [x, y]
        if (mistakes > 1) {
          return [false, undefined]
        }
      }
    }
  }

  return [mistakes === 0, coord]
}

const findMirrors = (
  g: Grid,
  mistakes: number = 0,
): [Array<number>, Array<number>] => {
  const horizontal = []
  for (let y = 0; y < g.g.length; y++) {
    const [result, coord] = hasHorizontalMirror(g, y)
    if (result && mistakes === 0) {
      horizontal.push(y)
    } else if (coord && mistakes > 0) {
      g.g[coord[1]][coord[0]] = g.g[coord[1]][coord[0]] === "." ? "#" : "."

      horizontal.push(y)
      mistakes--
    }
  }
  if (horizontal.length > 1) {
    console.log("more than one horizontal detected", horizontal)
  }

  const vertical = []
  for (let x = 0; x < g.g[0].length; x++) {
    const [result, coord] = hasVerticalMirror(g, x)
    if (result && mistakes === 0) {
      vertical.push(x)
    } else if (coord && mistakes > 0) {
      g.g[coord[1]][coord[0]] = g.g[coord[1]][coord[0]] === "." ? "#" : "."

      vertical.push(x)
      mistakes--
    }
  }
  if (vertical.length > 1) {
    console.log("more than one vertical detected", vertical)
  }

  return [horizontal, vertical]
}

const part1 = (rawInput: string) => {
  const grids = parseInput(rawInput)

  let total = 0
  for (const grid of grids) {
    const results = findMirrors(grid, 0)
    // console.log(results)
    results[0].forEach((v: number) => {
      total += (v + 1) * 100
    })
    results[1].forEach((v: number) => {
      total += v + 1
    })
  }

  return total
}

const part2 = (rawInput: string) => {
  const grids = parseInput(rawInput)

  let total = 0
  for (const grid of grids) {
    const results = findMirrors(grid, 1)
    console.log(results)
    results[0].forEach((v: number) => {
      total += (v + 1) * 100
    })
    results[1].forEach((v: number) => {
      total += v + 1
    })
  }

  return total
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
