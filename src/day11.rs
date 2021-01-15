use crate::utils::load_input;
use std::fmt::Debug;
use std::fmt::Display;
use std::fmt::Formatter;
use std::iter::repeat;
use std::ops::Index;

#[derive(PartialEq, Copy, Clone)]
enum Position {
    FLOOR = 0,
    EMPTY = 1,
    OCCUPIED = 2,
}

type Pair = (usize, usize);

impl Debug for Position {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), std::fmt::Error> {
        match *self {
            Position::FLOOR => write!(f, "."),
            Position::EMPTY => write!(f, "L"),
            Position::OCCUPIED => write!(f, "#"),
        }
    }
}

impl Display for Position {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), std::fmt::Error> {
        write!(f, "{:?}", self)
    }
}

impl From<char> for Position {
    fn from(c: char) -> Self {
        match c {
            'L' => Position::EMPTY,
            '#' => Position::OCCUPIED,
            _ => Position::FLOOR,
        }
    }
}

#[derive(PartialEq, Clone, Debug)]
struct Layout {
    data: Box<Vec<Position>>,
    width: usize,
    height: usize,
}

impl Layout {
    fn from_input(filename: &str) -> Self {
        let input = load_input(&filename);
        let positions = input
            .iter()
            .map(|line| line.chars().map(Position::from))
            .flatten();

        let data: Box<Vec<Position>> = Box::new(positions.collect());
        let width = data.len() / input.len();

        return Layout {
            data,
            width,
            height: input.len(),
        };
    }

    fn is_inbounds(&self, d0: &i8, d1: &i8) -> bool {
        &0 <= d0 && d0 < &(self.width as i8) && &0 <= d1 && d1 < &(self.height as i8)
    }
}

impl Display for Layout {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), std::fmt::Error> {
        write!(f, "({:?})", self.data)
    }
}

impl Index<Pair> for Layout {
    type Output = Position;
    fn index(&self, p: Pair) -> &<Self as Index<Pair>>::Output {
        let idx = p.1 * self.width + p.0;
        if idx >= self.data.len() {
            panic!("NO")
        } else {
            &self.data[idx]
        }
    }
}

struct LayoutIter<T>
where
    T: OccupiedCounter,
{
    layout: Layout,
    counter: T,
}

trait OccupiedCounter {
    fn count(&self, pair: &Pair, layout: &Layout) -> u8;
    fn next_for(&self, position: &Position, count: u8) -> Position;
}

impl<T> LayoutIter<T>
where
    T: OccupiedCounter,
{
    fn new(layout: Layout, counter: T) -> Self {
        return LayoutIter {
            layout: layout,
            counter,
        };
    }

    fn stable_occupied(self) -> usize {
        let last = self.last().unwrap();
        last.data
            .iter()
            .filter(|p| *p == &Position::OCCUPIED)
            .count()
    }

    fn count_for(&self, pair: &Pair) -> u8 {
        self.counter.count(pair, &self.layout)
    }

    fn next_for(&self, pair: &Pair, position: &Position) -> Position {
        let count = self.count_for(pair);
        self.counter.next_for(position, count)
    }
}

impl<T> Iterator for LayoutIter<T>
where
    T: OccupiedCounter,
{
    type Item = Layout;
    fn next(&mut self) -> Option<<Self as Iterator>::Item> {
        let next_data = Box::new(
            self.layout
                .data
                .iter()
                .enumerate()
                .map(|(idx, position)| {
                    let pair = (idx % self.layout.width, idx / self.layout.width);
                    let next = self.next_for(&pair, position);
                    return next;
                })
                .collect(),
        );

        if next_data == self.layout.data {
            None
        } else {
            self.layout = Layout {
                data: next_data,
                width: self.layout.width,
                height: self.layout.height,
            };
            Some(self.layout.clone())
        }
    }
}

struct AdjacentOccupied;
impl OccupiedCounter for AdjacentOccupied {
    fn count(&self, pair: &Pair, layout: &Layout) -> u8 {
        let adjacents = (-1..=1)
            .flat_map(|i| repeat(i).zip(-1..=1))
            .filter(|(d0, d1)| !(*d0 == 0 && *d1 == 0))
            .map(|(d0, d1)| (pair.0 as i8 + d0, pair.1 as i8 + d1))
            .filter(|(d0, d1)| layout.is_inbounds(d0, d1))
            .filter(|(c0, c1)| layout[(*c0 as usize, *c1 as usize)] == Position::OCCUPIED)
            .count() as u8;

        adjacents
    }
    fn next_for(&self, position: &Position, count: u8) -> Position {
        if *position == Position::EMPTY && count == 0 {
            Position::OCCUPIED
        } else if *position == Position::OCCUPIED && count >= 4 {
            Position::EMPTY
        } else {
            *position
        }
    }
}

struct VisibleOccupied;
impl OccupiedCounter for VisibleOccupied {
    fn count(&self, pair: &Pair, layout: &Layout) -> u8 {
        let mut count: u8 = 0;
        let (p0, p1) = (pair.0 as i8, pair.1 as i8);

        for (v0, v1) in (-1..=1)
            .flat_map(|i| repeat(i).zip(-1..=1))
            .filter(|(d0, d1)| !(*d0 == 0 && *d1 == 0))
        {
            for k in 1.. {
                let (d0, d1) = ((v0 * k + p0), (v1 * k + p1));
                if !layout.is_inbounds(&d0, &d1) {
                    break;
                }

                match layout[(d0 as usize, d1 as usize)] {
                    Position::FLOOR => {}
                    Position::EMPTY => {
                        break;
                    }
                    Position::OCCUPIED => {
                        count += 1;
                        break;
                    }
                };
            }
        }
        count
    }

    fn next_for(&self, position: &Position, count: u8) -> Position {
        if *position == Position::EMPTY && count == 0 {
            Position::OCCUPIED
        } else if *position == Position::OCCUPIED && count >= 5 {
            Position::EMPTY
        } else {
            *position
        }
    }
}

pub fn solution_1() -> usize {
    let layout = Layout::from_input("input/11.txt");
    let layout_seq = LayoutIter::new(layout, AdjacentOccupied);
    layout_seq.stable_occupied()
}

pub fn solution_2() -> usize {
    let layout = Layout::from_input("input/11.txt");
    let layout_seq = LayoutIter::new(layout, VisibleOccupied);
    layout_seq.stable_occupied()
}

#[cfg(test)]
mod test {
    use super::{AdjacentOccupied, Layout, LayoutIter, VisibleOccupied};
    #[cfg(test)]
    mod adjacent_occupied_seq {
        use super::{AdjacentOccupied, Layout, LayoutIter};
        #[test]
        fn single_step() {
            let layout = Layout::from_input("input/11ex1.txt");
            let mut layout_seq = LayoutIter::new(layout, AdjacentOccupied);
            layout_seq.next();
            assert_eq!(5, layout_seq.count_for(&(1, 0)));
        }
        #[test]
        fn stable() {
            let layout = Layout::from_input("input/11ex1.txt");
            let layout_seq = LayoutIter::new(layout, AdjacentOccupied);
            let stable_occupied = layout_seq.stable_occupied();
            assert_eq!(37, stable_occupied);
        }
    }
    #[cfg(test)]
    mod visible_occupied_seq {
        use super::*;
        #[test]
        fn stable() {
            let layout = Layout::from_input("input/11ex1.txt");
            let layout_seq = LayoutIter::new(layout, VisibleOccupied);
            let stable_occupied = layout_seq.stable_occupied();
            assert_eq!(26, stable_occupied);
        }
    }
}
