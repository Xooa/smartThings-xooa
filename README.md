
This page is a step-by-step tutorial to integrate the blockchain smartApp with Xooa “blockchain as a service”.
The repository in use for this example is <https://github.com/Xooa/smartThings-xooa>


# Overview
There are two components in this repo, the blockchain chaincode & smartthings smartapp. Blockchain chaincode is deployed in Xooa console and smartapp in smartthings IDE. 
Xooa exposes a permeant cloud end point to smarrthings to enable cloud-to-cloud integration while maintaining peer to peer capabilities of blockchain.

## Deploy  smartthings blockchain chaincode in xooa using github integration 
1. [optional] If you wish, fork repository from <https://github.com/Xooa/smartThings-xooa> to your Github account.
2. Login to Xooa (create an account in xooa if needed) at <https://xooa.com/blockchain
4. From the main dashboard under apps, use **Deploy New** on top left corner to start the deployment process. If it is  first time you deploy a blockchain chaincode app with Xooa you will need to authorize Xooa account to your Github account (Github integration is mandatory for deploying blockchain chaincode in Xooa). Click or tap **Connect to GitHub**. follow the instructions until you successfully authorized your Xooa account to your github account. 
[add screenshot here]
6. Enter **App Name** and **Description** for the app.  Then **Next**.
7. Search for **smartThings-xooa** repo, or if you forked under a different name use that name. A list of repositories matching the search criteria are shown. Select the repo the, then **Next**.
8. The next step is deployment details. Select **master** as the branch and **smartThings-xooa** as the chaincode from the dropdown available. Then **Deploy**.
9. Sit down and relax while Xooa does the blockchain heavy lifting for you. You will be redirected to app dashboard once the deployment is complete.
10. Copy the **App ID** from the dashboard under `Basic Information` tab. It will be required later to connect the smartapp.
11. From the app dashboard, navigate to **Identities** tab.
12. For the identity available, click **Show API token** and copy the token value. We will be using this for authorizing the API requests. This will be your `API Token` needed to authorize **smartApp**.

___

## SmartApp Setup (SmartThings IDE)
1. Login with your Samsung  SmartThings account to the SmartThings IDE at <https://graph.api.smartthings.com>.
2. Navigate to **My SmartApps** 
3. Now you will publish the app.  Steps depends if wish to use github integration or not. 
**Not using github integration:**
a. From top menu, select `New SmartApp`.
b. Naviagte to **From code** tab. 
c. Locate `blockchain-event-logger.groovy` in **smartThings-xooa** github repo.
c. Click or tap **Raw** 
d. Select all and Copy the code
e. Paste it in the **From Code** section in smartThings console, click or tap **create**, **save**, **Publish -> For me**.
**Using Github integration** (if you ahvnet already set it up here there is FAQ in the community https://community.smartthings.com/t/faq-github-integration-how-to-add-and-update-from-repositories/39046)
a. Sill under Smartapps tab, Select  **Settings**.
b. Select  **Add new repository**.
d. Add the Github Repo to your IDE with the following settings:
	* Owner: xooa (or your account if you forked it)
	* Name: smartThings-xooa (or the name you used to fork it)
	* Branch: master
e. Click on **Save**.
f. Click on **Update from Repo**.
g. Click on **smartThings-xooa (master)**.
g. Select **blockchain-event-logger.groovy** from **New(only in GitHub)** column.
i. Check the **Publish** checkbox
j. Click on **Execute Update**.

### SmartApp Configuration (still at smartthings IDE)
4. Select on the smartApp you just deployed.
5. Navigate to **App Settings**.
2. Navigate to **Settings** option.
[insert screenshot here]
3. Enter app Id provided in dashboard of Xooa under `Basic Information` tab.
4. Enter the API Token provided in Xooa dashboard under `Identities` tab.
5. Click or tap **Update**.

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
5. You should find `Blockchain Event Logger` here. Tap it and proceed to setup.
6. Select which devices' data you want to log into the Xooa Blockchain.
7. **Save**

### SmartThings new app
1. Tap on `automations` in lower bar.
2. Tap on `Add`.
3. If prompted, select The location you want to add the app to.
4. You should find `Blockchain Event Logger` at the very end of the page. It may take few seconds to show up. Tap it and continue to setup.
5. Select which devices' data you want to log into the Xooa Blockchain.

