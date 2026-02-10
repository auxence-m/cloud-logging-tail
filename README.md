## Overview

`cloudtail` is a lightweight cloud-native command-line tool written in Golang that allows users to view or tail logs from Google Cloud Logging (similar to kubectl logs).

## How to run

### Prerequisites
- Ensure you have [Golang](https://go.dev/doc/install) installed on your machine.


### Installation

1. Clone the repository
```
git clone https://github.com/auxence-m/cloudtail.git
```

2. Navigate to the Project Directory

```
cd cloudtail
```

3. Install Dependencies

```
go mod tidy
```

4. Build the Application

```
go build
``` 

After building, you'll find the `cloudtail` executable (`cloudtail.exe` on Windows) in your project directory.

5. Run the Application

```
cloudtail [command] --flag
```

## Documentation

Full documentation for `cloudtail`, including all available commands is available [here](https://cloudtail-docs.vercel.app/)

## License

This project is licensed under the MIT License.
