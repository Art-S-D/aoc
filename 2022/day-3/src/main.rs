use core::panic;
use itertools::Itertools;
use std::collections::HashSet;
use std::error::Error;
use std::fs::File;
use std::io;
use std::io::BufRead;

fn priority(c: char) -> u8 {
    if c >= 'a' && c <= 'z' {
        c as u8 - 'a' as u8 + 1
    } else {
        c as u8 - 'A' as u8 + 27
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    let file = File::open("input.txt")?;

    let all_possible_items: HashSet<char> = ('a'..='z').chain('A'..='Z').collect();

    let badges_sum: u32 = io::BufReader::new(file)
        .lines()
        .chunks(3)
        .into_iter()
        // group by intersection
        .map(|group| {
            group
                .into_iter()
                .fold(all_possible_items.clone(), |acc, rucksack| {
                    let rucksack = rucksack.unwrap();
                    let sack_items: HashSet<char> = rucksack.chars().collect();

                    let intersection: HashSet<char> =
                        acc.intersection(&sack_items).map(|c| *c).collect();

                    intersection
                })
        })
        // get badge priority
        .map(|intersection| {
            if intersection.len() != 1 {
                panic!("too many or not enough common items {:?}", intersection);
            }
            let badge = intersection.iter().next().unwrap();
            priority(*badge) as u32
        })
        .sum();
    println!("{}", badges_sum);

    // let line = line?;
    // let compartment1 = &line[..line.len() / 2];
    // let compartment2 = &line[line.len() / 2..];

    // let compartment1_set = compartment1.chars().collect::<HashSet<char>>();
    // let compartment2_set = compartment2.chars().collect::<HashSet<char>>();

    // let common_items = compartment1_set.intersection(&compartment2_set);
    // for item in common_items {
    //     sum += priority(*item) as u32;
    // }

    Ok(())
}
