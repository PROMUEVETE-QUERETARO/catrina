use std::io::{stdout, Write, stdin};

pub fn read_user_response() -> String {
    let mut user_response = String::new();
    let _  =stdout().flush();
    stdin().read_line(&mut user_response).expect("Error reading user input");
    user_response.trim().to_string()
}
