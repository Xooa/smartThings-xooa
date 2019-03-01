
This page provides a step-by-step tutorial to integrate a blockchain SmartApp with Xooa's blockchain platform-as-a-service (PaaS).



# Overview

This repository contains an event logger SmartApp. The SmartApp allows you to read and write events from a blockchain ledger. With these instructions, you can provide an immutable ledger for your connected devices. For example, providing a trusted audit of usage history. 

**To enable the app you need to:**

Set up the Event Logger SmartApp and the Event Viewer SmartApp via both the SmartThings IDE and your mobile device (see below).

## Get Xooa API token


1. Log in to Xooa console.

2. Go to **SmartThings > Get Started**.

3. Click **Generate Xooa API Token**.

4. Copy and store the API token.

___

## Set up the Event Logger SmartApp (IDE)

1. Log in with your Samsung SmartThings account to the SmartThings IDE at <https://graph.api.smartthings.com/>.

2. From the menu, click **My SmartApps**.

You now need to publish the app.  You can do this with or without GitHub integration:

**Without GitHub integration:**

1. Click **New SmartApp**.

2. Copy the *groovy* script from the **smartthings** folder of the  **samples** repo: <https://raw.githubusercontent.com/Xooa/samples/master/smartthings/blockchain-event-logger.groovy>
   
3. On the **From code** tab, paste the code.

4. Click **Create**.

5. Click **Save**, and then click **Publish > For me**.

**With GitHub integration** 

> **Tip:** If you haven't already set up GitHub to work with SmartThings, here's the community FAQ on the subject <https://community.smartthings.com/t/faq-github-integration-how-to-add-and-update-from-repositories/39046>

1. From the menu, click **My SmartApps**, and then click **Settings**.   
2. Click **Add new repository**.
3. Add the GitHub repo to your IDE with the following parameters:
   * `Owner`: xooa
   * `Name`: smartthings-xooa
   * `Branch`: master
4. Click **Save**, click **Update from Repo**, and then click **samples (master)**.
5. Select **blockchain-event-logger.groovy** from the **New (only in GitHub)** column.
6. Click **Publish**.
7. Click  **Execute Update**.




## Set up the Event Logger SmartApp (Android/iOS)
> **Tip:** There are two apps available for SmartThings in the Google Play store. We recommend the classic app over the new app.

Before you begin, ensure that you:

* Have the SmartThings app installed on your phone
* Have at least one location and one device defined
* Are using the same login ID as your developer account


### SmartThings classic app (preferred)

1. Open the SmartThings app and, in the bottom navigation bar, tap **Automation**.

2. Tap the SmartApps tab on top.

3. Tap **Add a SmartApp**.

4. Scroll to the last entry and then tap **My Apps**.

5. Tap the `Blockchain Event Logger` app.

6. Select the devices you want logging to your blockchain.

7. Enter the **Xooa API token** you noted earlier.

9. Click **Save**.

### SmartThings new app

1. Open the SmartThings app and, at the bottom navigation bar, tap **Automations**.

2. Tap **Add** (Android) or **plus** **+** (iOS).

3. If prompted, select the location you want to add the app to.

4. Tap **Done** (iOS).

5. Find `Blockchain Event Logger`, usually it will appear last and may take a few seconds to appear.

6. Tap it to continue setting it up.

7. Select which devices you want to log in the Xooa blockchain.

8. Enter the **Xooa API token** you noted earlier.

## Set up the Event Viewer SmartApp (IDE)

Follow the same steps as `Blockchain Event Logger SmartApp Setup (SmartThings IDE)` except:

1. For section **Without GitHub integration**, copy this *groovy* script from the **smartthings** folder of the  **samples** repo: <https://raw.githubusercontent.com/Xooa/samples/master/smartthings/blockchain-event-viewer.groovy>

2. For section **With GitHub integration**,in step 7, choose `blockchain-event-viewer.groovy` from **samples** GitHub repo.


## Set up the Event Viewer SmartApp (Android/iOS)

Follow the same steps as `Blockchain Event Logger SmartApp Setup (Android/iOS)` except use the Event Viewer SmartApp instead. 