import Algorithms

struct Day07: AdventDay {
  var data: String

  var operators: [Int: [Int]] {
    var ops: [Int: [Int]] = [:]
    data.split(separator: "\n").forEach {
      let parts = $0.split(separator: ":")
      let op = Int(parts[0])!
      let vals = parts[1].split(separator: " ").compactMap {
        return Int($0)
      }
      ops[op] = vals
    }
    return ops
  }

  func concatValues(a: Int, b: Int) -> Int {
    var nw = a
    nw *= 10
    if b >= 10 {
      nw *= 10
    }
    if b >= 100 {
      nw *= 10
    }
    if b >= 1000 {
      nw *= 10
    }
    if b >= 10000 {
      nw *= 10
    }
    return nw + b
  }

  func evalEquals(result: Int, vals: [Int], idx: Int, current: Int, concat: Bool) -> Bool {
    if current > result {
      // None of the operators actually decreases the result
      return false
    }

    if idx >= vals.count {
      print("index out of range >= length", idx)
      return false
    }

    if idx == vals.count - 1 {
      return current + vals[idx] == result 
          || current * vals[idx] == result
          || (concat && concatValues(a: current, b: vals[idx]) == result)
    }

    if evalEquals(result: result, vals: vals, idx: idx + 1, current: current + vals[idx], concat: concat) {
      return true
    }
    if evalEquals(result: result, vals: vals, idx: idx + 1, current: current * vals[idx], concat: concat) {
      return true
    }
    if concat && evalEquals(result: result, vals: vals, idx: idx + 1, current: concatValues(a: current, b: vals[idx]), concat: concat) {
      return true
    }
    return false
  }

  func part1() -> Any {
    let ops = operators

    var total = 0
    for (result, vals) in ops {
      if evalEquals(result: result, vals: vals, idx: 1, current: vals[0], concat: false) {
        total += result
      }
    }
    return total
  }

  func part2() -> Any {
    let ops = operators

    var total = 0
    for (result, vals) in ops {
      if evalEquals(result: result, vals: vals, idx: 1, current: vals[0], concat: true) {
        total += result
      }
    }
    return total
  }
}
