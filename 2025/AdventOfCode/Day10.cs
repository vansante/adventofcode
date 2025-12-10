using System.Text;

namespace AdventOfCode;

public class Day10 : BaseDay
{
    public enum LightState : ushort
    {
        Off = 0,
        On = 1,
    }

    public class Machine(Lights lights, List<Button> buttons, List<int> requirements)
    {
        public readonly Lights lights = lights;
        public readonly List<Button> buttons = buttons;
        public readonly List<int> requirements = requirements;
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

            List<int> reqs = [];
            string reqStr = line[(reqStart + 1) .. reqEnd];
            foreach (string req in reqStr.Split(','))
            {
                reqs.Add(int.Parse(req));
            }

            machines.Add(new(new([.. lights]), buttons, reqs));
        }
    }

    public int FindFewestButtonPresses(Machine m)
    {
        Lights outcome = m.lights;
        Lights state = outcome.AllOff();

        return FindFewestButtonPresses(outcome, state, [.. m.buttons]);
    }

    public int FindFewestButtonPresses(
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
            long presses = (long) FindFewestButtonPresses(m);
            Console.WriteLine($"m {m.lights} > {presses}");
            sum += (long) presses;
        }

        return new($"The sum of fewest button presses is {sum}");
    }

    public override ValueTask<string> Solve_2()
    {
        long largest = 0;
        return new($" {largest}");
    }
}
