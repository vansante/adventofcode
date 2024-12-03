import Algorithms

struct Day03: AdventDay {
  var data: String

  func part1() -> Any {
    let mulRegex = /mul\((\d+)\,(\d+)\)/
    let results = data.matches(of: mulRegex)

    var total = 0
    for result in results {
      let mul = Int(result.output.1)! * Int(result.output.2)!
      total += mul
    }

    return total
  }

  func part2() -> Any {
    var active = [Int](repeating: 1, count: data.count)

    let doRegex = /do\(\)/
    let doResults = data.matches(of: doRegex)
    for result in doResults {
      let start = result.range.lowerBound.utf16Offset(in: data)
      for idx in start...data.count-1 {
        active[idx] += 1
      }
    }

    let dontRegex = /don't\(\)/
    let dontResults = data.matches(of: dontRegex)
    for result in dontResults {
      let start = result.range.lowerBound.utf16Offset(in: data)
      for idx in start...data.count-1 {
        active[idx] -= 1
      }
    }

    var activeDiff = [Bool](repeating: true, count: data.count)
    for idx in 1...active.count-1 {
      let diff = active[idx] - active[idx-1]
      if diff == 0 {
        activeDiff[idx] = activeDiff[idx-1]
        continue
      }
      if diff < 0 {
        activeDiff[idx] = false
        continue
      }
      if diff > 0 {
        activeDiff[idx] = true
        continue
      }
    }

    let mulRegex = /mul\((\d+)\,(\d+)\)/
    let results = data.matches(of: mulRegex)

    var total = 0
    for result in results {
      let offset = result.range.lowerBound.utf16Offset(in: data)
      if !activeDiff[offset] {
        continue
      }
      let mul = Int(result.output.1)! * Int(result.output.2)!
      total += mul
    }

    return total
  }
}
