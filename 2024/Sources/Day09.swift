import Algorithms

struct Day09: AdventDay {
  class Space {
    var size: Int

    init(size: Int) {
      self.size = size
    }

    func toString() -> String {
      return String(repeating: ".", count: size)
    }
  }

  class File : Space {
    let id: Int

    init(size: Int, id: Int) {
      self.id = id
      super.init(size: size)
    }

    override func toString() -> String {
      return String(repeating: String(id), count: size)
    }
  }

  var data: String

  var disk: [Space] {
    var idx = 0
    return data.split(separator: "").compactMap { 
      let num = Int($0)
      if num! <= 0 {
        idx += 1
        return nil
      }
      var i: Space
      if idx % 2 == 0 {
        i = File(size: num!, id: idx / 2)
      } else {
        i = Space(size: num!)
      }
      idx += 1
      return i
    }
  }

  func previousFileId(dsk: [Space], fileId: Int = -1) -> (Int, Int) {
    for i in (0...dsk.count - 1).reversed() {
      let f = dsk[i] as? File
      if f == nil {
        continue
      }
      if fileId == f!.id || fileId == -1 {
        return (i, f!.id)
      }
    }
    return (-1, -1)
  }

  func findFirstFreeSpace(dsk: [Space], size: Int = 1) -> (Int, Bool) {
    var firstFreeSpace = -1
    var lastFile = -1

    for (i, s) in dsk.enumerated() {
      let f = s as? File
      if f != nil {
        lastFile = i
        continue
      }

      if s.size < size {
        continue
      }
      assert(s.size >= size, "1. not enough free space \(s.size) < \(size)")

      if firstFreeSpace == -1 {
        firstFreeSpace = i
      }
    }

    if firstFreeSpace != -1 {
      assert(dsk[firstFreeSpace].size >= size, "2. not enough space free")
    }
    return (firstFreeSpace, lastFile > firstFreeSpace)
  }

  func combineFreeSpace(dsk: inout [Space]) {
    var firstFreeIdx = -1
    var lastFreeIdx = -1
    for (i, s) in dsk.enumerated() {
      let f = s as? File
      if f == nil {
        if firstFreeIdx == -1 {
          firstFreeIdx = i
        }
        lastFreeIdx = i
        continue
      }
      
      if firstFreeIdx < 0 {
        continue
      }
      
      let count = lastFreeIdx - firstFreeIdx
      if count <= 1 {
        firstFreeIdx = -1
        continue
      }

      var space = dsk[lastFreeIdx].size
      for j in firstFreeIdx...lastFreeIdx {
        space += dsk[j].size
        if j < lastFreeIdx {
          dsk.remove(at: j)
        }
      }
      dsk[firstFreeIdx].size = space
      return
    }
  }

  func moveFiles1(dsk: inout [Space]) {
    let freeSpace = Space(size: 0)
    dsk += [freeSpace]

    var i = dsk.count-1
    while i > 0 {
      // printDisk(dsk: dsk)
      let s = dsk[i]
      let f = s as? File
      if f == nil {
        i -= 1
        continue
      }

      let (freeIdx, more) = findFirstFreeSpace(dsk: dsk)
      if !more {
        return
      }

      let free = dsk[freeIdx]
      if free.size > f!.size {
        free.size -= f!.size
        freeSpace.size += f!.size
        dsk.remove(at: i)
        dsk.insert(contentsOf: [f!], at: freeIdx)
        continue
      }
      if free.size == f!.size {
        dsk[freeIdx] = f!
        dsk.remove(at: i)
        freeSpace.size += f!.size
        i -= 1
        continue
      }

      // At this point, free size is smaller than the wanted file
      let fCopy = File(size: free.size, id: f!.id)
      f!.size -= fCopy.size
      dsk[freeIdx] = fCopy
      freeSpace.size += fCopy.size
    }
  }

  func moveFiles2(dsk: inout [Space]) {
    var count = 0
    var (idx, fid) = previousFileId(dsk: dsk)
    while fid >= 0 && idx >= 0 {
      // print("MOVE", fid, idx)
      // printDisk(dsk: dsk)

      let file = dsk[idx]
      let (freeIdx, _) = findFirstFreeSpace(dsk: dsk, size: file.size)
      if freeIdx < 0 || freeIdx >= idx {
        // No move possible
        if fid == 0 {
          return
        }
        (idx, fid) = previousFileId(dsk: dsk, fileId: fid - 1)
        count += 1
        continue
      }

      let free = dsk[freeIdx]
      assert(free.size >= file.size, "3. not enough free space \(file.size) > \(free.size)")
      // print("size", free.size, file.size, free.size >= file.size)
      // print("moving", idx, freeIdx, dsk.count)
      
      dsk[freeIdx] = file
      dsk[idx] = free
      if free.size == file.size {
        combineFreeSpace(dsk: &dsk)
        if fid == 0 {
          return
        }
        (idx, fid) = previousFileId(dsk: dsk, fileId: fid - 1)
        count += 1
        continue
      }
      assert(free.size - file.size > 0, "wrong free size \(free.size) - \(file.size)")
      dsk.insert(contentsOf: [Space(size: free.size - file.size)], at: freeIdx + 1)
      free.size = file.size

      if fid == 0 {
        return
      }
      (idx, fid) = previousFileId(dsk: dsk, fileId: fid - 1)
      count += 1
    }
  }

  func printDisk(dsk: [Space]) {
    var str = ""
    for (idx, s) in dsk.enumerated() {
      str += "|"
      if let f = s as? File {
        assert(f.size >= 0, "file \(idx) id \(f.id) negative size")
        str += String(repeating: String(f.id), count: f.size)
      } else {
        assert(s.size >= 0, "space \(idx) negative size")
        str += String(repeating: ".", count: s.size)
      }
    }
    print(str)
  }

  func checksum(dsk: [Space]) -> Int {
    var total = 0
    var idx = 0
    for s in dsk {
      let f = s as? File
      if f == nil {
        for _ in 0...s.size - 1 {
          idx += 1
        }
        continue
      }

      for _ in 0...f!.size - 1 {
        total += f!.id * idx
        idx += 1
      }
    }
    return total
  }

  func part1() -> Any {
    var dsk = disk
    moveFiles1(dsk: &dsk)
    return checksum(dsk: dsk)
  }

  func part2() -> Any {
    var dsk = disk
    // printDisk(dsk: dsk)
    moveFiles2(dsk: &dsk)
    printDisk(dsk: dsk)
    
    // 15689548622102 too high
    // == 6347435485773
    return checksum(dsk: dsk)
  }
}
