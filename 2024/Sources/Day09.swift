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

  func findFirstFreeSpace(dsk: [Space]) -> (Int, Bool) {
    var firstFreeSpace = -1
    var lastFreeSpace = -1
    var lastFile = -1

    for (i, s) in dsk.enumerated() {
      let f = s as? File
      if f != nil {
        lastFile = i
        continue
      }

      if s.size <= 0 {
        print("empty free space :(", i, s.size)
        continue
      }

      if firstFreeSpace == -1 {
        firstFreeSpace = i
      }

      lastFreeSpace = i
    }

    return (firstFreeSpace, lastFile > firstFreeSpace)
  }

  func moveFiles(dsk: inout [Space]) {
    var freeSpace = Space(size: 0)
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
        // print("bigger free")
        continue
      }
      if free.size == f!.size {
        dsk[freeIdx] = f!
        dsk.remove(at: i)
        freeSpace.size += f!.size
        i -= 1
        // print("equal free")
        continue
      }

      // At this point, free size is smaller than the wanted file
      let fCopy = File(size: free.size, id: f!.id)
      f!.size -= fCopy.size
      dsk[freeIdx] = fCopy
      freeSpace.size += fCopy.size

      // print("smaller free", fCopy.toString())
    }
  }

  func printDisk(dsk: [Space]) {
    var str = ""
    for s in dsk {
      str += "|"
      if let f = s as? File {
        str += String(repeating: String(f.id), count: f.size)
      } else {
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
        idx += 1
        continue
      }

      for i in 0...f!.size - 1 {
        print(">>", f!.id, idx, i, "==", f!.id * idx)
        total += f!.id * idx
        idx += 1
      }
    }
    return total
  }

  func part1() -> Any {
    var dsk = disk
    
    printDisk(dsk: dsk)
    moveFiles(dsk: &dsk)
    printDisk(dsk: dsk)

    return checksum(dsk: dsk)
  }

  func part2() -> Any {
    return 0
  }
}
