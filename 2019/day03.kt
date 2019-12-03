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
            var deltaX = 0
            var deltaY = 0
            val delta = direction.substring(1).toInt()
            when (direction.substring(0, 1)) {
                "U" -> deltaY -= delta
                "D" -> deltaY += delta
                "L" -> deltaX -= delta
                "R" -> deltaX += delta
            }

            if (deltaY < 0) {
                for (i in -1..deltaY) {
                    val closer = checkCoordinateDistance(coordinates, wireID, x, y+i)
                    if (closer in 0 until closest) {
                        closest = closer
                    }
                }
            }
            if (deltaY > 0) {
                for (i in 1..deltaY) {
                    val closer = checkCoordinateDistance(coordinates, wireID, x, y+i)
                    if (closer in 0 until closest) {
                        closest = closer
                    }
                }
            }
            if (deltaX < 0) {
                for (i in -1..deltaX) {
                    val closer = checkCoordinateDistance(coordinates, wireID, x+i, y)
                    if (closer in 0 until closest) {
                        closest = closer
                    }
                }
            }
            if (deltaX > 0) {
                for (i in 1..deltaX) {
                    val closer = checkCoordinateDistance(coordinates, wireID, x+i, y)
                    if (closer in 0 until closest) {
                        closest = closer
                    }
                }
            }

            x += deltaX
            y += deltaY
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