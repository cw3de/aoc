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
struct Action {
    increase: bool,
    value: i64,
}

fn read_actions(filename: &str) -> Vec<Action> {
    let mut actions: Vec<Action> = Vec::new();

    let lines = read_lines(filename);
    for line in lines {
        if line.starts_with('R') {
            let value = line[1..].parse().unwrap();
            actions.push(Action {
                increase: true,
                value: value,
            });
        }
        if line.starts_with('L') {
            let value = line[1..].parse().unwrap();
            actions.push(Action {
                increase: false,
                value: value,
            });
        }
    }
    actions
}

const WHEEL_SIZE: i64 = 100;

fn task1(filename: &str) {
    let mut result = 0;
    let mut position: i64 = 50;

    let actions = read_actions(filename);
    for action in actions {
        // println!("Inc: {}, Value {}", action.increase, action.value);

        match action.increase {
            true => {
                position = (position + action.value) % WHEEL_SIZE;
            }
            false => {
                position = (position + WHEEL_SIZE - action.value) % WHEEL_SIZE;
            }
        }
        if position == 0 {
            result += 1;
        }
    }
    println!("task 1 result: {}", result)
}

fn task2(filename: &str) {
    let mut result: i64 = 0;
    let mut position: i64 = 50;

    let actions = read_actions(filename);
    for action in actions {
        // println!("Inc: {}, Value {}", action.increase, action.value);

        let mut val = action.value;

        // count full turns
        while val >= WHEEL_SIZE {
            val -= WHEEL_SIZE;
            result += 1;
        }

        match action.increase {
            true => {
                if position + val < WHEEL_SIZE {
                    position += val; // no zero click
                } else {
                    position = (position + val) % WHEEL_SIZE;
                    result += 1;
                }
            }
            false => {
                if position == 0 {
                    position = WHEEL_SIZE - val;
                    // no zero click
                } else if val < position {
                    position -= val; // no zero click
                } else {
                    position = (position + WHEEL_SIZE - val) % WHEEL_SIZE;
                    result += 1;
                }
            }
        }
        // println!(
        //     "{} {}: {} {}",
        //     action.increase, action.value, position, result
        // )
    }

    println!("task 2 result: {}", result)
}
fn main() {
    task1("input.txt");
    task2("input.txt");
}
