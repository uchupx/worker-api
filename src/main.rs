use std::io;
use dotenv::dotenv;
use actix_web::{web, App, HttpServer};

mod handler;
mod db;
mod schema;
mod grpc;

#[actix_rt::main]
async fn main() -> io::Result<()> {
    dotenv().ok();
    
    let out_dir = env::var("OUT_DIR").unwrap_or_else(|_| "target/debug/build".to_string());
    println!("OUT_DIR: {}", out_dir);

    let _ =env::set_var("RUST_LOG", "info");

    env::set_var("OUT_DIR", out_dir);
    env_logger::init();

    db::db::connect_to_db().await.expect("Failed to connect to database");
    HttpServer::new(|| {
        App::new()
            .service(web::resource("/api/hello").to(handler::example::handle_request))
    })
        .bind("0.0.0.0:9091")?
        .run()
        .await

}
