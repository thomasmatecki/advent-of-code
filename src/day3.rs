use crate::utils::load_input;

fn count_trees(right: usize, down: usize) -> u32 {
    let mut count = 0;

    for (step, line) in load_input(&"input/3.txt").iter().enumerate() {
        if step % down == 0 {
            // Tricky arithmetic
            let shift = (step * right / down) % line.len();
            count += match line.chars().nth(shift) {
                Some('#') => 1,
                _ => 0,
            };
        }
    }
    return count;
}

pub fn solution_1() -> u32 {
    return count_trees(3, 1);
}

pub fn solution_2() -> u32 {
    count_trees(1, 1)
        * count_trees(3, 1)
        * count_trees(5, 1)
        * count_trees(7, 1)
        * count_trees(1, 2)
}
