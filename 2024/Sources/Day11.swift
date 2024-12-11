import Algorithms

struct Day11: AdventDay {
  struct Node {
    var next: Node?
    var val: Int

    func insert(val: Int) -> Node {
      let n = Node(next: self.next, val: val)
      self.next = n
      return n
    }
  }
  struct List {
    var count: Int
    var head: Node?

    func add(val: Int) -> Node {
      if head != nil {
        print("list not empty!")
        return Node(next: nil, val: val)
      }
      self.head = Node(next: nil, val: val)
      self.count += 1
    }
  }

  var data: String

  var stones: [Int] {
    return data
      .trimmingCharacters(in: .whitespacesAndNewlines)
      .split(separator: " ")
      .compactMap {
        return Int($0)
      }
  }

  func blink(stones: inout [Int]) {
    var idx = 0
    while idx < stones.count {
      // print(idx)
      if stones[idx] == 0 {
        stones[idx] = 1
        idx += 1
        continue
      }

      let label = String(stones[idx])
      if label.count % 2 == 0 {
        let half = label.count / 2
        let upperBound = label.index(label.startIndex, offsetBy: half)

        stones[idx] = Int(label[..<upperBound])!
        stones.insert(Int(label[upperBound...])!, at: idx + 1)
        // blink(stones: &stones, startIdx: idx + 2)
        idx += 2
        // FIXME: Skip next iteration
        // return
        continue
      }

      stones[idx] *= 2024
      idx += 1
    }
  }

  func part1() -> Any {
    var stones = stones
    print(stones)
    for i in 1...25 {
      blink(stones: &stones)
    }
    // print(stones)
    return stones.count
  }

  func part2() -> Any {
    var stones = stones
    for i in 1...75 {
      blink(stones: &stones)
    }
    return stones.count
  }
}
