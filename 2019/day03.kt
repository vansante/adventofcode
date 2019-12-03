import java.io.File
import kotlin.math.abs

fun main() {
    val wires: MutableList<List<String>> = ArrayList()
    File("day03.txt").forEachLine {
        wires += it.split(",")
    }

    val coordinates = mutableMapOf<String, Int>()
    var closest = Int.MAX_VALUE
    var wireID = 1
    
    wires.forEach { wire ->
        var x = 0
        var y = 0
        coordinates[coord(0, 0)] = wireID
        wire.forEach { direction ->
            var delta = direction.substring(1).toInt()
            var inc = 1
            when (direction.substring(0, 1)) {
                "U", "D" -> {
                    if (direction.substring(0, 1) == "D") {
                        delta = -delta
                        inc = -1
                    }
                    val origX = x
                    while (true) {
                        val closer = checkCoordinateDistance(coordinates, wireID, x, y)
                        if (closer in 0 until closest) {
                            closest = closer
                        }

                        if (x == origX + delta) {
                            break
                        }
                        x += inc
                    }
                }
                "L", "R" -> {
                    if (direction.substring(0, 1) == "L") {
                        delta = -delta
                        inc = -1
                    }
                    val origY = y
                    while (true) {
                        val closer = checkCoordinateDistance(coordinates, wireID, x, y)
                        if (closer in 0 until closest) {
                            closest = closer
                        }

                        if (y == origY + delta) {
                            break
                        }
                        y += inc
                    }
                }
            }
        }
        wireID++
    }

    println("Part I: Closest is $closest")
}

fun checkCoordinateDistance(coordinates: MutableMap<String, Int>, wireID: Int, x: Int, y: Int): Int {
    val wire = coordinates[coord(x, y)]
    coordinates[coord(x, y)] = wireID
    if (wire == null || wire == wireID) {
        return -1
    }
    return distance(x, y)
}

fun coord(x: Int, y: Int): String {
    return "$x|$y"
}

fun distance(x: Int, y: Int): Int {
    return abs(x) + abs(y)
}