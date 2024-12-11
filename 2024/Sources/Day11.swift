import Algorithms

struct Day11: AdventDay {
  class Node {
    var next: Node?
    var val: Int

    init(val: Int) {
      self.val = val
    }

    init(next: Node?, val: Int) {
      self.next = next
      self.val = val
    }

    func add(val: Int) -> Node {
      let n = Node(next: self.next, val: val)
      self.next = n
      return n
    }
  }

  class List {
    var count: Int
    var head: Node?

    init() {
      self.count = 0
    }

    func add(val: Int) -> Node {
      if head != nil {
        print("list not empty!")
        return Node(next: nil, val: val)
      }
      head = Node(next: nil, val: val)
      count += 1
      return head!
    }

    func printList() {
      var str = ""
      var n: Node? = self.head

      while n != nil {
        str += "\(n?.val), "
        n = n!.next
      }
    }
  }

  var data: String

  var stones: List {
    var l = List()
    var n: Node?
    data
      .trimmingCharacters(in: .whitespacesAndNewlines)
      .split(separator: " ")
      .forEach {
        if n == nil {
          n = l.add(val: Int($0)!)
        } else {
          n!.add(val: Int($0)!)
          l.count += 1
        }
      }
      return l
  }

  func blink(stones: inout List) {
    var n: Node? = stones.head
    while n != nil {
      if n!.val == 0 {
        n!.val = 1
        n = n!.next
        continue
      }

      let label = String(n!.val)
      if label.count % 2 == 0 {
        let half = label.count / 2
        let upperBound = label.index(label.startIndex, offsetBy: half)

        n!.val = Int(label[..<upperBound])!
        n = n!.add(val: Int(label[upperBound...])!)
        stones.count += 1
        n = n!.next
        continue
      }

      n!.val *= 2024
      n = n!.next
    }
  }

  func part1() -> Any {
    var stones: Day11.List = stones
    stones.printList()
    for i in 1...25 {
      blink(stones: &stones)
    }
    stones.printList()
    return stones.count
  }

  func part2() -> Any {
    var stones = stones
    for i in 1...75 {
      print(i)
      blink(stones: &stones)
    }
    return stones.count
  }
}
