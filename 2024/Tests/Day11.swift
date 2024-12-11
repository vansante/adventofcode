import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day11Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    125 17
    """

  @Test func testPartD111() async throws {
    var challenge = Day11(data: testData1)

    // var stones = [0, 1, 10, 99, 999]
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[1, 2024, 1, 0, 9, 9, 2021976]")

    // stones = [125, 17]
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[253000, 1, 7]")
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[253, 0, 2024, 14168]")
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[512072, 1, 20, 24, 28676032]")
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[512, 72, 2024, 2, 0, 2, 4, 2867, 6032]")
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32]")
    // challenge.blink(stones: &stones)
    // #expect(String(describing: stones) == "[2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2]")

    #expect(String(describing: challenge.part1()) == "55312")
  }

  @Test func testPartD112() async throws {
    var challenge = Day11(data: testData1)
    #expect(String(describing: challenge.part2()) == "0")
  }
}
