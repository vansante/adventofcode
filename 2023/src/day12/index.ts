import run from "aocrunner"

interface Condition {
  springs: Array<string>
  groups: Array<Number>
}

const parseInput = (rawInput: string): Array<Condition> => {
  const lines = rawInput.split("\n")

  const conds = [] as Array<Condition>
  for (const line of lines) {
    const cond = {} as Condition
    const parts = line.split(" ")

    cond.springs = parts[0].split("").map((v: string): string => {
      switch (v) {
        case "#":
        case ".":
        case "?":
          return v
      }
      console.error("unknown state", v)
      throw "unknown state"
    })

    cond.groups = parts[1].split(",").map((v: string) => parseInt(v, 10))
    conds.push(cond)
  }

  return conds
}

const count = (
  cond: Condition,
  strPos: number,
  groupPos: number,
  groupLen: number,
): number => {
  if (strPos === cond.springs.length) {
    if (groupPos === cond.groups.length && groupLen === 0) {
      return 1
    }
    if (
      groupPos === cond.groups.length - 1 &&
      cond.groups[groupPos] === groupLen
    ) {
      return 1
    }
    return 0
  }

  let total = 0
  if (cond.springs[strPos] === "." || cond.springs[strPos] === "?") {
    if (groupLen === 0) {
      total += count(cond, strPos + 1, groupPos, 0)
    } else if (
      groupPos < cond.groups.length &&
      cond.groups[groupPos] === groupLen
    ) {
      total += count(cond, strPos + 1, groupPos + 1, 0)
    }
  }
  if (cond.springs[strPos] === "#" || cond.springs[strPos] === "?") {
    total += count(cond, strPos + 1, groupPos, groupLen + 1)
  }

  return total
}

const part1 = (rawInput: string): number => {
  const conds = parseInput(rawInput)

  let total = 0
  for (const cond of conds) {
    console.log(cond, count(cond, 0, 0, 0))
    total += count(cond, 0, 0, 0)
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
        input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`,
        expected: 21,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`,
        expected: 525152,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
