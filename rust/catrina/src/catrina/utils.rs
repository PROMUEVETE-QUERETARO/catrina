use std::io::{stdout, Write, stdin};
use std::path::{PathBuf, Path};
use std::env;

pub fn read_user_response() -> String {
    let mut user_response = String::new();
    let _  = stdout().flush();
    stdin().read_line(&mut user_response).expect("Error reading user input");
    user_response.trim().to_string()
}

pub fn bin_dir() -> PathBuf {
    let bin = env::current_exe().expect("Error reading binary path");
    let dir = bin.parent().unwrap();
    PathBuf::from(dir)
}

pub fn getwd() -> PathBuf {
    env::current_dir().expect("Error reading execution path ")
}