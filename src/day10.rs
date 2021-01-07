use crate::utils::load_input;
use std::collections::HashMap;

fn parse_input(filename: &str) -> Vec<u32> {
    load_input(filename)
        .iter()
        .map(|line| line.parse().unwrap())
        .collect()
}

fn multiply_counts(voltages: &mut Vec<u32>, diff1: u32, diff2: u32) -> (u32, u32) {
    let mut diff_counts: HashMap<u32, u32> = HashMap::new();
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

pub fn solution_1() -> u32 {
    let mut voltages = parse_input("input/10.txt");
    let (a, b) = multiply_counts(&mut voltages, 1, 3);
    a * b
}

mod tests {

    use super::*;
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
