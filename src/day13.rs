use crate::utils::load_input;

fn parse_input(filename: &str) -> (u32, Vec<u32>) {
    let input = load_input(filename);
    if let [depart_time_str, buses_str] = &input[..] {
        let depart_time: u32 = depart_time_str.parse().unwrap();
        let buses: Vec<u32> = buses_str
            .split(',')
            .filter_map(|c| c.parse::<u32>().ok())
            .collect();

        return (depart_time, buses);
    };
    unreachable!();
}

fn earliest_bus(depart_time: u32, buses: Vec<u32>) -> (u32, u32) {
    let next_buses = buses.iter().map(|bus| {
        let mut v = 0;
        while v < depart_time {
            v += bus;
        }
        return (v, bus);
    });

    let result = next_buses
        .map(|(nb, b)| (nb - depart_time, *b))
        .min()
        .unwrap();

    return result;
}

///
///
///
///
pub fn solution_1() -> u32 {
    let (depart_time, buses) = parse_input("input/13.txt");
    let result = earliest_bus(depart_time, buses);
    return result.0 * result.1;
}

#[cfg(test)]
mod test {

    use super::*;
    #[test]
    fn example() {
        let (depart_time, buses) = parse_input("input/13ex.txt");
        let result = earliest_bus(depart_time, buses);
        assert_eq!(result.0, 5);
        assert_eq!(result.1, 59);
    }
}
