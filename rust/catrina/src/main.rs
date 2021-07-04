use std::env;
use crate::catrina::catrina_tool;

mod catrina;
#[macro_use] extern crate serde_derive;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("Not enough arguments");
        // TODO print manual
        return;
    }

    let r = catrina_tool(args);
    match  r {
        Err(e) => panic!(e),
        _ => println!("No errors")
    }

}
