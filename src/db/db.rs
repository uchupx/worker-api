use mongodb::{Client, Database};
use mongodb::bson::doc;
use std::env;
pub async fn connect_to_db() -> Result<Database, mongodb::error::Error> {
    let db_url = env::var("MONGODB_URL").unwrap_or_else(|_| "mongodb://localhost:27017".to_string());
    let client = Client::with_uri_str(&db_url).await?;
    client
        .database("admin")
        .run_command(doc! { "ping": 1 }).await?;
    println!("Connected to MongoDB");   
    
    let db_name = env::var("MONGODB_DB_NAME").unwrap_or_else(|_| "my_database".to_string());
    Ok(client.database(&db_name))
}
