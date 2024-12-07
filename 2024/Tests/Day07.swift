import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day07Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    190: 10 19
    3267: 81 40 27
    83: 17 5
    156: 15 6
    7290: 6 8 6 15
    161011: 16 10 13
    192: 17 8 14
    21037: 9 7 18 13
    292: 11 6 16 20
    """

  @Test func testPartD71() async throws {
    let challenge = Day07(data: testData1)
    #expect(String(describing: challenge.part1()) == "3749")
  }

  @Test func testPartD72() async throws {
    let challenge = Day07(data: testData1)
    #expect(String(describing: challenge.part2()) == "11387")
  }

  @Test func testD72Concat() async throws {
    let challenge = Day07(data: testData1)
    #expect(String(describing: challenge.concatValues(a: 15, b: 6)) == "156")
    #expect(String(describing: challenge.concatValues(a: 125, b: 71)) == "12571")
    #expect(String(describing: challenge.concatValues(a: 1, b: 724)) == "1724")
  }
}