using System.Reflection.Metadata;
using System.Runtime.CompilerServices;

namespace AdventOfCode;

public class Day08 : BaseDay
{
    public class Group
    {
        private static int nextGroupId = 0;

        public int id;

        public HashSet<Box> boxes = [];

        public Group()
        {
            id = ++nextGroupId;
        }

        public void AddBox(Box b)
        {
            boxes.Add(b);
        }
    }

    public class Box
    {
        public int x;
        public int y;
        public int z;

        public HashSet<Box> connectedTo = [];

        public Group group;

        public Box(int x, int y, int z)
        {
            this.x = x;
            this.y = y;
            this.z = z;

            group = new Group();
            groups.Add(group);
        }

        public bool Equals(Box other)
        {
            return x == other.x && y == other.y && z == other.z;
        }

        public double Distance(Box other)
        {
            long xs = x - other.x;
            long ys = y - other.y;
            long zs = z - other.z;
            double d = Math.Sqrt(
                xs * xs
                +
                ys * ys
                +
                zs * zs
            );
            if (double.IsNaN(d))
            {
                throw new Exception("nan");
            }
            return d;
        }

        public void Connect(Box other)
        {
            connectedTo.Add(other);
            other.connectedTo.Add(this);

            other.SetGroup(group);
        }

        private void SetGroup(Group grp)
        {
            SetGroup(grp, []);
        }

        private void SetGroup(Group grp, HashSet<Box> visited)
        {
            visited.Add(this);
            foreach (Box b in connectedTo)
            {
                if (visited.Contains(b))
                {
                    continue;
                }

                b.SetGroup(grp, visited);
            }

            grp.AddBox(this);
            if (grp != group) {
                groups.Remove(group);
            }
            group = grp;
        }
    }

    private readonly string _input;

    private readonly List<Box> boxes = [];

    public static List<Group> groups = [];

    public Day08()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        for (int y = 0; y < lines.Length; y++)
        {
            string[] coords = lines[y].Split(",");

            Box b = new(int.Parse(coords[0]), int.Parse(coords[1]), int.Parse(coords[2]));
            boxes.Add(b);
        }
    }

    public Dictionary<(Box, Box), double> CalculateDistances()
    {
        Dictionary<(Box, Box), double> dict = [];
        for (int i = 0; i < boxes.Count; i++)
        {
            for (int j = 0; j < boxes.Count; j++)
            {
                if (i == j)
                {
                    continue;
                }

                double dist = boxes[i].Distance(boxes[j]);
                dict.Add((boxes[i], boxes[j]), dist);
            }
        }
        return dict;
    }

    public (Box, Box) ConnectShortestDistances(int count = -1)
    {
        Dictionary<(Box, Box), double> distances = CalculateDistances();

        double currentShortest = 0;
        Box a = null, b = null;
        for (int i = 0; ; i++)
        {
            double shortest = double.MaxValue;

            foreach (KeyValuePair<(Box,Box), double> entry in distances)
            {
                if (entry.Value > currentShortest && entry.Value < shortest)
                {
                    a = entry.Key.Item1;
                    b = entry.Key.Item2;
                    shortest = entry.Value;
                }
            }

            if (a == null || b == null)
            {
                throw new Exception("boxes unset");
            }

            a.Connect(b);
            currentShortest = shortest;

            if (count > 0 && i >= count)
            {
                return (a, b);
            }

            if (groups.Count == 1)
            {
                return (a, b);
            }
        }
    }

    public Dictionary<int, int> GroupConnections()
    {
        Dictionary<int, int> groups = [];

        foreach (Box b in boxes)
        {
            if (b.group == null || groups.ContainsKey(b.group.id))
            {
                continue;
            }

            groups.Add(b.group.id, b.group.boxes.Count);
        }
        return groups;
    }

    public override ValueTask<string> Solve_1()
    {
        ConnectShortestDistances(1000);

        Dictionary<int, int> connections = GroupConnections();

        List<int> orderedSize = [.. connections.Values];
        orderedSize.Sort();
        orderedSize.Reverse();

        long answer = orderedSize[0] * orderedSize[1] * orderedSize[2];
        return new($"The three largest circuit are multiplied: {answer}");
    }

    public override ValueTask<string> Solve_2()
    {
        (Box, Box) connection = ConnectShortestDistances(-1);

        int answer = connection.Item1.x * connection.Item2.x;
        return new($"Multiplied X of last connection: {answer}");
    }
}
