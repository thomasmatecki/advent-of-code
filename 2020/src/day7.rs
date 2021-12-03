use aoc_lib::load_input;
use lazy_static::lazy_static;
use regex::Regex;
use std::collections::{HashMap, HashSet, VecDeque};

fn parse_container_map<'a>(lines: &'a Vec<String>) -> Box<HashMap<&'a str, HashSet<&'a str>>> {
    let mut container_map: Box<HashMap<&str, HashSet<&str>>> = Box::new(HashMap::new());

    for line in lines.iter() {
        let containers = line.split(" bags contain ").take(2).collect::<Vec<&str>>();

        if let [container, containees] = &containers[..] {
            let containees = containees.trim_end_matches(".");
            for containee in containees.split(", ") {
                let (_count, mut bag) = containee.split_at(1);

                bag = bag.trim_end_matches('s');
                bag = bag.trim_end_matches("bag");
                bag = bag.trim();

                if let Some(set) = container_map.get_mut(bag) {
                    set.insert(container);
                } else {
                    let mut set: HashSet<&str> = HashSet::new();
                    set.insert(container);
                    container_map.insert(bag, set);
                }
            }
        }
    }

    return container_map;
}

fn count_containers(container_map: &HashMap<&str, HashSet<&str>>) -> u32 {
    let mut holders: HashSet<&str> = HashSet::new();
    let mut queue: VecDeque<&str> = VecDeque::new();

    queue.push_back("shiny gold");
    while let Some(containee) = queue.pop_front() {
        holders.insert(containee);
        if let Some(containers) = container_map.get(containee) {
            for container in containers {
                if !holders.contains(container) {
                    queue.push_back(container);
                }
            }
        }
    }

    return (holders.len() - 1) as u32;
}

pub fn solution_1() -> u32 {
    let lines = load_input("input/7.txt");
    // A map from each bag to those that can contain it.
    let container_map = parse_container_map(&lines);
    return count_containers(&container_map);
}

struct BagCount<'a> {
    count: u32,
    bag: &'a str,
}

pub struct ContaineeCount<'a> {
    map: Box<HashMap<&'a str, Vec<BagCount<'a>>>>,
}

impl<'a> ContaineeCount<'a> {
    pub fn from_input(lines: &'a Vec<String>) -> Self {
        lazy_static! {
            static ref RE: Regex = Regex::new(r"(\d) (\w+ \w+)").unwrap();
        }
        let mut map: Box<HashMap<&str, Vec<BagCount>>> = Box::new(HashMap::new());

        for line in lines.iter() {
            let containers = line.split(" bags contain ").take(2).collect::<Vec<&str>>();

            if let [container, containees] = &containers[..] {
                let mut counts: Vec<BagCount> = Vec::new();

                for containee in RE.captures_iter(containees) {
                    let count = &containee[1];
                    let bag = containee.get(2).unwrap().as_str();
                    let bag_count = BagCount {
                        count: count.parse().unwrap(),
                        bag,
                    };
                    counts.push(bag_count);
                }

                map.insert(container, counts);
            }
        }
        return ContaineeCount { map };
    }

    fn count(&self, count: u32, bag: &str) -> u32 {
        let containees: &Vec<BagCount> = self.map.get(bag).unwrap();
        let sum: u32 = containees.iter().map(|c| self.count(c.count, c.bag)).sum();
        return count + (sum * count);
    }

    pub fn for_bag(&self, bag: &str) -> u32 {
        self.count(1, bag) - 1
    }
}

pub fn solution_2() -> u32 {
    let input = load_input("input/7.txt");
    // A map from each bag to those it can contain.
    let containee_count = ContaineeCount::from_input(&input);
    return containee_count.for_bag("shiny gold");
}

#[cfg(test)]
mod tests {

    use super::*;
    mod solution_1 {
        use super::*;
        #[test]
        fn example() {
            let lines = load_input("input/7ex1.txt");
            let container_map = parse_container_map(&lines);
            assert_eq!(4, count_containers(&container_map));
        }
    }

    mod solution_2 {
        use super::*;

        #[test]
        fn get_empty() {
            let input = load_input("input/7ex1.txt");
            let containee_count = ContaineeCount::from_input(&input);
            let dotted_black = containee_count.map.get("dotted black").unwrap();
            assert_eq!(dotted_black.is_empty(), true);
        }

        #[test]
        fn example_1() {
            let input = load_input("input/7ex1.txt");
            let containee_count = ContaineeCount::from_input(&input);
            let count = containee_count.for_bag("shiny gold");
            assert_eq!(count, 32)
        }

        #[test]
        fn example_2() {
            let input = load_input("input/7ex2.txt");
            let containee_count = ContaineeCount::from_input(&input);
            let count = containee_count.for_bag("shiny gold");
            assert_eq!(count, 126)
        }
    }
}
