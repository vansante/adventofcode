import Algorithms

struct Day02: AdventDay {
  var data: String

  var reports: [[Int]] {
    data.split(separator: "\n").map {
      $0.split(separator: " ").compactMap { Int($0) }
    }
  }

  func isSafe(report: [Int], dampen: Bool = false) -> Bool {
    var positive: Bool = true
    for i in 1...report.count-1 {
      let diff = report[i] - report[i-1]
      if diff == 0 || abs(diff) > 3 {
        return dampen && dampenedSafe(report: report, idx: i)
      }
      if i == 1 {
        positive = (diff > 0)
        continue
      }
      if positive && diff < 0 || !positive && diff > 0 {
        return dampen && dampenedSafe(report: report, idx: i)
      }
    }
    return true
  }

  func dampenedSafe(report: [Int], idx: Int) -> Bool {
    for j in max(idx-2, 0)...min(idx+2, report.count-1) {
      var newReport = report
      newReport.remove(at: j)
      if isSafe(report: newReport) {
        return true
      }
    }
    return false
  }

  func part1() -> Any {
    return reports.reduce(0, { total, report in
        total + (isSafe(report: report) ? 1 : 0)
    })
  }

  func part2() -> Any {
    return reports.reduce(0, { total, report in
        total + (isSafe(report: report, dampen: true) ? 1 : 0)
    })
  }
}
