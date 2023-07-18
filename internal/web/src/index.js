function moveElement(targetId, destinationId, prepend=false) {
    var target = document.getElementById(targetId);
    var destination = document.getElementById(destinationId);
    
    if(prepend){
        destination.prepend(target);
        return;
    }

    destination.append(target);
} 