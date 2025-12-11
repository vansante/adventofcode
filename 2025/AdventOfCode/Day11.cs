namespace AdventOfCode;

public class Day11 : BaseDay
{
    private readonly string _input;

    public class Node(string id)
    {
        public readonly string id = id;

        public readonly List<Node> connections = [];

        public override string ToString()
        {
            return id;
        }
    }

    private readonly Dictionary<string, Node> nodeMap = [];
    private readonly List<Node> nodes = [];

    public Day11()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        var lines = _input.Split("\n");

        foreach (var line in lines)
        {
            var id = line.Split(':')[0];

            Node n = new(id);
            nodeMap.Add(id, n);
            nodes.Add(n);
        }
        Node nOut = new("out");
        nodes.Add(nOut);
        nodeMap.Add(nOut.id, nOut);

        foreach (var line in lines)
        {
            var parts = line.Split(':');
            var id = parts[0];
            var conns = parts[1].Trim().Split(' ');
            
            var n = nodeMap[id] ?? throw new Exception("node not found");
            foreach (var conn in conns)
            {
                n.connections.Add(nodeMap[conn] ?? throw new Exception("connection node not found"));
            }
        }
    }

    // https://stackoverflow.com/a/59604254
    public static Dictionary<Node, HashSet<Node>> FindAllPredecessors(Node n)
    {
        Dictionary<Node, HashSet<Node>> predecessors = [];
        Queue<Node> queue = [];
        queue.Enqueue(n);

        while (queue.Count > 0)
        {
            var current = queue.Dequeue();
            foreach (var successor in current.connections)
            {
                if (!predecessors.TryGetValue(successor, out HashSet<Node> value))
                {
                    value = [];
                    predecessors[successor] = value;
                }

                value.Add(current);
                queue.Enqueue(successor);
            }
        }

        return predecessors;
    }

    public static List<List<Node>> FindAllPaths(Dictionary<Node, HashSet<Node>> predecessors, Node start, Node end)
    {
        List<List<Node>> paths = [];
        if (start == end)
        {
            paths.Add([start]);
            return paths;
        }

        if (!predecessors.TryGetValue(end, out var parents) || parents.Count == 0)
        {
            // If there are no parents, and we haven't reached the startNode, no path exists.
            return paths;
        }

        // Iterate over all parents
        foreach (var parent in parents)
        {
            // Recursively find all paths from startNode to the current parent
            var pathsFromParent = FindAllPaths(predecessors, start, parent);
            
            // Extend each path found by adding the current endNode
            foreach (var path in pathsFromParent)
            {
                // Must create a new list for the new path to avoid modifying the recursive results
                var newPath = new List<Node>(path)
                {
                    end
                }; 
                paths.Add(newPath);
            }
        }

        return paths;
    }


    public override ValueTask<string> Solve_1()
    {
        var paths = FindAllPaths(FindAllPredecessors(nodeMap["you"]), nodeMap["you"], nodeMap["out"]);

        return new($"Path count is {paths.Count}");
    }

    public static long CountPaths(Node start, Node end)
    {
        return CountPaths([], start, end);
    }

    public static long CountPaths(Dictionary<Node, long> cache, Node start, Node end)
    {
        if (cache.TryGetValue(start, out var count))
        {
            return count;
        }

        if (start.connections.Contains(end))
        {
            return 1;
        }

        var sum = start.connections.Sum(n => CountPaths(cache, n, end));
        cache[start] = sum;
        return sum;
    }

    public override ValueTask<string> Solve_2()
    {
        var pathsFft = CountPaths(nodeMap["svr"], nodeMap["fft"]);
        var pathsDac = CountPaths(nodeMap["fft"], nodeMap["dac"]);
        var pathsOut = CountPaths(nodeMap["dac"], nodeMap["out"]);

        var total = pathsFft * pathsDac * pathsOut;

        return new($"Path count is {total}");
    }
}
