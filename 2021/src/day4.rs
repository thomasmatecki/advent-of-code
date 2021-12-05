use aoc_lib::load_input;
use lazy_static::lazy_static;
use regex::Regex;
use std::collections::{HashMap, HashSet};
use std::iter::FromIterator;

lazy_static! {
    static ref INPUT: Vec<String> = load_input("input/4.txt");
    static ref WHITESPACE: Regex = Regex::new(r"\s+").unwrap();
}

pub fn parse_input<T: FromIterator<Vec<u8>>>() -> (Vec<u8>, T) {
    let number_order: Vec<u8> = INPUT[0].split(",").map(|s| s.parse().unwrap()).collect();

    let chunks: T = INPUT[2..]
        .chunks(6)
        .map(|chunk| {
            let joined = &chunk[..5].join(" ");
            let trimmed = joined.trim();
            let split: Vec<&str> = WHITESPACE.split(trimmed).collect();
            let parsed = split.iter().map(|s| s.parse().unwrap()).collect();
            parsed
        })
        .collect();

    return (number_order, chunks);
}

fn build_number_map(boards: &Vec<Vec<u8>>) -> HashMap<u8, Box<Vec<usize>>> {
    let mut map = HashMap::new();

    for (board_idx, board) in boards.iter().enumerate() {
        for (idx, number) in board.iter().enumerate() {
            let board_start = board_idx * 10;
            let v = map.entry(*number).or_insert(Box::new(Vec::new()));
            v.push(board_start + (idx / 5)); // row
            v.push(board_start + (idx % 5) + 5); // column
        }
    }
    return map;
}

fn calculate_score(board: &Vec<u8>, called_numbers: HashSet<u8>, last_number: u32) -> u32 {
    let (_, unmarked): (Vec<u8>, Vec<u8>) = board.iter().partition(|v| called_numbers.contains(v));

    let unmarked_sum = unmarked.iter().map(|v| *v as u32).sum::<u32>();
    return last_number * unmarked_sum;
}

pub fn solution_1() -> u32 {
    let (number_order, boards) = parse_input::<Vec<Vec<u8>>>();

    let row_col_map = build_number_map(&boards);
    let mut called_numbers: HashSet<u8> = HashSet::new();
    let mut remaining_squares_idx: [u8; 1000] = [5; 1000];

    for (_number_idx, number) in number_order.iter().enumerate() {
        called_numbers.insert(*number);
        for square in row_col_map.get(&number).unwrap().iter() {
            remaining_squares_idx[*square as usize] -= 1;
            if remaining_squares_idx[*square as usize] == 0 {
                let board_idx = square / 10;
                let winning_board = &boards[board_idx];
                return calculate_score(winning_board, called_numbers, *number as u32);
            }
        }
    }

    return 0;
}

pub fn solution_2() -> u32 {
    let (number_order, boards) = parse_input::<Vec<Vec<u8>>>();

    let row_col_map = build_number_map(&boards);
    let mut called_numbers: HashSet<u8> = HashSet::new();
    let mut remaining_squares_idx: [u8; 1000] = [5; 1000];
    let mut remaining_board_idxs: HashSet<usize> = (0..boards.len()).collect();

    for (_number_idx, number) in number_order.iter().enumerate() {
        called_numbers.insert(*number);
        for square in row_col_map.get(&number).unwrap().iter() {
            remaining_squares_idx[*square as usize] -= 1;

            let board_idx = square / 10;
            if remaining_squares_idx[*square as usize] == 0 {
                remaining_board_idxs.remove(&board_idx);
            }
            if remaining_board_idxs.is_empty() {
                let winning_board = &boards[board_idx];
                return calculate_score(winning_board, called_numbers, *number as u32);
            }
        }
    }

    return 0;
}
