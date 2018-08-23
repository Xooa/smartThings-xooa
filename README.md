
This page provides a step-by-step tutorial to integrate a blockchain SmartApp with xooa's blockchain-as-a-service (BaaS).

The repository used in this example is <https://github.com/Xooa/smartThings-xooa>

# Overview

There are two components in this repo, the blockchain chaincode (henceforth chaincode) and the SmartThings SmartApp. The chaincode is deployed via the xooa console; the SmartApp via the SmartThings IDE.

A permanent cloud end-point to SmartThings is provided by xooa to enable cloud-to-cloud integration while maintaining peer-to-peer capabilities of blockchain.

## Deploy the SmartThings chaincode using GitHub integration

Before you begin you may want to fork the [SmartThings-xooa](https://github.com/Xooa/smartThings-xooa) repository.

1. Log in or create a xooa account at <https://xooa.com/blockchain>

2. Click **apps**, then **Deploy New**. If this is your first time deploy a chaincode app with xooa, you will need to authorize xooa with your GitHub account.

3. Click **Connect to GitHub**.

4. Follow the onscreen instructions to complete the authorization process.

<img src="https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_logging.png" alt="github logging" width="500px"/>

<img src="https://github.com/Xooa/smartThings-xooa/blob/aa7a46efde038f15ad55cda8f606d9460b7c2ee4/screenshots/ScreenShot_authorizing.jpg" alt="github authorizing" width="500px"/>

5. Enter a name and description for your app, and then click **Next**.

6. Search for the **SmartThings-xooa** repo (or your fork). A list of repositories matching the search criteria are shown.

7. Select the repo, and then click **Next**. The deployment details appear.

8. Select **master** branch and **SmartThings-xooa** as the chaincode, and then click **Deploy**.

9. Relax while Xooa does the blockchain heavy lifting for you. You will be redirected to app dashboard when the deployment completes.

10. Copy the **Xooa app ID** from the `Basic Information` tab on the Apps dashboard. You will need this ID later to connect the SmartApp.

11. From the Apps dashboard, navigate to the **Identities** tab.

12. For the available identity, click **Show API token** and copy the token value. You need this token to authorize API requests to the **SmartApp**.

___

## Event Logger SmartApp Setup (SmartThings IDE)

1. Log in with your Samsung SmartThings account to the SmartThings IDE at <https://graph.api.SmartThings.com>.

2. Navigate to **My SmartApps**.

You now need to publish the app.  You can do this with or without GitHub integration:

**Without GitHub integration:**

1. From main menu, select `New SmartApp`.

2. Navigate to **From code**.

3. Locate `blockchain-event-logger.groovy` in the **SmartThings-xooa** GitHub repo.

4. Click **Raw**.

5. Select all and copy the code.

6. Paste your selection in the **From Code** section in the SmartThings console, and then click **create**.

7. Click **save**.

8. Click **Publish -> For me**.

**With GitHub integration** (if you haven't already set up GitHub to work with SmartThings, here is the community FAQ on the subject <https://community.SmartThings.com/t/faq-github-integration-how-to-add-and-update-from-repositories/39046>)

1. Still under SmartApps tab, select **Settings**.   
2. Select **Add new repository**.
3.  Add the GitHub repo to your IDE with the following parameters:
    * `Owner`: xooa
    * `Name`: SmartThings-xooa
    * `Branch`: master
4. Click **Save**.
5. Click **Update from Repo**.
6. Click **SmartThings-xooa (master)**.
7. Select **blockchain-event-logger.groovy** from the **New (only in GitHub)** column.
8.  Select **Publish**.
9.  Click  **Execute Update**.


There are two apps available for SmartThings in the Google Play store. We recommend the classic app over the new app.

## Event Logger SmartApp Setup (Smartphone)

1. Ensure you have the SmartThings app installed on your phone with at least one location and one device defined.

2. Ensure you are using the same login ID you used for your developer account to log in to the app.

3. Open your SmartThings app on your smartphone.

### SmartThings Classic App (Preferred)

1. Tap **automation** in the lower bar.

2. Tap the SmartApps tab on top.

3. Tap **Add a SmartApp**.

4. Scroll to the bottom and then tap **My Apps**.

5. Identify the `Blockchain Event Logger` app and tap it to proceed.

6. Select the devices you want to log in the xooa blockchain.

7. Scroll down and enter the **Xooa app ID** provided in the xooa dashboard under `Basic Information`.

8. Enter the **Xooa participant API token** provided in xooa dashboard under `Identities`.

9. Click **Save**

### SmartThings New App

1. Tap `automations` in lower bar.

2. Tap **Add** (Android) or **+** (iOS).

3. If prompted, select the location you want to add the app to.

4. Tap **Done** (iOS).

5. Find `Blockchain Event Logger`, usually it will appear last and may take a few seconds to appear.

6. Tap it to continue setting it up.

7. Select which devices you want to log in the xooa blockchain.

8. Scroll down and enter the **Xooa app ID** provided in the xooa dashboard under `Basic Information`.

9. Enter the **Xooa participant API token** provided in xooa dashboard under `Identities`.

## Event Viewer SmartApp Setup (SmartThings IDE)

Follow the same steps as `Event Logger SmartApp Setup (SmartThings IDE)` but:

1. Use `blockchain-event-viewer.groovy` instead of `blockchain-event-logger.groovy` from **SmartThings-xooa** GitHub repo.

2. Skip `Event Logger SmartApp Configuration` steps.

### Using Event Viewer SmartApp

#### SmartThings classic app (Preferred)

1. Tap **automation** in the lower bar.

2. Tap the SmartApps tab on top.

3. Tap **Add a SmartApp**.

4. Scroll to the bottom and tap **My Apps**.

5. Find `Blockchain Event Viewer` and tap it.

6. Enter the **Xooa app ID** provided in the xooa dashboard under `Basic Information`.

7. Enter the **Xooa participant API token** provided in the xooa dashboard under `Identities`.

8. Enter the **Location Id** for which you want to view events. Keep it as it is if you want to view events of your own location.

9. Click **Next** to proceed to view the devices logging to the blockchain.

10. Click any device to view the past logged events for that device.

11. Input the date for which you want to view the past logged events.(Last logged date is preset)

12. Click **Save** to store **Xooa app ID**, **API Token** and **Location Id** with SmartApp for future uses.

#### SmartThings new app

1. Tap **automations** in lower bar.

2. Tap **Add**(in android) or **+**(in IOS).

3. If prompted, select the location you want to add the app to.

4. Tap **Done**(in IOS).

5. Find `Blockchain Event Viewer`, usually it appears at the bottom of the page and may take a few seconds to appear.

6. Tap the app.

6. Enter the **Xooa app ID** provided in the xooa dashboard under `Basic Information`.

7. Enter the **Xooa participant API token** provided in xooa dashboard under `Identities`.

8. Enter the **Location Id** for which you want to view events. Keep it as it is if you want to view events of your own location.

9. Click **Next** to proceed to view devices logging to the blockchain.

10. Click on any device to view the past logged events for that device.

11. Input the date for which you want to view the past logged events.(Latest logged date is preset)

12. Click **Save** to store the **Xooa app ID**, **API Token** and **Location Id** with SmartApp for future uses.
