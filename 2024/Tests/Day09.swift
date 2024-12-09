import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day09Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    2333133121414131402
    """

  @Test func testPartD91() async throws {
    let challenge = Day09(data: testData1)
    #expect(String(describing: challenge.part1()) == "1928")
  }

  @Test func testPartD92() async throws {
    let challenge = Day09(data: testData1)
    #expect(String(describing: challenge.part2()) == "2858")
  }
}