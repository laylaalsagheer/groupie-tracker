$(document).ready(function(){
    // Function to toggle mobile navigation
    $('.burgermenu').on('click', function(){
        $('.mob-nav').toggle("slow");
    });
}); // Added closing parenthesis here

document.querySelectorAll('a.detailsbutton').forEach(function(button) {
    button.addEventListener('click', function(e) {
        e.preventDefault();
        fetch(e.target.href)
            .then(response => response.text())
            .then(html => {
                document.querySelector('.container').innerHTML = html;
            });
    });
}); // Added closing parenthesis here