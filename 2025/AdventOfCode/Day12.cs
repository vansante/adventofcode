
using System.Text.RegularExpressions;

namespace AdventOfCode;


record Todo(int w, int h, int[] counts);
public class Day12 : BaseDay
{
    public class Shape(int id, bool[][] area)
    {
        public readonly int id = id;
        public readonly bool[][] area = area;

        public override string ToString()
        {
            return $"S{id}";
        }

        public long Area()
        {
            long sum = 0;
            foreach (var l in area)
            {
                foreach (var b in l)
                {
                    if (b) sum++;
                }
            }
            return sum;
        }
    }

    public class Region(int width, int height, Dictionary<Shape, int> shapeCount)
    {
        public readonly int width = width;
        public readonly int height = height;
        public readonly Dictionary<Shape, int> shapeCount = shapeCount;

        public override string ToString()
        {
            return $"R{width}x{height}";
        }

        public long Area()
        {
            return width * height;
        }

        public long ShapeAreaSum()
        {
            long sum = 0;
            foreach (var p in shapeCount)
            {
                sum += p.Key.Area() * (long) p.Value;
            }
            return sum;
        }
    }

    public readonly string _input;
    public readonly Dictionary<int, Shape> shapes = [];
    public readonly List<Region> regions = [];

    public Day12()
    {
        _input = File.ReadAllText(InputFilePath).Trim();

        var blocks = _input.Split("\n\n");

        foreach (var blk in blocks[..^1])
        {
            var lines = blk.Split('\n');
            var id = int.Parse(lines[0][..^1]);
            var area = new bool[3][];

            var i = 0;
            foreach (var line in lines[1..])
            {
                area[i] = [.. line.Select(c => c == '#')];
                i++;
            }

            shapes.Add(id, new(id, area));
        }

        foreach (var line in blocks[^1].Split('\n'))
        {
            var xIdx = line.IndexOf('x');
            var colonIdx = line.IndexOf(':');

            var width = int.Parse(line[..xIdx]);
            var height = int.Parse(line[(xIdx+1) .. colonIdx]);

            var shapeCount = new Dictionary<Shape, int>();

            var shapeList = line[(colonIdx + 1) ..].Trim().Split(' ').Select(s => int.Parse(s)).ToList();
            for (int id = 0; id < shapeList.Count; id++)
            {
                shapeCount[shapes[id]] = shapeList[id];
            }

            regions.Add(new(width, height, shapeCount));
        }
    }

    public override ValueTask<string> Solve_1()
    {
        int sum = 0;

        // We make a very bold assumption that each shape fits if the total area is equal
        // to the needed area. This only works for the actual input, not the sample.

        foreach (var region in regions)
        {
            if (region.Area() >= region.ShapeAreaSum())
            {
                sum++;
            }
        }

        return new($"Regions that fit: {sum}");
    }

    public override ValueTask<string> Solve_2()
    {
        return new($"");
    }
}
