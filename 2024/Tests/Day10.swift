import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day10Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    0123
    1234
    8765
    9876
    """
  let testData2 = """
    ...0...
    ...1...
    ...2...
    6543456
    7.....7
    8.....8
    9.....9
    """
  let testData3 = """
    ..90..9
    ...1.98
    ...2..7
    6543456
    765.987
    876....
    987....
    """
  let testData4 = """
    89010123
    78121874
    87430965
    96549874
    45678903
    32019012
    01329801
    10456732
    """


  @Test func testPartD101() async throws {
    var challenge = Day10(data: testData1)
    #expect(String(describing: challenge.part1()) == "1")
    challenge.data = testData2
    #expect(String(describing: challenge.part1()) == "2")
    challenge.data = testData3
    #expect(String(describing: challenge.part1()) == "4")
    challenge.data = testData4
    #expect(String(describing: challenge.part1()) == "36")
  }

  @Test func testPartD102() async throws {
    var challenge = Day10(data: testData3)
    #expect(String(describing: challenge.part2()) == "13")
    challenge.data = testData4
    #expect(String(describing: challenge.part2()) == "81")
  }
}
