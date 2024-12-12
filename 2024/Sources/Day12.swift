import Algorithms

struct Day12: AdventDay {
  struct Point : Hashable {
    let x: Int
    let y: Int

    func add(p: Point) -> Point {
      return Point(x: x + p.x, y: y + p.y)
    }
  }

 struct Region {
  let label: String
  var points: Set<Point>
 }

  var data: String

  var map: [[String]] {
    return data.replacingOccurrences(of: ".", with: "1").split(separator: "\n").compactMap {
      $0.split(separator: "").compactMap {
        return String($0)
      }
    }
  }

  let directions: [Point] = [
    Point(x: 0, y: -1),
    Point(x: 1, y: 0),
    Point(x: 0, y: 1),
    Point(x: -1, y: 0),
  ]

  func inBounds(mp: [[String]], p: Point) -> Bool {
    return p.y >= 0 
      && p.y < mp.count 
      && p.x >= 0 
      && p.x < mp[p.y].count
  }

  func mapLabel(mp: [[String]], p: Point) -> String {
    if !inBounds(mp: mp, p: p) {
      return "?"
    }
    return mp[p.y][p.x]
  }

  func findNeighbours(mp: [[String]], p: Point, label: String) -> [Point] {
    var coords: [Point] = []
    for dir in directions {
      let n = p.add(p: dir)
      if mapLabel(mp: mp, p: n) == label {
        coords += [n]
      }
    }
    return coords
  }

  func expandRegion(mp: [[String]], p: Point, region: inout Region, visited: inout Set<Point>) {
    visited.insert(p)
    region.points.insert(p)

    let nbs = findNeighbours(mp: mp, p: p, label: region.label)
    for nb in nbs {
      if visited.contains(nb) {
        continue
      }
      expandRegion(mp: mp, p: nb, region: &region, visited: &visited)
    }
  }

  func findRegions(mp: [[String]], regions: inout [Point: Region], visited: inout Set<Point>) {
    for (y, line) in mp.enumerated() {
      for (x, _) in line.enumerated() {
        let p = Point(x: x, y: y)
        if visited.contains(p) {
          continue
        }

        var region = Region(label: mapLabel(mp: mp, p: p), points: Set<Point>())
        expandRegion(mp: mp, p: p, region: &region, visited: &visited)
        regions[p] = region
      }
    }
  }

  func regionFences(mp: [[String]], region: Region) -> Int {
    var fences = 0
    for p in region.points {
      let nbs = findNeighbours(mp: mp, p: p, label: region.label)
      fences += 4 - nbs.count
    }
    return fences
  }

  func regionPrice(mp: [[String]], region: Region) -> Int {
    return region.points.count * regionFences(mp: mp, region: region)
  }

  func part1() -> Any {
    let mp = map
    var regions: [Point: Region] = [:]
    var visited = Set<Point>()
    
    findRegions(mp: mp, regions: &regions, visited: &visited)

    var total = 0
    for (_, region) in regions {
      total += regionPrice(mp: mp, region: region)
    }
    return total
  }

  func part2() -> Any {
    return 0
  }
}
