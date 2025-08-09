use std::{fs, time};

use rand::random;

use axum::{Router, extract::State, http::StatusCode, response, routing::get};
use serde_json::json;

#[derive(Clone)]
struct AppState {
    hostname: String,
}

#[tokio::main()]
async fn main() {
    let hostname = fs::read_to_string("/etc/hostname")
        .unwrap()
        .trim()
        .to_string();

    let app_state = AppState { hostname: hostname };

    let app = Router::new()
        .route(
            "/",
            get(|State(state): State<AppState>| async move {
                return response::Json(json!({
                    "result": "Ok",
                    "hostname": state.hostname,
                }));
            }),
        )
        .route(
            "/db/",
            get(|State(state): State<AppState>| async move {
                return (
                    StatusCode::NOT_IMPLEMENTED,
                    response::Json(json!({
                        "result": "Not implemented!",
                        "hostname": state.hostname,
                    })),
                );
            }),
        )
        .route(
            "/chaos/",
            get(|State(state): State<AppState>| async move {
                return (
                    StatusCode::NOT_IMPLEMENTED,
                    response::Json(json!({
                        "result": "Not implemented!",
                        "hostname": state.hostname,
                    })),
                );
            }),
        )
        .route(
            "/health/",
            get(|State(state): State<AppState>| async move {
                return response::Json(json!({
                    "result": "Ok",
                    "hostname": state.hostname,
                }));
            }),
        )
        .route(
            "/chaos-server/",
            get(|State(state): State<AppState>| async move {
                let sleep_time = time::Duration::from_millis(random::<u64>() % 100 + 1);

                tokio::time::sleep(sleep_time).await;

                return response::Json(json!({
                    "result": "Ok",
                    "hostname": state.hostname,
                    "sleep_time": sleep_time.as_millis(),
                }));
            }),
        )
        .with_state(app_state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    println!("listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}
