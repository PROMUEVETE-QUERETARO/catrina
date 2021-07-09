use crate::catrina::lib::StdLib;
use crate::catrina::project::{auto_project, Project};
use crate::catrina::utils::{getwd, read_user_response};
use crate::catrina::wizard::run_wizard;
use eyre::Result as R;
use std::fs;
use std::fs::File;

mod lib;
mod project;
mod utils;
mod wizard;

const DEFAULT_PORT: &str = ":9095";
const CONFIG_FILE: &str = "catrina.config.json";
const VERSION_APP: &str = "v1.2.0";
const START_COMMAND: &str = "new";
const UPDATE_COMMAND: &str = "update";
const RUN_SERVER_COMMAND: &str = "run";
const BUILD_COMMAND: &str = "build";
const GET_LIB_VERSION_COMMAND: &str = "get";
const FLAG_SKIP: &str = "-s";

fn catrina_new(project_name: &str, flag: &str) {
    fs::create_dir(project_name).expect("Error creating project folder");
    let mut location = getwd();
    location.push(project_name);

    let std_lib = StdLib::new(VERSION_APP, location);
    match std_lib.get() {
        Ok(_x) => println!("The project has been created successfully!"),
        Err(e) => panic!("{:?}", e),
    }

    if flag == FLAG_SKIP {
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

fn project_from_location() -> R<Project> {
    let actual_path = getwd();
    let project_name = actual_path
        .file_name()
        .expect("Error reading current directory ");
    let project_name = project_name.to_str().expect("Error parsing directory name");

    let mut file_path = getwd();
    file_path.push(&CONFIG_FILE);

    let file = File::open(file_path)?;
    let project = Project::from(file, String::from(project_name))?;
    Ok(project)
}

fn catrina_update(flag: &str) -> R<()> {
    if flag == FLAG_SKIP {
        let project = project_from_location()?;
        project.update_lib()?;
        return Ok(());
    }

    println!("IMPORTANT! This command delete all additional libraries installed");
    println!("Do you want continue?(y/n)");
    if read_user_response() == "y" {
        let project = project_from_location()?;
        project.update_lib()?;
    }

    Ok(())
}

pub fn catrina_tool(args: Vec<String>) -> R<()> {
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
        &UPDATE_COMMAND => catrina_update(flag)?,
        _ => {
            println!("{}", &command);
        }
    }
    Ok(())
}
