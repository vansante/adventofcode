using System.Collections;
using System.ComponentModel;
using System.Numerics;

namespace AdventOfCode;

public class Day01 : BaseDay
{
    private readonly string _input;
    private readonly Input[] instructions;

    enum Direction
    {
        Left,
        Right
    }

    class Input
    {
        public Direction dir;
        public int clicks;
    }

    public Day01()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        instructions = new Input[lines.Length];
        for (int i = 0; i < lines.Length; i++)
        {
            string line = lines[i];
            Input instr = new Input
            {
                dir = Direction.Left,
                clicks = 0
            };

            if (line[0] == 'R')
            {
                instr.dir = Direction.Right;
            }

            try
            {
                instr.clicks = Int32.Parse(line[1..]);
            }
            catch (FormatException)
            {
                Console.WriteLine($"Unable to parse integer '{line[1..]}'");
                continue;
            }

            instructions[i] = instr;
        }
    }

    public override ValueTask<string> Solve_1() {
        // Console.WriteLine($"{instructions[2].dir} {instructions[2].clicks}");

        int zeroCount = 0;
        int dial = 50;
        for (int i = 0; i < instructions.Length; i++) {
            Input instr = instructions[i];

            dial += 100;
            if (instr.dir == Direction.Left)
            {
                dial -= instr.clicks;
            } else
            {
                dial += instr.clicks;
            }
            dial %= 100;

            if (dial == 0)
            {
                zeroCount++;
            }

            Console.WriteLine($"{i} dial {dial} zeroes: {zeroCount}");
        }

        return new($"Zero was reached {zeroCount} times");
    }

    public override ValueTask<string> Solve_2() {
        return new($"The answer is ");
    }
}
