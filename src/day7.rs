use crate::utils::load_input;
use std::collections::HashMap;
use std::collections::HashSet;
use std::collections::VecDeque;

fn parse_input<'a>(lines: &'a Vec<String>) -> Box<HashMap<&'a str, HashSet<&'a str>>> {
    let mut container_map: Box<HashMap<&str, HashSet<&str>>> = Box::new(HashMap::new());

    for line in lines.iter() {
        let containers = line.split(" bags contain ").take(2).collect::<Vec<&str>>();

        if let [container, containees] = &containers[..] {
            let containees = containees.trim_end_matches(".");
            for containee in containees.split(", ") {
                let (count, mut bag) = containee.split_at(1);

                bag = bag.trim_end_matches('s');
                bag = bag.trim_end_matches("bag");
                bag = bag.trim();
                match container_map.get_mut(bag) {
                    Some(set) => {
                        set.insert(container);
                    }
                    None => {
                        let mut set: HashSet<&str> = HashSet::new();
                        set.insert(container);
                        container_map.insert(bag, set);
                    }
                }
            }
        }
    }
    return container_map;
}

fn calc_possible_holders(container_map: &HashMap<&str, HashSet<&str>>) -> u32 {
    let mut holders: HashSet<&str> = HashSet::new();
    let mut queue: VecDeque<&str> = VecDeque::new();

    queue.push_back("shiny gold");

    loop {
        match queue.pop_front() {
            Some(containee) => {
                holders.insert(containee);
                match container_map.get(containee) {
                    Some(containers) => {
                        containers.iter().for_each(|container| {
                            if !holders.contains(container) {
                                queue.push_back(container);
                            }
                        });
                    }
                    None => {}
                }
            }
            None => break,
        };
    }
    return (holders.len() - 1) as u32;
}

pub fn solution_1() -> u32 {
    let lines = load_input("input/7.txt");
    // A map from each bag to those that can contain it.
    let container_map = parse_input(&lines);
    return calc_possible_holders(&container_map);
}

#[cfg(test)]
mod tests {

    use super::*;
    #[test]
    fn example() {
        let lines = load_input("input/7ex.txt");
        let container_map = parse_input(&lines);
        assert_eq!(4, calc_possible_holders(&container_map));
    }
}
