import java.io.File
import kotlin.math.abs

fun main() {
    val wires: MutableList<List<String>> = ArrayList()
    File("day03.txt").forEachLine {
        wires += it.split(",")
    }

    val coordinates = mutableSetOf<String>(coord(0, 0))
    var closest = Int.MAX_VALUE

    wires.forEach { wire ->
        var x = 0
        var y = 0
        wire.forEach { direction ->
            var deltaX = 0
            var deltaY   = 0
            val delta = direction.substring(1).toInt()
            when (direction.substring(0, 1)) {
                "U" -> deltaY -= delta
                "D" -> deltaY += delta
                "L" -> deltaX -= delta
                "R" -> deltaX += delta
            }

            if (deltaY < 0) {
                for (i in -1..deltaY) {
                    if (coordinates.contains(coord(x, y+i)) && distance(x, y+i) < closest) {
                        closest = distance(x, y+i)
                    }
                    coordinates.add(coord(x, y+i))
                }
            }
            if (deltaY > 0) {
                for (i in 1..deltaY) {
                    if (coordinates.contains(coord(x, y+i)) && distance(x, y+i) < closest) {
                        closest = distance(x, y+i)
                    }
                    coordinates.add(coord(x, y+i))
                }
            }
            if (deltaX < 0) {
                for (i in -1..deltaX) {
                    if (coordinates.contains(coord(x+i, y)) && distance(x+i, y) < closest) {
                        closest = distance(x+i, y)
                    }
                    coordinates.add(coord(x+i, y))
                }
            }
            if (deltaX > 0) {
                for (i in 1..deltaX) {
                    if (coordinates.contains(coord(x+i, y)) && distance(x+i, y) < closest) {
                        closest = distance(x+i, y)
                    }
                    coordinates.add(coord(x+i, y))
                }
            }

            x += deltaX
            y += deltaY
            println("X: $x, Y: $y")
        }
    }


    println("Part I: Closest is $closest")
}

fun coord(x: Int, y: Int): String {
    return "$x|$y"
}

fun distance(x: Int, y: Int): Int {
    return abs(x) + abs(y)
}