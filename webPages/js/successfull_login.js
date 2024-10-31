document.addEventListener("DOMContentLoaded", () => {
    const welcomeBox           = document.getElementById("welcomeBox");
// ----------------------------------------------------------------------------------------
    const ExtractIntroLine     = document.getElementById("ExtractIntroLine"); 
    const ExtractIntroDisp     = document.getElementById("ExtractIntroDisp");
// ----------------------------------------------------------------------------------------
    const BioMetrics           = document.getElementById("BioMetrics");
    const BioMetricsDisplay    = document.getElementById("BioMetricsDisplay");

    setTimeout(() => {
        welcomeBox.style.animation = "fadeOut 1s ease-in-out forwards";     // after 2 seconds the welcome box will fade out over 1 seconds using css fadeout animation 
        setTimeout(() => {
            welcomeBox.style.display = "none";                              // after the fade-out is done the welcomebox is hidden from the screen
            ExtractIntroLine.style.display = "inline-block";                // the ExtractIntroLine is now made visible  
            BioMetrics.style.display = "inline-block";                      // the BioMetrics section is also made visible after the fade-out completes 
        }, 1000);                                                           // will fade over 1 seconds using css fadeOut animation
    }, 2000);                                                               // wait for 2 seconds before srarting the fade-out animation 

    // Show and animate the "System Biometrics Extraction process" after 3 seconds
    setTimeout(() => {
        ExtractIntroDisp.textContent = "System Biometrics Extraction process";  // change the text of the ExtractIntroDisp element to say (....)
        ExtractIntroDisp.classList.add('typed');                                // adds a CSS class called typed to the ExtractIntroDisp element, triggering the typing animation for the text 
    }, 3000);                                                                   // 3-second delay for the first message
 
    setTimeout(() => {
        // fetch('/getSysBioMetx') - makes a HTTP request to a serve, sends a GET request to the '/getSysBioMetx' route on the gins server
        // .then is method for handling asynchronous operation like fetch, once request done and response received do something 

        // response => response.json() the response is beeing converted to a json format 
        fetch('/getSysBioMetx').then(response => response.json())
        .then( data => {BioMetricsDisplay.innerHTML = `
                <div class="biometrics-list">
                    <h2>System Biometrics</h2>
                    <ul>
                        <li><strong>SSID                          </strong> <span class="fixed-space">   </span>: ${data.SSID}                      </li>
                        <li><strong>MAC                           </strong> <span class="fixed-space">   </span>: ${data.MAC}                       </li>
                        <li><strong>SystemSerialNumber            </strong> <span class="fixed-space">   </span>: ${data.SystemSerialNumber}        </li>
                        <li><strong>UUID                          </strong> <span class="fixed-space">   </span>: ${data.UUID}                      </li>
                    </ul>
                </div>
                `;
                BioMetricsDisplay.classList.add('typed')
            }).catch( error => {
                console.error('Error fetching biometrics: ', error);
            });
    },3000);

    /* SIMPLE UNDERSTANDING OF THE ABOVE CODE 
        fetch('/someEndpoint')
            .then(response => response.text())
            .then(data => {
                console.log(data);  // Once the response is received, log the text data
        }); 
    */

});

/*
Client (Browser)
    |
    |----> fetch('/getSysBioMetx') ---->  Gin Server
                                             |
                                             |----> router.GET("/getSysBioMetx", controllers.GetSysBioMetrix)
                                                            |
                                                            |----> controllers.GetSysBioMetrix
                                                                           |
                                                                           |----> (Execute logic to fetch system biometrics)
                                                                           |
                                                            |<---- (Return JSON response with biometrics data)
                                             |
    |<---- Response (JSON with system biometrics) ----|
    |
Display data on web page
*/