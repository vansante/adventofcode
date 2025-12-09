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
        Unknown,
        Other,
        Red,
        Green,
    }

    private readonly string _input;

    private readonly List<Coord> coords = [];

    private readonly Dictionary<(int, int), Tile> grid = [];

    private readonly int minX = int.MaxValue;
    private readonly int maxX = 0;

    private readonly int minY = int.MaxValue;
    private readonly int maxY = 0;

    private readonly Dictionary<(int, int), bool?> containCache = [];

    private readonly Tile[][] tiles;

    public Day09()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        tiles = new Tile[100_000][];
        for (int i = 0; i < tiles.Length; i++)
        {
            tiles[i] = new Tile[100_000];
        }

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
            if (c.x < minX)
            {
                minX = c.x;
            }
            if (c.y > maxY)
            {
                maxY = c.y;
            }
            if (c.y < minY)
            {
                minY = c.y;
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
        tiles[coords[0].x][coords[0].y] = Tile.Red;
        grid.Add((coords[0].x, coords[0].y), Tile.Red);
        for (int i = 0; i < coords.Count; i++)
        {
            Coord cur = coords[i];
            Coord nxt = coords[(i + 1) % coords.Count];
            grid.TryAdd((nxt.x, nxt.y), Tile.Red);
            tiles[coords[0].x][coords[0].y] = Tile.Red;

            if (cur.x == nxt.x)
            {
                int yMin = Math.Min(cur.y, nxt.y);
                int yMax = Math.Max(cur.y, nxt.y);
                for (int y = yMin + 1; y < yMax; y++)
                {
                    grid.Add((cur.x, y), Tile.Green);
                    tiles[cur.x][y] = Tile.Green;
                }
            } else if (cur.y == nxt.y)
            {
                int xMin = Math.Min(cur.x, nxt.x);
                int xMax = Math.Max(cur.x, nxt.x);
                for (int x = xMin + 1; x < xMax; x++)
                {
                    grid.Add((x, cur.y), Tile.Green);
                    tiles[x][cur.y] = Tile.Green;
                }
            } else
            {
                throw new Exception("not a valid sequential point");
            }
        }
    }

    public void FloodFill()
    {
        Stack<(int, int)> stack = [];
        stack.Push((minX - 1, minY - 1));

        while (stack.Count > 0)
        {
            (int, int) c = stack.Pop();

            if (c.Item1 < minX - 1 || c.Item1 > maxX + 1)
            {
                continue;
            }

            if (c.Item2 < minY - 1 || c.Item2 > maxY + 1)
            {
                continue;
            }

            if (tiles[c.Item1][c.Item2] != Tile.Unknown)
            {
                continue;
            }

            if (tiles[c.Item1 - 1][c.Item2] == Tile.Unknown)
            {
                stack.Push((c.Item1 - 1, c.Item2)); // west
            }
            if (tiles[c.Item1 + 1][c.Item2] == Tile.Unknown)
            {
                stack.Push((c.Item1 + 1, c.Item2)); // east
            }
            if (tiles[c.Item1][c.Item2 - 1] == Tile.Unknown)
            {
                stack.Push((c.Item1, c.Item2 - 1)); // north
            }
            if (tiles[c.Item1][c.Item2 + 1] == Tile.Unknown)
            {
                stack.Push((c.Item1, c.Item2 + 1)); // south
            }

            tiles[c.Item1][c.Item2] = Tile.Other;
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

                if (!RectangleContained(a, b))
                {
                    // Console.WriteLine($"Invalid square: {a.x},{a.y} / {b.x},{b.y}");
                    continue;
                }

                long area = a.Area(b);
                // Console.WriteLine($"Found square: {a.x},{a.y} / {b.x},{b.y} : {area}");
                if (area > largest)
                {
                    largest = area;
                }
            }
        }
        return largest;
    }

    public bool RectangleContained(Coord a, Coord b)
    {
        int xMin = Math.Min(a.x, b.x);
        int xMax = Math.Max(a.x, b.x);
        int yMin = Math.Min(a.y, b.y);
        int yMax = Math.Max(a.y, b.y);

        bool cornersContained = PointIsContained(xMin, yMin)
            && PointIsContained(xMax, yMin)
            && PointIsContained(xMin, yMax)
            && PointIsContained(xMax, yMax)
        ;

        if (!cornersContained)
        {
            return false;
        }

        int xLen = Math.Abs(xMax - xMin);
        int yLen = Math.Abs(yMax - yMin);

        for (int x = xMin; x < xMax; x++)
        {
            if (!PointIsContained(x, yMin) || !PointIsContained(x, yMax))
            {
                return false;
            }
        }

        for (int y = yMin; y < yMax; y++)
        {
            if (!PointIsContained(xMin, y) || !PointIsContained(xMax, y))
            {
                return false;
            }
        }
        return true;
    }
    public bool PointIsContained(int px, int py)
    {
        bool? cache = containCache.GetValueOrDefault((px, py), null);
        if (cache != null)
        {
            return (bool) cache;
        }

        // If we are on the line, then quickly return
        Tile tile = grid.GetValueOrDefault((px, py), Tile.Other);
        if (tile != Tile.Other)
        {
            containCache.Add((px, py), true);
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
            // Console.WriteLine($"{px},{py} : Contained");
            containCache.Add((px, py), true);
            return true;
        }

        // Console.WriteLine($"{px},{py} : Not contained");
        containCache.Add((px, py), false);
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

        FloodFill();

        PrintGrid();

        long largest = 0;
        // long largest = FindLargestContainedSquare();

        // Console.WriteLine($"check max {PointIsContained(9,5)} {PointIsContained(2, 3)} {PointIsContained(9, 3)} {PointIsContained(9, 5)}");

        // < 4653414735
        // < 4615010043
        return new($"Largest contained area is {largest}");
    }

    public void PrintGrid()
    {
        for (int y = -2; y < 13; y++)
        {
            StringBuilder chars = new();
            for (int x = -2; x < 24; x++)
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
