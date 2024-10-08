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
                        <li><strong>MAC                         </strong> <span class="fixed-space">   </span>: ${data.MAC}                   </li>
                        <li><strong>CPU Serial                  </strong> <span class="fixed-space">   </span>: ${data.CPUSerial}             </li>
                        <li><strong>Hard Drive Serial           </strong> <span class="fixed-space">   </span>: ${data.HardDriveSerial}       </li>
                        <li><strong>Mother Board Serial         </strong> <span class="fixed-space">   </span>: ${data.MotherBoardSerial}     </li>
                        <li><strong>BIOS Serial                 </strong> <span class="fixed-space">   </span>: ${data.BIOSSerial}            </li>
                        <li><strong>SSD Serial                  </strong> <span class="fixed-space">   </span>: ${data.SSDSerial}             </li>
                        <li><strong>TPM chip ID                 </strong> <span class="fixed-space">   </span>: ${data.TPMChipID}             </li>
                        <li><strong>RAM Serial                  </strong> <span class="fixed-space">   </span>: ${data.RAMSerial}             </li>
                        <li><strong>GPU Serial                  </strong> <span class="fixed-space">   </span>: ${data.GPUSerial}             </li>
                        <li><strong>NIC ID                      </strong> <span class="fixed-space">   </span>: ${data.NICID}                 </li>
                        <li><strong>Base Board Product          </strong> <span class="fixed-space">   </span>: ${data.BaseBoardProduct}      </li>
                        <li><strong>System UUID                 </strong> <span class="fixed-space">   </span>: ${data.SystemUUID}            </li>
                        <li><strong>OSInstallationID            </strong> <span class="fixed-space">   </span>: ${data.OSInstallationID}      </li>
                        <li><strong>Disk Volume Serial Number   </strong> <span class="fixed-space">   </span>: ${data.DiskVolumeSerialNumber}</li>
                        <li><strong>BIOSVersion                 </strong> <span class="fixed-space">   </span>: ${data.BIOSVersion}           </li>
                        <li><strong>Boot ROM Version            </strong> <span class="fixed-space">   </span>: ${data.BootROMVersion}        </li>
                        <li><strong>Boot GPUVendorID            </strong> <span class="fixed-space">   </span>: ${data.GPUVendorID}           </li>
                        <li><strong>DeviceTreeIdentifier        </strong> <span class="fixed-space">   </span>: ${data.DeviceTreeIdentifier}  </li>
                        <li><strong>UEFIFirmwareVersion         </strong> <span class="fixed-space">   </span>: ${data.UEFIFirmwareVersion}   </li>
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