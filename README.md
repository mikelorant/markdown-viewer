# markdown-viewer

This is a simple Go application that serves a Markdown file as HTML using the
`goldmark` package.

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/mikelorant/markdown-viewer.git
   ```

1. Navigate to the project directory:

   ```shell
   cd markdown-viewer
   ```

## Usage

1. Place your Markdown file in the project directory and rename it to
   `content.md`.

1. Run the application:

   ```shell
   go run .
   ```

1. Open your web browser and visit `http://localhost:8080` to view the rendered
   Markdown file.

## Dependencies

This project uses the following third-party package:

- [goldmark](https://github.com/yuin/goldmark) - A Markdown parser written in Go

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you
have any improvements or bug fixes.
