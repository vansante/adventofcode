import java.io.File
import java.lang.Exception
import java.math.BigInteger

fun main() {
    val input = File("day09.txt").readText()
    val numbers = input.splitToSequence(",").map { it.toLong() }.toList().toLongArray()

    val machine1 = Machine09(numbers.copyOf(5000), mutableListOf(1))
    machine1.run()

    println("Part II: Keycode is ${machine1.lastOutput()}")

    val machine2 = Machine09(numbers.copyOf(5000), mutableListOf(2))
    machine2.run()

    println("Part II: Coordinates are ${machine2.lastOutput()}")
}

class InputRequired09(message: String) : Exception(message)

class Machine09(private val memory: LongArray, private var input: MutableList<Long>) {
    private var output = mutableListOf<Long>()
    private var idx = 0
    private var relBase = 0

    fun setOutput(outp: MutableList<Long>) {
        output = outp
    }

    fun done(): Boolean {
        return idx >= memory.count()
    }

    fun run() {
        while (idx < memory.count()) {
            idx = execute(idx)
        }
    }

    fun runUntilOutput() {
        val startLen = output.count()
        while (idx < memory.count()) {
            idx = execute(idx)
            if (output.count() > startLen) {
                break
            }
        }
    }

    fun runUntilInput(): Boolean {
        while (idx < memory.count()) {
            try {
                idx = execute(idx)
            } catch (e: InputRequired09) {
                return false
            }
        }
        return true
    }

    fun lastOutput(): Long {
        return output.last()
    }

    fun output(): MutableList<Long> {
        return output
    }

    fun execute(idx: Int): Int {
        val instruction = memory[idx]
        when (getOpcode(instruction)) {
            1 -> {
                val value1 = getValue(valueMode(instruction, 2), idx + 1)
                val value2 = getValue(valueMode(instruction, 1), idx + 2)
                setValue(valueMode(instruction, 0), idx + 3, value1 + value2)
                return idx + 4
            }
            2 -> {
                val value1 = getValue(valueMode(instruction, 2), idx + 1)
                val value2 = getValue(valueMode(instruction, 1), idx + 2)
                setValue(valueMode(instruction, 0), idx + 3, value1 * value2)
                return idx + 4
            }
            3 -> {
                if (input.count() == 0) {
                    throw InputRequired09("no input")
                }
                setValue(valueMode(instruction, 2), idx + 1, input.first())
                input.removeAt(0)
                return idx + 2
            }
            4 -> {
                val value = getValue(valueMode(instruction, 2), idx + 1)
                output.add(value)
                return idx + 2
            }
            5 -> {
                val value1 = getValue(valueMode(instruction, 2), idx + 1)
                val value2 = getValue(valueMode(instruction, 1), idx + 2)
                if (value1 != 0L) {
                    return value2.toInt()
                }
                return idx + 3
            }
            6 -> {
                val value1 = getValue(valueMode(instruction, 2), idx + 1)
                val value2 = getValue(valueMode(instruction, 1), idx + 2)
                if (value1 == 0L) {
                    return value2.toInt()
                }
                return idx + 3
            }
            7 -> {
                val value1 = getValue(valueMode(instruction, 2), idx + 1)
                val value2 = getValue(valueMode(instruction, 1), idx + 2)
                if (value1 < value2) {
                    setValue(valueMode(instruction, 0), idx + 3, 1)
                } else {
                    setValue(valueMode(instruction, 0), idx + 3, 0)
                }
                return idx + 4
            }
            8 -> {
                val value1 = getValue(valueMode(instruction, 2), idx + 1)
                val value2 = getValue(valueMode(instruction, 1), idx + 2)
                if (value1 == value2) {
                    setValue(valueMode(instruction, 0), idx + 3, 1)
                } else {
                    setValue(valueMode(instruction, 0), idx + 3, 0)
                }
                return idx + 4
            }
            9 -> {
                val value = getValue(valueMode(instruction, 2), idx + 1)
                relBase += value.toInt()
                return idx + 2
            }
            99 -> {
                return Int.MAX_VALUE
            }
        }
        throw Exception("invalid opcode at idx $idx: ${getOpcode(instruction)}")
    }

    private fun getValue(mode: Int, address: Int): Long {
        val value = memory[address]
        when (mode) {
            0 -> {
                return memory[value.toInt()]
            }
            1 -> {
                return value
            }
            2 -> {
                return memory[relBase + value.toInt()]
            }
        }
        throw Exception("unknown parameter mode $mode")
    }

    private fun setValue(mode: Int, address: Int, value: Long) {
        val addr = memory[address].toInt()
        when (mode) {
            0 -> {
                memory[addr] = value
            }
            2 -> {
                memory[relBase + addr] = value
            }
            else -> throw Exception("invalid parameter mode $mode")
        }
    }

    private fun valueMode(instruction: Long, position: Int): Int {
        val positions = instruction.toString().padStart(5, '0').substring(0, 3)
        return positions[position].toString().toInt()
    }

    private fun getOpcode(instruction: Long): Int {
        return instruction.toInt() % 100
    }
}