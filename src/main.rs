mod server;

use std::collections::HashMap;
use std::net::TcpListener;

fn main() {
    let listener = TcpListener::bind("127.0.0.1:3000").expect("Unable to bind to socket");

    let data = HashMap::new();
    let storage = whodb::server::Storage { data };
    let server = whodb::server::Server { listener, storage };

    server.run()
}
