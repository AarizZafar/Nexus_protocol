// successfull_login.js

document.addEventListener("DOMContentLoaded", () => {
  const welcomeBox = document.getElementById("welcomeBox");
  const biometricsContainer = document.getElementById("biometricsContainer"); // Now this will not be null
  const biometricsDisplay = document.getElementById("biometricsDisplay");

  // Set a timer to hide the welcome message after 2 seconds
  setTimeout(() => {
      console.log("fade");
      welcomeBox.style.animation = "fadeOut 1s ease-in-out forwards"; // Fade out the welcome message
      setTimeout(() => {
          welcomeBox.style.display = "none"; // Hide the welcome box after fading out
          biometricsContainer.style.display = "inline-block"; // Show the biometrics container
      }, 1000); // Wait for the fade out animation to finish
  }, 2000); // 2 seconds
  
  // After fading out, show the biometrics message
  setTimeout(() => {
      console.log("text print");
      biometricsDisplay.textContent = "System Biometrics Extraction......"; // Set the text
      biometricsDisplay.classList.add('typed'); // Add typing animation class
  }, 3000); // Start typing animation after 3 seconds
});

