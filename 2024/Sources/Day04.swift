import Algorithms

struct Day04: AdventDay {
  var data: String

  var word1: [String] {
    ["X", "M", "A", "S"]
  }

  var word2: [String] {
    ["M", "A", "S"]
  }

  var directions1: [(Int, Int)] {
    [
      (1, 0),
      (1, 1),
      (1, -1),
      (0, 1),
      (0, -1),
      (-1, 0),
      (-1, 1),
      (-1, -1),
    ]
  }

  var grid: [[String]] {
    data.split(separator: "\n").map {
      $0.split(separator: "").map {
        String($0)
      }
    }
  }

  func findWord(grid: [[String]], word: [String], xStart: Int, yStart: Int, xDelta: Int, yDelta: Int) -> Bool {
    var x = xStart, y = yStart
    for i in 0...word.count-1 {
      if y < 0 || y >= grid.count {
        return false
      }
      if x < 0 || x >= grid[y].count {
        return false
      }
      if grid[y][x] != word[i] {
        return false
      }
      x += xDelta
      y += yDelta
    }

    return true
  }

  func findWordDirections1(grid: [[String]], word: [String], xStart: Int, yStart: Int) -> Int {
    var total = 0
    for dir in directions1 {
      if findWord(grid: grid, word: word, xStart: xStart, yStart: yStart, xDelta: dir.0, yDelta: dir.1) {
        total += 1
      }
    }
    return total
  }

  func findWordCross(grid: [[String]], word: [String], xStart: Int, yStart: Int) -> Bool {
    var revWord = word
    revWord.reverse()

    let firstWordFound = 
      findWord(grid: grid, word: word, xStart: xStart, yStart: yStart, xDelta: 1, yDelta: 1)
      || findWord(grid: grid, word: revWord , xStart: xStart, yStart: yStart, xDelta: 1, yDelta: 1)

    if !firstWordFound {
      return false
    }

    return (
      findWord(grid: grid, word: word, xStart: xStart + 2, yStart: yStart, xDelta: -1, yDelta: 1)
      || findWord(grid: grid, word: revWord, xStart: xStart + 2, yStart: yStart, xDelta: -1, yDelta: 1)
    )
  }

  func part1() -> Any {
    let g = grid
    
    var total = 0
    for y in 0...g.count-1 {
      for x in 0...g[y].count-1 {
        total += findWordDirections1(grid: g, word: word1, xStart: x, yStart: y)
      }
    }
    return total
  }

  func part2() -> Any {
    let g = grid
    
    var total = 0    
    for y in 0...g.count-1 {
      for x in 0...g[y].count-1 {
        if findWordCross(grid: g, word: word2, xStart: x, yStart: y) {
          total += 1
        }
      }
    }
    return total
  }
}
