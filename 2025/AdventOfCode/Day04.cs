namespace AdventOfCode;

public class Day04 : BaseDay
{
    public const int Empty = 0;
    public const int Roll = 1;

    public const int MaxRolls = 3;

    public int[][] directions = [
        [-1, -1],
        [-1, 0],
        [-1, +1],
        [0, -1],
        [0, 1],
        [1, -1],
        [1, 0],
        [1, 1],
    ];

    private readonly string _input;

    public Day04()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
    }

    public int[][] GetGrid()
    {
        string[] lines = _input.Split("\n");
        int[][] grid = new int[lines.Length][];
        for (int y = 0; y < lines.Length; y++)
        {
            string line = lines[y];

            grid[y] = new int[line.Length];
            for (int x = 0; x < line.Length; x++)
            {
                grid[y][x] = Empty;
                if (line[x..(x+1)] == "@")
                {
                    grid[y][x] = Roll;
                }
            }
        }
        return grid;
    }

    public int GridValue(int[][] grid, int y, int x, int defValue = -1)
    {
        if (y < 0 || y >= grid.Length)
        {
            return defValue;
        }
        if (x < 0 || x >= grid[y].Length)
        {
            return defValue;
        }
        return grid[y][x];
    }

    public int AccessibleRolls(int[][] grid, bool remove = false)
    {
        int sum = 0;
        for (int y = 0; y < grid.Length; y++)
        {
            for (int x = 0; x < grid[y].Length; x++)
            {
                if (GridValue(grid, y, x) == Empty)
                {
                    continue;
                }

                int rolls = 0;
                foreach (int[] dir in directions)
                {
                    rolls += GridValue(grid, y + dir[0], x + dir[1], Empty);
                }

                if (rolls <= MaxRolls)
                {
                    sum++;
                    if (remove) {
                        grid[y][x] = Empty;
                    }
                }
            }
        }
        return sum;
    }

    public override ValueTask<string> Solve_1()
    {
        int[][] grid = GetGrid();

        int forkliftSum = AccessibleRolls(grid);

        return new($"Accessable by forklift: {forkliftSum}");
    }

    public override ValueTask<string> Solve_2()
    {
        int[][] grid = GetGrid();
        int forkliftSum = AccessibleRolls(grid, true);
        while (true)
        {
            int removed = AccessibleRolls(grid, true);
            if (removed == 0)
            {
                break;
            }
            forkliftSum += removed;
        }

        return new($"Accessable by forklift: {forkliftSum}");
    }
}
