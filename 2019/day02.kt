import java.io.File
import java.nio.file.Paths

fun main() {
    val input = File("day02.txt").readText()
    val numbers = input.splitToSequence(",").map{ it.toInt() }.toList().toIntArray()

    val input1 = numbers.copyOf()
    // Replace values:
    input1[1] = 12
    input1[2] = 2
    val result = executeProgram(input1)

    println("Part I: Position zero holds: $result")

    var value1 = 0
    var value2 = 0
    for (i in 1..99) {
        for (j in 1..99) {
            val input2 = numbers.copyOf()

            input2[1] = i
            input2[2] = j
            val result =executeProgram(input2)
            if (result == 19690720) {
                value1 = i
                value2 = j
                break
            }
        }
    }

    println("Part II: Values used are: $value1 and $value2")
    println("Part II: Result is: " + ((value1 * 100) + value2))
}

fun executeProgram(input: IntArray): Int {
    var i = 0;
    while (i >= 0 && i < input.count()) {
        if (input[i] == 99) {
            break
        }

        val pos1 = input[i+1]
        val pos2 = input[i+2]
        val pos3 = input[i+3]
        if (input[i] == 1) {
            input[pos3] = input[pos1] + input[pos2]
        } else if (input[i] == 2) {
            input[pos3] = input[pos1] * input[pos2]
        }

        i += 4
    }

    return input[0]
}