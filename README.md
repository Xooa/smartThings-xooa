# Integrating SmartThings App with Xooa
This document will provide step-by-step tutorial to integrate a smartApp with Xooa.
The objective is to make entries for each event of a device into the hyperledger blockchain and retrieve it.

The repository in use for this example is <https://github.com/Xooa/smartthings-xooa>.

## Xooa Setup
1. Fork repository from <https://github.com/Xooa/smartthings-xooa> to your account.
2. Go to Xooa backend webservice located at `https://d3u0qtw0cf1xi7.cloudfront.net/#`
3. Login to the platform. If you dont have an account, you can click on **Sign Up** to create a new account. Once logged in, you will be taken to the main Xooa dashboard.
4. From the main dashboard, click on **Deploy New** on top left corner to start the deployment process.
5. The first step is to connect your new Xooa login to your Github account. In the main dashboard you should see the **Connect to Git** option. Click it and in the new browser window that opens up, allow Xooa to access your github account. Select the Github account you hold in which you have forked smartthings-xooa.
6. Enter **App Name** and **Description** for the app. These fields are for you to identify a particular app from list so make sure your values are relevant to the app you are deploying. Click **Next**.
7. Select **Private** and search for **smartthings-xooa**. A list of repositories matching the search criteria are shown below. In our case, **smartthings-xooa** will show up. Click **Select** and then, click **Next**.
8. The next step is deployment details. Select **master** as the branch and **smartthings-xooa** as the chaincode from the dropdown available. Click **Deploy**.
9. Sit back and relax while Xooa sets up the backend. Once *Xooa* is finished with the process you will see Success for all the process listed on screen.
10. You will be redirected to app dashboard once the deployment is complete.

### Accessing Xooa dashboard
1. Take note of **App ID** from the dashboard under `Basic Information` tab. It will be required later to connect the **smartApp**.
2. From the app dashboard, navigate to **Identities** tab.
3. For the identity available, click **Show API token** and copy the token value. We will be using this for authorizing the API requests. This will be your `API Token` needed to authorize **smartApp**.

___

## SmartApp Setup (Server)
1. Create a developer account or login with your existing account at <https://graph.api.smartthings.com>.
2. Go to My SmartApps tab from top menu, then `New SmartApp`.
3. Go to **From code** option.
4. Copy the code from `xooa_logger.groovy` present in **smartthings-xooa** repository and paste it into the editor in developer portal of smartThings.
5. Click on **create**.
6. Click on **save**.
7. Click on **Publish -> For me**.
8. Click on App Settings.
9. Click on Settings option.
10. Enter app Id provided in dashboard of Xooa under `Basic Information` tab.
11. Enter the API Token provided in Xooa dashboard under `Identities` tab.
12. Click on **Update**.

The below document is currently for android devices (need to verify for iOS).
There are 2 apps available for smartThings in google play store . We recommend the classic app over new app due to frequent glitches in new app being faced currently.
## SmartApp Setup (Smartphone)
1. Make sure you have smartThings app installed on your phone with at least 1 location and 1 device defined.
2. Also make sure that you use same login ID to log into the app as used in developer account login.
3. Open your SmartThings app on your smartphone.

### SmartThings classic app (Preferred)
1. Tap on **automation** in the lower bar.
2. Tap on the SmartApps tab on top.
3. Tap on **Add a SmartApp**.
4. Scroll down to end and tap on **My Apps**.
5. You should find `Xooa Logger` here. Tap it and proceed to setup.
6. Select which devices' data you want to log into the Xooa Blockchain.

### SmartThings new app
1. Tap on `automations` in lower bar.
2. Tap on `Add`.
3. If prompted, select The location you want to add the app to.
4. You should find `Xooa logger` at the very end of the page. It may take few seconds to show up. Tap it and continue to setup.
5. Select which devices' data you want to log into the Xooa Blockchain.
___

## Log some events
1. Take some actions with the smartThings devices you associated with the smartApp.
2. If they are simulated devices, then go to your smartThings app and tap your simulated devices to make some events.


## Accessing the past events
1. In the Xooa dashboard, from the very top menu bar, go to **Explore** option.
2. Click on **Transactions** tab from the menu below.
3. You will find all your events logged in as transactions displaying latest transaction at top.
4. Click on Tx Id you want to see details for.
5. Under the writes column, you will find the data that has been logged into the blockchain for this particular event.