import run from "aocrunner"

interface Race {
  time: number
  distance: number
}

const parseInput = (rawInput: string, two: boolean = false): Array<Race> => {
  let lines = [] as Array<string>

  if (two) {
    lines = rawInput.replaceAll("\n", "|").replace(/\s+/g, "").trim().split("|")
  } else {
    lines = rawInput
      .replaceAll("\n", "|")
      .replace(/\s+/g, " ")
      .trim()
      .split("|")
  }

  const timeLine = lines[0]
  const times = timeLine
    .substring(timeLine.indexOf(":") + 1)
    .split(" ")
    .map((val: string) => parseInt(val.trim(), 10))
    .filter((val: number) => !Number.isNaN(val))

  const distLine = lines[1]
  const distances = distLine
    .substring(distLine.indexOf(":") + 1)
    .split(" ")
    .map((val: string) => parseInt(val.trim(), 10))
    .filter((val: number) => !Number.isNaN(val))

  const races = [] as Array<Race>
  for (let i = 0; i < times.length; i++) {
    races.push({
      time: times[i],
      distance: distances[i],
    })
  }

  return races
}

const raceDistance = (race: Race, time: number): number => {
  return (race.time - time) * time
}

const raceWins = (race: Race): number => {
  let wins = 0
  for (let i = 1; i < race.time; i++) {
    const dist = raceDistance(race, i)
    if (dist > race.distance) {
      wins++
    }
  }
  return wins
}

const solve = (races: Array<Race>): number => {
  const wins = [] as Array<number>
  for (const race of races) {
    wins.push(raceWins(race))
  }

  let result = wins[0]
  for (let i = 1; i < wins.length; i++) {
    result *= wins[i]
  }
  return result
}

const part1 = (rawInput: string): number => {
  const races = parseInput(rawInput)

  return solve(races)
}

const part2 = (rawInput: string): number => {
  const races = parseInput(rawInput, true)

  return solve(races)
}

run({
  part1: {
    tests: [
      {
        input: `Time:      7  15   30
Distance:  9  40  200`,
        expected: 288,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `Time:      7  15   30
        Distance:  9  40  200`,
        expected: 71503,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
