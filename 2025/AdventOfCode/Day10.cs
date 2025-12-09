namespace AdventOfCode;

public class Day10 : BaseDay
{
    private readonly string _input;

    public Day10()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");
    }

    public override ValueTask<string> Solve_1()
    {
        long largest = 0;

        return new($" {largest}");
    }

    public override ValueTask<string> Solve_2()
    {
        long largest = 0;
        return new($" {largest}");
    }
}
