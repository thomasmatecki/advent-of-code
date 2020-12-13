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
