# Ityfuzz ABI server
## Installation
### Prerequisites

- Go 1.20
- Ityfuzz with `LOCAL` (this repo/PR)

### Build

- Setup the relations in `relations.json`:
  - You need a json file that has key-value pairs of the form `contract_address: ABI_file.json`, like `relations_example.json`
- The ABI files (hardhat generates this in `artifacts/contracts/`, they are called `contract.json`) should be in the folder `ABI`.

- Then run the server:
```bash
go run .
```

The server listens port 8080, when ityfuzz requests an ABI, this mocks a response like the etherscan api.