from calendar import c
from dataclasses import dataclass, field
from doctest import debug


@dataclass
class Computer:
    A: int
    B: int
    C: int
    debug: bool = field(default=False)

    instrx_ptr: int = 0

    output: list[int] = field(default_factory=list)

    def _combo(self, operand):
        if 0 <= operand <= 3:
            return operand

        match operand:
            case 4:
                return self.A
            case 5:
                return self.B
            case 6:
                return self.C
            case _:
                raise ValueError(f"Invalid combo operand: {operand}")

    def run(self, instrx):
        self.instrx_ptr = 0
        while self.instrx_ptr < len(instrx):
            stor_instrx_ptr = self.instrx_ptr
            opcode, operand = instrx[self.instrx_ptr], instrx[self.instrx_ptr + 1]
            op_func = INSTRX[opcode]
            if opcode in [1, 3, 4]:
                _operand = operand
            else:
                _operand = self._combo(operand)
            op_func(self, _operand)
            if self.instrx_ptr == stor_instrx_ptr:
                self.instrx_ptr += 2

            reg_vals = f"--- {self.A} {self.B} {self.C}"
            if self.debug:
                print(
                    f"{stor_instrx_ptr:03}: {opcode} [{op_func.__name__}] ({operand}:{_operand})\t-> {self.instrx_ptr:03} {reg_vals}"
                )

    def flush_output(self):
        return ",".join(map(str, self.output))


def adv(cpu, operand):
    """
    The adv instruction (opcode 0) performs division. The numerator is the
    value in the A register. The denominator is found by raising 2 to the
    power of the instruction's combo operand. (So, an operand of 2 would
    divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result
    of the division operation is truncated to an integer and then written to
    the A register.
    """
    cpu.A = cpu.A // (2**operand)


def bxl(cpu, operand):
    """
    The bxl instruction (opcode 1) calculates the bitwise XOR of register B
    and the instruction's literal operand, then stores the result in
    register B.
    """
    cpu.B = cpu.B ^ operand


def bst(cpu, operand):
    """
    The bst instruction (opcode 2) calculates the value of its combo operand
    modulo 8 (thereby keeping only its lowest 3 bits), then writes that
    value to the B register.
    """
    cpu.B = operand % 8


def jnz(cpu, operand):
    """
    The jnz instruction (opcode 3) does nothing if the A register is 0.
    However, if the A register is not zero, it jumps by setting the
    instruction pointer to the value of its literal operand. If this
    instruction jumps, the instruction pointer is not increased by 2 after
    this instruction.
    """
    if cpu.A == 0:
        pass
    else:
        cpu.instrx_ptr = operand


def bxc(cpu, operand):
    """
    The bxc instruction (opcode 4) calculates the bitwise XOR of register B
    and register C, then stores the result in register B. (For legacy
    reasons, this instruction reads an operand but ignores it.)
    """
    cpu.B = cpu.B ^ cpu.C


def out(cpu, operand):
    """
    The out instruction (opcode 5) calculates the value of its combo operand
    modulo 8, then outputs that value. (If a program outputs multiple
    values, they are separated by commas.)
    """
    val = operand % 8
    cpu.output.append(val)


def bdv(cpu, operand):
    """
    The bdv instruction (opcode 6) works exactly like the adv instruction
    except that the result is stored in the B register. (The numerator is
    still read from the A register.)
    """

    cpu.B = cpu.A // (2**operand)


def cdv(cpu, operand):
    """
    The cdv instruction (opcode 7) works exactly like the adv instruction
    except that the result is stored in the C register. (The numerator is
    still read from the A register.)
    """

    cpu.C = cpu.A // (2**operand)


INSTRX = [adv, bxl, bst, jnz, bxc, out, bdv, cdv]


def load(filename):

    with open(f"2024/day17/{filename}") as f:
        lines = f.readlines()

    a_reg, b_reg, c_reg, _, instrx = lines
    regs = []
    for reg_str in [a_reg, b_reg, c_reg]:
        regs.append(int(reg_str.split(":")[1].strip()))

    _, instrx = instrx.split(":")

    instrx = [int(x) for x in instrx.split(",")]

    return regs, instrx


def test():
    #    # If register C contains 9, the program 2,6 would set register B to 1.
    c = Computer(-1, -1, 9)
    c.run([2, 6])
    assert c.B == 1

    # If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.
    c = Computer(10, -1, -1)
    c.run([5, 0, 5, 1, 5, 4])
    assert c.output == [0, 1, 2]

    # If register A contains 2024, the program 0,1,5,4,3,0 would output
    # 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.
    c = Computer(2024, -1, -1)
    c.run([0, 1, 5, 4, 3, 0])
    assert c.output == [4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0]
    assert c.A == 0

    # If register B contains 29, the program 1,7 would set register B to 26.
    c = Computer(-1, 29, -1)
    c.run([1, 7])
    assert c.B == 26

    # If register B contains 2024 and register C contains 43690, the program 4,0
    # would set register B to 44354.
    c = Computer(-1, 2024, 43690)
    c.run([4, 0])
    assert c.B == 44354


def part_one(filename):
    regs, instrx = load(filename)

    c = Computer(*regs)

    c.run(instrx)

    # 3,5,0,1,5,1,5,1,0
    print(f"part one ({filename}): {c.flush_output()}")


def part_two(filename):
    """
    000: 2 [bst] (4:22118)  -> 002 --- 22118 6 44237    ;; A % 8    -> B
    002: 1 [bxl] (5:5)      -> 004 --- 22118 3 44237    ;; B ^ 5    -> B
    004: 7 [cdv] (5:3)      -> 006 --- 22118 3 2764     ;; A // 2^B -> C
    006: 1 [bxl] (6:6)      -> 008 --- 22118 5 2764     ;; B ^ 6    -> B
    008: 4 [bxc] (1:1)      -> 010 --- 22118 2761 2764  ;; B ^ C    -> B
    010: 5 [out] (5:2761)   -> 012 --- 22118 2761 2764  ;; B % 8    -> OUT
    012: 0 [adv] (3:3)      -> 014 --- 2764 2761 2764   ;; A // 2^3 -> A
    014: 3 [jnz] (0:0)      -> 000 --- 2764 2761 2764
    """
    regs, instrx = load(filename)

    def step(A, v_):
        B = (A % 8) ^ 5
        C = A // (2**B)
        B = B ^ 6
        v = (B ^ C) % 8
        return v == v_

    A_inps = list(range(1, 8))

    for v in reversed([2, 4, 1, 5, 7, 5, 1, 6, 4, 1, 5, 5, 0, 3, 3, 0]):
        untruncted = []
        A_seeds = []
        for A0 in A_inps:
            if step(A0, v):
                A_seeds.append(A0)
                untruncted.extend(A0 * 8 + t for t in range(8))

        A_inps = untruncted

    min_a_seed = min(A_seeds)
    c = Computer(*regs)
    c.A = min_a_seed
    c.run(instrx)
    assert c.output == instrx
    print(f"part two ({filename}): {min_a_seed}")


if __name__ == "__main__":
    test()
    part_one("test1.txt")  # 4,6,3,5,6,3,5,2,1,0
    part_one("input.txt")
    part_one("test2.txt")

    part_two("input.txt")
