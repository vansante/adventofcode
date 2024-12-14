import Algorithms
import AppKit

struct Day14: AdventDay {
  struct Point : Hashable {
    let x: Int
    let y: Int

    func add(p: Point, width: Int, height: Int) -> Point {
      return Point(
        x: (x + p.x + width) % width,
        y: (y + p.y + height) % height
      )
    }
  }

  struct Robot {
    let id: Int
    var pos: Point
    let vec: Point

    mutating func move(width: Int, height: Int) {
      pos = pos.add(p: vec, width: width, height: height)
    }
  }

  var data: String

  func getRobots() -> [Robot] {
    var id = 0
    return data.split(separator: "\n").compactMap { 
      id += 1
      return parseRobot(id: id, line: String($0))
    }
  }

  func parseRobot(id: Int, line: String) -> Robot {
    let pattern = #"(-?\d+)"#
    let regex = try! NSRegularExpression(pattern: pattern)
    let range = NSRange(line.startIndex..., in: line)
    let matches = regex.matches(in: line, range: range)
    
    let numbers = matches.compactMap { match -> Int? in
        let range = Range(match.range, in: line)!
        return Int(line[range])
    }
    
    return Robot(id: id, pos: Point(x: numbers[0], y: numbers[1]), vec: Point(x: numbers[2], y: numbers[3]))
  }

  func printMap(bots: [Robot], width: Int, height: Int) {
    var str = ""
    for y in 0..<height {
      for x in 0..<width {
        var b = 0
        for bot in bots {
          if bot.pos.x != x || bot.pos.y != y {
            continue
          }
          b += 1
        }

        if b > 0 {
          str += String(b)
        } else {
          str += "."
        }
      }
      str += "\n"
    }
    print(str)
  }

  func moveBots(bots: inout [Robot], width: Int, height: Int) {
    for var(idx, bot) in bots.enumerated() {
      bot.move(width: width, height: height)
      // Wow, this is quite a stupid quirk:
      // https://augustatseattle.medium.com/mastering-struct-mutations-in-swift-a-deep-dive-into-value-types-and-collections-eaee69b7d876
      bots[idx] = bot
    }
  }

  func countQuadrants(bots: [Robot], width: Int, height: Int) -> [Int] {
    var q: [Int] = [0, 0, 0, 0]
    let midW = (width / 2)
    let midH = (height / 2)

    for b in bots {
      var idx = 0
      if b.pos.x == midW {
        continue
      }
      if b.pos.x > midW {
        idx += 1
      }

      if b.pos.y == midH {
        continue
      }
      if b.pos.y > midH {
        idx += 2
      }
      q[idx] += 1
    }
    return q
  }

  func move1(width: Int = 101, height: Int = 103) -> Any {
    var bots = getRobots()

    for i in 0..<100 {
      moveBots(bots: &bots, width: width, height: height)
    }
    return countQuadrants(bots: bots, width: width, height: height).reduce(1) {
      return $0 * $1
    }
  }

  func part1() -> Any {
    return move1()
  }

  func move2(width: Int = 101, height: Int = 103) -> Any {
    var bots = getRobots()

    for i in 0...7500 {
      moveBots(bots: &bots, width: width, height: height)

      if i > 5000 {
        createImage(frame: i, width: 101, height: 103, bots: bots)
      }
    }
    return 0
  }

  func part2() -> Any {
    var bots = getRobots()
    move2()
    return 0
  }
  
  // https://stackoverflow.com/questions/40205830/how-to-create-an-image-pixel-by-pixel
  func createImage(frame: Int, width: Int, height: Int, bots: [Robot]) {
    let colorSpace       = CGColorSpaceCreateDeviceRGB()
    let bytesPerPixel    = 4
    let bitsPerComponent = 8
    let bytesPerRow      = bytesPerPixel * width
    let bitmapInfo       = RGBA32.bitmapInfo

    guard let context = CGContext(data: nil, width: width, height: height, bitsPerComponent: bitsPerComponent, bytesPerRow: bytesPerRow, space: colorSpace, bitmapInfo: bitmapInfo) else {
        print("unable to create context")
        return
    }

    guard let buffer = context.data else {
        print("unable to get context data")
        return
    }

    let pixelBuffer = buffer.bindMemory(to: RGBA32.self, capacity: width * height)

    for y in 0..<height {
      for x in 0..<width {
        var b = 0
        for bot in bots {
          if bot.pos.x != x || bot.pos.y != y {
            continue
          }
          b += 1
        }

        if b == 0 {
          pixelBuffer[y * height + x] = .black
          continue
        }
        pixelBuffer[y * height + x] = RGBA32(red: 10 * UInt8(b), green: 255, blue: 10 * UInt8(b), alpha: 255)
      }
    }

    let cgImage = context.makeImage()!

    let image = NSImage(cgImage: cgImage, size: NSSize(width: Double(width), height: Double(height)))

    // or
    //
    // let image = UIImage(cgImage: cgImage, scale: UIScreen.main.scale, orientation: .up)

    if !image.save(as: "frame_"+String(frame), fileType: .png, at: URL(fileURLWithPath: FileManager.default.currentDirectoryPath + "/Day14_img")) {
      print("error saving")
    }
  }
   
  struct RGBA32: Equatable {
      private var color: UInt32

      var redComponent: UInt8 {
          return UInt8((color >> 24) & 255)
      }

      var greenComponent: UInt8 {
          return UInt8((color >> 16) & 255)
      }

      var blueComponent: UInt8 {
          return UInt8((color >> 8) & 255)
      }

      var alphaComponent: UInt8 {
          return UInt8((color >> 0) & 255)
      }

      init(red: UInt8, green: UInt8, blue: UInt8, alpha: UInt8) {
          color = (UInt32(red) << 24) | (UInt32(green) << 16) | (UInt32(blue) << 8) | (UInt32(alpha) << 0)
      }

      static let bitmapInfo = CGImageAlphaInfo.premultipliedLast.rawValue | CGBitmapInfo.byteOrder32Little.rawValue

      static func ==(lhs: RGBA32, rhs: RGBA32) -> Bool {
          return lhs.color == rhs.color
      }

      static let black = RGBA32(red: 0, green: 0, blue: 0, alpha: 255)
      static let red   = RGBA32(red: 255, green: 0, blue: 0, alpha: 255)
      static let green = RGBA32(red: 0, green: 255, blue: 0, alpha: 255)
      static let blue  = RGBA32(red: 0, green: 0, blue: 255, alpha: 255)
  }

}

// Why do you need to extend the language to do basic stuff? :S

extension NSImage {
    func save(as fileName: String, fileType: NSBitmapImageRep.FileType = .jpeg, at directory: URL = URL(fileURLWithPath: FileManager.default.currentDirectoryPath)) -> Bool {
      guard let tiffRepresentation = tiffRepresentation, directory.isDirectory, !fileName.isEmpty else { print("error dir"); return false }
        do {
            try NSBitmapImageRep(data: tiffRepresentation)?
                .representation(using: fileType, properties: [:])?
                .write(to: directory.appendingPathComponent(fileName).appendingPathExtension(fileType.pathExtension))
            return true
        } catch {
            print(error)
            return false
        }
    }
}

extension URL {
    var isDirectory: Bool {
       return (try? resourceValues(forKeys: [.isDirectoryKey]))?.isDirectory == true
    }
}

extension NSBitmapImageRep.FileType {
    var pathExtension: String {
        switch self {
        case .bmp:
            return "bmp"
        case .gif:
            return "gif"
        case .jpeg:
            return "jpg"
        case .jpeg2000:
            return "jp2"
        case .png:
            return "png"
        case .tiff:
            return "tif"
        default:
          print("unknown type")
          return ""
        }
    }
}

