from typing import Dict


def climb_stairs(steps: int) -> int:
    if steps < 3:
        return steps

    step_variations = [0] * (steps + 1)
    step_variations[1] = 1
    step_variations[2] = 2

    for i in range(3, steps+1):
        step_variations[i] = step_variations[i-1] + step_variations[i-2]

    return step_variations[steps]
