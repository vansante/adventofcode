import Testing

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
struct Day14Tests {
  let testData1 = """
    p=0,4 v=3,-3
    p=6,3 v=-1,-3
    p=10,3 v=-1,2
    p=2,0 v=2,-1
    p=0,0 v=1,3
    p=3,0 v=-2,-2
    p=7,6 v=-1,-3
    p=3,0 v=-1,-2
    p=9,3 v=2,3
    p=7,3 v=-1,2
    p=2,4 v=2,-3
    p=9,5 v=-3,-3
    """

  @Test func testPartD141() async throws {
    let challenge = Day14(data: testData1)
    #expect(String(describing: challenge.move1(width: 11, height: 7)) == "12")
  }

  @Test func testPartD142() async throws {
   var challenge = Day14(data: testData1)
   #expect(String(describing: challenge.part2()) == "?")
  }
}
