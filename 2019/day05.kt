import java.io.File
import java.lang.Exception
import java.nio.file.Paths

fun main() {
    val input = File("day05.txt").readText()
    val numbers = input.splitToSequence(",").map { it.toInt() }.toList().toIntArray()

    val machine = Machine(memory = numbers, input = 1)
    machine.run()
    println("Part I: Answer is: ${machine.output}")
}


class Machine(private val memory: IntArray, private val input: Int) {
    var output = Int.MIN_VALUE

    fun run() {
        var idx = 0
        while (idx < memory.count()) {
            println("executing idx $idx : ${memory[idx]}")
            idx = execute(idx)
        }
    }

    fun execute(idx: Int): Int {
        val instruction = memory[idx]
        when (getOpcode(instruction)) {
            1 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                println("add: $value1 + $value2")
                setValue(memory[idx+3], value1 + value2)
                return idx + 4
            }
            2 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                println("multiply: $value1 * $value2")
                setValue(memory[idx+3], value1 * value2)
                return idx + 4
            }
            3 -> {
                val param = getValue(true, idx + 1)
                setValue(param, input)
                return idx + 2
            }
            4 -> {
                output = getValue(isModeImmediate(instruction, 2), idx + 1)
                return idx + 2
            }
            99 -> {
                return Int.MAX_VALUE
            }
        }
        throw Exception("invalid opcode at idx $idx: ${getOpcode(instruction)}")
    }

    private fun getValue(immediate: Boolean, address: Int): Int {
        val value = memory[address]
        if (immediate) {
            return value
        }
        return memory[value]
    }

    private fun setValue(address: Int, value: Int) {
        println("storing $address -> $value")
        memory[address] = value
    }

    private fun isModeImmediate(instruction: Int, position: Int): Boolean {
        val positions = instruction.toString().padStart(5, '0').substring(0, 3)
        return positions[position].toString() == "1"
    }

    private fun getOpcode(instruction: Int): Int {
        return instruction % 100
    }
}