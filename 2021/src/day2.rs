use aoc_lib::load_input;
use lazy_static::lazy_static;
use regex::Regex;


lazy_static! {
    static ref RE: Regex = Regex::new(r"(\w+) (\d+)").unwrap();
}

fn parse_input() -> Vec<(i32, i32)> {
    let instructions: Vec<String> = load_input("input/2.txt");

    return instructions.iter().map(|instruction| {
        let cap = RE.captures(instruction).unwrap();
        let magnitude:i32 = cap[2].parse().unwrap();

        match &cap[1]{
            "up" => {(0, -magnitude)}
            "down" => {(0, magnitude)}
            "forward" => {(magnitude,0)}
            _ => panic!("invalid input")
        }
    }).collect();
}

pub fn solution_1() -> i32 {

    let final_position = parse_input().iter().fold((0,0),|acc, mov| (acc.0+ mov.0,acc.1 + mov.1) );
    return final_position.0 * final_position.1;
}

pub fn solution_2() -> i32 {
    return 0;
}
