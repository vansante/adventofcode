import java.io.File
import java.lang.Exception
import java.nio.file.Paths

fun main() {
    val input = File("day02.txt").readText()
    val numbers = input.splitToSequence(",").map { it.toInt() }.toList().toIntArray()

    val machine = Machine(memory = numbers)
    machine.run()
}


class Machine(private val memory: IntArray) {
    fun run() {
        var idx = 0
        while (idx < memory.count()) {
            idx = execute(idx)
        }
    }

    fun execute(idx: Int): Int {
        val instruction = memory[idx]
        when (getOpcode(instruction)) {
            1 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                setValue(memory[idx+3], value1 + value2)
                return idx + 4
            }
            2 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                setValue(memory[idx+3], value1 * value2)
                return idx + 4
            }
            99 -> {
                return Int.MAX_VALUE
            }
        }
        throw Exception("invalid opcode at idx $idx: ${getOpcode(idx)}")
    }

    private fun getValue(immediate: Boolean, address: Int): Int {
        val value = memory[address]
        if (immediate) {
            return value
        }
        return memory[value]
    }

    private fun setValue(address: Int, value: Int) {
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