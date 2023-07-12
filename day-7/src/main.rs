use std::collections::HashMap;
use std::fs;
use std::io::{BufRead, BufReader};

mod command;

#[derive(Debug)]
struct File {
    size: usize,
    name: String,
}

#[derive(Default, Debug)]
struct Dir {
    entries: HashMap<String, FsEntry>,
}
impl Dir {
    fn size(&self) -> usize {
        self.entries.values().map(|ent| ent.size()).sum()
    }
}

#[derive(Debug)]
enum FsEntry {
    FileEnt(File),
    DirEnt(Box<Dir>),
}
impl FsEntry {
    fn pretty_print(&self, indent: u32) {
        match self {
            FsEntry::FileEnt(file) => {
                println!(
                    "{}- {} {}",
                    (0..indent).map(|_| " ").collect::<String>(),
                    file.size,
                    file.name
                );
            }
            FsEntry::DirEnt(dir) => {
                for (key, value) in dir.entries.iter() {
                    println!(
                        "{}- dir {:?}",
                        (0..indent).map(|_| " ").collect::<String>(),
                        key
                    );
                    value.pretty_print(indent + 4)
                }
            }
        }
    }
    fn size(&self) -> usize {
        match self {
            FsEntry::FileEnt(file) => file.size,
            FsEntry::DirEnt(dir) => dir.size(),
        }
    }
    fn size_under(&self, max_size: usize) -> usize {
        match self {
            FsEntry::FileEnt(_) => 0,
            FsEntry::DirEnt(dir) => {
                let sum = dir
                    .entries
                    .values()
                    .map(|ent| ent.size_under(max_size))
                    .sum();
                let self_size = self.size();
                if self_size <= max_size {
                    sum + self_size
                } else {
                    sum
                }
            }
        }
    }
    fn smallest_dir_with_size(&self, min_size: usize) -> Option<&Dir> {
        match self {
            FsEntry::FileEnt(_) => None,
            FsEntry::DirEnt(dir) => {
                let mut sub_dirs: Vec<&Dir> = dir
                    .entries
                    .values()
                    .map(|v| v.smallest_dir_with_size(min_size))
                    .filter(|v| v.is_some())
                    .map(|v| v.unwrap())
                    .collect();

                if dir.size() >= min_size {
                    sub_dirs.push(dir);
                }

                sub_dirs.iter().min_by_key(|v| v.size()).copied()
            }
        }
    }
}

fn parse_fs(commands: &mut impl Iterator<Item = command::ParsedCommand>, current: &mut Dir) {
    use command::*;
    loop {
        match commands.next() {
            None => return,
            Some(ParsedCommand::Cd(Dirname::Parent)) => return,

            Some(ParsedCommand::Cd(Dirname::Child { name })) => {
                let mut child_dir = Dir {
                    entries: HashMap::new(),
                };
                parse_fs(commands, &mut child_dir);
                let child_entry = if current.entries.contains_key(&name) {
                    current.entries.get_mut(&name).unwrap()
                } else {
                    current
                        .entries
                        .insert(name.clone(), FsEntry::DirEnt(Box::new(Dir::default())));
                    current.entries.get_mut(&name).unwrap()
                };
                match child_entry {
                    FsEntry::DirEnt(dir) => dir.entries.extend(child_dir.entries),
                    FsEntry::FileEnt(_) => panic!("child dir already exists as a file"),
                };
            }

            Some(ParsedCommand::Ls(files)) => {
                for file in files {
                    match file {
                        FileType::Dir { name } => {
                            current
                                .entries
                                .entry(name)
                                .or_insert(FsEntry::DirEnt(Box::new(Dir::default())));
                        }

                        FileType::File { size, name } => {
                            current
                                .entries
                                .insert(name.clone(), FsEntry::FileEnt(File { size, name }));
                        }
                    };
                }
            }
        }
    }
}

fn main() {
    let input = fs::File::open("input.txt").unwrap();

    let mut reader = BufReader::new(input).lines().map(|l| l.unwrap());

    let mut commands = command::parse_commands(&mut reader);
    let mut root = Dir::default();
    parse_fs(&mut commands.drain(..), &mut root);

    let fs_root = FsEntry::DirEnt(Box::new(root));
    let size = fs_root.size();
    let unused_space = 70_000_000 - size;
    let extra_space_needed = 30_000_000 - unused_space;
    println!("total size {}", size);
    println!("need {} more bytes", extra_space_needed);
    let smaller_dir = fs_root.smallest_dir_with_size(extra_space_needed).unwrap();
    // println!("smallest dir with this size: {:?}", smaller_dir);
    println!("smallest removable dir size: {}", smaller_dir.size());
}
