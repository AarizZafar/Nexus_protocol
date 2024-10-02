document.addEventListener("DOMContentLoaded", () => {
    const welcomeBox           = document.getElementById("welcomeBox");

    const ExtractIntroLine     = document.getElementById("ExtractIntroLine"); 
    const ExtractIntroDisp     = document.getElementById("ExtractIntroDisp");

    const BioMetrics           = document.getElementById("BioMetrics");
    const BioMetricsDisplay    = document.getElementById("BioMetricsDisplay");

    setTimeout(() => {
        welcomeBox.style.animation = "fadeOut 1s ease-in-out forwards";  // after 2 seconds the welcome box will fade out over 1 seconds using css fadeout animation 
        setTimeout(() => {
            welcomeBox.style.display = "none";     // after the fade-out is done the welcomebox is hidden from the screen
            ExtractIntroLine.style.display = "inline-block"; // the ExtractIntroLine is now made visible  
            BioMetrics.style.display = "inline-block"; // the BioMetrics section is also made visible after the fade-out completes 
        }, 1000);  // will fade over 1 seconds using css fadeOut animation
    }, 2000);         // wait for 2 seconds before srarting the fade-out animation 

    // Show and animate the "System Biometrics Extraction process" after 3 seconds
    setTimeout(() => {
        ExtractIntroDisp.textContent = "System Biometrics Extraction process";  // change the text of the ExtractIntroDisp element to say (....)
        ExtractIntroDisp.classList.add('typed'); // adds a CSS class called typed to the ExtractIntroDisp element, triggering the typing animation for the text 
    }, 3000); // 3-second delay for the first message

    // Show and animate the "Display Biometrics" message with large content
    // setTimeout(() => {
    //     BioMetricsDisplay.innerHTML = "display biometrics"; // sets the inner content of the BioMetricDisplay element to the text (....)
    //     BioMetricsDisplay.classList.add('typed'); // adds the type animation 
    // }, 3000); // 3-second delay for the large content

    setTimeout(() => {
        fetch('/getSysBioMetx').then(response => response.json()).then(
            data => {
                BioMetricsDisplay.innerHTML = `
                <div class="biometrics-list">
                    <h2>System Biometrics</h2>
                    <ul>
                        <li><strong>MAC</strong> <span class="fixed-space"></span>: ${data.MAC}</li>
                        <li><strong>CPU Serial</strong> <span class="fixed-space"></span>: ${data.CPUSerial}</li>
                        <li><strong>Hard Drive Serial</strong> <span class="fixed-space"></span>: ${data.HardDriveSerial}</li>
                        <li><strong>Mother Board Serial</strong> <span class="fixed-space"></span>: ${data.MotherBoardSerial}</li>
                        <li><strong>BIOS Serial</strong> <span class="fixed-space"></span>: ${data.BIOSSerial}</li>
                        <li><strong>SSD Serial</strong> <span class="fixed-space"></span>: ${data.SSDSerial}</li>
                        <li><strong>TPM chip ID</strong> <span class="fixed-space"></span>: ${data.TPMChipID}</li>
                        <li><strong>RAM Serial</strong> <span class="fixed-space"></span>: ${data.RAMSerial}</li>
                        <li><strong>GPU Serial</strong> <span class="fixed-space"></span>: ${data.GPUSerial}</li>
                        <li><strong>NIC ID</strong> <span class="fixed-space"></span>: ${data.NICID}</li>
                    </ul>
                </div>
                `;
                BioMetricsDisplay.classList.add('typed')
            }).catch( error => {
                console.error('Error fetching biometrics: ', error);
            });
    },3000);
});
