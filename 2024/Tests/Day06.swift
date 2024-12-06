import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day06Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    ....#.....
    .........#
    ..........
    ..#.......
    .......#..
    ..........
    .#..^.....
    ........#.
    #.........
    ......#...
    """

  @Test func testPartD61() async throws {
    let challenge = Day06(data: testData1)
    #expect(String(describing: challenge.part1()) == "41")
  }

  @Test func testPartD62() async throws {
    let challenge = Day06(data: testData1)
    #expect(String(describing: challenge.part2()) == "6")
  }
}