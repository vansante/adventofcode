import Algorithms

struct Day11: AdventDay {
  var data: String

  var stones: [Int: Int] {
    var d: [Int: Int] = [:]
    data
      .trimmingCharacters(in: .whitespacesAndNewlines)
      .split(separator: " ")
      .forEach {
        let num = Int($0)!
        d[num] = (d[num] ?? 0) + 1
      }
      return d
  }

  func blink(stones: [Int: Int]) -> [Int: Int] {
    var nw:[Int: Int] = [:]
    for (stoneNum, amount) in stones {
      if stoneNum == 0 {
        // nw[0] = 0
        nw[1] = (nw[1] ?? 0) + amount
        continue
      }

      let label = String(stoneNum)
      if label.count % 2 == 0 {
        let half = label.count / 2
        let upperBound = label.index(label.startIndex, offsetBy: half)

        let firstHalf = Int(label[..<upperBound])!
        nw[firstHalf] = (nw[firstHalf] ?? 0) + amount
        let secondHalf = Int(label[upperBound...])!
        nw[secondHalf] = (nw[secondHalf] ?? 0) + amount
        continue
      }

      let multi = stoneNum * 2024
      nw[multi] = (nw[multi] ?? 0) + amount
    }
    return nw
  }

  func countStones(stones: [Int: Int]) -> Int {
    var total = 0
    for (_, amount) in stones {
      total += amount
    }
    return total
  }

  func part1() -> Any {
    var stones = stones
    for _ in 1...25 {
      stones = blink(stones: stones)
    }
    return countStones(stones: stones)
  }

  func part2() -> Any {
    var stones = stones
    for _ in 1...75 {
      stones = blink(stones: stones)
    }
    return countStones(stones: stones)
  }
}
