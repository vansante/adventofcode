import Algorithms

struct Day05: AdventDay {
  var data: String

  var fields: [String] {
    return data.split(separator: "\n\n").compactMap { String($0) }
  }

  var pageOrder: [(Int, Int)] {
    return fields[0].split(separator: "\n").compactMap {
      let nums = $0.split(separator: "|").compactMap { Int($0) }
      return (nums[0], nums[1])
    }
  }

  var pageUpdates: [[Int]] {
    return fields[1].split(separator: "\n").compactMap {
      return $0.split(separator: ",").compactMap { Int($0) }
    }
  }

  var pageRules: [Int: [Int]] {
    let order = pageOrder
    var rules: [Int: [Int]] = [:]
    for ord in order {
      if rules[ord.0] == nil {
        rules[ord.0] = [ord.1]
        continue
      }
      rules[ord.0]! += [ord.1]
    }
    return rules
  }

  func beforePage(order: [(Int,Int)], page: Int) -> [Int] {
    var before: [Int] = []
    for ord in order {
      if ord.1 == page {
        before += [ord.0]
      }
    }
    return before
  }

  func allBeforePage(update: [Int], order: [(Int,Int)]) -> Bool {
    for (idx, page) in update.enumerated() {
      let beforePages = update.prefix(upTo: idx)
      let orderPages = beforePage(order: order, page: page)

      for pg in beforePages {
        if !orderPages.contains(pg) {
          return false
        }
      }
    }

    return true
  }

  func middlePage(update: [Int]) -> Int {
    let idx = update.count / 2
    return update[idx]
  }

  func orderMiddlePage(update: [Int], pageRules: [Int: [Int]]) -> Int {
    var newOrder = update

    newOrder.sort {
      var rulesA = pageRules[$0]
      var rulesB = pageRules[$1]
      if rulesA == nil {
        rulesA = []
      }
      if rulesB == nil {
        rulesB = []
      }

      return rulesA!.contains($1) && !rulesB!.contains($0)
    }

    return middlePage(update: newOrder)
  }

  func part1() -> Any {
    let order = pageOrder
    let updates = pageUpdates

    var total = 0
    for update in updates {
      if !allBeforePage(update: update, order: order) {
        continue
      }
      total += middlePage(update: update)
    }

    return total
  }

  func part2() -> Any {
    let order = pageOrder
    let updates = pageUpdates
    let rules = pageRules

    var total = 0
    for update in updates {
      if allBeforePage(update: update, order: order) {
        continue
      }

      total += orderMiddlePage(update: update, pageRules: rules)
    }
    return total
  }
}
