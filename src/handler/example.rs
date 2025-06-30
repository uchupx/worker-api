use actix_web::{web, HttpResponse, Responder};
use serde_json::json;

pub async fn handle_request() -> impl Responder {     
    let data = json!({ "message": "Hello, World!" });
    HttpResponse::Ok().json(data) 
}