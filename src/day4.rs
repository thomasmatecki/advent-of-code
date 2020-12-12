use crate::utils::load_input;
use regex::Regex;
use std::collections::HashMap;
use std::mem::replace;

type Doc<'a> = Box<HashMap<&'a str, &'a str>>;
const PASSPORT_KEYS: [&str; 7] = [
    "ecl", "pid", "eyr", "hcl", "byr", "iyr", // "cid",
    "hgt",
];

const HAIR_COLORS: [&str; 7] = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];

fn is_valid_year(year_str: &str, min: u32, max: u32) -> bool {
    let year: u32 = year_str.parse().unwrap();
    return min <= year && year <= max;
}

fn is_valid_height(height_str: &str) -> bool {
    let re = Regex::new(r"^(\d+)(cm|in)$").unwrap();
    match re.captures(height_str) {
        Some(capture) => {
            let value: u32 = capture[1].parse().unwrap();
            if &capture[2] == "cm" {
                150 <= value && value <= 193
            } else if &capture[2] == "in" {
                59 <= value && value <= 76
            } else {
                false
            }
        }
        None => false,
    }
}

fn is_valid_hair_color(hair_color: &str) -> bool {
    let re = Regex::new(r"#(\d|[a-f]){6}").unwrap();
    re.is_match(hair_color)
}

fn is_valid_eye_color(eye_color: &str) -> bool {
    HAIR_COLORS.contains(&eye_color)
}

fn is_valid_passport(passpord_id: &str) -> bool {
    let re = Regex::new(r"^\d{9}$").unwrap();
    re.is_match(passpord_id)
}

fn has_required_fields(doc: &Doc) -> bool {
    PASSPORT_KEYS.iter().all(|key| doc.contains_key(key))
}

/// Convert lines of input files to _Docs_(HashMaps of 'fields')
fn parse_input<'a>(lines: &'a Vec<String>, docs: &mut Vec<Doc<'a>>) {
    let mut doc: Doc = Box::new(HashMap::new());
    for line in lines.iter() {
        if line == "" {
            docs.push(replace(&mut doc, Box::new(HashMap::new())));
        } else {
            for key_value in line.split(' ') {
                let kv = key_value.split(":").take(2).collect::<Vec<&str>>();
                if let [key, value] = &kv[..] {
                    doc.insert(key, value);
                }
            }
        }
    }
    docs.push(doc);
}

pub fn solution_1() -> u32 {
    let mut docs: Vec<Doc> = Vec::new();
    let lines: Vec<String> = load_input("input/4.txt");

    parse_input(&lines, &mut docs);

    let count = docs.iter().filter(|doc| has_required_fields(doc)).count();

    return count as u32;
}

pub fn solution_2() -> u32 {
    let mut docs: Vec<Doc> = Vec::new();
    let lines: Vec<String> = load_input("input/4.txt");

    parse_input(&lines, &mut docs);

    //cid (Country ID) - ignored, missing or not.
    let count = docs
        .iter()
        .filter(|doc| {
            has_required_fields(doc)
                //byr (Birth Year) - four digits; at least 1920 and at most 2002.
                && is_valid_year(doc["byr"], 1920, 2002)
                //iyr (Issue Year) - four digits; at least 2010 and at most 2020.
                && is_valid_year(doc["iyr"], 2010, 2020)
                //eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
                && is_valid_year(doc["eyr"], 2020, 2030)
                //hgt (Height) - a number followed by either cm or in:
                //If cm, the number must be at least 150 and at most 193.
                //If in, the number must be at least 59 and at most 76.
                && is_valid_height(doc["hgt"])
                //hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
                && is_valid_hair_color(doc["hcl"])
                //ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
                && is_valid_eye_color(doc["ecl"])
                //pid (Passport ID) - a nine-digit number, including leading zeroes.
                && is_valid_passport(doc["pid"])
        })
        .count();

    return count as u32;
}

#[cfg(test)]
mod tests {
    use super::*;
    mod birth_year {
        use super::*;

        #[test]
        fn valid() {
            assert_eq!(true, (is_valid_year("2002", 1920, 2002)));
        }
        #[test]
        fn invalid_after() {
            assert_eq!(false, (is_valid_year("2003", 1920, 2002)));
        }
        #[test]
        fn invalid_before() {
            assert_eq!(false, (is_valid_year("1919", 1920, 2002)));
        }
    }

    mod height {
        use super::*;

        #[test]
        fn valid_in() {
            assert_eq!(true, is_valid_height("60in"));
        }

        #[test]
        fn valid_cm() {
            assert_eq!(true, is_valid_height("190cm"));
        }

        #[test]
        fn invalid_in() {
            assert_eq!(false, is_valid_height("190in"));
        }

        #[test]
        fn invalid() {
            assert_eq!(false, is_valid_height("190"));
        }
    }

    mod hair_color {
        use super::*;

        #[test]
        fn valid() {
            assert_eq!(true, is_valid_hair_color("#123abc"));
        }
        #[test]
        fn valid_letters() {
            assert_eq!(true, is_valid_hair_color("#ffffff"));
        }
        #[test]
        fn invalid_char() {
            assert_eq!(false, is_valid_hair_color("#123abz"));
        }
        #[test]
        fn invalid_missing_hash() {
            assert_eq!(false, is_valid_hair_color("123abc"));
        }
        #[test]
        fn invalid_only_numbers() {
            assert_eq!(false, is_valid_hair_color("123456"));
        }
    }

    mod eye_color {
        use super::*;

        #[test]
        fn valid() {
            assert_eq!(true, is_valid_eye_color("brn"));
        }
        #[test]
        fn invalid() {
            assert_eq!(false, is_valid_eye_color("wat"));
        }
    }
    mod passport {
        use super::*;

        #[test]
        fn valid() {
            assert_eq!(true, is_valid_passport("000000001"));
        }
        #[test]
        fn invalid_long() {
            assert_eq!(false, is_valid_passport("0123456789"));
        }
        #[test]
        fn invalid_short() {
            assert_eq!(false, is_valid_passport("23456789"));
        }
    }
}
