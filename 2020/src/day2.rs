use aoc_lib::load_input;

use regex::Regex;

struct Record {
    mino: usize,
    maxo: usize,
    letter: char,
    password: String,
}

/// Parse rows like: 2-6 c: fcpwjqhcgtffzlbj
fn parse_input() -> Vec<Record> {
    let re = Regex::new(r"(\d+)-(\d+) (\w): (\w+)").unwrap();
    let lines = load_input(&"input/2.txt");
    let captures = lines.iter().map(|line| re.captures(line).unwrap());

    let records = captures.map(|capture| {
        let mino: usize = capture[1].parse().unwrap();
        let maxo: usize = capture[2].parse().unwrap();
        let letter: char = capture[3].parse().unwrap();
        let password: String = capture[4].parse().unwrap();

        return Record {
            mino,
            maxo,
            letter,
            password,
        };
    });

    return records.collect();
}

pub fn solution_1() -> u32 {
    let password_records: Vec<Record> = parse_input();
    let mut count: u32 = 0;

    for record in password_records {
        let letter_count: usize = record.password.matches(record.letter).count();

        if (record.mino..record.maxo + 1).contains(&letter_count) {
            count += 1;
        };
    }

    return count;
}

pub fn solution_2() -> u32 {
    let password_records: Vec<Record> = parse_input();
    let mut count: u32 = 0;

    for record in password_records {
        let passsword_chars: Vec<char> = record.password.chars().collect();
        let min_char = passsword_chars.get(record.mino - 1);
        let max_char = passsword_chars.get(record.maxo - 1);

        if min_char.is_none() || min_char.is_none() {
            break;
        }

        if (min_char.unwrap() == &record.letter) ^ (max_char.unwrap() == &record.letter) {
            count += 1;
        };
    }

    return count;
}
