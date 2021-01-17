use std::collections::HashMap;
use std::io::Read;
use std::io::Write;
use std::net::{TcpListener, TcpStream};
use std::str;
use std::thread;

pub struct Storage<'a> {
    pub data: HashMap<&'a str, &'a str>,
}

pub struct Server<'a> {
    pub listener: TcpListener,
    pub storage: Storage<'a>,
}

impl<'a> Server<'a> {
    pub fn run(&self) {
        let addr = self
            .listener
            .local_addr()
            .expect("unable to get the local port?");

        println!("Listening on port {}", addr.port());

        for connection in self.listener.incoming() {
            match connection {
                Ok(stream) => {
                    thread::spawn(|| {
                        handle_client(stream);
                    });
                }
                Err(e) => panic!(e),
            }
        }
    }
}

fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 4];

    while let Ok(read) = stream.read(&mut buffer) {
        if read == 0 {
            break;
        }

        let cmd = str::from_utf8(&buffer).unwrap();
        if cmd == "PING" && stream.write("PONG\n".as_bytes()).is_err() {
            break;
        }
    }

    println!("disconnected")
}
