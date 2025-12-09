using System.Text;

namespace AdventOfCode;

public class Day09 : BaseDay
{
    public class Coord(int x, int y)
    {
        public int x = x;
        public int y = y;

        public bool Equals(Coord other)
        {
            return x == other.x && y == other.y;
        }

        public long Area(Coord other)
        {
            return ( (long) Math.Abs( x - other.x) + 1) * ((long) Math.Abs(y - other.y) + 1);
        }

        public bool AreaContains(Coord square, Coord point)
        {
            int xMin = Math.Min(x, square.x);
            int xMax = Math.Max(x, square.x);
            int yMin = Math.Min(y, square.y);
            int yMax = Math.Max(y, square.y);

            return (
                point.x > xMin && point.x < xMax
                && point.y > yMin && point.y < yMax
            );
        }
    }

    public enum Tile
    {
        Other,
        Red,
        Green,
    }

    private readonly string _input;

    private readonly List<Coord> coords = [];

    private readonly Dictionary<(int, int), Tile> grid = [];

    private readonly int maxX = 0;

    public Day09()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        for (int i = 0; i < lines.Length; i++)
        {
            int pos = lines[i].IndexOf(',');
            Coord c = new(
                int.Parse(lines[i][..pos]),
                int.Parse(lines[i][(pos+1)..])
            );
            coords.Add(c);

            if (c.x > maxX)
            {
                maxX = c.x;
            }
        }
    }
    public override ValueTask<string> Solve_1()
    {
        long largest = 0;

        foreach (Coord a in coords)
        {
            foreach (Coord b in coords)
            {
                if (a.Equals(b))
                {
                    continue;
                }

                long area = a.Area(b);
                if (area > largest)
                {
                    largest = area;
                }
            }
        }
        return new($"Largest area is {largest}");
    }

    public void DrawBorderTiles()
    {
        grid.Add((coords[0].x, coords[0].y), Tile.Red);
        for (int i = 0; i < coords.Count; i++)
        {
            Coord cur = coords[i];
            Coord nxt = coords[(i + 1) % coords.Count];
            grid.TryAdd((nxt.x, nxt.y), Tile.Red);

            if (cur.x == nxt.x)
            {
                int yMin = Math.Min(cur.y, nxt.y);
                int yMax = Math.Max(cur.y, nxt.y);
                for (int y = yMin + 1; y < yMax; y++)
                {
                    grid.Add((cur.x, y), Tile.Green);
                }
            } else if (cur.y == nxt.y)
            {
                int xMin = Math.Min(cur.x, nxt.x);
                int xMax = Math.Max(cur.x, nxt.x);
                for (int x = xMin + 1; x < xMax; x++)
                {
                    grid.Add((x, cur.y), Tile.Green);
                }
            } else
            {
                throw new Exception("not a valid sequential point");
            }
        }
    }

    public long FindLargestContainedSquare()
    {
        long largest = 0;

        foreach (Coord a in coords)
        {
            foreach (Coord b in coords)
            {
                if (a.Equals(b))
                {
                    continue;
                }

                if (!CornersContained(a, b))
                {
                    // Console.WriteLine($"Invalid square: {a.x},{a.y} / {b.x},{b.y}");
                    continue;
                }

                bool valid = true;
                foreach (Coord c in coords)
                {
                    if (c.Equals(a) || c.Equals(b))
                    {
                        continue;
                    }

                    if (a.AreaContains(b, c))
                    {
                        valid = false;
                        break;
                    }
                }
                if (!valid)
                {
                    // Console.WriteLine($"Invalid square: {a.x},{a.y} / {b.x},{b.y}");
                    continue;
                }

                long area = a.Area(b);
                Console.WriteLine($"Found square: {a.x},{a.y} / {b.x},{b.y} : {area}");
                if (area > largest)
                {
                    largest = area;
                }
            }
        }
        return largest;
    }

    public bool CornersContained(Coord a, Coord b)
    {
        int xMin = Math.Min(a.x, b.x);
        int xMax = Math.Max(a.x, b.x);
        int yMin = Math.Min(a.y, b.y);
        int yMax = Math.Max(a.y, b.y);

        return PointIsContained(xMin, yMin)
            && PointIsContained(xMax, yMin)
            && PointIsContained(xMin, yMax)
            && PointIsContained(xMax, yMax)
        ;
    }

    public bool PointIsContained(int px, int py)
    {
        // If we are on the line, then quickly return
        Tile tile = grid.GetValueOrDefault((px, py), Tile.Other);
        if (tile != Tile.Other)
        {
            return true;
        }

        int changes = 0;
        bool border;
        bool prevBorder = false;
        int x = 0;
        int y = py;

        while (x < px)
        {
            x++;

            tile = grid.GetValueOrDefault((x, y), Tile.Other);
            border = tile != Tile.Other;
            // Console.WriteLine($"checkY {px},{py} | {x},{y}: {border}");
            if (!prevBorder && border)
            {
                changes++;
            }
            prevBorder = border;
        }

        int pointChanges = changes;

        // Contained == odd number
        if (pointChanges % 2 != 0)
        {
            return true;
        }
return false;
        while (x < maxX)
        {
            x++;

            tile = grid.GetValueOrDefault((x, y), Tile.Other);
            border = tile != Tile.Other;
            // Console.WriteLine($"checkY {px},{py} | {x},{y}: {border}");
            if (!prevBorder && border)
            {
                changes++;
            }
            prevBorder = border;
        }

        return (changes - pointChanges) % 2 != 0;
    }

    public override ValueTask<string> Solve_2()
    {
        DrawBorderTiles();

        PrintGrid();

        long largest = FindLargestContainedSquare();

        Console.WriteLine($"check max {PointIsContained(9,5)} {PointIsContained(2, 3)} {PointIsContained(9, 3)} {PointIsContained(9, 5)}");

        // < 4653414735
        return new($"Largest contained area is {largest}");
    }

    public void PrintGrid()
    {
        for (int y = 0; y < 10; y++)
        {
            StringBuilder chars = new();
            for (int x = 0; x < 20; x++)
            {
                Tile tile = grid.GetValueOrDefault((x, y), Tile.Other);
                switch (tile) {
                    case Tile.Other:
                        chars.Append('.');
                        break;
                    case Tile.Red:
                        chars.Append('X');
                        break;
                    case Tile.Green:
                        chars.Append('O');
                        break;
                }
            }
            Console.WriteLine(chars.ToString());
        }
    }
}
