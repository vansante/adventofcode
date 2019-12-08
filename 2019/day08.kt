import java.io.File

fun main() {
    val width = 25
    val height = 6
    val input = File("day08.txt").readText()
    val layers = input.chunked(width*height)
    var least = Int.MAX_VALUE
    var leastIndex = -1
    var index = 0
    layers.forEach {layer ->
        val count = layer.count {char -> char == '0' }
        if (count < least) {
            least = count
            leastIndex = index
        }
        index++
    }

    val ones = layers[leastIndex].count {it == '1'}
    val twos = layers[leastIndex].count {it == '2'}
    val answer = ones * twos
    println("Part I: Ones: $ones, Twos: $twos, answer: $answer")

    println("Part II: Image")
    for (y in 0 until height) {
        var line = ""
        for (x in 0 until width) {
            line += if (getPixelColor(layers, width, x, y) == 1) {
                "X"
            } else {
                " "
            }
        }
        println(line)
    }
}

fun getPixelColor(layers: List<String>, width: Int, x: Int, y: Int): Int {
    var color = 2;
    var layerIndex = 0
    while (color == 2 && layerIndex < layers.count()) {
        color = layers[layerIndex][(y*width)+x].toString().toInt()
        if (color < 2) {
            return color
        }

        layerIndex++
    }

    throw Exception("color $color found at $x, $y")
}