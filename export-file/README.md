## <span style="color:magenta; size : 20px">*Ascii-Art-Web* (Dockerize)</span>

### <span style="color:pink">Description

Ascii-Art-Web (Dockerize) is a web site that allows you to convert text into ASCII art using different banner styles, with a more appealing, intuitive, and user-friendly interface.
Dockerization: To make it easy to deploy and disseminate the Ascii-Art-Web site, it is presented in a Docker format. This makes it easy to run and test the site in different environments of having to set up the site.

### <span style="color:pink">Authors

- Ayoub Benramdane
- Anas Sebbar
- Hasnae Lamrani


### <span style="color:pink">Usage

1. Clone the repository:
``` 
git clone https://learn.zone01oujda.ma/git/abenramd/ascii-art-web-dockerize
```
2. Navigate to the project directory:
```
cd ascii-art-web-dockerize
```
3. Build the Docker image && Run the container:
```
./commands.sh
```
4. Open your web browser and go to http://localhost:8081

### <span style="color:pink">Implementation Details

The web site is build using the Go programming language. The web server handles HTTP requests, processes user input, and generates the ASCII art using a dedicated package.

The main steps of the algorithm are:

1. Retrieve the text entered by the user and the selected banner style from the web form.
2. Pass the input text and the banner style to the ASCII art generation package.
3. The package reads the corresponding banner file, processes the input text, and generates the ASCII art.
4. The web server receives the generated ASCII art and displays it on the web page.
5. The web site must be packaged as a Docker container, making it easy to deploy and run in various environments.
6. The Docker image should be self-contained and include all the necessary dependencies.


### <span style="color:pink">Requirements

1. The web server must be written in Go and use the standard Go packages.
2. Dockerization:
   - The application must be packaged as a Docker container, making it easy to deploy and run in various environments.
   - The Docker image should be self-contained and include all the necessary dependencies.