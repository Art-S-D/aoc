use std::error::Error;
use std::fs::File;
use std::io::{BufRead, BufReader};

struct SectionRange(u32, u32);

impl SectionRange {
    #[allow(dead_code)]
    fn contains(&self, other: &SectionRange) -> bool {
        self.0 <= other.0 && self.1 >= other.1
    }
    fn overlaps(&self, other: &SectionRange) -> bool {
        !((self.0 < other.0 && self.1 < other.0) || (self.0 > other.1 && self.1 > other.1))
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    let file = File::open("input.txt")?;

    let mut contains_count = 0;

    for line in BufReader::new(file).lines() {
        let line = line?;
        let elves: Vec<SectionRange> = line
            .split(",")
            .map(|elf| {
                let mut range = elf.split("-");
                let min = range.next().unwrap().parse().unwrap();
                let max = range.next().unwrap().parse().unwrap();
                assert!(range.next().is_none());
                SectionRange(min, max)
            })
            .collect();
        assert!(elves.len() == 2);
        if elves[0].overlaps(&elves[1]) {
            contains_count += 1;
        }
    }

    println!("{}", contains_count);

    Ok(())
}
