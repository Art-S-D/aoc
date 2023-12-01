use std::convert::TryFrom;
use std::error::Error;
use std::fmt::{self, Display};
use std::fs::File;
use std::io::{self, BufRead};

#[derive(Debug, PartialEq, Eq, Copy, Clone)]
enum Hand {
    Rock,
    Paper,
    Scissors,
}
#[derive(Debug)]
struct UnknownHand {
    hand: String,
}
impl Display for UnknownHand {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:?}", self)
    }
}
impl Error for UnknownHand {}
impl TryFrom<&str> for Hand {
    type Error = UnknownHand;
    fn try_from(s: &str) -> Result<Hand, UnknownHand> {
        match s {
            "A" => Ok(Hand::Rock),
            "B" => Ok(Hand::Paper),
            "C" => Ok(Hand::Scissors),
            // "X" => Ok(Hand::Rock),
            // "Y" => Ok(Hand::Paper),
            // "Z" => Ok(Hand::Scissors),
            s => Err(UnknownHand { hand: s.into() }),
        }
    }
}
impl TryFrom<(Hand, &str)> for Hand {
    type Error = UnknownHand;
    fn try_from(input: (Hand, &str)) -> Result<Hand, UnknownHand> {
        let (opponent_hand, player_move) = input;
        match player_move {
            "X" => Ok(opponent_hand.beats()),
            "Y" => Ok(opponent_hand),
            "Z" => Ok(opponent_hand.beated_by()),
            s => Err(UnknownHand { hand: s.into() }),
        }
    }
}

impl Hand {
    fn score(&self) -> u32 {
        match self {
            Hand::Rock => 1,
            Hand::Paper => 2,
            Hand::Scissors => 3,
        }
    }
    fn round(h1: &Hand, h2: &Hand) -> (u32, u32) {
        use Hand::*;
        match (h1, h2) {
            // tie
            (Rock, Rock) | (Paper, Paper) | (Scissors, Scissors) => {
                (h1.score() + 3, h2.score() + 3)
            }
            // h1 won
            (Rock, Scissors) | (Scissors, Paper) | (Paper, Rock) => (h1.score() + 6, h2.score()),
            // h2 won
            (Rock, Paper) | (Paper, Scissors) | (Scissors, Rock) => (h1.score(), h2.score() + 6),
        }
    }
    fn beats(&self) -> Hand {
        use Hand::*;
        match self {
            Rock => Scissors,
            Paper => Rock,
            Scissors => Paper,
        }
    }
    fn beated_by(&self) -> Hand {
        use Hand::*;
        match self {
            Rock => Paper,
            Paper => Scissors,
            Scissors => Rock,
        }
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    let file = File::open("input.txt")?;

    let mut total_player_score = 0;

    for line in io::BufReader::new(file).lines() {
        let line = line?;
        let mut splitted = line.split(" ");
        let opponent_hand: Hand = splitted.next().unwrap_or("").try_into()?;
        let player_strategy = splitted.next().unwrap_or("");
        let player_hand: Hand = (opponent_hand, player_strategy).try_into()?;

        let (_, player_score) = Hand::round(&opponent_hand, &player_hand);
        total_player_score += player_score;
    }

    println!("{}", total_player_score);
    Ok(())
}
