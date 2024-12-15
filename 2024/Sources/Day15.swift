import Algorithms

struct Day15: AdventDay {
  let EMPTY = 0
  let WALL = 1
  let BOX = 2
  let ROBOT = 3

  struct Point : Hashable {
    let x: Int
    let y: Int

    func add(p: Point) -> Point {
      return Point(x: x + p.x, y: y + p.y)
    }
  }
 
  var data: String

  let directions: [String : Point] = [
    "^": Point(x: 0, y: -1),
    ">": Point(x: 1, y: 0),
    "v": Point(x: 0, y: 1),
    "<": Point(x: -1, y: 0),
  ]
  
  func getMap() -> [[Int]] {
    let inputs = data.split(separator: "\n\n")
    return inputs[0].split(separator: "\n").compactMap {
      $0.split(separator: "").compactMap {
        switch String($0) {
          case ".":
            return EMPTY
          case "#":
            return WALL
          case "O":
            return BOX
          case "@":
            return ROBOT
          default:
            print("unknown map value: \($0)")
            return EMPTY
        }
      }
    }
  }

  func getInstructions() -> [Point] {
    let inputs = data.split(separator: "\n\n")
    return inputs[1]
      .replacingOccurrences(of: "\n", with: "")
      .split(separator: "")
      .compactMap {
        return directions[String($0)]
      }
  }

  func isEmpty(map: [[Int]], p: Point) -> Bool {
    return p.y >= 0 
      && p.y < map.count 
      && p.x >= 0 
      && p.x < map[p.y].count
      && map[p.y][p.x] != WALL
  }

  func isBox(map: [[Int]], p: Point) -> Bool {
    return map[p.y][p.x] == BOX
  }

  func moveRobot(map: inout [[Int]], start: Point, dir: Point) -> Point {
    var next = start.add(p: dir)
    var emptySpot: Point?
    var moveBox = next
    
    while emptySpot == nil {
      if !isEmpty(map: map, p: next) {
        return start
      }
      if isBox(map: map, p: next) {
        next = next.add(p: dir)
        continue
      }
      
      emptySpot = next
      break
      // We found an empty spot, and can do our scheduled movals
    }

    if emptySpot != nil {
      map[emptySpot!.y][emptySpot!.x] = BOX
      map[moveBox.y][moveBox.x] = EMPTY
    }

    return moveBox
  }

  func grabRobot(map: inout [[Int]]) -> Point {
    for (y, line) in map.enumerated() {
      for (x, val) in line.enumerated() {
        if val == ROBOT {
          map[y][x] = EMPTY
          return Point(x: x, y: y)
        }
      }
    }
    print("robot not found")
    return Point(x: 0, y: 0)
  }

  func sumBoxGPS(map: [[Int]]) -> Int {
    var total = 0
    for (y, line) in map.enumerated() {
      for (x, val) in line.enumerated() {
        if val != BOX {
          continue
        }
        total += (y * 100) + x
      }
    }
    return total
  }

  func part1() -> Any {
    var mp = getMap()
    var instrs = getInstructions()

    printMap(map: mp)

    var robot = grabRobot(map: &mp)
    print(robot)

    for instr in instrs {
      robot = moveRobot(map: &mp, start: robot, dir: instr)
    }

    printMap(map: mp)
    return sumBoxGPS(map: mp)
  }

  func part2() -> Any {
    return 0
  }

  func printMap(map: [[Int]]) {
    var str = ""
    for line in map {
      for val in line {
        switch val {
          case EMPTY:
            str += "."
          case WALL:
            str += "#"
          case BOX:
            str += "o"
          case ROBOT:
            str += "@"
          default:
            print("unknown map value: \(val)")
            str += "."
        }
      }
      str += "\n"
    }
    print(str)
  }
}
