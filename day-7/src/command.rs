use std::fmt;

pub enum Dirname {
    Parent,                 // cd ..
    Child { name: String }, // cd somefolder
}
impl Dirname {
    fn from_str(str: String) -> Dirname {
        let dir = &str[5..];
        if dir == ".." {
            Dirname::Parent
        } else {
            Dirname::Child { name: dir.into() }
        }
    }
}
impl fmt::Display for Dirname {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Dirname::Parent => write!(f, ".."),
            Dirname::Child { name } => write!(f, "{}", name),
        }
    }
}

pub enum FileType {
    Dir { name: String },
    File { size: usize, name: String },
}
impl FileType {
    fn from_str(str: String) -> FileType {
        let splitted: Vec<&str> = str.split(' ').collect();
        if splitted[0] == "dir" {
            FileType::Dir {
                name: splitted[1].into(),
            }
        } else {
            FileType::File {
                size: splitted[0].parse().unwrap(),
                name: splitted[1].into(),
            }
        }
    }
}
impl fmt::Display for FileType {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            FileType::Dir { name } => write!(f, "dir {}", name),
            FileType::File { size, name } => write!(f, "{} {}", size, name),
        }
    }
}

pub enum ParsedCommand {
    Cd(Dirname),
    Ls(Vec<FileType>),
}
impl fmt::Display for ParsedCommand {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            ParsedCommand::Cd(name) => write!(f, "$ cd {}\n", name),
            ParsedCommand::Ls(files) => {
                write!(f, "$ ls\n")?;
                for file in files {
                    write!(f, "{}\n", file)?;
                }
                Ok(())
            }
        }
    }
}

pub fn parse_commands(input: impl Iterator<Item = String>) -> Vec<ParsedCommand> {
    let mut res = Vec::new();

    for line in input {
        match &line[0..4] {
            "$ cd" => res.push(ParsedCommand::Cd(Dirname::from_str(line))),
            "$ ls" => res.push(ParsedCommand::Ls(Vec::new())),
            _ => match res.last_mut() {
                None => panic!("file displayed with no commands"),
                Some(ParsedCommand::Cd(_)) => panic!("file displayed after a cd"),
                Some(ParsedCommand::Ls(ref mut v)) => v.push(FileType::from_str(line)),
            },
        }
    }

    res
}
