import run from "aocrunner"

interface Input {
  seeds: Array<number>
  maps: Array<Mapping>
}

interface Mapping {
  name: string
  from: string
  to: string
  ranges: Array<Range>
}

interface Range {
  dstStart: number
  srcStart: number
  len: number
}

const parseInput = (rawInput: string): Input => {
  const inp = {
    seeds: [],
    maps: [],
  } as Input

  const lines = rawInput.split("\n")

  let curMap = null as Mapping|null

  for (const line of lines) {
    if (line.indexOf('seeds:') === 0) {
      inp.seeds = line.substring(line.indexOf(':') + 2).split(' ').map((val: string): number => parseInt(val, 10))
      continue
    }

    if (line === '') {
      curMap = null
      continue
    }

    if (line.indexOf(':') > 0) {
      const name = line.substring(0, line.indexOf(':'))
      curMap = {
        name: name,
        from: name.substring(0, name.indexOf('-')),
        to: name.substring(name.lastIndexOf('-') + 1, name.lastIndexOf(' ')),
        ranges: [],
      }
      inp.maps.push(curMap)
      continue
    }

    const range = line.split(' ').map((val: string): number => parseInt(val, 10))
    curMap?.ranges.push({
      dstStart: range[0],
      srcStart: range[1],
      len: range[2],
    })
  }

  return inp
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput)
  console.log(input.maps[1])

  return
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)

  return
}

run({
  part1: {
    tests: [
      {
        input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
        expected: 35,
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
