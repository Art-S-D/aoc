use std::error::Error;
use std::fs::File;
use std::io::{self, BufRead};

fn main() -> Result<(), Box<dyn Error>> {
    let file = File::open("input.txt")?;

    let mut elves: Vec<u32> = Vec::new();
    let mut current_calories = 0;

    for line in io::BufReader::new(file).lines() {
        let line = line?;
        if line == "" {
            elves.push(current_calories);
            current_calories = 0;
        } else {
            let calories: u32 = line.parse()?;
            current_calories += calories;
        }
    }

    elves.sort();

    let top_elves: u32 = elves.iter().rev().take(3).sum();

    println!("{}", top_elves);
    Ok(())
}
