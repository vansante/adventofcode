namespace AdventOfCode;

public class Day01 : BaseDay
{
    private readonly string _input;

    public Day01()
    {
        _input = File.ReadAllText(InputFilePath);
    }

    public override ValueTask<string> Solve_1() {
        int floor = 0;
        for (int i = 0; i < _input.Length; i++) {
            char c = _input[i];

            if (c == '(') {
                floor++;
            }
            if (c == ')') {
                floor--;
            }
        }

        return new($"The last floor is {floor}");
    }

    public override ValueTask<string> Solve_2() {
        int floor = 0;
        int position = 0;
        for (int i = 0; i < _input.Length; i++) {
            position++;
            char c = _input[i];

            if (c == '(') {
                floor++;
            }
            if (c == ')') {
                floor--;
            }

            if (floor == -1) {
                break;
            }
        }

        return new($"The position is {position}");
    }
}
