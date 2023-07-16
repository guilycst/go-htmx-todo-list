htmx.defineExtension('hxe-mv-to', {
    onEvent: function (name, evt) {
        if (name !== "htmx:beforeSwap"){
            return;
        }

        var triggerElement = evt.target;
        var destinationId = triggerElement.getAttribute("hxe-mv-to-target")
        if (!destinationId){
            return
        }
        
        var prepend = triggerElement.getAttribute("hxe-mv-to-append") === "start"
        
        var targetId = evt.detail.elt.id
        var target = document.getElementById(targetId);

        var destination = document.querySelector(destinationId);
        
        if(prepend){
            destination.prepend(target);
        }else{
            destination.appendChild(target);
        }
    },
    handleSwap : function(swapStyle, target, fragment, settleInfo) {
        return false;
    },

})