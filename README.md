
This page is a step-by-step tutorial to integrate blockchain smartApps with Xooa “blockchain as a service”.
The repository in use for this example is <https://github.com/Xooa/smartThings-xooa>


# Overview
There are two components in this repo, the blockchain chaincode & smartthings smartapps. Blockchain chaincode is deployed in Xooa console and smartapps in smartthings IDE. 
Xooa exposes a permanent cloud end point to smarrthings to enable cloud-to-cloud integration while maintaining peer to peer capabilities of blockchain.

## Deploy  smartthings blockchain chaincode in xooa using github integration 
1. [optional] If you wish, fork repository from <https://github.com/Xooa/smartThings-xooa> to your Github account.
2. Login to Xooa (create an account in xooa if needed) at <https://xooa.com/blockchain>
3. From the main dashboard under apps, use **Deploy New** on top left corner to start the deployment process. If it is  first time you deploy a blockchain chaincode app with Xooa you will need to authorize Xooa account to your Github account (Github integration is mandatory for deploying blockchain chaincode in Xooa).
Click or tap **Connect to GitHub**. follow the instructions until you successfully authorized your Xooa account to your github account.
![logging](https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_logging.png | width=200 "Github logging")
![authorizing](https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_authorizing.jpg | width=200 "Github Authorizing")
4. Enter **App Name** and **Description** for the app.  Then **Next**.
5. Search for **smartThings-xooa** repo, or if you forked under a different name use that name. A list of repositories matching the search criteria are shown. Select the repo the, then **Next**.
6. The next step is deployment details. Select **master** as the branch and **smartThings-xooa** as the chaincode from the dropdown available. Then **Deploy**.
7. Sit down and relax while Xooa does the blockchain heavy lifting for you. You will be redirected to app dashboard once the deployment is complete.
8. Copy the **App ID** from the dashboard under `Basic Information` tab. It will be required later to connect the smartapp.
9. From the app dashboard, navigate to **Identities** tab.
10. For the identity available, click **Show API token** and copy the token value. We will be using this for authorizing the API requests. This will be your `API Token` needed to authorize **smartApp**.

___

## Event Logger SmartApp Setup (SmartThings IDE)
1. Login with your Samsung  SmartThings account to the SmartThings IDE at <https://graph.api.smartthings.com>.
2. Navigate to **My SmartApps** 
3. Now you will publish the app.  Steps depends if wish to use github integration or not.

   **Not using github integration:**

   * From top menu, select `New SmartApp`.
   * Naviagte to **From code** tab. 
   * Locate `blockchain-event-logger.groovy` in **smartThings-xooa** github repo.
   * Click or tap **Raw** 
   * Select all and Copy the code
   * Paste it in the **From Code** section in smartThings console, click or tap **create**, **save**, **Publish -> For me**.

   **Using Github integration** (if you haven't already set it up here there is FAQ in the community <https://community.smartthings.com/t/faq-github-integration-how-to-add-and-update-from-repositories/39046>)
   * Sill under Smartapps tab, Select  **Settings**.
   * Select  **Add new repository**.
   * Add the Github Repo to your IDE with the following settings:
	   * Owner: xooa
	   * Name: smartThings-xooa
	   * Branch: master
   * Click on **Save**.
   * Click on **Update from Repo**.
   * Click on **smartThings-xooa (master)**.
   * Select **blockchain-event-logger.groovy** from **New(only in GitHub)** column.
   * Check the **Publish** checkbox
   * Click on **Execute Update**.

### Event Logger SmartApp Configuration (still at smartthings IDE)
4. Select on the smartApp you just deployed.
5. Navigate to **App Settings**.
6. Navigate to **Settings** option.
![appsettings](https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_appSettings.png "App Settings")
7. Enter app Id provided in dashboard of Xooa under `Basic Information` tab.
8. Enter the API Token provided in Xooa dashboard under `Identities` tab.
9. Click or tap **Update**.

There are 2 apps available for smartThings in google play store . We recommend the classic app over the new app due to frequent glitches in the new app being faced currently.

## Event Logger SmartApp Setup (Smartphone)
1. Make sure you have smartThings app installed on your phone with at least 1 location and 1 device defined.
2. Also make sure that you use same login ID to log into the app as used in developer account login.
3. Open your SmartThings app on your smartphone.

### SmartThings classic app (Preferred)
1. Tap on **automation** in the lower bar.
2. Tap on the SmartApps tab on top.
3. Tap on **Add a SmartApp**.
4. Scroll down to end and tap on **My Apps**.
5. You should find `Blockchain Event Logger` here. Tap it and proceed to setup.
6. Select which devices' data you want to log into the Xooa Blockchain.
7. **Save**

### SmartThings new app
1. Tap on `automations` in lower bar.
2. Tap on `Add`(in android) or `+`(in IOS).
3. If prompted, select The location you want to add the app to.
4. Tap on `Done`(in IOS).
5. You should find `Blockchain Event Logger` at the very end of the page. It may take few seconds to show up. Tap it and continue to setup.
6. Select which devices' data you want to log into the Xooa Blockchain.

## Event Viewer SmartApp Setup (SmartThings IDE)
Follow the same steps as followed for `Event Logger SmartApp Setup (SmartThings IDE)` with following exceptions:
1. use `blockchain-event-viewer.groovy` instead of `blockchain-event-logger.groovy` from **smartThings-xooa** github repo.
2. Skip `Event Logger SmartApp Configuration` steps.

### Using Event Viewer smartApp
#### SmartThings classic app (Preferred)
1. Tap on **automation** in the lower bar.
2. Tap on the SmartApps tab on top.
3. Tap on **Add a SmartApp**.
4. Scroll down to end and tap on **My Apps**.
5. You should find `Blockchain Event Viewer` here. Tap it.
6. Enter **App Id** provided in dashboard of Xooa under `Basic Information` tab.
7. Enter the **API Token** provided in Xooa dashboard under `Identities` tab.
8. Click **Next** to proceed to view devices logged with the blockchain.
9. Click on any device to view all the past logged events of that particular device.
10. Click on **Save** when done to store **App Id** and **API Token** with smartApp for future uses.

#### SmartThings new app
1. Tap on `automations` in lower bar.
2. Tap on `Add`(in android) or `+`(in IOS).
3. If prompted, select The location you want to add the app to.
4. Tap on `Done`(in IOS).
5. You should find `Blockchain Event Viewer` at the very end of the page. It may take few seconds to show up. Tap it.
6. Enter **App Id** provided in dashboard of Xooa under `Basic Information` tab.
7. Enter the **API Token** provided in Xooa dashboard under `Identities` tab.
8. Click **Next** to proceed to view devices logged with the blockchain.
9. Click on any device to view all the past logged events of that particular device.
10. Click on **Save** when done to store **App Id** and **API Token** with smartApp for future uses.
