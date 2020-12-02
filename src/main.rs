use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;
use std::path::Path;

const FILENAME: &str = "inputs/1.txt";

fn load_input() -> Vec<i32> {
    let path = Path::new(FILENAME);

    let mut file: File = match File::open(&path) {
        Err(why) => panic!("couldn't open {}", why),
        Ok(file) => file,
    };

    let reader = BufReader::new(&file).lines();
    let mut numbers: Vec<i32> = reader
        .map(|line| line.unwrap().parse::<i32>().unwrap())
        .collect();

    return numbers;
}

fn day1_1() -> i32 {
    let mut numbers = load_input();
    for pivot in (1..numbers.len()).rev() {
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

fn day1_2() -> i32 {
    let numbers = load_input();
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

fn main() {
    println!("the answer is {}", day1_2());
}
