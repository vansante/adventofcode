import java.io.File
import kotlin.math.abs

fun main() {
    val inputLower = 273025
    val inputUpper = 767253

    var part1Matches = 0
    var part2Matches = 0

    for (i in inputLower..inputUpper) {
        var highest = 0
        var match = true
        var doubleChar = false
        var exclusiveDoubleChar = false
        var lastChar = 'a' // never matches
        var charMatches = 1
         i.toString().forEach loop@{ char ->
            if (char.toInt() < highest) {
                match = false
                return@loop
            }
            highest = char.toInt()
            if (lastChar == char) {
                charMatches++
            } else {
                if (charMatches > 2) {
                    doubleChar = true
                } else if (charMatches == 2) {
                    doubleChar = true
                    exclusiveDoubleChar = true
                }
                charMatches = 1
            }
            lastChar = char
        }
        if (charMatches > 2) {
            doubleChar = true
        } else if (charMatches == 2) {
            doubleChar = true
            exclusiveDoubleChar = true
        }

        if (match && doubleChar) {
            part1Matches++
        }
        if (match && exclusiveDoubleChar) {
            part2Matches++
        }
    }

    println("Part I: Total matches: $part1Matches")
    println("Part II: Total matches: $part2Matches")
}
