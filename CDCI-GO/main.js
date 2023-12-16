const eye = document.getElementById("eye");
const password= document.getElementById("password");
const image1 = "./images/eye-slash.svg";
const image2 = "./images/eye.svg"


eye.addEventListener("click", changing)

function changing() {
    var currentImage = eye.getAttribute("src")
    if (currentImage === image1) {
        eye.setAttribute("src", image2);
        password.setAttribute("type", "text");

    }
     else {
        eye.setAttribute("src", image1)
        password.setAttribute("type", "password");
   }
}



