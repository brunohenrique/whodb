use std::collections::HashMap;
use std::io::Read;
use std::io::Write;
use std::net::TcpListener;
use std::net::TcpStream;
use std::net::{IpAddr, Ipv4Addr, SocketAddr};
use std::time::Duration;

use std::str;

use whodb;

#[test]
fn ping_command_is_replyed_with_pong() {
    let listener = TcpListener::bind("127.0.0.1:3000").expect("Unable to bind to socket");

    let data = HashMap::new();
    let storage = whodb::server::Storage { data: data };
    let server = whodb::server::Server { listener, storage };
    std::thread::spawn(move || {
        server.run();
    });

    let mut buffer = [0; 4];
    let socket = SocketAddr::new(IpAddr::V4(Ipv4Addr::new(127, 0, 0, 1)), 3000);
    let mut client = TcpStream::connect_timeout(&socket, Duration::new(0, 1000))
        .expect("Unable to bind to socket");
    if client.write(b"PING").is_err() {
        panic!("Failed to write")
    }
    let mut cmd = "";
    if let Ok(_) = client.read(&mut buffer) {
        cmd = str::from_utf8(&buffer).unwrap();
    };

    assert_eq!(cmd, "PONG");
}
