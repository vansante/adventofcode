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

    public List<(double, (Box, Box))> orderedDistances = [];

    public Day08()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        for (int i = 0; i < lines.Length; i++)
        {
            string[] coords = lines[i].Split(",");

            Box b = new(int.Parse(coords[0]), int.Parse(coords[1]), int.Parse(coords[2]));
            boxes.Add(b);
        }

        OrderDistances(CalculateDistances());
    }

    public Dictionary<(Box, Box), double> CalculateDistances()
    {
        Dictionary<(Box, Box), double> distances = [];
        for (int i = 0; i < boxes.Count; i++)
        {
            for (int j = 0; j < boxes.Count; j++)
            {
                if (i == j)
                {
                    continue;
                }

                if (distances.ContainsKey((boxes[j], boxes[i]))) {
                    continue;
                }

                double dist = boxes[i].Distance(boxes[j]);
                distances.Add((boxes[i], boxes[j]), dist);
            }
        }
        return distances;
    }

    public void OrderDistances(Dictionary<(Box, Box), double> distances)
    {
        foreach (KeyValuePair<(Box, Box), double> itm in distances)
        {
            orderedDistances.Add((itm.Value, (itm.Key.Item1, itm.Key.Item2)));
        }
        orderedDistances = [.. orderedDistances.OrderBy(v => v.Item1)];
    }

    public (Box, Box) ConnectShortestDistances(int count = -1)
    {
        Box a = null, b = null;
        for (int i = 0; ; i++)
        {
            (_, (a, b)) = orderedDistances[i];

            a.Connect(b);

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
