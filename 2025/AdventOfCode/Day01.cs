using System.Collections;
using System.ComponentModel;
using System.Numerics;

namespace AdventOfCode;

public class Day01 : BaseDay
{
    public const int DialSize = 100;
    public const int DialStart = 50;

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

            instr.clicks = Int32.Parse(line[1..]);

            instructions[i] = instr;
        }
    }

    public override ValueTask<string> Solve_1()
    {
        int zeroCount = 0;
        int dial = 50;
        for (int i = 0; i < instructions.Length; i++) {
            Input instr = instructions[i];

            dial += DialSize;
            if (instr.dir == Direction.Left)
            {
                dial -= instr.clicks;
            } else
            {
                dial += instr.clicks;
            }
            dial %= DialSize;

            if (dial == 0)
            {
                zeroCount++;
            }
        }

        return new($"Zero was reached {zeroCount} times");
    }

    public static int DialModulo(int dial)
    {
        return ((dial % DialSize) + DialSize) % DialSize;
    }

    public override ValueTask<string> Solve_2()
    {
        int zeroCount = 0;
        int dial = DialStart;
        for (int i = 0; i < instructions.Length; i++) {
            Input instr = instructions[i];

            // Add the amount of full cycles (will definitely hit zero)
            zeroCount += instr.clicks / DialSize;

            if (instr.dir == Direction.Left)
            {
                // If we are not currently on zero, and we will pass zero in the current pass remainder, add zero
                if (dial != 0 && dial - DialModulo(instr.clicks) < 0)
                {
                    zeroCount++;
                }
                // Add the full amount of clicks
                dial -= instr.clicks;
            } else
            {
                // If we will pass zero in the current pass remainder, add zero
                if (dial + DialModulo(instr.clicks) > 100)
                {
                    zeroCount++;
                }
                // Add the full amount of clicks
                dial += instr.clicks;
            }

            // Set the dial back to 0-99
            dial = DialModulo(dial);

            // If we end on a zero, add it:
            if (dial == 0)
            {
                zeroCount++;
            }
        }


        // < 7426
        // > 6350
        return new($"Zero was hit {zeroCount} times");
    }
}
