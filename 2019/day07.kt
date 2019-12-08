import java.io.File
import java.lang.Exception
import java.nio.file.Paths
import java.util.*
import kotlin.collections.ArrayList

fun main() {
    val input = File("day07.txt").readText()
    val numbers = input.splitToSequence(",").map { it.toInt() }.toList().toIntArray()

    var a = -1
    var b = -1
    var c = -1
    var d = -1
    var e = -1
    var highest = 0
    for (A in 0..4) {
        val machineA = Machine07(memory = numbers.copyOf(), input = mutableListOf(A, 0))
        machineA.run()

        for (B in 0..4) {
            if (B == A) { continue }

            val machineB = Machine07(memory = numbers.copyOf(), input = mutableListOf(B, machineA.lastOutput()))
            machineB.run()

            for (C in 0..4) {
                if (C == A || C == B) { continue }

                val machineC = Machine07(memory = numbers.copyOf(), input = mutableListOf(C, machineB.lastOutput()))
                machineC.run()

                for (D in 0..4) {
                    if (D == A || D == B || D == C) { continue }

                    val machineD = Machine07(memory = numbers.copyOf(), input = mutableListOf(D, machineC.lastOutput()))
                    machineD.run()

                    for (E in 0..4) {
                        if (E == A || E == B || E == C || E == D) { continue }

                        val machineE = Machine07(memory = numbers.copyOf(), input = mutableListOf(E, machineD.lastOutput()))
                        machineE.run()
                        if (machineE.lastOutput() > highest) {
                            highest = machineE.lastOutput()
                            a = A
                            b = B
                            c = C
                            d = D
                            e = E
                        }
                    }
                }
            }
        }
    }

    println("Part I: Highest signal: $highest. Settings: a$a b$b c$c d$d e$e ")

    a = -1
    b = -1
    c = -1
    d = -1
    e = -1
    highest = 0
    for (A in 5..9) {
        for (B in 5..9) {
            if (B == A) { continue }
            for (C in 5..9) {
                if (C == A || C == B)  {continue }
                for (D in 5..9) {
                    if (D == A || D == B || D == C) { continue }
                    for (E in 5..9) {
                        if (E == A || E == B || E == C || E == D) { continue }

                        val temp = mutableListOf<Int>()
                        val machineA = Machine07(numbers.copyOf(), temp)
                        machineA.output().add(B)
                        val machineB = Machine07(numbers.copyOf(), machineA.output())
                        machineB.output().add(C)
                        val machineC = Machine07(numbers.copyOf(), machineB.output())
                        machineC.output().add(D)
                        val machineD = Machine07(numbers.copyOf(), machineC.output())
                        machineD.output().add(E)
                        val machineE = Machine07(numbers.copyOf(), machineD.output())
                        machineE.setOutput(temp)
                        machineE.output().add(A)
                        machineE.output().add(0)

                        val list = listOf(machineA, machineB, machineC, machineD, machineE)
                        var running = false
                        while (true) {
                            list.forEach {
                                if (!it.done()) {
                                    running = true
                                    if (it.runUntilInput()) {
                                        running = false
                                    }
                                }
                            }
                            if (!running) {
                                break
                            }
                        }

                        if (machineE.lastOutput() > highest) {
                            highest = machineE.lastOutput()
                            a = A
                            b = B
                            c = C
                            d = D
                            e = E
                        }
                    }
                }
            }
        }
    }

    println("Part II: Highest signal: $highest. Settings: a$a b$b c$c d$d e$e ")
}

class InputRequired(message: String) : Exception(message)

class Machine07(private val memory: IntArray, private var input: MutableList<Int>) {
    private var output = mutableListOf<Int>()
    private var idx = 0

    fun setOutput(outp: MutableList<Int>) {
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
            } catch (e: InputRequired) {
                return false
            }
        }
        return true
    }

    fun lastOutput(): Int {
        return output.last()
    }

    fun output(): MutableList<Int> {
        return output
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
            3 -> {
                val param = getValue(true, idx + 1)
                if (input.count() == 0) {
                    throw InputRequired("no input")
                }
                setValue(param, input.first())
                input.removeAt(0)
                return idx + 2
            }
            4 -> {
                val value = getValue(isModeImmediate(instruction, 2), idx + 1)
                output.add(value)
                return idx + 2
            }
            5 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                if (value1 != 0) {
                    return value2
                }
                return idx + 3
            }
            6 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                if (value1 == 0) {
                    return value2
                }
                return idx + 3
            }
            7 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                if (value1 < value2) {
                    setValue(memory[idx+3], 1)
                } else {
                    setValue(memory[idx+3], 0)
                }
                return idx + 4
            }
            8 -> {
                val value1 = getValue(isModeImmediate(instruction, 2), idx + 1)
                val value2 = getValue(isModeImmediate(instruction, 1), idx + 2)
                if (value1 == value2) {
                    setValue(memory[idx+3], 1)
                } else {
                    setValue(memory[idx+3], 0)
                }
                return idx + 4
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