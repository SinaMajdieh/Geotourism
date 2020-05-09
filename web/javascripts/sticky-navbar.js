window.onscroll = function() {scrollFunction()};

function scrollFunction() {
  if (document.body.scrollTop > 80 || document.documentElement.scrollTop > 80) {
    document.getElementById("navbar").style.height = "10%";
    document.getElementById("navbar-logo").style.height = "80px";
    document.getElementById("navbar-logo").style.width = "80px";
  } else {
    document.getElementById("navbar").style.height = "15%";
    document.getElementById("navbar-logo").style.height = "100px";
    document.getElementById("navbar-logo").style.width = "100px";
  }
}
