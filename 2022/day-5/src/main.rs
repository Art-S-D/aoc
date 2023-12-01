use std::fmt;
use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;

type Stack = Vec<char>;

fn parse_stacks() -> Vec<Stack> {
    let file = File::open("stacks.txt").unwrap();

    let buf = BufReader::new(file);
    let lines = buf.lines().collect::<Vec<_>>();
    let mut reversed_lines = lines.iter().rev();

    let stack_count: usize = reversed_lines
        .next()
        .unwrap()
        .as_ref()
        .unwrap()
        .split(" ")
        .filter(|s| s.len() > 0)
        .map(|s| s.parse().unwrap())
        .reduce(|acc: usize, i| acc.max(i))
        .unwrap();

    let mut res = Vec::new();
    for _ in 0..stack_count {
        res.push(Vec::new());
    }

    for line in reversed_lines {
        let line = line.as_ref().unwrap();

        let mut chars = line.chars();
        chars.next();
        for (i, stack_crate) in chars.step_by(4).enumerate() {
            if stack_crate != ' ' {
                res[i].push(stack_crate);
            }
        }
    }

    return res;
}

struct Instruction {
    count: u32,
    from: usize,
    to: usize,
}
impl fmt::Display for Instruction {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "move {} from {} to {}", self.count, self.from, self.to)
    }
}
fn parse_instructions() -> Vec<Instruction> {
    let file = File::open("instructions.txt").unwrap();

    BufReader::new(file)
        .lines()
        .map(|line| {
            let line = line.unwrap();
            let splitted: Vec<&str> = line.split(' ').collect();
            Instruction {
                count: splitted[1].parse().unwrap(),
                from: splitted[3].parse().unwrap(),
                to: splitted[5].parse().unwrap(),
            }
        })
        .collect()
}

fn main() {
    let mut stacks = parse_stacks();

    let instructions = parse_instructions();

    for instruction in instructions {
        let moved: Vec<char> = {
            let from = &mut stacks[instruction.from - 1];
            from.drain((from.len() - instruction.count as usize)..(from.len()))
                .collect()
        };
        stacks[instruction.to - 1].extend(moved);
    }

    let res: String = stacks.iter().map(|stack| stack.last().unwrap()).collect();
    println!("{}", res);
}
