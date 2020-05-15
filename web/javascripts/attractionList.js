function display(v){
    let x = v.value;
    let itemboxs = document.getElementsByClassName("container");
    for(var i = 0 ; i < itemboxs.length ; i++){
        itemboxs[i].style.display = "none";
    }
    let itemtexts = document.getElementsByClassName("item-text")
    for(var i = 0 ; i < itemtexts.length ; i++){
        itemtexts[i].style.backgroundColor  = "transparent";
    }
    v.style.backgroundColor = "goldenrod"
    document.getElementById(x).style.display = "flex";
}