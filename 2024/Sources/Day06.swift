import Algorithms

struct Day06: AdventDay {
  struct Point: Hashable {
    let x: Int
    let y: Int

    func add(p: Point) -> Point {
      return Point(x: p.x + x, y: p.y + y)
    }
    func rotate() -> Point {
      return Point(x: -y, y: x)
    }
  }

  var data: String

  var map: [[Bool]] {
    return data.split(separator: "\n").compactMap {
      return $0.split(separator: "").compactMap { $0 == "#" }
    }
  }

  var guardPosition: Point {
    for (y, line) in data.split(separator: "\n").enumerated() {
      if !line.contains("^") {
        continue
      }
      let x = line.firstIndex(of: "^")?.utf16Offset(in: line)
      return Point(x: x!, y: y)
    }
    print("guard not found!")
    return Point(x: -1, y: -1)
  }

  var guardDirection: Point {
    return Point(x: 0, y: -1)
  }

  func onMap(mp: [[Bool]], p: Point) -> Bool {
    if p.y < 0 || p.y >= mp.count {
      return false
    }
    if p.x < 0 || p.x >= mp[p.y].count {
      return false
    }
    return true
  }

  func hasObstacle(mp: [[Bool]], p: Point) -> Bool {
    return onMap(mp: mp, p: p) && mp[p.y][p.x]
  }

  func walk(mp: [[Bool]], start: Point, direction: Point) -> Int {
    var p = start
    var dir = direction
    var positions = Set<Point>()
    var visited: [Point: Set<Point>] = [:]

    while true {
      if visited[p] == nil {
        visited[p] = Set<Point>()
        visited[p]!.insert(dir)
      } else if visited[p]!.contains(dir) {
        // Circle complete
        print("complete")
        break
      } else {
        visited[p]!.insert(dir)
      }

      let next = p.add(p: dir)
      if !onMap(mp: mp, p: next) {
        print("leaving map")
        break
      }
      if hasObstacle(mp: mp, p: next){
        print("rotate")
        dir = dir.rotate()
        continue
      }
      print("walk", p, next)
      p = next
      positions.insert(p)
    }

    return positions.count
  }

  func part1() -> Any {
    let mp = map
    let pos = guardPosition
    let dir = guardDirection

    return walk(mp: mp, start: pos, direction: dir)
  }

  func part2() -> Any {

    return 0
  }
}
