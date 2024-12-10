import Algorithms

struct Day10: AdventDay {
  struct Point : Hashable {
    let x: Int
    let y: Int

    func add(p: Point) -> Point {
      return Point(x: x + p.x, y: y + p.y)
    }
  }
 
  var data: String

  let top = 9

  let directions: [Point] = [
    Point(x: 0, y: -1),
    Point(x: 1, y: 0),
    Point(x: 0, y: 1),
    Point(x: -1, y: 0),
  ]

  var map: [[Int]] {
    return data.replacingOccurrences(of: ".", with: "1").split(separator: "\n").compactMap {
      $0.split(separator: "").compactMap {
        return Int($0)
      }
    }
  }

  func mapHeight(mp: [[Int]], p: Point) -> Int {
    if !inBounds(mp: mp, p: p) {
      return -1
    }
    return mp[p.y][p.x]
  }

  func inBounds(mp: [[Int]], p: Point) -> Bool {
    return p.y >= 0 
      && p.y < mp.count 
      && p.x >= 0 
      && p.x < mp[p.y].count
  }

  func findNeighbours(mp: [[Int]], p: Point, height: Int) -> [Point] {
    var coords: [Point] = []
    for dir in directions {
      let n = p.add(p: dir)
      if mapHeight(mp: mp, p: n) == height {
        coords += [n]
      }
    }
    return coords
  }

  func findHeight(mp: [[Int]], height: Int) -> [Point] {
    var coords: [Point] = []
    for (y, line) in mp.enumerated() {
      for (x, hght) in line.enumerated() {
        if hght == height {
          coords += [Point(x: x, y: y)]
        }
      }
    }
    return coords
  }

  func walkToTop(mp: [[Int]], start: Point, tops: inout Set<Point>) -> Int {
    var routes = 0
    let height = mapHeight(mp: mp, p: start)
    if height == top {
      tops.insert(start)
      routes += 1
      return routes
    }
    
    var p = start
    for h in height...top {
      if h == top {
        tops.insert(p)
        routes += 1
        return routes
      }

      let ns = findNeighbours(mp: mp, p: p, height: h + 1)
      if ns.count > 1 {
        for n in ns {
          routes += walkToTop(mp: mp, start: n, tops: &tops)
        }
        // Let our little spawned functions do the rest of the work
        return routes
      }
      if ns.count > 0 {
        p = ns[0]
        continue
      }
      return routes
    }
    return routes
  }

  func part1() -> Any {
    let mp = map
    let starts = findHeight(mp: mp, height: 0)

    var total = 0
    for start in starts {
      var tops: Set<Point> = []
      walkToTop(mp: mp, start: start, tops: &tops)
      total += tops.count
    }

    return total
  }

  func part2() -> Any {
    let mp = map
    let starts = findHeight(mp: mp, height: 0)

    var total = 0
    for start in starts {
      var tops: Set<Point> = []
      total += walkToTop(mp: mp, start: start, tops: &tops)
    }

    return total
  }
}
