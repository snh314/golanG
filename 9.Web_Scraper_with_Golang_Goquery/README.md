# Go Blog Title Scraper

A fast and efficient web scraper built with Go (Golang) and the [goquery](https://github.com/PuerkitoBio/goquery) library.

## Project Description

This project demonstrates how to perform HTML scraping in Go. It sends an HTTP GET request to the ScrapingBee blog (`https://www.scrapingbee.com/blog/`), parses the HTML response body, and extracts all the text contained within `<h4>` tags (which represent the article headlines on that specific page). The extracted titles are then formatted into a clean, hyphenated list and printed directly to the terminal.

## Features

- **HTTP GET Requests:** Fetches live web pages using Go's standard `net/http` package.
- **HTML Parsing:** Parses HTML documents easily using the jQuery-like `goquery` package.
- **Targeted Extraction:** Iterates over specific HTML elements (`<h4>` tags) to extract clean text.
- **Robust Error Handling:** Includes basic panic and error handling for network requests and parsing failures.

## Prerequisites

- Go installed on your local machine (version 1.16 or higher recommended).
- An active internet connection to fetch the target website and download dependencies.

## Installation & Setup

1. **Create the project directory:**

   ```bash
   mkdir go-blog-scraper
   cd go-blog-scraper

   ```

2. **Initialize the Go module:**

   ```bash
    go mod init go-blog-scraper

   ```

3. **Install the goquery dependency:**

   ```bash
    go get [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)

   ```

4. **Usage- Run the application directly from terminal:**

   ```bash
    go run main.go

   ```

5. **Expected Output:The terminal will display a list of blog titles extracted from the website**
   Plaintext
   Blog Titles:

- Web Scraping with Python
- How to scrape Amazon
- Best Web Scraping Tools
  ...

**How It Works**

- **GetLatestBlogTitles(url):** Makes an HTTP GET request to the provided URL.
- **goquery.NewDocumentFromReader:** Reads the HTTP response body and converts it into a searchable HTML document.
- **doc.Find("h4").Each(...):** Traverses the DOM to find every <h4> tag, extracts the text using s.Text(), and appends it to a formatted string variable.
- **fmt.Printf:** Safely prints the final formatted string to the console.
