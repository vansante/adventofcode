import run from "aocrunner"

const start = 'AAA'
const end = 'ZZZ'

interface Input {
  instr: string
  nodes: Map<string, Node>
}

interface Node {
  name: string
  lft: string
  rgt: string
}

const parseInput = (rawInput: string): Input => {
  const inp = {
    instr: '',
    nodes: new Map(),
  } as Input

  const lines = rawInput.split("\n")

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i]

    if (i === 0) {
      inp.instr = line.trim()
      continue
    }

    if (line.indexOf('=') < 0) {
      continue
    }

    const nd = {
      name: line.substring(0, 3),
      lft: line.substring(7, 10),
      rgt: line.substring(12, 15),
    } as Node

    inp.nodes.set(nd.name, nd)
  }

  return inp
}

const nextNode = (instr: string, nd: Node, nodes: Map<string, Node>): Node|null => {
  switch (instr) {
    case 'L':
      nd = nodes.get(nd.lft) as Node
      if (!nd) {
        console.log('left node not found')
        return null
      }
      break
    case 'R':
      nd = nodes.get(nd.rgt) as Node
      if (!nd) {
        console.log('right node not found')
        return null
      }
      break
    default:
      console.error('invalid instruction', instr)
      return null
  }

  return nd
}

const solve = (inp: Input, nd: Node, isEnd: Function): number => {
  let steps = 0
  let instr = 0
  while (!isEnd(nd)) {
    const instruction = inp.instr.charAt(instr)
    nd = nextNode(instruction, nd, inp.nodes) as Node

    instr++
    if (instr >= inp.instr.length) {
      instr = 0
    }
    steps++
  }

  return steps
}

const part1 = (rawInput: string): number => {
  const input = parseInput(rawInput)

  let nd = input.nodes.get(start) as Node
  if (!nd) {
    console.error('start not found', start)
    return 0
  }

  return solve(input, nd, (nd: Node) => { return nd.name === end})
}

const isStart = (nd: Node): boolean => {
  return nd.name.endsWith('A')
}

const isEnd = (nd: Node): boolean => {
  return nd.name.endsWith('Z')
}

const greatestCommonDivisor = (a: number, b: number): number => {
  if (!b) {
    return b === 0 ? a : NaN;
  }
  return greatestCommonDivisor(b, a%b);
}

const leastCommonMultiple = (a: number, b: number): number => {
  return a * b / greatestCommonDivisor(a, b);
}

const leastCommonMultipleArray = (array: Array<number>): number => {
  let n = 1;
  for (let i = 0; i < array.length; ++i) {
    n = leastCommonMultiple(array[i], n);
  }
  return n;
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)

  let nodes = [] as Array<Node>
  input.nodes.forEach((nd: Node) => {
    if (isStart(nd)) {
      nodes.push(nd)
    }
  })

  const periods = [] as Array<number>
  for (const nd of nodes) {
    periods.push(solve(input, nd, isEnd))
  }

  return leastCommonMultipleArray(periods)
}

run({
  part1: {
    tests: [
      {
        input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
        expected: 2,
      },
      {
        input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
        expected: 6,
      }
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
        expected: 6,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
