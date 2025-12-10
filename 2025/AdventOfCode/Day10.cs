using Microsoft.Z3;
using System.Text;

namespace AdventOfCode;

public class Day10 : BaseDay
{
    public enum LightState : ushort
    {
        Off = 0,
        On = 1,
    }

    public class Machine(Lights lights, List<Button> buttons, Counters counters)
    {
        public readonly Lights lights = lights;
        public readonly List<Button> buttons = buttons;
        public readonly Counters counters = counters;
    }

    public class Lights(LightState[] state)
    {
        public readonly LightState[] state = state;

        public bool Equals(Lights other)
        {
            if (state.Length != other.state.Length)
            {
                return false;
            }

            for (int i = 0; i < state.Length; i++)
            {
                if (state[i] != other.state[i])
                {
                    return false;
                }
            }
            return true;
        }

        public override String ToString()
        {
            StringBuilder str = new();
            str.Append('[');
            foreach (LightState s in state)
            {
                str.Append(s == LightState.Off ? '.' : '#');
            }
            str.Append(']');
            return str.ToString();
        }

        public Lights AllOff()
        {
            return new(new LightState[state.Length]);
        }

        public Lights Push(Button btn)
        {
            LightState[] nw = new LightState[state.Length];
            for (int i = 0; i < state.Length; i++)
            {
                nw[i] = state[i];
            }
            foreach (int toggle in btn.toggles)
            {
                nw[toggle] = nw[toggle] == LightState.Off ? LightState.On : LightState.Off;
            }
            return new(nw);
        }
    }

    public class Counters(int[] state)
    {
        public readonly int[] state = state;

        public bool Equals(Counters other)
        {
            if (state.Length != other.state.Length)
            {
                return false;
            }

            for (int i = 0; i < state.Length; i++)
            {
                if (state[i] != other.state[i])
                {
                    return false;
                }
            }
            return true;
        }

        public override String ToString()
        {
            StringBuilder str = new();
            str.Append('{');
            foreach (int s in state)
            {
                str.Append(s);
                str.Append(',');
            }
            str.Append('}');
            return str.ToString();
        }

        public Counters AllOff()
        {
            return new(new int[state.Length]);
        }

        public Counters Push(Button btn)
        {
            int[] nw = new int[state.Length];
            for (int i = 0; i < state.Length; i++)
            {
                nw[i] = state[i];
            }
            foreach (int toggle in btn.toggles)
            {
                nw[toggle]++;
            }
            return new(nw);
        }

        public bool Overcharged(Counters other)
        {
            if (state.Length != other.state.Length)
            {
                throw new Exception("invalid correct state");
            }

            for (int i = 0; i < state.Length; i++)
            {
                if (state[i] < other.state[i])
                {
                    return true;
                }
            }

            return false;
        }
    }

    public class Button(int[] toggles)
    {
        public int[] toggles = toggles;

        public override String ToString()
        {
            StringBuilder str = new();
            str.Append('(');
            foreach (int t in toggles)
            {
                str.Append(t);
                str.Append(',');
            }
            str.Append(')');
            return str.ToString();
        }
    }

    private readonly string _input;

    private readonly List<Machine> machines = [];

    public Day10()
    {
        _input = File.ReadAllText(InputFilePath).Trim();
        string[] lines = _input.Split("\n");
        foreach (string line in lines)
        {
            int lightStart = line.IndexOf('[');
            int lightEnd = line.IndexOf(']');
            int wiringStart = line.IndexOf('(');
            int wiringEnd = line.LastIndexOf(')');
            int reqStart = line.IndexOf('{');
            int reqEnd = line.IndexOf('}');

            List<LightState> lights = [];
            string lightStr = line[(lightStart + 1) .. lightEnd];
            foreach (char ch in lightStr.ToArray())
            {
                switch (ch)
                {
                    case '.':
                        lights.Add(LightState.Off);
                        break;
                    case '#':
                        lights.Add(LightState.On);
                        break;
                    default:
                        throw new Exception("unexpected light state");
                }
            }

            List<Button> buttons = [];
            string wiringStr = line[wiringStart .. (wiringEnd+1)];
            foreach (string wiring in wiringStr.Split(' '))
            {
                List<int> toggles = [];
                foreach (string itm in wiring[1 .. (wiring.Length - 1)].Split(','))
                {
                    toggles.Add(int.Parse(itm));
                }
                buttons.Add(new([.. toggles]));
            }

            List<int> counters = [];
            string reqStr = line[(reqStart + 1) .. reqEnd];
            foreach (string req in reqStr.Split(','))
            {
                counters.Add(int.Parse(req));
            }

            machines.Add(new(new([.. lights]), buttons, new([.. counters])));
        }
    }

    public int FindFewestLightButtonPresses(Machine m)
    {
        Lights outcome = m.lights;
        Lights state = outcome.AllOff();

        return FindFewestLightButtonPresses(outcome, state, [.. m.buttons]);
    }

    public int FindFewestLightButtonPresses(
        Lights outcome,
        Lights state,
        Button[] buttons
    ) {
        Queue<(Lights, Button, int)> queue = [];
        foreach (Button btn in buttons)
        {
            queue.Enqueue((state, btn, 1));
        }

        while(queue.Count > 0)
        {
            (Lights current, Button btn, int presses) = queue.Dequeue();
            Lights newState = current.Push(btn);
            if (newState.Equals(outcome))
            {
                return presses;
            }

            foreach (Button next in buttons)
            {
                queue.Enqueue((newState, next, presses + 1));
            }
        }

        return int.MaxValue;
    }

    public override ValueTask<string> Solve_1()
    {
        long sum = 0;
        foreach (Machine m in machines)
        {
            long presses = (long) FindFewestLightButtonPresses(m);
            // Console.WriteLine($"m {m.lights} > {presses}");
            sum += (long) presses;
        }

        return new($"The sum of fewest button presses is {sum}");
    }

    public long FindFewestJoltageButtonPresses(Machine m)
    {
        // use the boring Z3 solving strategy, I am no math wizard :'(
        using var ctx = new Context();
        using var opt = ctx.MkOptimize();

        var presses = Enumerable.Range(0, m.buttons.Count)
            .Select(i => ctx.MkIntConst($"p{i}"))
            .ToArray()
        ;

        foreach (var press in presses)
        {
            opt.Add(ctx.MkGe(press, ctx.MkInt(0)));
        }

        for (int i = 0; i < m.counters.state.Length; i++)
        {
            var affecting = presses.Where((_, j) => m.buttons[j].toggles.Contains(i)).ToArray();
            if (affecting.Length > 0)
            {
                var sum = affecting.Length == 1 ? affecting[0] : ctx.MkAdd(affecting);
                opt.Add(ctx.MkEq(sum, ctx.MkInt(m.counters.state[i])));
            }
        }

        opt.MkMinimize(presses.Length == 1 ? presses[0] : ctx.MkAdd(presses));
        opt.Check();

        var model = opt.Model;
        return presses.Sum(p => ((IntNum)model.Evaluate(p, true)).Int64);
    }

    public override ValueTask<string> Solve_2()
    {
        long sum = 0;
        foreach (Machine m in machines)
        {
            long presses = (long) FindFewestJoltageButtonPresses(m);
            Console.WriteLine($"m {m.lights} > {presses}");
            sum += (long) presses;
        }

        return new($"The sum of fewest button presses is {sum}");
    }
}
