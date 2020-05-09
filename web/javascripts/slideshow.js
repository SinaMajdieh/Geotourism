var slideIndex = 1;
showDivs(0 , false)
setTimeout(showSlides , 5000); // Change image every 2 seconds
function plusDivs(n) {
  showDivs(n , true);
}
function showSlides(){
  showSlideDivs(1 , true)
}
function showSlideDivs(n) {
  var x = document.getElementsByClassName("slide"); 
  
  var elem1 = document.getElementById('slide'+slideIndex);
  slideIndex += n;
  if (slideIndex > x.length) {slideIndex = 1}
  if (slideIndex < 1) {slideIndex = x.length}
  elem2 = document.getElementById('slide'+slideIndex);
  slideleft = true;
  if(n < 0){
    slideleft = false;
  }
  myMove(elem1 , elem2 , slideleft);
  setTimeout(showSlides , 7000); // Change image every 2 seconds
  
      
}
function showDivs(n , slide) {
  var x = document.getElementsByClassName("slide"); 
  if (slide == true){
    var elem1 = document.getElementById('slide'+slideIndex);
    slideIndex += n;
    if (slideIndex > x.length) {slideIndex = 1}
    if (slideIndex < 1) {slideIndex = x.length}
    elem2 = document.getElementById('slide'+slideIndex);
    slideleft = true;
    if(n < 0){
      slideleft = false;
    }
    myMove(elem1 , elem2 , slideleft);
  }else{
    for(var i = 0 ; i < x.length ; i++){
      x[i].style.display = "none"
    }
    var elem = document.getElementById('slide'+slideIndex)
    elem.style.display = "block"
    elem.style.top = "0%"
    elem.style.left = "0%"
  }
      
}
function myMove(elem1 , elem2 , slideleft) {
  elem1.style.display = "block"
  elem2.style.display = "block"
  var factor = 1
  var pos = 100;
  var id = setInterval(frame, 5);
  if(slideleft == true){
    factor *= -1
  }
  function frame() {
    if (pos == 0) {
      clearInterval(id);
    } else {
      pos--;
      elem2.style.left = factor*pos + "%"; 
      elem1.style.left = factor*(pos - 100) + "%";
      elem1.style.top = 0 + "%";
      elem2.style.top = 0 + "%"; 
    }
  }
}