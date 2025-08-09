use axum::{Router, http::StatusCode, response, routing::get};

use serde_json::json;

#[tokio::main()]
async fn main() {
    let app = Router::new()
        .route(
            "/",
            get(|| async {
                return response::Json(json!({
                    "result": "Ok"
                }));
            }),
        )
        .route(
            "/db/",
            get(|| async {
                return (
                    StatusCode::NOT_IMPLEMENTED,
                    response::Json(json!({
                        "result": "Not implemented!"
                    })),
                );
            }),
        )
        .route(
            "/chaos/",
            get(|| async {
                return (
                    StatusCode::NOT_IMPLEMENTED,
                    response::Json(json!({
                        "result": "Not implemented!"
                    })),
                );
            }),
        )
        .route(
            "/health/",
            get(|| async {
                return response::Json(json!({
                    "result": "Ok"
                }));
            }),
        );

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    println!("listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}
