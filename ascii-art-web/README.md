## ascii-art-web

### Description

Ascii-Art-Web is a web site that allows you to convert text into ASCII art using different banner styles.

### Authors

- Ayoub Benramdane
- Anas Sebbar
- Hasnae Lamrani


### Usage

1. Clone the repository:
``` 
git clone https://learn.zone01oujda.ma/git/abenramd/ascii-art-web
```
2. Navigate to the project directory:
```
cd ascii-art-web
```
3. Start the server:
```
go run . 
```
4. Open your web browser and go to http://localhost:8081

### Implementation Details

The web site is build using the Go programming language. The web server handles HTTP requests, processes user input, and generates the ASCII art using a dedicated package.

The main steps of the algorithm are:

1. Retrieve the text entered by the user and the selected banner style from the web form.
2. Pass the input text and the banner style to the ASCII art generation package.
3. The package reads the corresponding banner file, processes the input text, and generates the ASCII art.
4. The web server receives the generated ASCII art and displays it on the web page.


### Requirements

1. The web server must be written in Go and use the standard Go packages.
2. The HTML templates must be placed in the templates directory at the root of the project.
3. The code should follow good programming practices, such as:
    - Modular design with a clear separation of responsibilities
    - Appropriate error handling and the use of HTTP status codes
    - Well-documented and readable code

