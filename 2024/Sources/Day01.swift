import Algorithms

struct Day01: AdventDay {
  // Save your data in a corresponding text file in the `Data` directory.
  var data: String

  // Splits input data into its component parts and convert from string.
  var lists: [[Int]] {
    var l: [[Int]] = [[], []]
    data.split( separator: "\n").forEach { line in
      let nums: [Int] = line.split(separator: "   ").compactMap { str in Int(String(str)) }
      for (idx, num) in nums.enumerated() {
        l[idx] += [num]
      }
    }

    return l
  }

  func countOccurrences(list: [Int], num: Int) -> Int {
    var count: Int = 0
    list.forEach { n in
      if n == num {
        count += 1
      }
    }
    return count
  }

  func part1() -> Any {
    var l1 = lists[0]
    var l2 = lists[1]

    l1.sort()
    l2.sort()

    if l1.count != l2.count {
      return "not even arrays"
    }

    var totalDiff: Int = 0
    for (idx, num) in l1.enumerated() {
      totalDiff += abs(num - l2[idx])
    }

    return totalDiff
  }

  func part2() -> Any {
    let l1 = lists[0]
    let l2 = lists[1]
    var total: Int = 0
    l1.forEach { num in
      total += self.countOccurrences(list: l2, num: num) * num
    }
    return total
  }
}
