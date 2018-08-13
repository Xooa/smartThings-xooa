
This page is a step-by-step tutorial to integrate blockchain smartApps with Xooa “blockchain as a service”.
The repository in use for this example is <https://github.com/Xooa/smartThings-xooa>


# Overview
There are two components in this repo, the blockchain chaincode & smartthings smartapps. Blockchain chaincode is deployed in Xooa console and smartapps in smartthings IDE. 
Xooa exposes a permanent cloud end-point to smarrthings to enable cloud-to-cloud integration while maintaining peer-to-peer capabilities of blockchain.

## Deploy smartthings blockchain chaincode in xooa using github integration 
1. [optional] fork the repository from <https://github.com/Xooa/smartThings-xooa>.
2. Log in to your xooa account at <https://xooa.com/blockchain>
3. Click **apps**, then **Deploy New** on. 
If this is your first time deploy a chaincode app with xooa, you will need to authorize xooa with your Github account.
Click **Connect to GitHub**. Follow the onscreen instructions to complete the authorization process.
<img src="https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_logging.png" alt="github logging" width="500px"/>
<img src="https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_authorizing.jpg" alt="github authorizing" width="500px"/>
4. Enter a name and description for your app, and then click **Next**.
5. Search for the **smartThings-xooa** repo, or your fork of that repo. A list of repositories matching the search criteria are shown. Select the repo, and then click **Next**.
The deployment details appear.
6. Select **master**  branch and **smartThings-xooa** as the chaincode from the dropdown available, then click **Deploy**.
7. Relax while Xooa does the blockchain heavy lifting for you. You will be redirected to app dashboard when the deployment completes.
8. Copy the **App ID** from the dashboard under `Basic Information` tab. You will need it later to connect the smartapp.
9. From the app dashboard, navigate to the **Identities** tab.
10. For the available identity, click **Show API token** and copy the token value. You need this token to authorize API requests to the **smartApp**.

___

## Event Logger SmartApp Setup (SmartThings IDE)
1. Log in with your Samsung  SmartThings account to the SmartThings IDE at <https://graph.api.smartthings.com>.
2. Navigate to **My SmartApps** 
Now you need to publish the app.  You can do this with or wiithout github integration: 

   **Without github integration:**

   * From main menu, select `New SmartApp`.
   * Navigate to **From code**. 
   * Locate `blockchain-event-logger.groovy` in **smartThings-xooa** github repo.
   * Click **Raw** 
   * Select all and copy the code
   * Paste your selection in the **From Code** section in smartThings console, click **create**, **save**, **Publish -> For me**.

   **With Github integration** (if you haven't already set up Github to work with smartthings, here is the community FAQ on the subject  <https://community.smartthings.com/t/faq-github-integration-how-to-add-and-update-from-repositories/39046>)
   * Sill under Smartapps tab, select  **Settings**.
   * Select  **Add new repository**.
   * Add the Github repo to your IDE with the following parameters:
	   * Owner: xooa
	   * Name: smartThings-xooa
	   * Branch: master
   * Click **Save**.
   * Click **Update from Repo**.
   * Click **smartThings-xooa (master)**.
   * Select **blockchain-event-logger.groovy** from **New(only in GitHub)** column.
   * Select **Publish**.
   * Click  **Execute Update**.

### Event Logger SmartApp Configuration (still at smartthings IDE)
4. Select  the smartApp you just deployed.
5. Navigate to **App Settings**.
6. Navigate to **Settings** option.
![appsettings](https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_appSettings.png "App Settings")
7. Enter app Id provided in dashboard of xooa under `Basic Information` tab.
8. Enter the API Token provided in xooa dashboard under `Identities` tab.
9. Click **Update**.

There are two apps available for smartThings in the Google Play store. We recommend the classic app over the new app.

## Event Logger SmartApp Setup (Smartphone)
1. Ensure you have the smartThings app installed on your phone with at least one location and one device defined.
2. Ensure you are using the same login ID you used for your developer account to log in to the app.
3. Open your SmartThings app on your smartphone.

### SmartThings classic app (Preferred)
1. Tap **automation** in the lower bar.
2. Tap the SmartApps tab on top.
3. Tap  **Add a SmartApp**.
4. Scroll to the bottom and then tap **My Apps**.
5. Identify the `Blockchain Event Logger` app, and tap it to proceed.
6. Select the devices you want to log in the xooa blockchain.
7. Click **Save**

### SmartThings new app
1. Tap `automations` in lower bar.
2. Tap `Add`(in android) or `+`(in IOS).
3. If prompted, select The location you want to add the app to.
4. Tap `Done`(in IOS).
5. Find `Blockchain Event Logger`, usually it will appear last and may take a few seconds to appear. 
6. Tap it to continue setting it up.
6. Select which devices you want to log in the xooa blockchain.

## Event Viewer SmartApp Setup (SmartThings IDE)
Follow the same steps as `Event Logger SmartApp Setup (SmartThings IDE)` but:
1. Use `blockchain-event-viewer.groovy` instead of `blockchain-event-logger.groovy` from **smartThings-xooa** github repo.
2. Skip `Event Logger SmartApp Configuration` steps.

### Using Event Viewer smartApp
#### SmartThings classic app (Preferred)
1. Tap **automation** in the lower bar.
2. Tap the SmartApps tab on top.
3. Tap **Add a SmartApp**.
4. Scroll to the bottom and tap  **My Apps**.
5. Find `Blockchain Event Viewer` and tap it.
6. Enter the **App Id** provided in the xooa dashboard under `Basic Information`.
7. Enter the **API Token** provided in the xooa dashboard under `Identities`.
8. Click **Next** to proceed to view the devices logging to the blockchain.
9. Click any device to view all the past logged events for that particular device.
10. Click **Save** to store **App Id** and **API Token** with smartApp for future uses.

#### SmartThings new app
1. Tap **automations** in lower bar.
2. Tap **Add**(in android) or **+**(in IOS).
3. If prompted, select The location you want to add the app to.
4. Tap **Done**(in IOS).
5. Find `Blockchain Event Viewer,` usually it appears at the bottom of the page and may take a few seconds to apepar. 
6. Tap the app.
6. Enter the **App Id** provided in the xooa dashboard under `Basic Information`.
7. Enter the **API Token** provided in xooa dashboard under `Identities`.
8. Click **Next** to proceed to view devices logging to the blockchain.
9. Click on any device to view all the past logged events for that particular device.
10. Click **Save** to store the **App Id** and **API Token** with smartApp for future uses.
