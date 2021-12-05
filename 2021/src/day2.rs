use aoc_lib::load_input;
use lazy_static::lazy_static;
use regex::Captures;
use regex::Regex;

lazy_static! {
    static ref RE: Regex = Regex::new(r"(\w+) (\d+)").unwrap();
    static ref INSTRUCTIONS: Vec<String> = load_input("input/2.txt");
}

fn parse_input<T>(mapper: fn(cap: &Captures, magnitude: &i32) -> T) -> Vec<T> {
    INSTRUCTIONS
        .iter()
        .map(|instruction| {
            let cap = RE.captures(instruction).unwrap();
            mapper(&cap, &cap[2].parse().unwrap())
        })
        .collect()
}

pub fn solution_1() -> i32 {
    let final_position = parse_input(|cap: &Captures, magnitude: &i32| match &cap[1] {
        "up" => (0, -magnitude),
        "down" => (0, magnitude.clone()),
        "forward" => (magnitude.clone(), 0),
        _ => panic!("invalid input"),
    })
    .iter()
    .fold((0, 0), |acc, mov| (acc.0 + mov.0, acc.1 + mov.1));
    return final_position.0 * final_position.1;
}

pub fn solution_2() -> i32 {
    let mut aim = 0;

    let mut position = 0;
    let mut depth = 0;

    for (movement, aim_change) in parse_input(|cap: &Captures, magnitude: &i32| match &cap[1] {
        "up" => (0, -magnitude),
        "down" => (0, magnitude.clone()),
        "forward" => (magnitude.clone(), 0),
        _ => panic!("invalid input"),
    })
    .iter()
    {
        aim += aim_change;
        position += movement;
        depth += aim * movement;
    }
    return position * depth;
}
