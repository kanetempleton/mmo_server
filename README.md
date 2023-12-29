
# MMO Game Server

Multiplayer game server written in Golang for high performance. The client I wrote that syncs with this communication protocol is written in C++ and can be found [here.](https://github.com/kanetempleton/mmo_client)


## Getting Started

### Prerequisites

- [Go](https://golang.org/)

### Running the Server

1. Clone this repository:

   ```bash
   git clone https://github.com/kanetempleton/mmo_server.git
   ```

2. Navigate to the project directory:

   ```bash
   cd mmo_server
   ```

3. Build the server:

   ```bash
   go build
   ```

4. Run the server:

   ```bash
   ./mmo_server
   ```

## Specification

### Connecting Clients

Client-server protocol is written over TCP. The communication protocol details will be specified here as they develop.

### CLI Commands

- **connections**: Display information about connected clients.
- **kick <ConnectionID>**: Kick a client based on their ID.
- **message <ConnectionID> <msg>**: Message a client based on their ID.