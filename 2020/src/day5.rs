use aoc_lib::load_input;

fn radix2str_to_int(column_str: &str, one: char, zero: char) -> u16 {
    let binary_str = column_str.replace(zero, "0").replace(one, "1");
    return u16::from_str_radix(&binary_str, 2).unwrap();
}

fn parse_seat_bsp(bsp: &str) -> u32 {
    let (row_str, column_str) = bsp.split_at(7);
    let row = radix2str_to_int(row_str, 'B', 'F');
    let column = radix2str_to_int(column_str, 'R', 'L');
    let seat_id = ((row * 8) + column).into();
    return seat_id;
}

pub fn solution_1() -> u32 {
    let max = load_input("input/5.txt")
        .iter()
        .map(|line| parse_seat_bsp(line))
        .max();
    return max.unwrap().into();
}

pub fn solution_2() -> u32 {
    let mut seat_ids: Vec<u32> = load_input("input/5.txt")
        .iter()
        .map(|line| parse_seat_bsp(line))
        .collect();

    seat_ids.sort();

    for window in seat_ids.windows(2) {
        if window[0] + 1 != window[1] {
            return window[0] + 1;
        }
    }

    //There should be a missing seat
    unreachable!();
}

#[cfg(test)]
mod tests {

    use super::*;
    #[test]
    fn example_row() {
        assert_eq!(44, radix2str_to_int("FBFBBFF", 'B', 'F'));
        assert_eq!(70, radix2str_to_int("BFFFBBF", 'B', 'F'));
        assert_eq!(14, radix2str_to_int("FFFBBBF", 'B', 'F'));
        assert_eq!(102, radix2str_to_int("BBFFBBF", 'B', 'F'));
    }
    #[test]
    fn example_column() {
        assert_eq!(5, radix2str_to_int("RLR", 'R', 'L'));
        assert_eq!(7, radix2str_to_int("RRR", 'R', 'L'));
        assert_eq!(7, radix2str_to_int("RRR", 'R', 'L'));
        assert_eq!(4, radix2str_to_int("RLL", 'R', 'L'));
    }

    #[test]
    fn examples() {
        assert_eq!(567, parse_seat_bsp("BFFFBBFRRR"));
        assert_eq!(119, parse_seat_bsp("FFFBBBFRRR"));
        assert_eq!(820, parse_seat_bsp("BBFFBBFRLL"));
    }
}
