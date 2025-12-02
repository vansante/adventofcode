using System.Collections;
using System.ComponentModel;
using System.Numerics;

namespace AdventOfCode;

public class Day02 : BaseDay
{
    private readonly string _input;

    private readonly Range[] ranges;

    class Range
    {
        public long start;
        public long end;
    }

    public Day02()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] rngs = _input.Split(",");

        ranges = new Range[rngs.Length];

        for (int i = 0; i < rngs.Length; i++)
        {
            int pos = rngs[i].IndexOf("-");
            if (pos <= 0)
            {
                throw new Exception("invalid dash position");
            }

            ranges[i] = new Range
            {
                start = long.Parse(rngs[i][..pos]),
                end = long.Parse(rngs[i][(pos + 1)..]),
            };

            if (ranges[i].start > ranges[i].end)
            {
                throw new Exception("invalid range");
            }
        }
    }

    public static bool IsValidID(long id)
    {
        int digits = 0;
        for (long t = id; 0 < t; t /= 10)
        {
            digits++;
        }
        // Check if we have an even amount of digits
        if (digits % 2 == 1)
        {
            return false;
        }

        int halfDigits = digits / 2;
        long a = id / ((long) Math.Pow(10, halfDigits));
        long b = id - (a * (long) Math.Pow(10, halfDigits));

        return a == b;
    }

    public override ValueTask<string> Solve_1()
    {
        long invalidSum = 0;

        for (int i = 0; i < ranges.Length; i++)
        {
            Range r = ranges[i];
            for (long j = r.start; j <= r.end; j++)
            {
                if (!IsValidID(j))
                {
                    continue;
                }

                invalidSum += j;
                if (invalidSum < 0)
                {
                    throw new Exception("negative");
                }
                Console.WriteLine($"Invalid ID: {j}");
            }
        }

        return new($"Invalid ID sum is {invalidSum}");
    }

    public override ValueTask<string> Solve_2()
    {
        return new($"");
    }
}
