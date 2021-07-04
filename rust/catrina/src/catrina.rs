use std::io::{Error, ErrorKind};
use crate::catrina::lib::StdLib;
use std::fs;
use crate::catrina::utils::{getwd, read_user_response};
use crate::catrina::wizard::{run_wizard};
use crate::catrina::project::auto_project;

mod project;
mod utils;
mod wizard;
mod lib;

const DEFAULT_PORT: &str = ":9095";
const CONFIG_FILE: &str = "catrina.config.json";
const VERSION_APP: &str = "v1.2.0";
const START_COMMAND: &str = "new";
const RUN_SERVER_COMMAND: &str = "run";
const BUILD_COMMAND: &str = "build";
const GET_LIB_VERSION_COMMAND: &str = "get";
const FLAG_SKIP: &str = "-s";

#[derive(Debug)]
pub struct CatrinaError {
    details: String
}

impl CatrinaError {

}

fn catrina_new(project_name: &str, flag: &str) {
    fs::create_dir(project_name).expect("Error creating project folder");
    let mut location = getwd();
    location.push(project_name);

    let std_lib  = StdLib::new(VERSION_APP, location);
    match std_lib.get() {
        Ok(x) => println!("The project has been created successfully!"),
        Err(e) => panic!("{:?}", e),
    }

    if flag == FLAG_SKIP  {
        auto_project(&project_name.to_string());
        return;
    }

    println!("Do you want to start the setup wizard?(y/n)");

    let r = read_user_response();
    if r == String::from("y") {
        run_wizard(&project_name.to_string())
    } else {
        auto_project(&project_name.to_string())
    }
}



pub fn catrina_tool(args: Vec<String>) -> Result<(), CatrinaError> {
    let mut command = "";
    let mut arg = "";
    let mut flag = "";
    let mut flag_value = "";

    match args.get(1) {
       Some(x) => command = x,
       _ => {
           println!("Error with arguments");
           // TODO print manual
           return Ok(());
       }
    }

    match args.get(2) {
        Some(x) => arg = x,
        _ => {
            println!("Error with arguments");
            // TODO print manual
            return Ok(());
        }
    }

    match args.get(3) {
        Some(x) => flag = x,
        _ => {}
    }

   match &command {
       &START_COMMAND => catrina_new(arg, flag),
       _=> {
           println!("{}", &command);
       }
   }
   Ok(())
}
