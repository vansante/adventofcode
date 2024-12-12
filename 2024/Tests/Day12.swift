import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day12Tests {
  let testData1 = """
    AAAA
    BBCD
    BBCC
    EEEC
    """

  let testData2 = """
    OOOOO
    OXOXO
    OOOOO
    OXOXO
    OOOOO
    """
  
  let testData3 = """
    RRRRIICCFF
    RRRRIICCCF
    VVRRRCCFFF
    VVRCCCJFFF
    VVVVCJJCFE
    VVIVCCJJEE
    VVIIICJJEE
    MIIIIIJJEE
    MIIISIJEEE
    MMMISSJEEE
    """

  @Test func testPartD121() async throws {
    var challenge = Day12(data: testData1)
    #expect(String(describing: challenge.part1()) == "140")

    challenge = Day12(data: testData2)
    #expect(String(describing: challenge.part1()) == "772")
    
    challenge = Day12(data: testData3)
    #expect(String(describing: challenge.part1()) == "1930")
  }

  @Test func testPartD122() async throws {
    var challenge = Day12(data: testData1)
    #expect(String(describing: challenge.part2()) == "0")
  }
}
