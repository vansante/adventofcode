import java.io.File

fun main() {
    val parentToChild = mutableMapOf<String, String>()
    File("day06.txt").forEachLine {
        val input = it.split(")")
        parentToChild[input[1]] = input[0]
    }

    var orbitCount = 0
    parentToChild.forEach { parent, _ ->
        var orbit = parent
        while (orbit != "COM") {
            orbitCount++
            orbit = parentToChild[orbit]!!
        }
    }

    println("Part I: Total orbits: $orbitCount")

    var myCount = 0
    var orbit = parentToChild["YOU"]!!
    val visited = mutableMapOf<String, Int>()
    while (orbit != "COM") {
        myCount++
        visited[orbit] = myCount
        orbit = parentToChild[orbit]!!
    }

    orbit = parentToChild["SAN"]!!
    var sanCount = 0
    while (orbit != "COM") {
        if (visited[parentToChild[orbit]] != null) {
            sanCount += visited[parentToChild[orbit]]!!
            break
        }
        sanCount++
        orbit = parentToChild[orbit]!!
    }

    println("Part II: Total transfers: $sanCount")
}
