import java.io.File
import java.nio.file.Paths

fun main() {
    var totalFuel = 0
    File("day01.txt").forEachLine {
        val mass = it.toInt()
        val fuel = (mass / 3) - 2
        totalFuel += fuel
    }
    println("Part I total = $totalFuel\n")

    totalFuel = 0
    File("day01.txt").forEachLine {
        val mass = it.toInt()
        var fuel = (mass / 3) - 2
        var moduleFuel = 0
        while (fuel > 0) {
            moduleFuel += fuel
            fuel = (fuel / 3) - 2
        }
        totalFuel += moduleFuel
    }
    println("Part II total = $totalFuel")
}