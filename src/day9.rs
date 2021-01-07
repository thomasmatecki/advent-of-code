use crate::utils::load_input;
use std::cmp::max;
use std::cmp::min;

fn check_pair_sum(preceding: &[u64], sum: u64) -> bool {
    for (idx, i) in preceding.iter().enumerate() {
        for j in &preceding[idx + 1..25] {
            if i + j == sum {
                return true;
            }
        }
    }

    false
}

fn parse_input() -> Vec<u64> {
    return load_input("input/9.txt")
        .iter()
        .map(|line| line.parse().unwrap())
        .collect();
}

fn first_invaid(xmas_data: &Vec<u64>) -> u64 {
    for (idx, window) in xmas_data.windows(25).enumerate() {
        let target_value = xmas_data[idx + 25];
        if !check_pair_sum(window, target_value) {
            return target_value;
        }
    }

    unreachable!()
}

pub fn solution_1() -> u64 {
    let xmas_data = parse_input();
    return first_invaid(&xmas_data);
}

pub fn solution_2() -> u64 {
    let xmas_data = parse_input();
    let target_sum = first_invaid(&xmas_data);
    let mut fr: usize = 0;
    let mut to: usize = 1;

    let mut window_sum: u64 = xmas_data[fr..=to].iter().sum();

    while window_sum != target_sum {
        if fr == to {
            to += 1;
            window_sum += xmas_data[to];
        } else if window_sum > target_sum {
            window_sum -= xmas_data[fr];
            fr += 1;
        } else {
            to += 1;
            window_sum += xmas_data[to];
        }
    }
    let min = xmas_data[fr..to]
        .iter()
        .fold(xmas_data[to], |acc, i| min(acc, *i));

    let max = xmas_data[fr..to]
        .iter()
        .fold(xmas_data[to], |acc, i| max(acc, *i));
    return min + max;
}
