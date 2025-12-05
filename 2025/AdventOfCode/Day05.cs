using System.Collections;
using System.ComponentModel;
using System.Numerics;

namespace AdventOfCode;

public class Day05 : BaseDay
{
    class Range
    {
        public long start;
        public long end;

        public bool IsInRange(long id)
        {
            return id >= start && id <= end;
        }

        public bool CanMerge(Range other)
        {
            return other.IsInRange(start) || other.IsInRange(end);
        }

        public Range Merge(Range other)
        {
            if (!CanMerge(other))
            {
                throw new Exception("cannot merge");
            }

            return new Range
            {
                start = Math.Min(start, other.start),
                end = Math.Max(end, other.end),
            };
        }

        public long Count()
        {
            return end - start + 1;
        }
    }

    private readonly string _input;
    private readonly List<Range> idRanges = [];
    private readonly long[] ids;

    public Day05()
    {
        _input = File.ReadAllText(InputFilePath).Trim();

        string[] parts = _input.Split("\n\n");
        string[] rangeStr = parts[0].Split("\n");
        string[] idStr = parts[1].Split("\n");

        idRanges = [];
        for (int i = 0; i < rangeStr.Length; i++)
        {
            int pos = rangeStr[i].IndexOf('-');
            if (pos <= 0)
            {
                throw new Exception("invalid dash position");
            }

            Range rng = new()
            {
                start = long.Parse(rangeStr[i][..pos]),
                end = long.Parse(rangeStr[i][(pos + 1)..]),
            };

            if (rng.start > rng.end)
            {
                throw new Exception("invalid range");
            }
            idRanges.Add(rng);
        }

        ids = new long[idStr.Length];
        for (int i = 0; i < idStr.Length; i++)
        {
            ids[i] = long.Parse(idStr[i]);
        }
    }

    public override ValueTask<string> Solve_1()
    {
        int sum = 0;

        foreach (long id in ids)
        {
            foreach (Range rng in idRanges)
            {
                if (rng.IsInRange(id))
                {
                    sum++;
                    break;
                }
            }
        }

        return new($"Amount of fresh ingredients: {sum}");
    }

    private static List<Range> MergeRanges(List<Range> ranges)
    {
        while (true)
        {
            int merges = 0;

            foreach (Range a in ranges)
            {
                HashSet<Range> nwRanges = [a];
                foreach (Range b in ranges.ToList())
                {
                    if (a == b)
                    {
                        continue;
                    }

                    if (a.CanMerge(b))
                    {
                        Range nw = a.Merge(b);
                        nwRanges.Remove(a);
                        nwRanges.Remove(b);
                        nwRanges.Add(nw);

                        merges++;
                    } else
                    {
                        nwRanges.Add(b);
                    }
                }
                ranges = [.. nwRanges];
            }
            if (merges == 0)
            {
                break;
            }
        }
        return ranges;
    }

    public override ValueTask<string> Solve_2()
    {
        List<Range> ranges = MergeRanges(idRanges);
        
        long sum = 0;
        foreach (Range rng in ranges)
        {
            sum += rng.Count();
        }

        return new($"Amount of fresh ingredients IDs: {sum}");
    }
}
