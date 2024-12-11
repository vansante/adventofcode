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

    var stones = [0: 1, 1: 1, 10: 1, 99: 1, 999: 1]
    stones = challenge.blink(stones: stones)
    print("test1", stones)
    #expect(stones[0] == 1)
    #expect(stones[2024] == 1)
    #expect(stones[1] == 2)
    #expect(stones[9] == 2)
    #expect(stones[2021976] == 1)

    stones = [125: 1, 17: 1]
    stones = challenge.blink(stones: stones)
    #expect(stones[253000] == 1)
    #expect(stones[1] == 1)
    #expect(stones[7] == 1)
    stones = challenge.blink(stones: stones)
    #expect(stones[253] == 1)
    #expect(stones[0] == 1)
    #expect(stones[2024] == 1)
    #expect(stones[14168] == 1)
    stones = challenge.blink(stones: stones)
    #expect(stones[512072] == 1)
    #expect(stones[1] == 1)
    #expect(stones[20] == 1)
    #expect(stones[24] == 1)
    #expect(stones[28676032] == 1)
    stones = challenge.blink(stones: stones)
    #expect(stones[512] == 1)
    #expect(stones[72] == 1)
    #expect(stones[2024] == 1)
    #expect(stones[2] == 2)
    #expect(stones[0] == 1)
    #expect(stones[4] == 1)
    #expect(stones[2867] == 1)
    #expect(stones[6032] == 1)

    #expect(String(describing: challenge.part1()) == "55312")
  }

  @Test func testPartD112() async throws {
    // var challenge = Day11(data: testData1)
    // #expect(String(describing: challenge.part2()) == "0")
  }
}
