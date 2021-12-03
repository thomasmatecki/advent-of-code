use aoc_lib::load_input;
use std::collections::HashSet;
use std::iter::FromIterator;
use std::mem::replace;

type Group = HashSet<char>;

pub fn solution_1() -> u32 {
    let mut groups: Vec<Group> = Vec::new();
    let lines: Vec<String> = load_input("input/6.txt");

    groups.push(HashSet::new());

    for line in lines.iter() {
        if line == "" {
            groups.push(HashSet::new());
        } else {
            let group = groups.last_mut().unwrap();
            line.chars().for_each(|c| {
                group.insert(c);
            });
        }
    }

    return groups.iter().map(|group| group.len() as u32).sum::<u32>();
}

const ALPHA: &str = "abcdefghijklmnopqrstuvwxyz";

pub fn solution_2() -> u32 {
    let mut groups: Vec<Group> = Vec::new();
    let lines: Vec<String> = load_input("input/6.txt");
    let mut idx = 0;

    groups.push(HashSet::from_iter(ALPHA.chars()));

    for line in lines.iter() {
        if line == "" {
            idx += 1;
            groups.push(HashSet::from_iter(ALPHA.chars()));
        } else {
            let both: HashSet<char> = line
                .chars()
                .collect::<HashSet<char>>()
                .intersection(groups.last().unwrap())
                .map(|c| c.clone())
                .collect();

            let _existing = replace(&mut groups[idx], both);
        }
    }

    return groups.iter().map(|group| group.len() as u32).sum::<u32>();
}
