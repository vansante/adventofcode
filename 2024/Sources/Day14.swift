import Algorithms
import Foundation

struct Day14: AdventDay {
  struct Point : Hashable {
    let x: Int
    let y: Int

    func add(p: Point, width: Int, height: Int) -> Point {
      return Point(
        x: (x + p.x + width) % width,
        y: (y + p.y + height) % height
      )
    }
  }

  struct Robot {
    let id: Int
    var pos: Point
    let vec: Point

    mutating func move(width: Int, height: Int) {
      pos = pos.add(p: vec, width: width, height: height)
    }
  }

  var data: String

  func getRobots() -> [Robot] {
    var id = 0
    return data.split(separator: "\n").compactMap { 
      id += 1
      return parseRobot(id: id, line: String($0))
    }
  }

  func parseRobot(id: Int, line: String) -> Robot {
    let pattern = #"(-?\d+)"#
    let regex = try! NSRegularExpression(pattern: pattern)
    let range = NSRange(line.startIndex..., in: line)
    let matches = regex.matches(in: line, range: range)
    
    let numbers = matches.compactMap { match -> Int? in
        let range = Range(match.range, in: line)!
        return Int(line[range])
    }
    
    return Robot(id: id, pos: Point(x: numbers[0], y: numbers[1]), vec: Point(x: numbers[2], y: numbers[3]))
  }

  func printMap(bots: [Robot], width: Int, height: Int) {
    var str = ""
    for y in 0..<height {
      for x in 0..<width {
        var b = 0
        for bot in bots {
          if bot.pos.x != x || bot.pos.y != y {
            continue
          }
          b += 1
        }

        if b > 0 {
          str += String(b)
        } else {
          str += "."
        }
      }
      str += "\n"
    }
    print(str)
  }

  func moveBots(bots: inout [Robot], width: Int, height: Int) {
    for var(idx, bot) in bots.enumerated() {
      bot.move(width: width, height: height)
      // Wow, this is quite a stupid quirk:
      // https://augustatseattle.medium.com/mastering-struct-mutations-in-swift-a-deep-dive-into-value-types-and-collections-eaee69b7d876
      bots[idx] = bot
    }
  }

  func countQuadrants(bots: [Robot], width: Int, height: Int) -> [Int] {
    var q: [Int] = [0, 0, 0, 0]
    let midW = (width / 2)
    let midH = (height / 2)

    for b in bots {
      var idx = 0
      if b.pos.x == midW {
        continue
      }
      if b.pos.x > midW {
        idx += 1
      }

      if b.pos.y == midH {
        continue
      }
      if b.pos.y > midH {
        idx += 2
      }
      q[idx] += 1
    }
    return q
  }

  func move1(width: Int = 101, height: Int = 103) -> Any {
    var bots = getRobots()

    for i in 0..<100 {
      moveBots(bots: &bots, width: width, height: height)
    }
    return countQuadrants(bots: bots, width: width, height: height).reduce(1) {
      return $0 * $1
    }
  }

  func part1() -> Any {
    return move1()
  }

  func part2() -> Any {
    return 0
  }
}
