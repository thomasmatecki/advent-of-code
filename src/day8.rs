use crate::utils::load_input;
use lazy_static::lazy_static;
use regex::Regex;
use std::collections::HashMap;
use std::collections::HashSet;

lazy_static! {
    static ref RE: Regex = Regex::new(r"(\w{3}) ([+|-]\d+)").unwrap();
}

pub fn solution_1() -> i32 {
    let mut executed_ops: HashSet<i32> = HashSet::new();
    let ops: Vec<String> = load_input("input/8.txt");
    let mut op_idx: i32 = 0;
    let mut acc: i32 = 0;

    while !executed_ops.contains(&op_idx) {
        executed_ops.insert(op_idx);
        let op_str = ops.get(op_idx as usize).unwrap();
        let cap = RE.captures(op_str).unwrap();

        match &cap[1] {
            "acc" => {
                // acc increases or decreases a single global value called the
                // accumulator by the value given in the argument. For example,
                // acc +7 would increase the accumulator by 7. The accumulator
                // starts at 0. After an acc instruction, the instruction
                // immediately below it is executed next.
                let inc: i32 = cap[2].parse().unwrap();
                acc += inc;
                op_idx += 1;
            }
            "jmp" => {
                // jmp jumps to a new instruction relative to itself. The next
                // instruction to execute is found using the argument as an
                // offset from the jmp instruction; for example, jmp +2 would
                // skip the next instruction, jmp +1 would continue to the
                // instruction immediately below it, and jmp -20 would cause
                // the instruction 20 lines above to be executed next.
                let jmp: i32 = cap[2].parse().unwrap();
                op_idx = op_idx + jmp;
            }
            "nop" => {
                // nop stands for No OPeration - it does nothing. The
                // instruction immediately below it is executed next.
                op_idx += 1;
            }
            _ => {}
        }
    }

    return acc;
}

enum OpCode {
    ACC,
    JMP,
    NOP,
}

impl From<&str> for OpCode {
    fn from(op_code: &str) -> Self {
        match op_code {
            "acc" => OpCode::ACC,
            "jmp" => OpCode::JMP,
            "nop" => OpCode::NOP,
            _ => unreachable!(),
        }
    }
}

struct Instruction {
    idx: usize,
    opcode: OpCode,
    arg: i32,
    corrected: bool,
}

impl Instruction {
    fn from_str(idx: usize, op_str: &str) -> Self {
        let capture = RE.captures(op_str).unwrap();
        let opcode: OpCode = OpCode::from(capture.get(1).unwrap().as_str());
        Instruction {
            idx,
            opcode,
            arg: 0,
            corrected: false,
        }
    }
    fn next_idx(&self) -> i32 {
        0
    }
}

pub fn solution_2() -> u32 {
    let mut idx_preds: HashMap<i32, HashSet<Instruction>> = HashMap::new();
    let ops: Vec<String> = load_input("input/8.txt");
    for (idx, op) in ops.iter().enumerate() {
        let op_str = ops.get(idx as usize).unwrap();
        let instr = Instruction::from_str(idx, op);
        idx_preds.insert(instr.next_idx(), ...);
    }

    0
}
