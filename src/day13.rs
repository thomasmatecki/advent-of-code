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

pub fn solution_1() -> u32 {
    let (depart_time, buses) = parse_input("input/13.txt");
    let result = earliest_bus(depart_time, buses);
    return result.0 * result.1;
}

fn ext_euclid(a: i64, b: i64) -> (i64, i64, i64) {
    let mut r = (a, b);
    let mut s = (1, 0);
    let mut t = (0, 1);

    while r.1 != 0 {
        let q = r.0 / r.1;
        r = (r.1, r.0 - q * r.1);
        s = (s.1, s.0 - q * s.1);
        t = (t.1, t.0 - q * t.1);
    }

    (r.0, s.0, t.0)
}

fn bezout_solve(a: (i64, i64), o: i64) -> (i64, i64) {
    let (gcd, s, _t) = ext_euclid(a.0, a.1);
    let (q, r) = (o / gcd, o % gcd);
    if r != 0 {
        panic!("GCD does not divide offset")
    }
    let lcm = (a.0 * a.1).abs() / gcd;
    let c = (s * -q).rem_euclid(lcm / a.0);

    ((a.0 * c) % lcm, lcm)
}
/// Given a0, a1
///
/// Return n, m
/// Such that s = n + mk , Satisfy
///    s + a0_1 mod a0_0 = 0
///    s + a1_1 mod a1_0 = 0
///
fn offset_bezout(a0: (i64, i64), a1: (i64, i64)) -> (i64, i64) {
    let d = a1.1 - a0.1;
    let s = bezout_solve((a0.0, a1.0), d);
    return (s.0 - a0.1, s.1);
}

fn parse_buses(s: &str) -> Vec<(i64, i64)> {
    s.split(',')
        .enumerate()
        .filter_map(|(idx, bus)| bus.parse::<i64>().ok().map(|p| (p, idx as i64)))
        .collect()
}

/// Whew... this deserves some explanation
///
/// We're utilizing the chinese remainder theorem to solve a system of modular
/// congruences. The buses(k) each of have a period of a_k (the modulo) and we
/// are looking to solve for a value, n where each bus departs some x_k minutes
/// later. Yielding the system:
///
///    n       mod a_0 = 0
///    n + x_1 mod a_1 = 0
///      .
///      .
///      .
///    n + x_1 mod a_1 = 0
///
/// We can solve an two of these using Extended Euclids Algorithm, that is
/// for
///
///    n       mod a_i = 0
///    n + x_j mod a_j = 0
///
/// We can find Bezout coefficents c_1 & c_2 satisying:
///    c_1 a_i + c_2 a_j = GCD(a_i, a_j) = 1
///
/// All of the bus intervals are co-prime so GCD(a_i, a_j) = 1. From this we
/// obtain the timestamp at which two buses will depart some m minutes
/// apart: m (c_1 a_i + c_2 a_j) = m
///        =>  m c_1 a_i = m (1 - c_2 a_j)
///
/// ... this is another periodic occurence. Every P = lcm(a_i, a_j) the i'th
/// and j'th buses will depart m minutes are part. So we've take two periodic
/// buses, figured out the period in which they coincide. For B buses, we now
/// have reduced are our system of congruences to B-1 periodic buses. We thus
/// can repeat this process _reducing_ over our set of buses until we have
/// a single period in which they all coincide.
pub fn solution_2() -> i64 {
    let input = load_input("input/13.txt");
    let buses = &input[1];
    let schedule = parse_buses(&buses);
    let mut redux = schedule[0];
    let mut next = redux;

    for bus in schedule.iter().skip(1) {
        redux = offset_bezout(next, *bus);
        next = (redux.1, -redux.0);
    }

    return redux.0;
}

#[cfg(test)]
mod test {

    use super::*;

    mod part_1 {
        use super::*;
        #[test]
        fn example() {
            let (depart_time, buses) = parse_input("input/13ex1.txt");
            let result = earliest_bus(depart_time, buses);
            assert_eq!(result.0, 5);
            assert_eq!(result.1, 59);
        }
    }

    mod test_ext_euclid {

        use super::ext_euclid;
        #[test]
        fn test_240_46() {
            let a = 240;
            let b = 46;
            let result = ext_euclid(a, b);
            assert_eq!(result, (2, -9, 47));
            let (gcd, s, t) = result;
            assert_eq!(s * a + t * b, gcd)
        }

        #[test]
        fn test_7_13() {
            let a = 7;
            let b = 13;
            let result = ext_euclid(a, b);
            let (gcd, s, t) = result;
            assert_eq!(s * a + t * b, gcd)
        }
    }
    mod test_bezout_solve {

        use super::bezout_solve;
        #[test]
        fn test_7_13_offset_1() {
            let a = 7;
            let b = 13;
            let offset = 1;
            let result = bezout_solve((a, b), offset);
            let (solution, modulo) = result;
            assert_eq!(solution, 77);
            assert_eq!(modulo, 91)
        }
    }

    mod test_offset_bezout {

        use super::*;
        #[test]
        fn test_7_13_offset_1_step_1() {
            let s = offset_bezout((7, 1), (13, 2));
            assert_eq!(s.0, 76);
            assert_eq!(s.1, 91);
        }

        #[test]
        fn test_221_13_offset_1_step() {
            let a0 = (221, -77);
            let a1 = (59, 2);
            let s = offset_bezout(a0, a1);

            //    n + a0_1 mod a0_0 = 0
            assert_eq!(0, (s.0 + a0.1) % a0.0);
            //    n + a1_1 mod a1_0 = 0
            assert_eq!(0, (s.0 + a1.1) % a1.0);
        }
    }

    mod part_2 {
        use super::*;
        #[test]
        fn example_1() {
            // Solve
            // 0 = n + 0 mod 17
            // 0 = n + 2 mod 13
            let s = offset_bezout((17, 0), (13, 2));
            assert_eq!(s.0, 102);
            assert_eq!(s.1, 221);

            // Solve
            // 102 = n mod 221 => 0 = n - 102 mod 221
            //                    0 = n + 3   mod 19
            let s = offset_bezout((221, -102), (19, 3));
            assert_eq!(s.0, 3417);
            assert_eq!(s.1, 221 * 19);
        }

        #[test]
        fn example_2() {
            let s = offset_bezout((67, 0), (7, 1));
            let s = offset_bezout((s.1, -s.0), (59, 2));
            let s = offset_bezout((s.1, -s.0), (61, 3));
            assert_eq!(s.0, 754018);
        }

        #[test]
        fn example_3() {
            let schedule = parse_buses("67,x,7,59,61");
            let mut redux = schedule[0];
            let mut next = redux;

            for bus in schedule.iter().skip(1) {
                redux = offset_bezout(next, *bus);
                next = (redux.1, -redux.0);
            }

            assert_eq!(redux.0, 779210);
        }
    }
}
