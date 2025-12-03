using System.Collections;
using System.ComponentModel;
using System.Numerics;

namespace AdventOfCode;

public class Day03 : BaseDay
{
    private readonly string _input;

    private readonly int[][] banks;

    public Day03()
    {
        _input = File.ReadAllText(InputFilePath).Trim();

        string[] bnks = _input.Split("\n");
        banks = new int[bnks.Length][];
        for (int i = 0; i < bnks.Length; i++)
        {
            string line = bnks[i];
            banks[i] = new int[line.Length];
            for (int j = 0; j < line.Length; j++)
            {
                banks[i][j] = Int32.Parse(line[j..(j+1)]);
            }
        }
    }

    public static long FindHighestJoltage(int[] bank, int batteries)
    {
        long sum = 0;
        int curIdx = -1;

        for (int b = batteries; b > 0; b--)
        {
            int curHigh = 0;
            for (int i = curIdx + 1; i < bank.Length - b + 1; i++)
            {
                if (curHigh < bank[i])
                {
                    curIdx = i;
                    curHigh = bank[i];
                }
            }
            sum += (long) Math.Pow(10, b - 1) * curHigh;
        }

        return sum;
    }

    public override ValueTask<string> Solve_1()
    {
        long joltageSum = 0;
        for (int i = 0; i < banks.Length; i++)
        {
            joltageSum += FindHighestJoltage(banks[i], 2);
        }

        return new($"Highest joltage sum is {joltageSum}");
    }

    public override ValueTask<string> Solve_2()
    {
        long joltageSum = 0;
        for (int i = 0; i < banks.Length; i++)
        {
            joltageSum += FindHighestJoltage(banks[i], 12);
        }

        return new($"Highest joltage sum is {joltageSum}");
    }
}
