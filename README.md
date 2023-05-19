# Ityfuzz ABI server
## Installation
### Prerequisites

- Go 1.20

### Build

- Setup the relations in `relations.json`:
  - You need a json file that has key-value pairs of the form `contract_address: ABI_file.json`, like `relations_example.json`
  - The ABI fileS should be in a folder called `ABI`.

- Run the server:
```bash
go run .
```

The server listens port 8080