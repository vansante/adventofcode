import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day08Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    
    """

  @Test func testPartD81() async throws {
    let challenge = Day08(data: testData1)
    #expect(String(describing: challenge.part1()) == "0")
  }

  @Test func testPartD82() async throws {
    let challenge = Day08(data: testData1)
    #expect(String(describing: challenge.part2()) == "0")
  }
}