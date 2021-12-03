use aoc_lib::load_input;
use std::collections::HashMap;

fn parse_input(filename: &str) -> Vec<u64> {
    load_input(filename)
        .iter()
        .map(|line| line.parse().unwrap())
        .collect()
}

fn multiply_counts(voltages: &mut Vec<u64>, diff1: u64, diff2: u64) -> (u64, u64) {
    let mut diff_counts: HashMap<u64, u64> = HashMap::new();
    voltages.sort();

    for w in voltages.windows(2) {
        let diff = w[1] - w[0];
        *diff_counts.entry(diff).or_insert(1) += 1;
    }
    (
        *diff_counts.get(&diff1).unwrap_or(&0),
        *diff_counts.get(&diff2).unwrap_or(&0),
    )
}

pub fn solution_1() -> u64 {
    let mut voltages = parse_input("input/10.txt");
    let (a, b) = multiply_counts(&mut voltages, 1, 3);
    a * b
}

fn calc_arrangements(voltages: &mut Vec<u64>) -> u64 {
    voltages.push(0);
    voltages.sort();
    let mut prefixs_counts = vec![1; voltages.len()];

    for (idx, voltage) in voltages.iter().enumerate().skip(1) {
        prefixs_counts[idx] = voltages[..idx]
            .iter()
            .enumerate()
            .filter(|(_, prev_v)| *prev_v + 3 >= *voltage)
            .map(|(idx, _)| prefixs_counts[idx])
            .sum();
    }

    return *prefixs_counts.last().unwrap();
}

pub fn solution_2() -> u64 {
    let mut voltages = parse_input("input/10.txt");
    return calc_arrangements(&mut voltages);
}
mod tests {

    use super::{calc_arrangements, multiply_counts, parse_input};
    mod solution_1 {
        use super::{multiply_counts, parse_input};
        #[test]
        fn example_1() {
            let mut voltages = parse_input("input/10ex1.txt");
            let (a, b) = multiply_counts(&mut voltages, 1, 3);
            assert_eq!(a, 7);
            assert_eq!(b, 5);
        }

        #[test]
        fn example_2() {
            let mut voltages = parse_input("input/10ex2.txt");
            let (a, b) = multiply_counts(&mut voltages, 1, 3);
            assert_eq!(a, 22);
            assert_eq!(b, 10);
        }
    }

    mod solution_2 {
        use super::{calc_arrangements, parse_input};
        #[test]
        fn example_1() {
            let mut voltages = parse_input("input/10ex1.txt");
            let answer = calc_arrangements(&mut voltages);
            assert_eq!(answer, 8);
        }
        #[test]
        fn example_2() {
            let mut voltages = parse_input("input/10ex2.txt");
            let answer = calc_arrangements(&mut voltages);
            assert_eq!(answer, 19208);
        }
    }
}
