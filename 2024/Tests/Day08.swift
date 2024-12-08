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
    #expect(String(describing: challenge.part2()) == "34")
  }

  @Test func testD7AntiNodes() async throws {
    let challenge = Day08(data: testData1)
    let mp = challenge.antennaMap

    #expect(String(describing: challenge.antiNodes(map: mp, c1: (4, 3), c2: (5, 5))) == "[(6, 7), (3, 1)]")
    #expect(String(describing: challenge.antiNodes(map: mp, c1: (5, 5), c2: (4, 3))) == "[(3, 1), (6, 7)]")
    #expect(String(describing: challenge.antiNodes(map: mp, c1: (4, 3), c2: (8, 4))) == "[(12, 5), (0, 2)]")
    #expect(String(describing: challenge.antiNodes(map: mp, c1: (5, 2), c2: (8, 1))) == "[(11, 0), (2, 3)]")
    #expect(String(describing: challenge.antiNodes(map: mp, c1: (8, 1), c2: (5, 2))) == "[(2, 3), (11, 0)]")

    #expect(String(describing: challenge.antiNodes(map: mp, c1: (0, 0), c2: (1, 2), resHarm: true)) == "[(2, 4), (3, 6), (4, 8), (5, 10)]")
    #expect(String(describing: challenge.antiNodes(map: mp, c1: (0, 0), c2: (3, 1), resHarm: true)) == "[(6, 2), (9, 3)]")
  }
}