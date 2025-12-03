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

    public static int FindHighestJoltage(int[] bank)
    {
        int highestFirst = 0;
        int idxUsed = 0;
        for (int i = 0; i < bank.Length - 1; i++)
        {
            if (highestFirst < bank[i])
            {
                idxUsed = i;
                highestFirst = bank[i];
            }
        }

        int highestSecond = 0;
        for (int i = idxUsed + 1; i < bank.Length; i++)
        {
            if (highestSecond < bank[i])
            {
                highestSecond = bank[i];
            }
        }
        return highestFirst * 10 + highestSecond;
    }

    public override ValueTask<string> Solve_1()
    {
        int joltageSum = 0;
        for (int i = 0; i < banks.Length; i++)
        {
            joltageSum += FindHighestJoltage(banks[i]);
        }

        return new($"Highest joltage sum is {joltageSum}");
    }

    public override ValueTask<string> Solve_2()
    {
        return new($"Invalid ID sum is ");
    }
}
