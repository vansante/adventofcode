using System.Collections;
using System.ComponentModel;
using System.Net.NetworkInformation;
using System.Numerics;

namespace AdventOfCode;

public class Day06 : BaseDay
{
    public enum Operator
    {
        Add,
        Multiply,
    }

    private static Operator FromChar(char c)
    {
        return c switch {
            '+' => Operator.Add,
            '*' => Operator.Multiply,
            _ => throw new Exception("not an operator"),
        };
    }

    private static long Operate(Operator op, long a, long b)
    {
        return op switch
        {
            Operator.Add => a + b,
            Operator.Multiply => a * b,
            _ => throw new Exception("invalid operator"),
        };
    }

    public class Problem
    {
        public List<long> nums;
        public Operator op;

        public long Solve()
        {
            long total = nums[0];
            for (int i = 1; i < nums.Count; i++)
            {
                total = Operate(op, total, nums[i]);
            }
            return total;
        }
    }

    private readonly string _input;

    public Day06()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
    }

    public List<Problem> GetProblems1() 
    {
        List<Problem> problems = [];
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
        return problems;
    }

    public List<Problem> GetProblems2() 
    {
        List<Problem> problems = [];
        string[] lines = _input.Split("\n");

        List<int> opIdx = [];

        // First find the indexes of all operators:
        int opLineIdx = lines.Length - 1;
        string opLine = lines[opLineIdx];

        for (int i = 0; i < opLine.Length; i++)
        {
            char c = opLine[i];

            if (c.Equals('+') || c.Equals('*'))
            {
                opIdx.Add(i);
                problems.Add(new Problem
                {
                    op = FromChar(c),
                    nums = [],
                });
            }
        }

        // Loop over all problems (denoted by the operator in the bottom line)
        for (int op = 0; op < opIdx.Count; op++)
        {
            int curOpIdx = opIdx[op];
            // Guard for the last problem, default to the lines max length
            int nextOpIdx = op < opIdx.Count - 1 ? opIdx[op + 1] : lines[0].Length;

            // Loop over all columns in the current problem
            for (int col = curOpIdx; col < nextOpIdx; col++)
            {
                List<char> chars = [];
                // Loop over the characters in the current column
                for (int l = 0; l < opLineIdx; l++)
                {
                    if (lines[l][col].Equals(' '))
                    {
                        continue;
                    }
                    chars.Add(lines[l][col]);
                }
                
                // Skip lines with just whitespace
                if (chars.Count == 0)
                {
                    continue;
                }

                // Add the found number to the current problem
                problems[op].nums.Add(
                    long.Parse(
                        new string([.. chars])
                    )
                );
            }
        }
        return problems;
    }

    public override ValueTask<string> Solve_1()
    {
        long sum = 0;
        foreach (Problem p in GetProblems1())
        {
            sum  += p.Solve();
        }

        return new($"The grand total is {sum}");
    }

    public override ValueTask<string> Solve_2()
    {
        long sum = 0;
        foreach (Problem p in GetProblems2())
        {
            sum  += p.Solve();
        }

        return new($"The grand total is {sum}");
    }
}
