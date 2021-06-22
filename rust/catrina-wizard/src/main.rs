use std::io::{stdin, stdout, Write, BufWriter, BufReader, Read};
use std::fs::File;
use std::env;
use std::path::PathBuf;

extern crate serde;
extern crate serde_json;

const DEFAULT_PORT: &str = ":9095";
const CONFIG_FILE: &str = "catrina.config.json";
#[macro_use] extern crate serde_derive;

#[derive(Serialize, Deserialize, Debug)]
struct Config {
    pub inputFileJs: String,
    pub inputFileCSS: String,
    pub deployPath: String,
    pub finalFileJS: String,
    pub finalFileCSS: String,
    pub serverPort: String,

}

impl Config {
    pub fn create_file(&self) {
        let data = serde_json::to_string_pretty(&self).unwrap();
        let file = File::create(file_from_location(CONFIG_FILE)).expect("Error creating config file");
        BufWriter::new(file)
            .write_all(data.as_bytes())
            .expect("Error writing config file");
    }
}

fn standard_config() -> Config {
    Config{
        inputFileJs: "./input.js".to_string(),
        inputFileCSS: "./input.css".to_string(),
        deployPath: "./deploy".to_string(),
        finalFileJS: "main.js".to_string(),
        finalFileCSS: "styles.css".to_string(),
        serverPort: DEFAULT_PORT.to_string()
    }
}

fn read_user_response() -> String {
    let mut user_response = String::new();
    let _  =stdout().flush();
    stdin().read_line(&mut user_response).expect("Error reading user input");
    user_response.trim().to_string()
}

fn print_config_file_result(config: &Config) {
    let data = serde_json::to_string_pretty(config).unwrap();
    let file = File::create(CONFIG_FILE).expect("Error creating config file");
    BufWriter::new(file)
        .write_all(data.as_bytes())
        .expect("Error writing config file");
    your_file_config_content();
}

fn your_file_config_content() {
    let mut data = String::new();
    let reference = File::open(file_from_location(CONFIG_FILE)).expect("Error reading config file");
    let mut br = BufReader::new(reference);
    br.read_to_string(&mut data).expect("Error parsing data");
    println!("\nYour project file:\n{}", data);
    println!("You can edit this configuration in file {}", CONFIG_FILE);
}



fn file_from_location(file: &str) -> PathBuf {
    let mut p = env::current_dir().unwrap();
    p.push(file);
    p
}


fn setup_wizard() {
    const EXIT_MSJ: &str = "(type 'exit' to close)";
    const EXIT_ORDER: &str = "exit";
    let mut config = standard_config();
    let standard_config = standard_config();

    println!("Set deploy path:{}", EXIT_MSJ);
    config.deployPath = read_user_response();
    if config.deployPath == EXIT_ORDER {
        standard_config.create_file();
        your_file_config_content();
        return;
    }

    println!("Set final javascript filename:{}", EXIT_MSJ);
    config.finalFileJS = read_user_response();
    if config.finalFileJS == EXIT_ORDER {
        config.finalFileJS = standard_config.finalFileJS;
        print_config_file_result(&config);
        return;
    }

    println!("Set final css filename:{}", EXIT_MSJ);
    config.finalFileCSS = read_user_response();
    if config.finalFileCSS == EXIT_ORDER {
        config.finalFileCSS = standard_config.finalFileCSS;
        print_config_file_result(&config);
        return;
    }

    println!("Set path of input javascript filename:{}", EXIT_MSJ);
    config.inputFileJs = read_user_response();
    if config.inputFileJs == EXIT_ORDER {
        config.inputFileJs = standard_config.inputFileJs;
        print_config_file_result(&config);
        return;
    }

    println!("Set path of input css filename:{}", EXIT_MSJ);
    config.inputFileCSS = read_user_response();
    if config.inputFileCSS == EXIT_ORDER {
        config.inputFileCSS = standard_config.inputFileCSS;
        print_config_file_result(&config);
        return;
    }


    println!("Set port of trial server:{}", EXIT_MSJ);
    config.serverPort = read_user_response();
    if config.serverPort == EXIT_ORDER {
        config.serverPort = standard_config.serverPort;
        print_config_file_result(&config);
        return;
    }

    print_config_file_result(&config);
}

fn main() {
    setup_wizard();
}

