use crate::utils::load_input;
use std::cmp::max;
use std::cmp::min;
use std::collections::HashSet;
use std::iter::repeat;

#[derive(Debug)]
struct PocketDimension {
    space: HashSet<(i32, i32, i32)>,
}

impl PocketDimension {
    fn new() -> Self {
        return PocketDimension {
            space: HashSet::new(),
        };
    }

    fn from_input(input: &str) -> Self {
        let mut dim = PocketDimension::new();
        for (y, line) in load_input(input).iter().enumerate() {
            for (x, c) in line.chars().enumerate() {
                if c == '#' {
                    dim.space.insert((x as i32, y as i32, 0));
                }
            }
        }

        return dim;
    }

    fn get(&self, x: i32, y: i32, z: i32) -> bool {
        match self.space.get(&(x, y, z)) {
            None => false,
            Some(v) => true,
        }
    }

    fn xyz_ranges(&self) -> Box<dyn Iterator<Item = (i32, i32, i32)>> {
        let mut xs = (0, 1);
        let mut ys = (0, 1);
        let mut zs = (0, 1);
        for (x, y, z) in self.space.iter() {
            xs = (min(xs.0, *x), max(xs.1, *x));
            ys = (min(ys.0, *y), max(ys.1, *y));
            zs = (min(zs.0, *z), max(zs.1, *z));
        }

        xs = (xs.0 - 1, xs.1 + 1);
        ys = (ys.0 - 1, ys.1 + 1);
        zs = (zs.0 - 1, zs.1 + 1);

        return Box::new(
            (xs.0..=xs.1)
                .flat_map(move |x| repeat(x).zip(ys.0..=ys.1))
                .flat_map(move |y| repeat(y).zip(zs.0..=zs.1))
                .map(|((x, y), z)| (x, y, z)),
        );
    }

    fn next_active(&self, x: i32, y: i32, z: i32) -> bool {
        // 5 * 5 * 3
        let mut count = 0;
        let active = self.space.contains(&(x, y, z));

        for xf in -1..=1 {
            for yf in -1..=1 {
                for zf in -1..=1 {
                    if !(xf == 0 && yf == 0 && zf == 0) {
                        if self.space.contains(&(x + xf, y + yf, z + zf)) {
                            count += 1;
                        };
                    }
                }
            }
        }
        let state = if active {
            // exactly 2 or 3 of its neighbors are also active, the cube remains active
            count == 2 || count == 3
        } else {
            // exactly 3 of its neighbors are active, the cube becomes active
            count == 3
        };

        return state;
    }

    fn tick(&mut self) {
        let mut space: HashSet<(i32, i32, i32)> = HashSet::new();
        for (x, y, z) in self.xyz_ranges() {
            if self.next_active(x, y, z) {
                space.insert((x, y, z));
            }
        }

        self.space = space;
    }
}

pub fn solution_1() -> u32 {
    let mut dim = PocketDimension::from_input("input/17.txt");
    for _ in 0..6 {
        dim.tick();
    }

    dim.space.len() as u32
}
