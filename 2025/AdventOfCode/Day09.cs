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

        public Coord North()
        {
            return new(x, y - 1);
        }
        public Coord East()
        {
            return new (x + 1, y);
        }
        public Coord South()
        {
            return new(x, y + 1);
        }
        public Coord West()
        {
            return new(x - 1, y);
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

    private readonly int maxX = 0;

    private readonly int maxY = 0;

    private readonly Dictionary<int, int> xMap = [];
    private readonly Dictionary<int, int> yMap = [];

    private readonly Tile[][] tiles;

    public Day09()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        HashSet<int> xCoords = [];
        HashSet<int> yCoords = [];
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
            if (c.y > maxY)
            {
                maxY = c.y;
            }

            xCoords.Add(c.x);
            xCoords.Add(c.x + 1);
            yCoords.Add(c.y);
            yCoords.Add(c.y + 1);
        }

        // Add some extra coordinates to keep an outer layer for our floodfill
        List<int> sortedX = [.. xCoords, 0, maxX + 2];
        sortedX.Sort();
        List<int> sortedY = [.. yCoords, 0, maxY + 2];
        sortedY.Sort();

        // Create mappings for our compressed coordinate system
        for (int i = 0; i < sortedX.Count; i++)
        {
            xMap[sortedX[i]] = i;
        }

        for (int i = 0; i < sortedY.Count; i++)
        {
            yMap[sortedY[i]] = i;
        }

        tiles = new Tile[yMap.Count][];
        for (int i = 0; i < yMap.Count; i++)
        {
            tiles[i] = new Tile[xMap.Count];
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

    public Coord GetMappedCoords(Coord c)
    {
        return new(
            xMap[c.x],
            yMap[c.y]
        );
    }

    public void DrawBorderTiles()
    {
        Coord cur = GetMappedCoords(coords[0]);

        tiles[cur.y][cur.x] = Tile.Red;
        for (int i = 0; i < coords.Count; i++)
        {
            cur = coords[i];
            Coord nxt = coords[(i + 1) % coords.Count];

            cur = GetMappedCoords(cur);
            nxt = GetMappedCoords(nxt);
            tiles[nxt.y][nxt.x] = Tile.Red;

            if (cur.x == nxt.x)
            {
                int yMin = Math.Min(cur.y, nxt.y);
                int yMax = Math.Max(cur.y, nxt.y);
                for (int y = yMin + 1; y < yMax; y++)
                {
                    tiles[y][cur.x] = Tile.Green;
                }
            } else if (cur.y == nxt.y)
            {
                int xMin = Math.Min(cur.x, nxt.x);
                int xMax = Math.Max(cur.x, nxt.x);
                for (int x = xMin + 1; x < xMax; x++)
                {
                    tiles[cur.y][x] = Tile.Green;
                }
            } else
            {
                throw new Exception("not a valid sequential point");
            }
        }
    }

    public void FloodFill()
    {
        Stack<Coord> stack = [];
        stack.Push(new(0, 0));

        while (stack.Count > 0)
        {
            Coord c = stack.Pop();
            if (tiles[c.y][c.x] != Tile.Unknown)
            {
                continue;
            }

            List<Coord> ngb = [c.West(), c.East(), c.North(), c.South()];
            foreach (Coord n in ngb)
            {
                if (n.y < 0 || n.y >= tiles.Length)
                {
                    continue;
                }
                if (n.x < 0 || n.x >= tiles[n.y].Length)
                {
                    continue;
                }
                if (tiles[n.y][n.x] == Tile.Unknown)
                {
                    stack.Push(n);
                }
            }

            tiles[c.y][c.x] = Tile.Other;
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

                long area = a.Area(b);
                if (area < largest)
                {
                    continue;
                }

                Coord mapA = GetMappedCoords(a);
                Coord mapB = GetMappedCoords(b);
                if (!RectangleContained(mapA, mapB))
                {
                    continue;
                }

                largest = area;
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

        for (int x = xMin; x <= xMax; x++)
        {
            if (!PointIsContained(x, yMin) || !PointIsContained(x, yMax))
            {
                return false;
            }
        }

        for (int y = yMin; y <= yMax; y++)
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
        return tiles[py][px] != Tile.Other;
    }

    public override ValueTask<string> Solve_2()
    {
        DrawBorderTiles();

        FloodFill();

        // PrintGrid();

        long largest = FindLargestContainedSquare();

        // < 4653414735
        // < 4615010043
        return new($"Largest contained area is {largest}");
    }

    public void PrintGrid()
    {
        for (int y = 0; y < tiles.Length; y++)
        {
            StringBuilder chars = new();
            for (int x = 0; x < tiles[y].Length; x++)
            {
                switch (tiles[y][x]) {
                    case Tile.Unknown:
                        chars.Append('?');
                        break;
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
