import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day08Tests {
  // Smoke test data provided in the challenge question
  let testData1 = """
    ............
    ........0...
    .....0......
    .......0....
    ....0.......
    ......A.....
    ............
    ............
    ........A...
    .........A..
    ............
    ............
    """

  @Test func testPartD81() async throws {
    let challenge = Day08(data: testData1)
    #expect(String(describing: challenge.part1()) == "14")
  }

  @Test func testPartD82() async throws {
    let challenge = Day08(data: testData1)
    #expect(String(describing: challenge.part2()) == "0")
  }

  @Test func testD7AntiNodes() async throws {
    let challenge = Day08(data: testData1)
    #expect(String(describing: challenge.antiNodes(c1: (4, 3), c2: (5, 5))) == "[(3, 1), (6, 7)]")
    #expect(String(describing: challenge.antiNodes(c1: (5, 5), c2: (4, 3))) == "[(3, 1), (6, 7)]")
    #expect(String(describing: challenge.antiNodes(c1: (4, 3), c2: (8, 4))) == "[(0, 2), (12, 5)]")
    #expect(String(describing: challenge.antiNodes(c1: (5, 2), c2: (8, 1))) == "[(2, 3), (11, 0)]")
    #expect(String(describing: challenge.antiNodes(c1: (8, 1), c2: (5, 2))) == "[(2, 3), (11, 0)]")
  }
}