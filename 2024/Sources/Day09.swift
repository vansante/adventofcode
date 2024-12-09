import Algorithms

struct Day09: AdventDay {
  class File {
    let id: Int

    init(id: Int) {
      self.id = id
    }
  }

  var data: String

  var disk: [File?] {
    var dsk: [File?] = []
    var idx = 0
    data.split(separator: "").forEach {
      let num = Int($0)
      if num! <= 0 {
        idx += 1
        return
      }

      if idx % 2 == 0 {
        dsk += Array(repeating: File(id: idx / 2), count: num!)
        idx += 1
        return
      }

      dsk += Array(repeating: nil, count: num!)
      idx += 1
    }
    return dsk
  }

  var maxId: Int {
    data.count / 2
  }

  func moveFiles1(dsk: inout [File?]) {
    for idx in (0...dsk.count - 1).reversed() {
      if dsk[idx] == nil {
        continue
      }

      var newIdx = -1
      for j in (0...idx) {
        if dsk[j] == nil {
          newIdx = j
          break
        }
      }
      if newIdx < 0 || newIdx >= idx {
        continue
      }

      dsk[newIdx] = dsk[idx]
      dsk[idx] = nil
    }
  }

  func findSpace2(dsk: inout [File?], size: Int, beforeIdx: Int) -> Int {
    var idx = 0
    while true {
      if idx >= beforeIdx {
        return -1
      }

      if dsk[idx] != nil {
        idx += 1
        continue
      }

      let start = idx
      for _ in 0...size - 1 {
        idx += 1
        if dsk[idx] != nil {
          break
        }
      }

      if idx - start == size {
        return start
      }
    }
  }

  func moveFile2(dsk: inout [File?], startIdx: Int, endIdx: Int) {
    let size = endIdx - startIdx + 1
    let toIdx = findSpace2(dsk: &dsk, size: size, beforeIdx: startIdx)
    if toIdx < 0 {
      return
    }

    for i in 0...size - 1 {
      dsk[toIdx + i] = dsk[startIdx + i]
      dsk[startIdx + i] = nil
    }
  }

  func findFile(dsk: [File?], id: Int) -> (Int, Int) {
    var rng = (-1, -1)
    for (idx, f) in dsk.enumerated() {
      if f == nil {
        if rng.0 != -1 {
          return rng
        }
        continue
      }

      if f!.id != id {
        if rng.0 != -1 {
          return rng
        }
        continue
      }
      if rng.0 == -1 {
        rng.0 = idx
      }
      rng.1 = idx
    }
    assert(rng.0 >= 0, "file id \(id) not found")
    return rng
  }
  
  func moveFiles2(dsk: inout [File?]) {
    for id in (0...maxId).reversed() {
      let rng = findFile(dsk: dsk, id: id)
      moveFile2(dsk: &dsk, startIdx: rng.0, endIdx: rng.1)
    }
  }

  func checksum(dsk: [File?]) -> Int {
    var total = 0
    var idx = 0
    for f in dsk {
      if f == nil {
        idx += 1
        continue
      }

      total += f!.id * idx
      idx += 1
    }
    return total
  }

  func printDisk(dsk: [File?]) {
    var str = ""
    for f in dsk {
      if f == nil {
        str += "."
        continue
      }
      str += String(f!.id)
    }
    print(str)
  }

  func part1() -> Any {
    var dsk = disk
    // printDisk(dsk: dsk)
    moveFiles1(dsk: &dsk)
    // printDisk(dsk: dsk)
    return checksum(dsk: dsk)
  }

  func part2() -> Any {
    var dsk = disk
    // printDisk(dsk: dsk)
    moveFiles2(dsk: &dsk)
    // printDisk(dsk: dsk)
    return checksum(dsk: dsk)
  }
}
