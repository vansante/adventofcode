using Microsoft.Z3;
using System.Text;

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

    Dictionary<string, Node> nodeMap = [];
    private readonly List<Node> nodes = [];

    private readonly Node start, end;

    public Day11()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        foreach (string line in lines)
        {
            string id = line.Split(':')[0];

            Node n = new(id);
            if (id == "you")
            {
                start = n;
            }
            nodeMap.Add(id, n);
            nodes.Add(n);
        }
        end = new("out");
        nodes.Add(end);
        nodeMap.Add(end.id, end);

        foreach (string line in lines)
        {
            string[] parts = line.Split(':');
            string id = parts[0];
            string[] conns = parts[1].Trim().Split(' ');
            
            Node n = nodeMap[id] ?? throw new Exception("node not found");
            foreach (string conn in conns)
            {
                n.connections.Add(nodeMap[conn] ?? throw new Exception("connection node not found"));
            }
        }
    }

    // https://stackoverflow.com/a/59604254
    public Dictionary<Node, HashSet<Node>> FindAllPredecessors(Node n)
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

    public List<List<Node>> FindAllPaths(Dictionary<Node, HashSet<Node>> predecessors, Node start, Node end)
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
        foreach (Node parent in parents)
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
        var paths = FindAllPaths(FindAllPredecessors(start), start, end);

        return new($"Path count is {paths.Count}");
    }

    public override ValueTask<string> Solve_2()
    {
        long sum = 0;

        return new($"{sum}");
    }
}
