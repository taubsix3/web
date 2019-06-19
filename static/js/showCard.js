function showCard(id) {
    elements = document.getElementsByClassName("showBox");
    for( i in elements ) {
        if ( elements[i].id != id ) {
            elements[i].style.display = "none";
        }
        else {
            elements[i].style.display = "block";
        }
    }
}