document.addEventListener("DOMContentLoaded", () => {
    const welcomeBox = document.getElementById("welcomeBox");

    const ExtractIntroLine = document.getElementById("ExtractIntroLine"); 
    const ExtractIntroDisp = document.getElementById("ExtractIntroDisp");

    const BioMetrics = document.getElementById("BioMetrics");
    const BioMetricsDisplay = document.getElementById("BioMetricsDisplay");

    // Hide welcome box after 2 seconds and display the biometrics sections
    setTimeout(() => {
        welcomeBox.style.animation = "fadeOut 1s ease-in-out forwards"; 
        setTimeout(() => {
            welcomeBox.style.display = "none"; 
            ExtractIntroLine.style.display = "inline-block"; 
            BioMetrics.style.display = "inline-block";
        }, 1000); // Wait for fade out to finish
    }, 2000);

    // Show and animate the "System Biometrics Extraction process" after 3 seconds
    setTimeout(() => {
        ExtractIntroDisp.textContent = "System Biometrics Extraction process";
        ExtractIntroDisp.classList.add('typed');
    }, 3000); // 3-second delay for the first message

    // Show and animate the "Display Biometrics" message with large content
    setTimeout(() => {
        BioMetricsDisplay.innerHTML = "display biometrics";
        BioMetricsDisplay.classList.add('typed');
    }, 3000); // 3-second delay for the large content
});
