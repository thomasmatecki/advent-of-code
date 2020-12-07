use crate::utils::load_input;

fn parse_input() -> Vec<i32> {
    load_input(&"input/1.txt")
        .iter()
        .map(|l| l.parse().unwrap())
        .collect()
}

pub fn solution_1() -> i32 {
    let mut numbers = parse_input();

    for pivot in (1..numbers.len()).rev() {
        // Iterate like a bubblesort, after each iteration all entries to the
        // right of `pivot` are sorted AND greater than ALL those to the left.
        for idx in 0..pivot {
            if numbers[idx] > numbers[idx + 1] {
                numbers.swap(idx, idx + 1);
            }
        }

        // Binary search from pivot -> to the end of the array for number that
        // adds up to 2020;
        let operand0 = numbers[pivot - 1];
        let operand1 = 2020 - operand0;
        let s = &numbers[pivot..numbers.len()];
        match s.binary_search(&operand1) {
            Err(_) => {}
            Ok(_) => {
                return operand0 * operand1;
            }
        };
    }
    return 0;
}

pub fn solution_2() -> i32 {
    let numbers = parse_input();

    for i in &numbers {
        for j in &numbers {
            for k in &numbers {
                if i + j + k == 2020 {
                    // This is so gross but still fast.
                    return i * j * k;
                }
            }
        }
    }
    return 0;
}
