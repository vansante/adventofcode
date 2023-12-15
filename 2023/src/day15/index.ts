import run from "aocrunner"

interface Step {
  str: Array<string>
  label: string
  operator: string
  number: number
}

const parseInput = (rawInput: string): Array<Step> => {
  const steps = rawInput.split(",").map((v: string): Step => {
    const step = {
      str: v.split(""),
      label: "",
      number: 0,
    } as Step

    let numbers = [] as Array<string>
    for (const char of step.str) {
      const code = char.charCodeAt(0)
      if (code >= 97 && code <= 122) {
        step.label += char
      }
      if (char === "-" || char === "=") {
        if (step.operator) {
          console.error("multiple operators!")
          throw "multiple operators!"
        }
        step.operator = char
      }
      if (code >= 48 && code <= 57) {
        numbers.push(char)
      }
    }
    if (numbers.length) {
      step.number = parseInt(numbers.join(""), 10)
    }

    return step
  })

  return steps
}

const hash = (str: string): number => {
  let current = 0
  for (const char of str) {
    current += char.charCodeAt(0)
    current *= 17
    current %= 256
  }
  return current
}

const hashes = (steps: Array<Step>): number => {
  let total = 0
  for (const step of steps) {
    let current = 0
    for (const char of step.str) {
      current += char.charCodeAt(0)
      current *= 17
      current %= 256
    }
    total += current
  }

  return total
}

type BoxSet = Map<number, Array<Step>>

const newBoxSet = (): BoxSet => {
  const b = new Map<number, Array<Step>>()

  for (let i = 0; i < 256; i++) {
    b.set(i, [] as Array<Step>)
  }
  return b
}

const hashMap = (steps: Array<Step>, boxes: BoxSet) => {
  for (const step of steps) {
    hashMapStep(step, boxes)
  }
}

const hashMapStep = (step: Step, boxes: BoxSet) => {
  const hsh = hash(step.label)
  const box = boxes.get(hsh)
  if (!box) {
    throw `box ${step.label} | ${hsh} not found`
  }

  const idx = box.findIndex((v: Step): boolean => v.label === step.label)
  switch (step.operator) {
    case "-":
      if (idx >= 0) {
        box.splice(idx, 1)
      }
      break
    case "=":
      if (idx > -1) {
        box[idx] = step
      } else {
        box.push(step)
      }
      break
    default:
      throw `unknown operator ${step.operator}`
  }
}

const focusPower = (boxes: BoxSet): number => {
  let total = 0
  boxes.forEach((lenses: Array<Step>, idx: number) => {
    lenses.forEach((lens: Step, idx2: number) => {
      total += lens.number * (idx + 1) * (idx2 + 1)
    })
  })
  return total
}

const part1 = (rawInput: string): number => {
  const steps = parseInput(rawInput)

  return hashes(steps)
}

const part2 = (rawInput: string): number => {
  const steps = parseInput(rawInput)

  const b = newBoxSet()
  hashMap(steps, b)

  return focusPower(b)
}

run({
  part1: {
    tests: [
      {
        input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
        expected: 1320,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
        expected: 145,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
