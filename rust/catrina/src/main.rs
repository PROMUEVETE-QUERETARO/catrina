use std::env;
use crate::catrina::catrina_tool;

mod catrina;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("Not enough arguments");
        // TODO print manual
        return;
    }

    let mut command = "";
    let mut flag = "";
    let mut flag_value = "";

    match args.get(1) {
        Some(x) => command = x,
        _ => {
            println!("Error with arguments");
            // TODO print manual
            return;
        }
    }

    match args.get(2) {
        Some(x) => flag = x,
        _ => {
            println!("Error with arguments");
            // TODO print manual
            return;
        }
    }

    match args.get(3) {
        Some(x) => flag_value = x,
        _ => {
            println!("Error with arguments");
            // TODO print manual
            return;
        }
    }
    let r = catrina_tool(command, flag, flag_value);
    match  r {
        Err(e) => panic!(e),
        _ => panic!("Unknown error")
    }

}
