#![feature(min_const_generics)]

mod day1;
mod day17;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod utils;

fn main() {
    println!("Day 1, part 1: {}", day1::solution_1());
    println!("Day 1, part 2: {}", day1::solution_2());
    println!("Day 2, part 1: {}", day2::solution_1());
    println!("Day 2, part 2 {}", day2::solution_2());
    println!("Day 3, part 1 {}", day3::solution_1());
    println!("Day 3, part 2 {}", day3::solution_2());
    println!("Day 4, part 1 {}", day4::solution_1());
    println!("Day 4, part 2 {}", day4::solution_2());
    println!("Day 5, part 1 {}", day5::solution_1());
    println!("Day 5, part 2 {}", day5::solution_2());
    println!("Day 6, part 1 {}", day6::solution_1());
    println!("Day 6, part 2 {}", day6::solution_2());
    println!("Day 7, part 1 {}", day7::solution_1());
    println!("Day 7, part 2 {}", day7::solution_2());
    println!("Day 8, part 1 {}", day8::solution_1());
    println!("Day 8, part 2 {}", day8::solution_2());
    println!("Day 17, part 1 {}", day17::solution_1());
    println!("Day 17, part 2 {}", day17::solution_2());
}
