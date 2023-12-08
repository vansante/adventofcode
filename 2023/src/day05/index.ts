import run from "aocrunner"

interface Input {
  seeds: Array<number>
  maps: Map<string, Mapping>
  revMaps: Map<string, Mapping>
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
    maps: new Map(),
    revMaps: new Map(),
  } as Input

  const lines = rawInput.split("\n")

  let curMap = null as Mapping | null

  for (const line of lines) {
    if (line.indexOf("seeds:") === 0) {
      inp.seeds = line
        .substring(line.indexOf(":") + 2)
        .split(" ")
        .map((val: string): number => parseInt(val, 10))
      continue
    }

    if (line === "") {
      curMap = null
      continue
    }

    if (line.indexOf(":") > 0) {
      const name = line.substring(0, line.indexOf(":"))
      curMap = {
        name: name,
        from: name.substring(0, name.indexOf("-")),
        to: name.substring(name.lastIndexOf("-") + 1, name.lastIndexOf(" ")),
        ranges: [],
      }
      inp.maps.set(curMap.from, curMap)
      inp.revMaps.set(curMap.to, curMap)
      continue
    }

    const range = line
      .split(" ")
      .map((val: string): number => parseInt(val, 10))
    curMap?.ranges.push({
      dstStart: range[0],
      srcStart: range[1],
      len: range[2],
    })
  }

  return inp
}

const translate = (map: Mapping, id: number): number => {
  for (const range of map.ranges) {
    if (id >= range.srcStart && id < range.srcStart + range.len) {
      return range.dstStart + (id - range.srcStart)
    }
  }

  return id
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput)

  const start = "seed"
  const end = "location"

  let item = start
  let ids = input.seeds
  while (item !== end) {
    const map = input.maps.get(item) as Mapping

    const newIds = [] as Array<number>
    for (const id of ids) {
      newIds.push(translate(map, id))
    }

    ids = newIds
    item = map.to
  }

  return ids.reduce((acc: number, val: number): number => {
    return acc < val ? acc : val
  })
}

const seedExists = (input: Input, seed: number): boolean => {
  for (let i = 0; i < input.seeds.length; ) {
    const start = input.seeds[i]
    const len = input.seeds[i + 1]

    if (seed >= start && seed < start + len) {
      return true
    }

    i += 2
  }
  return false
}

const translateReverse = (map: Mapping, id: number): number => {
  for (const range of map.ranges) {
    if (id >= range.dstStart && id < range.dstStart + range.len) {
      return range.srcStart + (id - range.dstStart)
    }
  }

  return id
}

const part2 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  const end = "seed"
  const start = "location"

  const max = 10_000_000
  for (let i = 0; i < max; i++) {
    let item = start
    let id = i
    while (item !== end) {
      const map = input.revMaps.get(item) as Mapping

      id = translateReverse(map, id)
      item = map.from
    }

    if (seedExists(input, id)) {
      return i
    }
  }

  console.error("not found")
  return -999_999_999
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
        expected: 46,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
