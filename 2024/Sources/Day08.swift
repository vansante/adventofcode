import Algorithms

struct Day08: AdventDay {
  struct AntennaType {
    let antenna: String
    var coords: [(Int, Int)]
  }

  struct Location: Hashable {
    let antenna: String
    var antiNode: Bool
  }

  var data: String

  var antennaMap: [[Location]] {
    return data.split(separator: "\n").compactMap { 
      $0.split(separator: "").compactMap {
        if String($0) == "." {
          return Location(antenna: "", antiNode: false)
        }
        return Location(antenna: String($0), antiNode: false)
      }
    }
  }

  var antennaTypes: [String:AntennaType] {
    var mapping:[String:AntennaType] = [:]
    for (y, line) in antennaMap.enumerated() {
      for (x, loc) in line.enumerated() {
        if loc.antenna == "" {
          continue
        }

        var ant = mapping[loc.antenna]
        if ant != nil {
          ant!.coords += [(x, y)]
          mapping[loc.antenna] = ant
          continue
        }

        let newAnt = AntennaType(antenna: loc.antenna, coords: [(x, y)])
        mapping[loc.antenna] = newAnt
      }
    }
    return mapping
  }

  func antiNodes(c1: (Int, Int), c2: (Int, Int)) -> [(Int, Int)] {
    return [
      (2 * c2.0 - c1.0, 2 * c2.1 - c1.1),
      (2 * c1.0 - c2.0, 2 * c1.1 - c2.1),
    ]
  }

  func markAntiNode(map: inout [[Location]], c: (Int, Int)) {
    if c.1 < 0 || c.1 >= map.count {
      return
    }
    if c.0 < 0 || c.0 >= map[c.1].count {
      return
    }
    map[c.1][c.0].antiNode = true
  }

  func markAntiNodes(antenna: AntennaType, map: inout [[Location]]) {
    for (idx1, c1) in antenna.coords.enumerated() {
      for (idx2, c2) in antenna.coords.enumerated() {
        if idx1 == idx2 {
          continue
        }
        for anti in antiNodes(c1: c1, c2: c2) {
          markAntiNode(map: &map, c: anti)
        }
      }
    }
  }

  func countAntiNodes(map: [[Location]]) -> Int {
    var total = 0
    for line in map {
      for loc in line {
        if loc.antiNode {
          total += 1
          continue
        }
      }
    }
    return total
  }

  func printMap(map: [[Location]], anti: Bool) {
    var mp = ""
    for line in map {
      for loc in line {
        if !anti && loc.antenna != "" {
          mp += loc.antenna
          continue
        }
        if loc.antiNode {
         mp += "#"
         continue
        }
        mp += "."
      }
      mp += "\n"
    }
    print(mp)
  }

  func part1() -> Any {
    var map = antennaMap
    let tps = antennaTypes

    // printMap(map: map, anti: true)

    for (_, tp) in tps {
      markAntiNodes(antenna: tp, map: &map)
    }
    // printMap(map: map, anti: true)
    // printMap(map: map, anti: false)
    // 356 too low
    return countAntiNodes(map: map)
  }

  func part2() -> Any {
    return 0
  }
}
