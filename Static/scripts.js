$(document).ready(function(){
    // Function to toggle mobile navigation
    $('.burgermenu').on('click', function(){
        $('.mob-nav').toggle("slow");
    });
}); 
document.querySelectorAll('a.detailsbutton').forEach(function(button) {
    button.addEventListener('click', function(e) {
        e.preventDefault();
        // Fetch the HTML content
        fetch(e.target.href)
            .then(response => response.text())
            .then(html => {
                // Update the content of the .container element
                document.querySelector('.container').innerHTML = html;
                // Update the browser's URL without reloading the page
                history.pushState({}, '', e.target.href);
            });
    });
});