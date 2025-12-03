// use std::fs::File;
use std::fs::read_to_string;
// use std::io::BufRead;
// use std::io::BufReader;
// use std::num::ParseIntError;

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

#[derive(Debug)]
struct Range {
    start: i64,
    end: i64,
}

fn read_ranges(filename: &str) -> Vec<Range> {
    let mut ranges: Vec<Range> = Vec::new();

    let mut text = read_to_string(filename).unwrap();
    // remove new-line char at the end
    text.truncate(text.len() - 1);
    let parts = text.split(',');
    for part in parts {
        if let Some((a, b)) = part.split_once('-') {
            let va = a.parse::<i64>().unwrap();
            let vb = b.parse::<i64>().unwrap();
            ranges.push(Range { start: va, end: vb });
        }
    }
    ranges
}

fn check_bad_id(txt: &String, pattern_length: usize) -> bool {
    let full_length = txt.len();
    let str = txt.as_str();
    let pattern = &str[0..pattern_length];
    if full_length % pattern_length != 0 {
        return false;
    }
    for pos in (pattern_length..full_length).step_by(pattern_length) {
        if pattern != &str[pos..pos + pattern_length] {
            return false;
        }
    }
    return true;
}

fn task1(filename: &str, verbose: bool) {
    let mut result = 0;

    let ranges = read_ranges(filename);
    for range in ranges {
        if verbose {
            println!("Range from {} to {}", range.start, range.end);
        }
        for id in range.start..range.end + 1 {
            let txt = id.to_string();
            let txt_len = txt.len();
            if txt_len % 2 == 0 && check_bad_id(&txt, txt_len / 2) {
                if verbose {
                    println!("  {id} ({txt_len})");
                }
                result += id;
            }
        }
    }
    println!("task 1 result: {}", result)
}

fn task2(filename: &str, verbose: bool) {
    let mut result = 0;

    let ranges = read_ranges(filename);
    for range in ranges {
        if verbose {
            println!("Range from {} to {}", range.start, range.end);
        }
        for id in range.start..range.end + 1 {
            let txt = id.to_string();
            let half_len = txt.len() / 2;
            for pattern_length in 1..half_len + 1 {
                if check_bad_id(&txt, pattern_length) {
                    if verbose {
                        println!("  {id} ({pattern_length})");
                    }
                    result += id;
                    break;
                }
            }
        }
    }
    println!("task 2 result: {}", result)
}

fn main() {
    let verbose = false;
    task1("input.txt", verbose);
    task2("input.txt", verbose);
}
