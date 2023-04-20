# URL Shortener

This is a URL shortener implemented in Go, designed with SOLID principles and design patterns in mind to optimize time and space complexity. It also features efficient handling of collisions by retrying up to 100 times for unique URL generation and real-time response generation in a distributed environment. The code has been optimized to handle all bottlenecks and includes custom error handling and custom response generator.

## Features

- Shorten long URLs to short URLs
- Expand short URLs to long URLs
- Retry up to 100 times for unique URL generation
- Real-time response generation
- Custom error handling and response generation

## Getting Started

### Prerequisites

- Go v1.16 or later

### Installation

1. Clone the repository

```
git clone https://github.com/<your_username>/url-shortener.git
```

2. Change into the directory

```
cd UrlShortener
```

3. Build the binary

```
go build
```

4. Run the binary

```
Press the run button of the compiler or press 'control+shift+R'
```

The URL shortener is now running on `http://localhost:8080`.

## Usage

### Shortening URLs

To shorten a URL, make a POST request to `http://localhost:8080/shorten` with a JSON body containing the long URL:

```json
{
  "url": "https://www.example.com/very/long/url/that/needs/to/be/shortened"
}
```

The response will contain the short URL:

```json
{
    "status_code": 200,
    "short_url": "https://bit.ly/1elsg5561d",
    "error": null
}
```

# Successful short-url generation
<img width="1728" alt="Screenshot 2023-04-20 at 11 49 18 AM" src="https://user-images.githubusercontent.com/36428256/233275268-90ff20ef-530c-432f-bafa-241bc2433d7e.png">



# Validating insecured urls
<img width="1728" alt="Screenshot 2023-04-20 at 11 51 12 AM" src="https://user-images.githubusercontent.com/36428256/233275670-09119b2a-2040-4b0b-b9ed-366461fd3188.png">



# Validating invalid urls
<img width="1728" alt="Screenshot 2023-04-20 at 11 53 59 AM" src="https://user-images.githubusercontent.com/36428256/233276224-b021bf9e-94d5-454f-80e6-7bade066f49a.png">
