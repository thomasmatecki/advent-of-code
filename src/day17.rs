use crate::utils::load_input;
use std::cmp::max;
use std::cmp::min;
use std::collections::HashSet;
#[derive(Debug)]
struct PocketDimension<const N: usize> {
    space: HashSet<[i32; N]>,
}

struct DimensionIter<const N: usize> {
    spans: [(i32, i32); N],
    state: [i32; N],
}

impl<const N: usize> DimensionIter<N> {
    fn new(spans: [(i32, i32); N]) -> Self {
        let mut state = [0; N];
        for i in 0..N {
            state[i] = spans[i].0;
        }

        DimensionIter { spans, state }
    }

    fn from_space(space: &HashSet<[i32; N]>) -> Self {
        let mut spans = [(0, 0); N];

        for coord in space.iter() {
            for i in 0..N {
                spans[i] = (min(spans[i].0, coord[i]), max(spans[i].1, coord[i]))
            }
        }

        for i in 0..N {
            spans[i] = (spans[i].0 - 1, spans[i].1 + 1);
        }

        return Self::new(spans);
    }
}

impl<const N: usize> Iterator for DimensionIter<N> {
    type Item = [i32; N];

    fn next(&mut self) -> Option<<Self as Iterator>::Item> {
        // The _last_ dimension has left it's bound
        if self.state[N - 1] > self.spans[N - 1].1 {
            return None;
        }

        let current_state = self.state;

        for i in 0..N {
            self.state[i] += 1;
            if self.state[i] <= self.spans[i].1 {
                break;
            } else if i < N - 1 {
                self.state[i] = self.spans[i].0;
            }
        }

        return Some(current_state);
    }
}

impl<const N: usize> PocketDimension<N> {
    fn new() -> Self {
        return PocketDimension {
            space: HashSet::new(),
        };
    }

    fn from_input(input: &str) -> Self {
        let mut dim = PocketDimension::new();
        for (y, line) in load_input(input).iter().enumerate() {
            for (x, c) in line.chars().enumerate() {
                let mut val: [i32; N] = [0; N];
                val[0] = x as i32;
                val[1] = y as i32;
                if c == '#' {
                    dim.space.insert(val);
                }
            }
        }

        return dim;
    }

    fn span_iter(&self) -> Box<dyn Iterator<Item = [i32; N]>> {
        return Box::new(DimensionIter::from_space(&self.space));
    }

    fn next_active(&self, coord: [i32; N]) -> bool {
        let active = self.space.contains(&coord);

        let mut neighbor_span: [(i32, i32); N] = [(0, 0); N];
        for dim in 0..N {
            neighbor_span[dim] = (coord[dim] - 1, coord[dim] + 1);
        }

        let neighbors = DimensionIter::new(neighbor_span).filter(|neighbor| *neighbor != coord);
        let active_neighbors = neighbors
            .filter(|neighbor| self.space.contains(neighbor))
            .count();

        let state = if active {
            // exactly 2 or 3 of its neighbors are also active, the cube remains active
            active_neighbors == 2 || active_neighbors == 3
        } else {
            // exactly 3 of its neighbors are active, the cube becomes active
            active_neighbors == 3
        };

        return state;
    }

    fn tick(&mut self) {
        let mut space: HashSet<[i32; N]> = HashSet::new();
        for coord in self.span_iter() {
            if self.next_active(coord) {
                space.insert(coord);
            }
        }
        self.space = space;
    }
}

pub fn solution_1() -> u32 {
    let mut dim = PocketDimension::<3>::from_input("input/17.txt");
    for _ in 0..6 {
        dim.tick();
    }

    dim.space.len() as u32
}

pub fn solution_2() -> u32 {
    let mut dim = PocketDimension::<4>::from_input("input/17.txt");
    for _ in 0..6 {
        dim.tick();
    }

    dim.space.len() as u32
}
