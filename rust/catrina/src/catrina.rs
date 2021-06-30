use std::fmt::Error;

mod project;
mod utils;
mod wizard;

const DEFAULT_PORT: &str = ":9095";
const CONFIG_FILE: &str = "catrina.config.json";

const START_COMMAND: &str = "new";
const RUN_SERVER_COMMAND: &str = "run";
const BUILD_COMMAND: &str = "build";
const GET_LIB_VERSION_COMMAND: &str = "get";

pub fn catrina_tool(order: &str, flag: &str, flag_value: &str) -> Result<String, Error> {
    let s = String::from("catrina is running...");
    Ok(s)
}