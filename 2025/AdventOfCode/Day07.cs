namespace AdventOfCode;

public class Day07 : BaseDay
{
    public enum Type
    {
        Empty,
        Origin,
        Splitter
    }

    public class Coord
    {
        public Type tp;
        public int x;
        public int y;
    }

    
    public class Line
    {
        public List<Coord> coords;
    }

    public class Grid
    {
        public Coord origin;

        public List<Line> lines;
    }

    private readonly string _input;

    private readonly Grid grid = new()
    {
        lines = [],
    };

    public Day07()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        for (int y = 0; y < lines.Length; y++)
        {
            Line l = new()
            {
                coords = [],
            };

            for (int x = 0; x < lines[y].Length; x++)
            {
                Type tp = Type.Empty;
                switch (lines[y][x])
                {
                    case 'S':
                        tp = Type.Origin;
                        break;
                    case '^':
                        tp = Type.Splitter;
                        break;
                }

                Coord c = new()
                {
                    tp = tp,
                    x = x,
                    y = y,
                };

                l.coords.Add(c);
                if (tp == Type.Origin)
                {
                    grid.origin = c;
                }
            }

            grid.lines.Add(l);
        }
    }

    public static int TraceBeams(Grid g)
    {
        int splitCount = 0;
        HashSet<int> beamX = [g.origin.x];

        foreach (Line l in g.lines)
        {
            HashSet<int> newBeamX = [];
            foreach (Coord c in l.coords)
            {
                if (!beamX.Contains(c.x))
                {
                    continue;
                }

                if (c.tp == Type.Splitter)
                {
                    splitCount++;
                    newBeamX.Add(c.x - 1);
                    newBeamX.Add(c.x + 1);
                    continue;
                }

                newBeamX.Add(c.x);
            }
            beamX = newBeamX;
        }

        return splitCount;
    }

    public override ValueTask<string> Solve_1()
    {
        long sum = TraceBeams(grid);

        return new($"The beam was split {sum} times");
    }

    public override ValueTask<string> Solve_2()
    {
        long sum = TraceBeams(grid);

        return new($"The beam was split {sum} times");
    }
}
