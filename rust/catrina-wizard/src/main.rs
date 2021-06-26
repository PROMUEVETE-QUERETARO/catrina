use std::io::{stdin, stdout, Write, BufWriter, BufReader, Read};
use std::fs::File;
use std::{env, fs};
use std::path::{Path, PathBuf};

extern crate serde;
extern crate serde_json;

const DEFAULT_PORT: &str = ":9095";
const CONFIG_FILE: &str = "catrina.config.json";

struct Project {
    pub config:Config,
    pub name: String,
}

impl Project {
    fn create_environment(&self) {
        fs::create_dir_all(&format!("{}/{}",&self.name, &self.config.deployPath));
        Project::create_input_file(&self.config.inputFileJs, &self.name);
        Project::create_input_file(&self.config.inputFileCSS, &self.name);
        File::create(&format!("{}/{}/{}",&self.name, &self.config.deployPath, &self.config.finalFileJS));
        File::create(&format!("{}/{}/{}",&self.name, &self.config.deployPath, &self.config.finalFileCSS));
    }

    fn create_input_file(file: &String, project: &String){
        let mut project_path = PathBuf::from(&project);
        let parent_file = Path::new(&file).parent().unwrap();

        if parent_file.to_str().unwrap().to_string() != String::from("") {
            project_path.push(parent_file);
            fs::create_dir_all(&project_path);

            let mut file_location = PathBuf::from(&project);
            file_location.push(file);
            println!("file location {:?}", &file_location);
            File::create(file_location);
            return;
        }

        project_path.push(file);
        File::create(project_path);
    }

    fn your_file_config_content(project: &String) {
        let mut data = String::new();
        let reference = File::open(&format!("{}/{}", project, CONFIG_FILE))
            .expect("Error reading config file");
        let mut br = BufReader::new(reference);
        br.read_to_string(&mut data).expect("Error parsing data");
        println!("\nYour project configuration:\n{}", data);
        println!("You can edit this configuration in file {}", CONFIG_FILE);
    }

    pub fn start(&self) {
        &self.config.create_file(&self.name);
        &self.create_environment();
        Project::your_file_config_content(&self.name);
    }
}

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

    pub fn create_file(&self, project: &String) {
        fs::create_dir_all(project);

        let data = serde_json::to_string_pretty(&self).unwrap();
        let file = File::create(&format!("{}/{}", project, CONFIG_FILE))
            .expect("Error creating config file");

        BufWriter::new(file)
            .write_all(data.as_bytes())
            .expect("Error writing config file");
    }


}

fn standard_config() -> Config {
    Config{
        inputFileJs: "input.js".to_string(),
        inputFileCSS: "input.css".to_string(),
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

fn setup_wizard(project_name: &String) {
    const EXIT_MSJ: &str = "(type 'exit' to close)";
    const EXIT_ORDER: &str = "exit";
    let mut project = Project{
        config: standard_config(),
        name: project_name.to_string()
    };

    let mut config = standard_config();
    let standard_config = standard_config();

    println!("Set deploy path:{}", EXIT_MSJ);
    project.config.deployPath = read_user_response();
    if project.config.deployPath == EXIT_ORDER {
        project.config.deployPath = standard_config.deployPath;
        project.start();
        return;
    }

    println!("Set final javascript filename:{}", EXIT_MSJ);
    project.config.finalFileJS = read_user_response();
    if project.config.finalFileJS == EXIT_ORDER {
        project.config.finalFileJS = standard_config.finalFileJS;
        project.start();
        return;
    }

    println!("Set final css filename:{}", EXIT_MSJ);
    project.config.finalFileCSS = read_user_response();
    if project.config.finalFileCSS == EXIT_ORDER {
        project.config.finalFileCSS = standard_config.finalFileCSS;
        project.start();
        return;
    }

    println!("Set path of input javascript filename:{}", EXIT_MSJ);
    project.config.inputFileJs = read_user_response();
    if project.config.inputFileJs == EXIT_ORDER {
        project.config.inputFileJs = standard_config.inputFileJs;
        project.start();
        return;
    }

    println!("Set path of input css filename:{}", EXIT_MSJ);
    project.config.inputFileCSS = read_user_response();
    if project.config.inputFileCSS == EXIT_ORDER {
        project.config.inputFileCSS = standard_config.inputFileCSS;
        project.start();
        return;
    }


    println!("Set port of trial server:{}", EXIT_MSJ);
    project.config.serverPort = read_user_response();
    if project.config.serverPort == EXIT_ORDER {
        project.config.serverPort = standard_config.serverPort;
    }

    project.start();
}

fn main() {
    let args: Vec<String> = env::args().collect();
    match args.get(1) {
        Some(x) => setup_wizard(x),
        _ => println!("Not enough arguments")
    }
}

