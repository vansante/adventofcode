namespace AdventOfCode;

public class Day08 : BaseDay
{
    public class Box
    {
        private static int nextGroupId = 0;

        public int x;
        public int y;
        public int z;

        public HashSet<Box> connectedTo = [];

        public int groupId = 0;

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

            if (groupId != 0)
            {
                other.SetGroupId(groupId);
            } else if (other.groupId != 0)
            {
                groupId = other.groupId;
            } else
            {
                SetGroupId(++nextGroupId);
            }
        }

        public void SetGroupId(int id)
        {
            SetGroupId(id, []);
        }

        private void SetGroupId(int id, HashSet<Box> visited)
        {
            visited.Add(this);
            foreach (Box b in connectedTo)
            {
                if (visited.Contains(b))
                {
                    continue;
                }

                b.SetGroupId(id, visited);
            }

            groupId = id;
        }

        public int CountBoxes()
        {
            return CountBoxes([]);
        }

        public int CountBoxes(HashSet<Box> visited)
        {
            int sum = 1;
            visited.Add(this);
            foreach (Box b in connectedTo)
            {
                if (visited.Contains(b))
                {
                    continue;
                }

                sum += b.CountBoxes(visited);
            }
            return sum;
        }
    }

    private readonly string _input;

    private readonly List<Box> boxes = [];

    public Day08()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        for (int y = 0; y < lines.Length; y++)
        {
            string[] coords = lines[y].Split(",");

            Box b = new()
            {
                x = int.Parse(coords[0]),
                y = int.Parse(coords[1]),
                z = int.Parse(coords[2]),
            };

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

    public void ConnectShortestDistances(int count)
    {
        Dictionary<(Box, Box), double> distances = CalculateDistances();

        double currentShortest = 0;
        for (int i = 0; i < count; i++)
        {
            double shortest = double.MaxValue;
            Box a = null, b = null;

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
        }
    }

    public Dictionary<int, int> GroupConnections()
    {
        Dictionary<int, int> groups = [];

        foreach (Box b in boxes)
        {
            if (b.groupId == 0 || groups.ContainsKey(b.groupId))
            {
                continue;
            }

            groups.Add(b.groupId, b.CountBoxes());
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
        long sum = 0;

        return new($"{sum} ");
    }
}
