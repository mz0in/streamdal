use std::str::from_utf8;
use protos::sp_wsm::{WASMExitCode, WASMRequest};
use infers_jsonschema::infer;
use serde_json::json;

#[no_mangle]
pub extern "C" fn f(ptr: *mut u8, length: usize) -> *mut u8 {
    // Read request
    let wasm_request = match common::read_request(ptr, length, false) {
        Ok(req) => req,
        Err(e) => {
            return common::write_response(
                None,
                None,
                WASMExitCode::WASM_EXIT_CODE_INTERNAL_ERROR,
                e.to_string(),
            );
        }
    };

    // Validate request
    if let Err(err) = validate_wasm_request(&wasm_request) {
        return common::write_response(
            None,
            None,
            WASMExitCode::WASM_EXIT_CODE_INTERNAL_ERROR,
            format!("invalid step: {}", err.to_string()),
        );
    }


    println!("input payload: {}", from_utf8(wasm_request.input_payload.as_slice()).unwrap());
    let input_payload = match serde_json::from_slice(wasm_request.input_payload.as_slice()) {
        Ok(v) => v,
        Err(e) => {
            return common::write_response(
                None,
                None,
                WASMExitCode::WASM_EXIT_CODE_INTERNAL_ERROR,
                format!("invalid json: {}", e),
            );
        },
    };

    let parsed_json = json!(input_payload);

    println!("parsed_json: {}", parsed_json.to_string());
    let schema = infer(&parsed_json);
    println!("schema: {}", schema.to_string());

    // If no previous schema is provided, infer and return the generated schema
    if wasm_request.step.infer_schema().current_schema.is_empty() {
        return common::write_response(
            Some(wasm_request.input_payload.as_slice()),
            Some(schema.to_string().as_bytes()),
            WASMExitCode::WASM_EXIT_CODE_SUCCESS,
            "inferred fresh schema".to_string(),
        );
    }

    let current_schema = serde_json::to_value(&wasm_request.step.infer_schema().current_schema).unwrap();

    // Perform diff
    let diff = serde_json_diff::values(
        schema.clone(),
        current_schema,
    );

    if diff.is_some() {
        return common::write_response(
            Some(wasm_request.input_payload.as_slice()),
            Some(schema.clone().to_string().as_bytes()),
            WASMExitCode::WASM_EXIT_CODE_FAILURE,
            "schema has changed".to_string(),
        )
    }

    return common::write_response(
        Some(wasm_request.input_payload.as_slice()),
        None,
        WASMExitCode::WASM_EXIT_CODE_FAILURE,
        "schema has not changed".to_string(),
    )
}

fn validate_wasm_request(req: &WASMRequest) -> Result<(), String> {
    if !req.step.has_infer_schema() {
        return Err("httprequest is required".to_string());
    }

    if req.input_payload.is_empty() {
        return Err("payload cannot be empty".to_string());
    }

    Ok(())
}

/// # Safety
///
/// This is unsafe because it operates on raw memory; see `common/src/lib.rs`.
#[no_mangle]
pub unsafe extern "C" fn alloc(size: i32) -> *mut u8 {
    common::alloc(size)
}

/// # Safety
///
/// This is unsafe because it operates on raw memory; see `common/src/lib.rs`.
#[no_mangle]
pub unsafe extern "C" fn dealloc(pointer: *mut u8, size: i32) {
    common::dealloc(pointer, size)
}
