using System.Collections;
using System.ComponentModel;
using System.Net.NetworkInformation;
using System.Numerics;

namespace AdventOfCode;

public class Day06 : BaseDay
{
    enum Operator
    {
        Add,
        Multiply,
    }

    private static long Operate(Operator op, long a, long b)
    {
        switch (op)
        {
            case Operator.Add:
                return a + b;
            case Operator.Multiply:
                return a * b;
            default:
                throw new Exception("invalid operator");
        }
    }

    class Problem
    {
        public List<long> nums;
        public Operator op;

        public long Solve1()
        {
            long total = nums[0];
            for (int i = 1; i < nums.Count; i++)
            {
                total = Operate(op, total, nums[i]);
            }
            return total;
        }

        public static long GetNthDigit(long num, int n) {
            if (n <= 0)
            {
                return 0;
            }

            long divisor = (long) Math.Pow(10, n - 1);

            // Get the digit: shift right by n-1 places, then get the last digit
            return (num / divisor) % 10;
        }

        public List<long> CephNums()
        {
            List<long> cnums = [];

            for (int digit = 0; digit < 10; digit++)
            {
                long cnum = 0;
                long curDigit = 0;
                for (int i = nums.Count - 1; i >= 0 ; i--)
                {
                    long num = nums[i];
                    long cdigit = GetNthDigit(num, digit);
                    if (cdigit != 0)
                    {
                        curDigit++;
                        cnum += GetNthDigit(num, digit) * (long) Math.Pow(10, curDigit - 1);
                    }
                }

                if (cnum > 0)
                {
                    cnums.Add(cnum);
                    continue;
                }
            }

            return cnums;
        }

        public long Solve2()
        {
            List<long> cnums = CephNums();
            long total = cnums[0];
            for (int i = 1; i < cnums.Count; i++)
            {
                total = Operate(op, total, cnums[i]);
            }
            return total;
        }
    }

    private readonly string _input;

    private readonly List<Problem> problems = [];

    public Day06()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");

        bool first = true;
        for (int i = 0; i < lines.Length; i++)
        {
            string[] itms = lines[i].Split(" ");
            int idx = 0;
            foreach (string it in itms)
            {
                if (it.IsWhiteSpace() || it.Equals(""))
                {
                    continue;
                }

                if (it.Equals("+"))
                {
                    problems[idx].op = Operator.Add;
                    idx++;
                    continue;
                }
                if (it.Equals("*"))
                {
                    problems[idx].op = Operator.Multiply;
                    idx++;
                    continue;
                }

                long num = long.Parse(it);
                if (first)
                {
                    problems.Add(new Problem
                    {
                        nums = [num],
                    });
                    idx++;
                    continue;
                }

                problems[idx].nums.Add(num);
                idx++;
            }
            first = false;
        }

    }

    public override ValueTask<string> Solve_1()
    {
        long sum = 0;
        foreach (Problem p in problems)
        {
            sum  += p.Solve1();
        }

        return new($"The grand total is {sum}");
    }

    public override ValueTask<string> Solve_2()
    {
        long sum = 0;
        foreach (Problem p in problems)
        {
            sum  += p.Solve2();
        }

        return new($"The grand total is {sum}");
    }
}
