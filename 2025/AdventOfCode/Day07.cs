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
        public long timelines = 0;
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
                    c.timelines = 1;
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

    public static long TraceTimelines(Grid g)
    {
        foreach (Line l in g.lines)
        {
            foreach (Coord c in l.coords)
            {
                // Skip the last line
                if (c.y == g.lines.Count - 1)
                {
                    continue;
                }

                if (c.tp == Type.Splitter)
                {
                    g.lines[c.y + 1].coords[c.x - 1].timelines += c.timelines;
                    g.lines[c.y + 1].coords[c.x + 1].timelines += c.timelines;
                    continue;
                }
                
                g.lines[c.y + 1].coords[c.x].timelines += c.timelines;
            }
        }

        long timelineSum = 0;
        foreach (Coord c in g.lines[g.lines.Count - 1].coords)
        {
            timelineSum += c.timelines;
        }

        return timelineSum;
    }

    public override ValueTask<string> Solve_1()
    {
        long sum = TraceBeams(grid);

        return new($"The beam was split {sum} times");
    }

    public override ValueTask<string> Solve_2()
    {
        long sum = TraceTimelines(grid);

        return new($"There are  {sum} timelines");
    }
}
