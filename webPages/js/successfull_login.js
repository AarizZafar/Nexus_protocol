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
                  <ul>
                    <li>MAC                 : ${data.MAC}               </li>
                    <li>CPU Serial          : ${data.CPUSerial}         </li>
                    <li>Hard Drive Serial   : ${data.HardDriveSerial}   </li>
                    <li>Mother Board Serial : ${data.MotherBoardSerial} </li>
                    <li>BIOS Serial         : ${data.MotherBoardSerial} </li>
                    <li>SSD Serial          : ${data.BIOSSerial}        </li>
                    <li>TPM chip ID         : ${data.TPMChipID}         </li>
                    <li>RAM Serial          : ${data.RAMSerial}         </li>
                    <li>GPU Serial          : ${data.GPUSerial}         </li>
                    <li>NICID               : ${data.NICID}             </li>
                  </ul>
                `;
                BioMetricsDisplay.classList.add('typed')
            }).catch( error => {
                console.error('Error fetching biometrics: ', error);
            });
    },3000);
});
