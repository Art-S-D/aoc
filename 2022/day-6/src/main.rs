use std::{collections::HashSet, fs};

fn start_of_packet_offset(s: &String, distincts: usize) -> usize {
    for i in distincts..s.len() {
        let set: HashSet<char> = s[(i - distincts)..i].chars().collect();
        if set.len() >= distincts {
            return i;
        }
    }

    return s.len();
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let start = start_of_packet_offset(&input, 14);

    // let started_input = &input[start..input.len()];
    println!("{}", start);
}
