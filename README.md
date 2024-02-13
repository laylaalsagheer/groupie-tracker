# groupie-tracker

<h1 align="center">Groupie Tracker</h1>

<h2 align="left">About The Project</h2>
Music discovery meets concert planning. This project delivers instant updates and concert information for 52 top bands, empowering music lovers to easily stay informed and never miss a beat. Dive into artist news, explore concert details, and stay ahead of the musical curve with this convenient and comprehensive resource.

## Getting Started

Clone the project

```bash
git clone https://github.com/laylaalsagheer/groupie-tracker.git
```

Go to the project directory

```bash
  cd groupie-tracker
```

### Directory Structure

```console
─ Groupie Tracker/
│
├── templates/
│   ├── index.html
│   ├── details.html
│   
├── static/
│   ├── css/
│   │   └── style.css
│   │   └── details.css
│   ├── javascript/
│ 
├── main.go
├── model.go
├── go.mod
├── README.md
└── ...
```

## Usage

_make sure you are in project directory_

```bash
go run main.go
```
1. Open http://localhost:8080/ in a browser .
2. Choose your favorite band to see the full details. 

## Additional information
- The project is written in Go.
- Only standard go packages were used.
- Future Plan:
  It features a user-friendly interface with the following capabilities:

  - **Search Bar:** Find your favorite bands quickly by searching for their name.
  - **Filters:** Sort and filter bands based on various criteria, such as creation date, first album, number of members and locations of concerts.
  - **Geolocation Maps (Using Google API):** Explore the concert locations of your favorite bands on an interactive map.

