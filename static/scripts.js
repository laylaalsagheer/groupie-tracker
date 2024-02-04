$(document).ready(function(){
    // Function to toggle mobile navigation
    $('.burgermenu').on('click', function(){
        $('.mob-nav').toggle("slow");
    });

    // Fetch data from the artist API
    fetch('http://localhost:8080/artists')
    .then(response => response.json())
    .then(artists => {
        // Clear existing content
        $('#band-container').empty();

        // For each artist, create a new div and append it to the container
        artists.forEach(artist => {
            let div = document.createElement('div');
            div.className = 'col';
            div.innerHTML = `
                <img src="${artist.Image}" alt="${artist.Name}">
                <p>Artist: ${artist.Name}</p>
                <p>Members: ${artist.Members.join(', ')}</p>
                <p>Creation Date: ${artist.CreationDate}</p>
                <p>First Album: ${artist.FirstAlbum}</p>
            `;

            // Fetch additional data for the artist (locations, dates, relations)
            fetchAdditionalData(artist.ID, div);

            document.querySelector('#band-container').appendChild(div);
        });
    })
    .catch(error => console.error('Error:', error));

    // Function to fetch additional data for an artist
    function fetchAdditionalData(artistId, container) {
        // Fetch location data for the artist
        fetch(`http://localhost:8080/locations?id=${artistId}`)
        .then(response => response.json())
        .then(locations => {
            const locationString = locations.Index.map(loc => loc.Locations.join(', ')).join('; ');
            container.querySelector('p:nth-child(3)').textContent = `Locations: ${locationString}`;
        })
        .catch(error => console.error('Error fetching locations:', error));

        // Fetch date data for the artist
        fetch(`http://localhost:8080/dates?id=${artistId}`)
        .then(response => response.json())
        .then(dates => {
            const dateString = dates.Index.map(date => date.Dates.join(', ')).join('; ');
            container.querySelector('p:nth-child(4)').textContent = `Dates: ${dateString}`;
        })
        .catch(error => console.error('Error fetching dates:', error));

        // Fetch relation data for the artist
        fetch(`http://localhost:8080/relations?id=${artistId}`)
        .then(response => response.json())
        .then(relations => {
            const relationString = relations.Index.map(rel => {
                const entries = Object.entries(rel.DatesLocations);
                return entries.map(([place, date]) => `${place}: ${date.join(', ')}`).join('; ');
            }).join('; ');
            container.querySelector('p:nth-child(5)').textContent = `Relations: ${relationString}`;
        })
        .catch(error => console.error('Error fetching relations:', error));
    }
});
