use crate::utils::load_input;
use lazy_static::lazy_static;
use regex::Regex;
use std::collections::HashMap;
use std::collections::HashSet;
use std::collections::VecDeque;
use std::fmt;
use std::hash::Hash;
use std::rc::Rc;
use std::str::FromStr;

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
#[derive(Eq, PartialEq, Hash, Debug)]
enum OpCode {
    ACC,
    JMP,
    NOP,
}

impl fmt::Display for OpCode {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}", self)
    }
}

impl FromStr for OpCode {
    type Err = ();
    fn from_str(op_str: &str) -> Result<Self, <Self as FromStr>::Err> {
        match op_str {
            "acc" => return Ok(OpCode::ACC),
            "jmp" => Ok(OpCode::JMP),
            "nop" => Ok(OpCode::NOP),
            _ => return Err(()),
        }
    }
}

#[derive(Eq, PartialEq, Hash, Debug)]
struct Instruction {
    idx: usize,
    opcode: OpCode,
    arg: i32,
    corrected: bool,
}

impl Instruction {
    fn from_str(idx: usize, op_str: &str) -> Self {
        let capture = RE.captures(op_str).unwrap();
        let op_str = capture.get(1).unwrap().as_str();
        let opcode: OpCode = op_str.parse().unwrap();
        let arg: i32 = capture[2].parse().unwrap();
        Instruction {
            idx,
            opcode,
            arg,
            corrected: false,
        }
    }
    fn next_idx(&self) -> usize {
        match self.opcode {
            OpCode::ACC => self.idx + 1,
            OpCode::JMP => (self.idx as i32 + self.arg) as usize,
            OpCode::NOP => self.idx + 1,
        }
    }

    fn exec(&self, op_idx: &mut usize, acc: &mut i32) {
        match self.opcode {
            OpCode::ACC => {
                *op_idx += 1;
                *acc += self.arg;
            }
            OpCode::JMP => {
                *op_idx = (*op_idx as i32 + self.arg) as usize;
            }
            OpCode::NOP => {
                *op_idx += 1;
            }
        };
    }

    fn corrected(&self) -> Option<Instruction> {
        if self.arg == 1 || self.opcode == OpCode::ACC {
            None
        } else {
            let opcode = match self.opcode {
                OpCode::NOP => OpCode::JMP,
                OpCode::JMP => OpCode::NOP,
                OpCode::ACC => OpCode::ACC,
            };

            Some(Instruction {
                idx: self.idx,
                opcode: opcode,
                arg: self.arg,
                corrected: true,
            })
        }
    }
}

struct IdxPreds {
    map: HashMap<usize, HashSet<Instruction>>,
}

impl IdxPreds {
    fn add(&mut self, instr: Instruction) {
        let next_idx = instr.next_idx();

        if let Some(set) = self.map.get_mut(&next_idx) {
            set.insert(instr);
        } else {
            let mut set: HashSet<Instruction> = HashSet::new();
            set.insert(instr);
            self.map.insert(next_idx, set);
        }
    }

    fn from_input(input: Vec<String>) -> Self {
        let mut idx_preds = IdxPreds {
            map: HashMap::new(),
        };

        for (idx, op) in input.iter().enumerate() {
            let instr = Instruction::from_str(idx, op);
            if let Some(corr_instr) = instr.corrected() {
                idx_preds.add(corr_instr);
            };
            idx_preds.add(instr);
        }

        idx_preds
    }
}

#[derive(Debug)]
struct TraceStep<'a> {
    corrected: bool,
    next_step: Option<Rc<TraceStep<'a>>>,
    instr: &'a Instruction,
}

fn determine_traceback<'a>(idx_preds: &'a IdxPreds, idx: usize) -> Option<Rc<TraceStep<'a>>> {
    let mut queue: VecDeque<Rc<TraceStep>> = VecDeque::new();
    let terms = idx_preds.map.get(&idx).unwrap();

    for instr in terms {
        queue.push_back(Rc::new(TraceStep {
            corrected: instr.corrected,
            next_step: None,
            instr,
        }));
    }

    // Traverse backwards
    while let Some(next_step) = queue.pop_front() {
        let next_instr = next_step.instr;
        // Get instructions that may be a predecessor
        if let Some(preds) = idx_preds.map.get(&next_instr.idx) {
            // For all instructions that may precede the current...
            for prev_instr in preds {
                let step = Rc::new(TraceStep {
                    // Either a previous instruction has been corrected or this
                    // one is the correction.
                    corrected: prev_instr.corrected || next_step.corrected,
                    next_step: Some(next_step.clone()),
                    instr: prev_instr,
                });
                if prev_instr.idx == 0 {
                    // We're done!
                    return Some(step);
                } else if !(next_step.corrected && prev_instr.corrected) {
                    // Push the next step in the backtrace to the queue.
                    queue.push_back(step);
                }
            }
        } else {
            panic!();
        }
    }
    return None;
}

pub fn solution_2() -> i32 {
    let ops: Vec<String> = load_input("input/8.txt");
    let term_idx = ops.len();

    let idx_preds = IdxPreds::from_input(ops);
    let mut trace_step: Option<Rc<TraceStep>> = determine_traceback(&idx_preds, term_idx);
    let mut acc: i32 = 0;
    let mut idx: usize = 0;
    while let Some(step) = trace_step {
        step.instr.exec(&mut idx, &mut acc);
        trace_step = step.next_step.clone();
    }

    return acc;
}
