function checkPassword(){
    var password = document.getElementById("userPassword"); 
    var confirmPassword = document.getElementById("confirmPassword");

    var message = document.getElementById("message");

    var matchingColorPass = "#66aa44"
    var matchingColorNotPassed = "#66aa44"

    if (password == confirmPassword){
        confirmPassword.style.color = matchingColorPass;
        message.innerText = "Password confirmed";

        matchingPasswords = true;
    } else {
        confirmPassword.style.color = matchingColorNotPassed;
        message.innerText = "No match";
    }
}

var matchingPasswords = false;

function passwordMatching(){
    return matchingPasswords;
}