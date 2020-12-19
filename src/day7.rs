use std::collections::HashSet;
use std::collections::HashMap;
use crate::utils::load_input;

fn parse_input() {

    let lines: Vec<String> = load_input("input/7.txt");
    let mut m: HashMap<&str, HashSet<&str>> = HashMap::new();
    
    for line in lines.iter() {
        let containers = line.split(" bags contain ").take(2).collect::<Vec<&str>>();

        if let [container, containees] = &containers[..] {
            let containees = containees.trim_end_matches(".");  
            for containee in containees.split(", "){
                println!("{}", containee);
                let (count, bag) = containee.split_at(1);
                match m.get_mut(bag){
                    Some(set)=>{
                        set.insert(container);
                    },
                    None =>{}
                }
            }
            println!("{}", container);
        }
    }
}

pub fn solution_1() -> u32 {
    parse_input();

    0
}
