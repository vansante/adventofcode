import Algorithms
import Foundation

struct Day13: AdventDay {
  struct Point : Hashable {
    let x: Int
    let y: Int
  }

  struct Button {
    let move: Point
    let cost: Int
  }

  struct Machine {
    let buttons: [Button]
    let prize: Point
  }

  var data: String

  func parseLine(line: String) -> Point {
    let pattern = #"\d+"#
    let regex = try! NSRegularExpression(pattern: pattern)
    let range = NSRange(line.startIndex..., in: line)
    let matches = regex.matches(in: line, range: range)
    
    let numbers = matches.compactMap { match -> Int? in
        let range = Range(match.range, in: line)!
        return Int(line[range])
    }
    
    return Point(x: numbers[0], y: numbers[1])
  }

  func getMachines() -> [Machine] {
    let machines = data.split(separator: "\n\n").compactMap { 
      let lines = String($0).split(separator: "\n")

      let btnA = parseLine(line: String(lines[0]))
      let btnB = parseLine(line: String(lines[1]))
      let prize = parseLine(line: String(lines[2]))

      return Machine(buttons: [
        Button(move: btnA, cost: 3),
        Button(move: btnB, cost: 1)
      ], prize: prize)  
    }
    return machines
  }

  func solve(mc: Machine, prizeAddition: Int64 = 0) -> Int64 {
    let btnA = mc.buttons[0].move
    let btnB = mc.buttons[1].move

    let pX = Float64(Int64(mc.prize.x) + prizeAddition)
    let pY = Float64(Int64(mc.prize.y) + prizeAddition)

    let a = Float64(pX * Float64(btnB.y) - pY * Float64(btnB.x)) / Float64(btnA.x * btnB.y - btnA.y * btnB.x)
    
    let b = Float64(Float64(btnA.x) * pY - Float64(btnA.y) * pX) / Float64(btnA.x * btnB.y - btnA.y * btnB.x)

    if floor(a) == a && floor(b) == b {
      return (Int64(a) * Int64(mc.buttons[0].cost)) + (Int64(b) * Int64(mc.buttons[1].cost))
    }
    return 0
  }

  func part1() -> Any {
    let mcs = getMachines()
    var total: Int64 = 0

    for mc in mcs {
      total += solve(mc: mc)
    }
    return total
  }

  func part2() -> Any {
    let mcs = getMachines()
    var total: Int64 = 0

    for mc in mcs {
      total += solve(mc: mc, prizeAddition: 10000000000000)
    }

    return total
  }
}
