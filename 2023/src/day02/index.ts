import run from "aocrunner"

interface Game {
  id: number
  sets: Array<Set>
}

interface Set {
  red: number
  green: number
  blue: number
}

const parseInput = (rawInput: string): Array<Game> => {
  const lines = rawInput.split("\n")

  const games = []
  for (const line of lines) {
    const game = {
      id: NaN,
      sets: []
    } as Game
    game.id = parseInt(line.substring(5, line.indexOf(':')))

    const sets = line.substring(line.indexOf(':') + 2).split('; ')
    for (const set of sets) {
      const s = {
        red: 0,
        green: 0,
        blue: 0
      } as Set

      const dice = set.split(', ')
      for (const die of dice) {
        const num = parseInt(die.substring(0, die.indexOf(' ')), 10)

        if (die.indexOf('red') >= 0) {
          s.red = num
        } else if (die.indexOf('green') >= 0) {
          s.green = num
        } else {
          s.blue = num
        }
      }

      game.sets.push(s)      
    }

    games.push(game)
  }

  return games
}

const part1 = (rawInput: string): number => {
  const games = parseInput(rawInput)
  
  const maxRed = 12
  const maxGreen = 13
  const maxBlue = 14

  let total = 0

  for (const game of games) {
    let possible = true
    for (const set of game.sets) {
      if (set.red > maxRed || set.green > maxGreen || set.blue > maxBlue) {
        possible = false
      }
    }

    if (possible) {
      total += game.id
    }
  }

  return total
}

const minimumDice = (game: Game, color: string): number => {
  let minimum = 0
  for (const set of game.sets) {
    const s = set as any
    if (s[color] > minimum) {
      minimum = s[color]
    } 
  }
  return minimum
}

const part2 = (rawInput: string): number => {
  const games = parseInput(rawInput)
  
  let total = 0
  for (const game of games) {
    const power = minimumDice(game, 'red') * minimumDice(game, 'green') * minimumDice(game, 'blue')

    total += power
  }

  return total
}

run({
  part1: {
    tests: [
      {
        input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
        expected: 8,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
        expected: 2286,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
})
