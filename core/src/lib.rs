use zxcvbn::zxcvbn;
use extism_pdk::{plugin_fn, FnResult};

#[plugin_fn]
pub fn strength_for_password(password: String) -> FnResult<String> {
    Ok(zxcvbn(&password, &[]).guesses_log10().to_string())
}